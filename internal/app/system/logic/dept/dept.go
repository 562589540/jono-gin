package dept

import (
	"context"
	"fmt"
	"github.com/562589540/jono-gin/internal/app/common/dal"
	"github.com/562589540/jono-gin/internal/app/system/dto"
	"github.com/562589540/jono-gin/internal/app/system/model"
	"github.com/562589540/jono-gin/internal/app/system/service"
	"github.com/562589540/jono-gin/internal/constants"
)

var deptService service.IDeptService

type Service struct{}

func New() service.IDeptService {
	if deptService == nil {
		deptService = &Service{}
	}
	return deptService
}

func (m *Service) Dao(ctx context.Context) dal.IDeptDo {
	return dal.Dept.WithContext(ctx)
}

func (m *Service) Create(ctx context.Context, data *dto.DeptAddReq) error {
	err := m.Dao(ctx).Create(data.ToModel(&model.Dept{}))
	if err != nil {
		return err
	}
	return nil
}

func (m *Service) Delete(ctx context.Context, id uint) error {
	return dal.Q.Transaction(func(tx *dal.Query) error {
		err := m.deleteDeptRecursively(ctx, id, tx)
		if err != nil {
			return err
		}
		return nil
	})
}

func (m *Service) Update(ctx context.Context, data *dto.DeptAddReq) error {
	dao := dal.Dept
	mModel, err := dao.WithContext(ctx).Where(dao.ID.Eq(data.ID)).First()
	if err != nil {
		return fmt.Errorf(constants.NoDataFound)
	}
	return dao.WithContext(ctx).Save(data.ToModel(mModel))
}

func (m *Service) List(ctx context.Context, _ *dto.DeptSearchReq) ([]dto.Dept, int64, error) {
	dao := dal.Dept
	list, err := dao.WithContext(ctx).Order(dao.Sort.Asc()).Find()
	if err != nil {
		return nil, 0, err
	}
	mDTOList := make([]dto.Dept, len(list))
	for i, item := range list {
		mDTOList[i] = dto.Dept{}.FromModel(item)
	}
	return mDTOList, 0, nil
}

// 递归删除部门
func (m *Service) deleteDeptRecursively(ctx context.Context, deptId uint, tx *dal.Query) error {
	d := tx.Dept
	//查询下级部门
	subDepts, err := d.WithContext(ctx).Where(d.ParentID.Eq(deptId)).Find()
	if err != nil {
		return err
	}

	//全删
	for _, dept := range subDepts {
		if err = m.deleteDeptRecursively(ctx, dept.ID, tx); err != nil {
			return err
		}
	}

	//清除关联
	dept := new(model.Dept)
	dept.ID = deptId
	if err = d.Admins.WithContext(ctx).Model(dept).Clear(); err != nil {
		return err
	}

	//删除自己
	if _, err = d.WithContext(ctx).DeleteByID(deptId); err != nil {
		return err
	}

	return nil
}

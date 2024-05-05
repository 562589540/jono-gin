package dictData

import (
	"context"
	"fmt"
	"github.com/562589540/jono-gin/ghub/gutils"
	"github.com/562589540/jono-gin/internal/app/common/dal"
	"github.com/562589540/jono-gin/internal/app/system/dto"
	"github.com/562589540/jono-gin/internal/app/system/model"
	"github.com/562589540/jono-gin/internal/app/system/service"
	"github.com/562589540/jono-gin/internal/constants"
)

var dictDataService service.IDictDataService

type Service struct {
	contextService service.IContextService
}

func New(contextService service.IContextService) service.IDictDataService {
	if dictDataService == nil {
		dictDataService = &Service{
			contextService,
		}
	}
	return dictDataService
}

func (m *Service) Dao(ctx context.Context) dal.IDictDataDo {
	return dal.DictData.WithContext(ctx)
}

func (m *Service) Create(ctx context.Context, data *dto.DictDataAddReq) error {
	mModel := new(model.DictData)
	if err := gutils.Copy(mModel, data); err != nil {
		return err
	}
	mModel.SetStatus(data.Status)
	mModel.SetIsDefault(data.IsDefault)

	userModel, err := m.contextService.GetLoginUserModel(ctx)
	if err == nil {
		mModel.CreateBy = int64(userModel.ID)
	}

	return m.Dao(ctx).Create(mModel)
}

func (m *Service) Delete(ctx context.Context, ids []uint) error {
	dao := dal.DictData
	info, err := dao.WithContext(ctx).Where(dao.DictCode.In(ids...)).Delete()
	if err != nil {
		return err
	}
	if info.RowsAffected == 0 {
		return fmt.Errorf(constants.DeleteError)
	}
	return nil
}

func (m *Service) Update(ctx context.Context, data *dto.DictDataUpdateReq) error {
	return dal.Q.Transaction(func(tx *dal.Query) error {
		dao := tx.DictData
		mModel, err := dao.WithContext(ctx).Where(dao.DictCode.Eq(data.ID)).First()
		if err != nil {
			return fmt.Errorf(constants.NoDataFound)
		}
		if err = gutils.Copy(mModel, data); err != nil {
			return err
		}
		userModel, err := m.contextService.GetLoginUserModel(ctx)
		if err == nil {
			mModel.UpdateBy = int64(userModel.ID)
		}
		mModel.SetStatus(data.Status)
		mModel.SetIsDefault(data.IsDefault)
		if mModel.IsDefault {
			_, err = dao.WithContext(ctx).Where(dao.DictCode.Neq(mModel.DictCode), dao.DictType.Eq(mModel.DictType)).
				Update(dao.IsDefault, false)
			if err != nil {
				return err
			}
		}
		return dao.WithContext(ctx).Save(mModel)
	})
}

func (m *Service) List(ctx context.Context, search *dto.DictDataSearchReq) ([]dto.DictData, int64, error) {
	dao := dal.DictData
	q := dao.WithContext(ctx)

	if search.DictLabel != "" {
		q = q.Where(dao.DictLabel.Eq(search.DictLabel))
	}
	if search.Status != "" {
		q = q.Where(dao.Status.Is(search.Status == "1"))
	}
	if search.DictType != "" {
		q = q.Where(dao.DictType.Eq(search.DictType))
	}

	list := make([]dto.DictData, 0)
	count, err := q.Order(dao.DictCode.Asc()).ScanByPage(&list, search.GetOffset(), search.GetLimit())
	if err != nil {
		return nil, 0, err
	}
	return list, count, nil
}

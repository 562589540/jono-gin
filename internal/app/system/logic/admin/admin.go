package admin

import (
	"context"
	"fmt"
	"github.com/562589540/jono-gin/ghub"
	"github.com/562589540/jono-gin/ghub/gutils"
	"github.com/562589540/jono-gin/internal/app/common/dal"
	"github.com/562589540/jono-gin/internal/app/system/dto"
	"github.com/562589540/jono-gin/internal/app/system/model"
	"github.com/562589540/jono-gin/internal/app/system/service"
	"github.com/562589540/jono-gin/internal/constants"
	"time"
)

var adminService service.IAdminService

type Service struct{}

func New() service.IAdminService {
	if adminService == nil {
		adminService = &Service{}
	}
	return adminService
}

func (m *Service) Dao(ctx context.Context) dal.IAdminDo {
	return dal.Admin.WithContext(ctx)
}

func (m *Service) Create(ctx context.Context, data *dto.AdminAddReq) error {
	mModel, err := data.ToModel(&model.Admin{})
	if err != nil {
		return err
	}
	return m.Dao(ctx).Create(mModel)
}

func (m *Service) Delete(ctx context.Context, id uint) error {
	return dal.Q.Transaction(func(tx *dal.Query) error {
		dao := tx.Admin
		mModel := &model.Admin{}
		mModel.ID = id
		err := dao.RoleSign.WithContext(ctx).Model(mModel).Clear()
		if err != nil {
			return err
		}
		_, err = dao.WithContext(ctx).Where(dao.ID.Eq(id)).Delete()
		if err != nil {
			return err
		}
		return nil
	})
}

func (m *Service) BatchDelete(ctx context.Context, ids []uint) error {
	return dal.Q.Transaction(func(tx *dal.Query) error {
		dao := tx.Admin
		for i := 0; i < len(ids); i++ {
			mModel := &model.Admin{}
			mModel.ID = ids[i]
			err := dao.RoleSign.WithContext(ctx).Model(mModel).Clear()
			if err != nil {
				return err
			}
		}
		_, err := dao.WithContext(ctx).Where(dao.ID.In(ids...)).Delete()
		if err != nil {
			return err
		}
		return nil
	})
}

func (m *Service) Update(ctx context.Context, data *dto.AdminUpdateReq) error {
	mModel, err := m.Dao(ctx).Where(dal.Admin.ID.Eq(data.ID)).First()
	if err != nil {
		return fmt.Errorf(constants.NoDataFound)
	}
	return m.Dao(ctx).Save(data.ToModel(mModel))
}

func (m *Service) UpdateStatus(ctx context.Context, data *dto.AdminUpdateStatusReq) error {
	dao := dal.Admin
	_, err := m.Dao(ctx).Where(dao.ID.Eq(data.ID)).Update(dao.Status, data.GetStatus())
	if err != nil {
		return err
	}
	return nil
}

func (m *Service) UpdatePassword(ctx context.Context, data *dto.AdminUpdatePassReq) error {
	password, err := data.SetPassword()
	if err != nil {
		return err
	}
	dao := dal.Admin
	_, err = m.Dao(ctx).Where(dao.ID.Eq(data.ID)).Update(dao.Password, password)
	if err != nil {
		return err
	}
	return nil
}

func (m *Service) UpdateAvatar(ctx context.Context, data *dto.AdminUpdateAvatar) error {
	image, err := gutils.SaveBase64Image(data.Avatar, "")
	if err != nil {
		return err
	}
	dao := dal.Admin
	_, err = m.Dao(ctx).Where(dao.ID.Eq(data.ID)).Update(dao.Avatar, image)
	if err != nil {
		return err
	}
	return nil
}

func (m *Service) UpdateRole(ctx context.Context, data *dto.AdminUpdateRoleReq) error {
	return dal.Q.Transaction(func(tx *dal.Query) error {
		dao := tx.Admin
		mModel := &model.Admin{}
		mModel.ID = data.ID
		//先清空所有关联
		if err := dao.RoleSign.WithContext(ctx).Model(mModel).Clear(); err != nil {
			return err
		}
		if len(data.RoleIds) > 0 {
			// 如果提供了新的角色ID，查询这些角色
			roles, err := tx.Roles.WithContext(ctx).Where(tx.Roles.ID.In(data.RoleIds...)).Find()
			if err != nil {
				return err
			}
			// 为管理员添加新的角色关系
			if err = dao.RoleSign.WithContext(ctx).Model(mModel).Append(roles...); err != nil {
				return err
			}
		}
		return nil
	})
}

func (m *Service) GetUserRoleIds(ctx context.Context, id uint) ([]uint, error) {
	mModel := &model.Admin{}
	mModel.ID = id
	roles, err := dal.Admin.RoleSign.WithContext(ctx).Model(mModel).Find()
	if err != nil {
		return nil, err
	}
	roleIDs := make([]uint, len(roles))
	for i, role := range roles {
		roleIDs[i] = role.ID
	}
	return roleIDs, nil
}

func (m *Service) List(ctx context.Context, search *dto.AdminSearchReq) ([]dto.Admin, int64, error) {
	dao := dal.Admin
	q := dao.WithContext(ctx).Preload(dao.Dept)

	if search.DeptId != 0 {
		q = q.Where(dao.DeptID.Eq(search.DeptId))
	}

	if search.Mobile != "" {
		q = q.Where(dao.Mobile.Eq(search.Mobile))
	}

	if search.Status != "" {
		q = q.Where(dao.Status.Is(search.Status == "1"))
	}

	if search.UserName != "" {
		q = q.Where(dao.UserName.Like(fmt.Sprintf("%%%s%%", search.UserName)))
	}

	list, total, err := q.Order(dao.ID.Desc()).FindByPage(search.GetOffset(), search.GetLimit())
	if err != nil {
		return nil, 0, err
	}
	mDTOList := make([]dto.Admin, len(list))
	for i, item := range list {
		mDTOList[i] = dto.Admin{}.FromModel(item)
	}
	return mDTOList, total, nil
}

func (m *Service) SetLogin(ctx context.Context, userName, ip string) {
	t := time.Now()
	dao := dal.Admin
	mModel := new(model.Admin)
	mModel.LastLoginIp = ip
	mModel.LastLoginTime = &t
	_, err := m.Dao(ctx).Where(dao.UserName.Eq(userName)).Updates(&mModel)
	if err != nil {
		ghub.Log.Error(err.Error())
	}
}

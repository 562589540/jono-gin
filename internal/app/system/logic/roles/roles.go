package roles

import (
	"context"
	"fmt"
	"github.com/562589540/jono-gin/internal/app/common/dal"
	"github.com/562589540/jono-gin/internal/app/system/dto"
	"github.com/562589540/jono-gin/internal/app/system/model"
	"github.com/562589540/jono-gin/internal/app/system/service"
	"github.com/562589540/jono-gin/internal/constants"
)

var rolesService service.IRolesService

type Service struct{}

func New() service.IRolesService {
	if rolesService == nil {
		rolesService = &Service{}
	}
	return rolesService
}

func (m *Service) Dao(ctx context.Context) dal.IRolesDo {
	return dal.Roles.WithContext(ctx)
}
func (m *Service) Create(ctx context.Context, data *dto.RolesAddReq) error {
	return m.Dao(ctx).Create(data.ToModel(&model.Roles{}, true))
}

func (m *Service) Delete(ctx context.Context, id uint) error {
	if id == 1 {
		return fmt.Errorf("禁止删除最高权限管理员")
	}
	return dal.Q.Transaction(func(tx *dal.Query) error {
		mRolesMode := &model.Roles{}
		mRolesMode.ID = id
		if err := tx.Roles.Menus.WithContext(ctx).Model(mRolesMode).Clear(); err != nil {
			return err
		}
		if _, err := tx.Roles.WithContext(ctx).DeleteByID(id); err != nil {
			return err
		}
		return nil
	})
}

func (m *Service) Update(ctx context.Context, data *dto.RolesUpdateReq) error {
	if data.ID == 1 {
		return fmt.Errorf("禁止修改最高权限管理员")
	}
	dao := dal.Roles
	mModel, err := dao.WithContext(ctx).Where(dao.ID.Eq(data.ID)).First()
	if err != nil {
		return fmt.Errorf(constants.NoDataFound)
	}
	return dao.WithContext(ctx).Save(data.ToModel(mModel, false))
}

func (m *Service) List(ctx context.Context, search *dto.RolesSearchReq) ([]dto.Role, int64, error) {
	dao := dal.Roles
	q := dao.WithContext(ctx)
	if search.Code != "" {
		q = q.Where(dao.Code.Eq(search.Code))
	}

	if search.Name != "" {
		q = q.Where(dao.Name.Eq(search.Name))
	}

	if search.Status != "" {
		q = q.Where(dao.Status.Is(search.Status == "1"))
	}

	list, total, err := q.FindByPage(search.GetOffset(), search.GetLimit())

	if err != nil {
		return nil, 0, err
	}

	mDTOList := make([]dto.Role, len(list))
	for i, item := range list {
		mDTOList[i] = dto.Role{}.FromModel(item)
	}
	return mDTOList, total, nil
}

// UpdateRoleMenusAuth 更新角色的菜单访问权限
func (m *Service) UpdateRoleMenusAuth(ctx context.Context, data *dto.RolesPowerReq) error {
	if data.ID == 1 {
		return fmt.Errorf("禁止修改最高权限管理员")
	}
	return dal.Q.Transaction(func(tx *dal.Query) error {
		dao := tx.Roles
		mRolesMode := &model.Roles{}
		mRolesMode.ID = data.ID
		if err := dao.Menus.WithContext(ctx).Model(mRolesMode).Clear(); err != nil {
			return err
		}
		if len(data.IDS) > 0 {
			menus, err := tx.Menu.WithContext(ctx).Where(tx.Menu.ID.In(data.IDS...)).Find()
			if err != nil {
				return err
			}
			// 为角色添加新的菜单关系
			if err = dao.Menus.WithContext(ctx).Model(mRolesMode).Append(menus...); err != nil {
				return err
			}
		}
		return nil
	})
}

// GetRoleMenuIds 获取角色的菜单Ids
func (m *Service) GetRoleMenuIds(ctx context.Context, id uint) ([]uint, error) {
	dao := dal.Roles
	mRolesMode := &model.Roles{}
	mRolesMode.ID = id
	menus, err := dao.Menus.WithContext(ctx).Model(mRolesMode).Find()
	if err != nil {
		return nil, err
	}
	menuIDs := make([]uint, len(menus))
	for i, menu := range menus {
		menuIDs[i] = menu.ID
	}
	return menuIDs, nil
}

// GetAllRoleList 获取所有角色信息 用户菜单使用的
func (m *Service) GetAllRoleList(ctx context.Context) ([]dto.Role, error) {
	dao := dal.Roles
	mModeList, err := dao.WithContext(ctx).Where(dao.Status.Is(true)).Find()
	if err != nil {
		return nil, err
	}
	mDTOList := make([]dto.Role, len(mModeList))
	for i, roles := range mModeList {
		mDTOList[i] = dto.Role{
			ID:   roles.ID,
			Name: roles.Name,
		}
	}
	return mDTOList, nil
}

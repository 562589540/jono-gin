package menu

import (
	"context"
	"fmt"
	"github.com/562589540/jono-gin/ghub"
	"github.com/562589540/jono-gin/ghub/gbootstrap"
	"github.com/562589540/jono-gin/internal/app/system/dal"
	"github.com/562589540/jono-gin/internal/app/system/dto"
	"github.com/562589540/jono-gin/internal/app/system/logic/bizctx"
	"github.com/562589540/jono-gin/internal/app/system/model"
	"github.com/562589540/jono-gin/internal/app/system/service"
	"github.com/562589540/jono-gin/internal/constants"
	"github.com/gin-gonic/gin"
	"sort"
)

var menuService service.IMenuService

type Service struct{}

func New() service.IMenuService {
	if menuService == nil {
		menuService = &Service{}
	}
	return menuService
}

func (m *Service) Dao(ctx context.Context) dal.IMenuDo {
	return dal.Menu.WithContext(ctx)
}

func (m *Service) Create(ctx context.Context, data *dto.AddMenuReq) error {
	mModel := new(model.Menu)
	if err := ghub.Copy(mModel, data); err != nil {
		return err
	}
	if err := mModel.CheckSaveParam(); err != nil {
		return err
	}
	mModel.CheckParentID()
	return m.Dao(ctx).Create(mModel)

}

func (m *Service) Update(ctx context.Context, data *dto.UpdateMenuReq) error {
	mn := dal.Menu
	mModel, err := mn.WithContext(ctx).Where(mn.ID.Eq(data.ID)).First()
	if err != nil {
		return err
	}
	if err = ghub.Copy(&mModel, data); err != nil {
		return err
	}
	if err = mModel.CheckSaveParam(); err != nil {
		return err
	}
	mModel.CheckParentID()
	return mn.WithContext(ctx).Save(mModel)
}

func (m *Service) Delete(ctx context.Context, id uint) error {
	return dal.Q.Transaction(func(tx *dal.Query) error {
		// 删除子菜单
		if err := m.deleteSubMenus(ctx, tx, id); err != nil {
			return err
		}
		// 最后删除指定的菜单
		return m.deleteMenu(ctx, tx, id)
	})
}

func (m *Service) List(ctx context.Context) ([]dto.Menu, error) {
	mn := dal.Menu
	mDTOs := make([]dto.Menu, 0)
	err := mn.WithContext(ctx).Order(mn.Rank, mn.ID).Scan(&mDTOs)
	if err != nil {
		return nil, err
	}
	return mDTOs, nil
}

func (m *Service) GetRoleMenu(ctx context.Context) ([]dto.RoleMenu, error) {
	mn := dal.Menu
	var mDTOs []dto.RoleMenu
	err := mn.WithContext(ctx).Order(mn.Rank, mn.ID).Scan(&mDTOs)
	if err != nil {
		return nil, err
	}
	return mDTOs, nil
}

func (m *Service) GetRoutes(ctx *gin.Context) ([]*dto.MenuList, error) {
	mn := dal.Menu
	mAdmin, err := bizctx.New().GetLoginUserModel(ctx)
	if err != nil {
		return nil, fmt.Errorf("权限错误2")
	}
	list, err := mn.WithContext(ctx).Preload(mn.Roles).Find()
	if err != nil {
		return nil, err
	}

	roles := make([]uint, len(mAdmin.RoleSign))
	for i := 0; i < len(mAdmin.RoleSign); i++ {
		roles[i] = mAdmin.RoleSign[i].ID
	}

	return m.buildMenu(list, roles, mAdmin.ID), nil
}

func (m *Service) buildMenu(items []*model.Menu, adminRoles []uint, userId uint) []*dto.MenuList {
	var (
		roots       []*dto.MenuList                  //基础菜单
		childrenMap = make(map[uint][]*dto.MenuList) //子菜单
	)

	// 第一步：筛选并构建映射
	for _, item := range items {
		//是菜单
		if item.MenuType != constants.Button {

			hasRole := false
			//检查角色权限
			for _, role := range item.Roles {
				if role.Status && ghub.Contains(adminRoles, role.ID) {
					hasRole = true
					break
				}
			}

			//角色符合一项即可
			if ghub.Contains(gbootstrap.Cfg.System.NotCheckAuthAdminIds, userId) || hasRole {
				if item.ParentID == nil {
					roots = append(roots, m.conv2DTO(item, true))
				} else {
					childrenMap[*item.ParentID] = append(childrenMap[*item.ParentID], m.conv2DTO(item, false))
				}
			}

		}
	}

	// 第二步：排序根菜单和子菜单
	m.sortMenuItems(roots)
	for _, children := range childrenMap {
		m.sortMenuItems(children)
	}

	// 第三步：构建层级关系
	for _, item := range roots {
		if _, ok := childrenMap[item.ID]; ok {
			item.Children = childrenMap[item.ID]
			// 递归为每个子菜单构建子树
			m.buildChildren(item, childrenMap)
		}
	}

	return roots
}

func (m *Service) buildChildren(parent *dto.MenuList, childrenMap map[uint][]*dto.MenuList) {
	for _, child := range parent.Children {
		if _, ok := childrenMap[child.ID]; ok {
			child.Children = childrenMap[child.ID]
			m.buildChildren(child, childrenMap)
		}
	}
}

func (m *Service) sortMenuItems(items []*dto.MenuList) {
	sort.Slice(items, func(i, j int) bool {
		if items[i].Rank == items[j].Rank {
			return items[i].ID < items[j].ID
		}
		return items[i].Rank < items[j].Rank
	})
}

func (m *Service) conv2DTO(mModel *model.Menu, isRoot bool) *dto.MenuList {
	d := &dto.MenuList{}
	d.ID = mModel.ID
	d.Path = mModel.Path
	d.Rank = mModel.Rank
	d.Component = mModel.Component
	d.Meta = &dto.MenuMeta{
		Icon:         mModel.Icon,
		Title:        mModel.Title,
		Api:          mModel.Api,
		ShowLink:     mModel.ShowLink,
		ShowParent:   mModel.ShowParent,
		KeepAlive:    mModel.KeepAlive,
		FrameSrc:     mModel.FrameSrc,
		FrameLoading: mModel.FrameLoading,
		HiddenTag:    mModel.HiddenTag,
		//DynamicLevel:  mModel.DynamicLevel,
		EnterTransition: mModel.EnterTransition,
		LeaveTransition: mModel.LeaveTransition,
		FixedTag:        mModel.FixedTag,
		ExtraIcon:       mModel.ExtraIcon,
		ActivePath:      mModel.ActivePath,
		Rank:            mModel.Rank,
	}
	if !isRoot {
		d.Name = mModel.Name //首菜单不能代名字
		//d.Meta.Roles = []string{"admin", "common"}
		d.Meta.Rank = 0 //只能首菜单带排序
	}
	return d
}

// 查找所有子菜单
func (m *Service) deleteSubMenus(ctx context.Context, tx *dal.Query, parentID uint) error {
	mn := tx.Menu
	subMenus, err := mn.WithContext(ctx).Where(mn.ParentID.Eq(parentID)).Find()
	if err != nil {
		return err
	}
	for _, menu := range subMenus {
		if err = m.deleteSubMenus(ctx, tx, menu.ID); err != nil {
			return err
		}
		if err = m.deleteMenu(ctx, tx, menu.ID); err != nil {
			return err
		}
	}
	return nil
}

// 删除中间表以及菜单
func (m *Service) deleteMenu(ctx context.Context, tx *dal.Query, menuID uint) error {
	mn := tx.Menu
	menu := new(model.Menu)
	menu.ID = menuID
	// 清除与菜单相关联的所有角色
	if err := mn.Roles.WithContext(ctx).Model(menu).Clear(); err != nil {
		return err
	}
	// 删除菜单本身
	_, err := mn.WithContext(ctx).DeleteByID(menuID)
	if err != nil {
		return err
	}
	return nil
}

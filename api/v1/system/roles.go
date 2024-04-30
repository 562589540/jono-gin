package system

import (
	"github.com/562589540/jono-gin/ghub"
	"github.com/562589540/jono-gin/ghub/glibrary/gdto"
	"github.com/562589540/jono-gin/ghub/glibrary/gres"
	"github.com/562589540/jono-gin/internal/app/system/dto"
	"github.com/562589540/jono-gin/internal/app/system/service"
	"github.com/562589540/jono-gin/internal/constants"
	"github.com/gin-gonic/gin"
)

type RolesApi struct {
	rolesService service.IRolesService
}

func NewRolesApi(rolesService service.IRolesService) *RolesApi {
	return &RolesApi{
		rolesService: rolesService,
	}
}

func (m RolesApi) Create(c *gin.Context, req dto.RolesAddReq) (any, error) {
	if err := m.rolesService.Create(c, &req); err != nil {
		return nil, err
	}
	return gres.Response{
		Message: "创建成功",
	}, nil
}

func (m RolesApi) List(c *gin.Context, mDto dto.RolesSearchReq) (any, error) {
	list, total, err := m.rolesService.List(c, &mDto)
	if err != nil {
		return nil, err
	}
	return gres.Response{
		Message: "查询成功",
		Data: gdto.ListRes{
			List:  list,
			Total: total,
		},
	}, nil
}

func (m RolesApi) Update(c *gin.Context, req dto.RolesUpdateReq) (any, error) {
	if err := m.rolesService.Update(c, &req); err != nil {
		return nil, err
	}
	//更细菜单规则缓存
	ghub.EventBus.Publish(constants.RefreshPathToRoles, nil)
	return gres.Response{
		Message: "更新成功",
	}, nil
}

func (m RolesApi) Delete(c *gin.Context, req gdto.IDReq) (any, error) {
	if err := m.rolesService.Delete(c, req.ID); err != nil {
		return nil, err
	}
	//更细菜单规则缓存
	ghub.EventBus.Publish(constants.RefreshPathToRoles, nil)
	return gres.Response{
		Message: "删除成功",
	}, nil
}

func (m RolesApi) UpdateRoleMenusAuth(c *gin.Context, mDto dto.RolesPowerReq) (any, error) {
	if err := m.rolesService.UpdateRoleMenusAuth(c, &mDto); err != nil {
		return nil, err
	}

	//更细菜单规则缓存
	ghub.EventBus.Publish(constants.RefreshPathToRoles, nil)
	return gres.Response{
		Message: "更新成功",
	}, nil
}

func (m RolesApi) GetRoleMenuIds(c *gin.Context, req gdto.IDReq) (any, error) {
	list, err := m.rolesService.GetRoleMenuIds(c, req.ID)
	if err != nil {
		return nil, err
	}
	return gres.Response{
		Message: "查询成功",
		Data:    list,
	}, nil
}

func (m RolesApi) GetAllRoleList(c *gin.Context, _ gdto.EmptyReq) (any, error) {
	list, err := m.rolesService.GetAllRoleList(c)
	if err != nil {
		return nil, err
	}
	return gres.Response{
		Message: "查询成功",
		Data:    list,
	}, nil
}

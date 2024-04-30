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

type MenuApi struct {
	menusService service.IMenuService
}

func NewMenuApi(menusService service.IMenuService) *MenuApi {
	return &MenuApi{
		menusService: menusService,
	}
}

func (m MenuApi) Create(c *gin.Context, req dto.AddMenuReq) (any, error) {
	if err := m.menusService.Create(c, &req); err != nil {
		return nil, err
	}
	//刷新菜单缓存
	ghub.EventBus.Publish(constants.RefreshPathToRoles, nil)
	return gres.Response{
		Message: "创建成功",
	}, nil
}

func (m MenuApi) Update(c *gin.Context, req dto.UpdateMenuReq) (any, error) {
	if err := m.menusService.Update(c, &req); err != nil {
		return nil, err
	}
	//刷新菜单缓存
	ghub.EventBus.Publish(constants.RefreshPathToRoles, nil)
	return gres.Response{
		Message: "更新成功",
	}, nil
}

func (m MenuApi) Delete(c *gin.Context, req gdto.IDReq) (any, error) {
	if err := m.menusService.Delete(c, req.ID); err != nil {
		return nil, err
	}
	//刷新菜单缓存
	ghub.EventBus.Publish(constants.RefreshPathToRoles, nil)
	return gres.Response{
		Message: "删除成功",
	}, nil
}

func (m MenuApi) List(c *gin.Context, _ gdto.EmptyReq) (any, error) {
	list, err := m.menusService.List(c)
	if err != nil {
		return nil, err
	}
	return gres.Response{
		Message: "获取成功",
		Data:    list,
	}, nil
}

func (m MenuApi) GetRoutes(c *gin.Context, _ gdto.EmptyReq) (any, error) {
	list, err := m.menusService.GetRoutes(c)
	if err != nil {
		return nil, err
	}
	return gres.Response{
		Success: true,
		Message: "获取成功",
		Data:    list,
	}, nil
}

func (m MenuApi) GetRoleMenu(c *gin.Context, _ gdto.EmptyReq) (any, error) {
	list, err := m.menusService.GetRoleMenu(c)
	if err != nil {
		return nil, err
	}
	return gres.Response{
		Message: "获取成功",
		Data:    list,
	}, nil
}

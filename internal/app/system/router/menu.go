package router

import (
	"github.com/562589540/jono-gin/api/v1/system"
	"github.com/562589540/jono-gin/ghub/glibrary/grouter"
	"github.com/562589540/jono-gin/internal/app/system/logic/menu"
	"github.com/gin-gonic/gin"
)

func (r *Router) BindMenuController(_ *gin.RouterGroup, auth *gin.RouterGroup) {
	apiController := system.NewMenuApi(menu.New())
	protectedRouter := auth.Group("menu")
	{
		protectedRouter.POST("add", grouter.HandlerFunc(apiController.Create))
		protectedRouter.DELETE("delete", grouter.HandlerFunc(apiController.Delete))
		protectedRouter.PUT("update", grouter.HandlerFunc(apiController.Update))
		protectedRouter.GET("list", grouter.HandlerFunc(apiController.List))
		protectedRouter.GET("getRoleMenu", grouter.HandlerFunc(apiController.GetRoleMenu))
		protectedRouter.GET("getRoutes", grouter.HandlerFunc(apiController.GetRoutes)) //获取所有菜单
	}
}

package router

import (
	"github.com/562589540/jono-gin/api/v1/system"
	"github.com/562589540/jono-gin/ghub/glibrary/grouter"
	"github.com/562589540/jono-gin/internal/app/system/logic/roles"
	"github.com/gin-gonic/gin"
)

func (r *Router) BindRolesController(_ *gin.RouterGroup, auth *gin.RouterGroup) {
	apiController := system.NewRolesApi(roles.New())
	protectedRouter := auth.Group("roles")
	{
		protectedRouter.POST("add", grouter.HandlerFunc(apiController.Create))
		protectedRouter.GET("list", grouter.HandlerFunc(apiController.List))
		protectedRouter.DELETE("delete", grouter.HandlerFunc(apiController.Delete))
		protectedRouter.PUT("update", grouter.HandlerFunc(apiController.Update))
		protectedRouter.PUT("updateRoleMenusAuth", grouter.HandlerFunc(apiController.UpdateRoleMenusAuth))
		protectedRouter.GET("getRoleMenuIds", grouter.HandlerFunc(apiController.GetRoleMenuIds))
		protectedRouter.GET("getAllRoleList", grouter.HandlerFunc(apiController.GetAllRoleList))
	}
}

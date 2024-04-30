package router

import (
	"github.com/562589540/jono-gin/api/v1/system"
	"github.com/562589540/jono-gin/ghub/glibrary/grouter"
	"github.com/562589540/jono-gin/internal/app/system/logic/admin"
	"github.com/gin-gonic/gin"
)

func (r *Router) BindAdminController(_ *gin.RouterGroup, auth *gin.RouterGroup) {
	apiController := system.NewAdminApi(admin.New())
	protectedRouter := auth.Group("admin")
	{
		protectedRouter.POST("add", grouter.HandlerFunc(apiController.Create))
		protectedRouter.DELETE("delete", grouter.HandlerFunc(apiController.Delete))
		protectedRouter.PUT("update", grouter.HandlerFunc(apiController.Update))
		protectedRouter.GET("list", grouter.HandlerFunc(apiController.List))
		protectedRouter.PUT("updateStatus", grouter.HandlerFunc(apiController.UpdateStatus))
		protectedRouter.PUT("updatePassword", grouter.HandlerFunc(apiController.UpdatePassword))
		protectedRouter.PUT("updateAvatar", grouter.HandlerFunc(apiController.UpdateAvatar))
		protectedRouter.PUT("updateRole", grouter.HandlerFunc(apiController.UpdateAdminRole))
		protectedRouter.GET("getAdminRoleIds", grouter.HandlerFunc(apiController.GetAdminRoleIds))
	}
}

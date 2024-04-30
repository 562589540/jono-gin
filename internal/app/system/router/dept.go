package router

import (
	"github.com/562589540/jono-gin/api/v1/system"
	"github.com/562589540/jono-gin/ghub/glibrary/grouter"
	"github.com/562589540/jono-gin/internal/app/system/logic/dept"
	"github.com/gin-gonic/gin"
)

func (r *Router) BindDeptController(_ *gin.RouterGroup, auth *gin.RouterGroup) {
	apiController := system.NewDeptApi(dept.New())
	protectedRouter := auth.Group("dept")
	{
		protectedRouter.POST("add", grouter.HandlerFunc(apiController.Create))
		protectedRouter.DELETE("delete", grouter.HandlerFunc(apiController.Delete))
		protectedRouter.PUT("update", grouter.HandlerFunc(apiController.Update))
		protectedRouter.GET("list", grouter.HandlerFunc(apiController.List))
	}
}

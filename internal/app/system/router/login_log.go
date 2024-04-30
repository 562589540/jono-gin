package router

import (
	"github.com/562589540/jono-gin/api/v1/system"
	"github.com/562589540/jono-gin/ghub/glibrary/grouter"
	"github.com/562589540/jono-gin/internal/app/system/logic/login_log"
	"github.com/gin-gonic/gin"
)

func (r *Router) BindLoginLogController(_ *gin.RouterGroup, auth *gin.RouterGroup) {
	apiController := system.NewLoginLogApi(login_log.New())
	protectedRouter := auth.Group("login_log")
	{
		protectedRouter.DELETE("delete", grouter.HandlerFunc(apiController.Delete))
		protectedRouter.GET("list", grouter.HandlerFunc(apiController.List))
		protectedRouter.DELETE("deleteAll", grouter.HandlerFunc(apiController.DeleteAll))
	}
}

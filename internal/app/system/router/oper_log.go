package router

import (
	"github.com/562589540/jono-gin/api/v1/system"
	"github.com/562589540/jono-gin/ghub/glibrary/grouter"
	"github.com/562589540/jono-gin/internal/app/system/logic/oper_log"
	"github.com/gin-gonic/gin"
)

func (r *Router) BindOperLogController(_ *gin.RouterGroup, auth *gin.RouterGroup) {
	apiController := system.NewOperLogApi(oper_log.New())
	protectedRouter := auth.Group("oper_log")
	{
		protectedRouter.DELETE("delete", grouter.HandlerFunc(apiController.Delete))
		protectedRouter.DELETE("deleteAll", grouter.HandlerFunc(apiController.DeleteAll))
		protectedRouter.GET("list", grouter.HandlerFunc(apiController.List))
	}
}

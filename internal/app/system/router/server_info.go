package router

import (
	"github.com/562589540/jono-gin/api/v1/system"
	"github.com/562589540/jono-gin/ghub/glibrary/grouter"
	"github.com/562589540/jono-gin/internal/app/system/logic/server_info"
	"github.com/gin-gonic/gin"
)

func (r *Router) BindServerInfoController(_ *gin.RouterGroup, auth *gin.RouterGroup) {
	apiController := system.NewServerInfoApi(server_info.New())
	protectedRouter := auth.Group("serverInfo")
	{
		protectedRouter.GET("getServerInfo", grouter.HandlerFunc(apiController.GetServerInfo))
	}
}

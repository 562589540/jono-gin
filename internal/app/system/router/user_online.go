package router

import (
	"github.com/562589540/jono-gin/api/v1/system"
	"github.com/562589540/jono-gin/ghub/glibrary/grouter"
	"github.com/562589540/jono-gin/internal/app/system/logic/user_online"
	"github.com/gin-gonic/gin"
)

func (r *Router) BindUserOnlineController(_ *gin.RouterGroup, auth *gin.RouterGroup) {
	apiController := system.NewUserOnlineApi(user_online.New())
	protectedRouter := auth.Group("online")
	{
		protectedRouter.DELETE("delete", grouter.HandlerFunc(apiController.Delete))
		protectedRouter.GET("list", grouter.HandlerFunc(apiController.List))
	}
}

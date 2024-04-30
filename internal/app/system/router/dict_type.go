package router

import (
	"github.com/562589540/jono-gin/api/v1/system"
	"github.com/562589540/jono-gin/ghub/glibrary/grouter"
	"github.com/562589540/jono-gin/internal/app/system/logic/dict_type"
	"github.com/gin-gonic/gin"
)

func (r *Router) BindDictTypeController(_ *gin.RouterGroup, auth *gin.RouterGroup) {
	apiController := system.NewDictTypeApi(dictType.New())
	protectedRouter := auth.Group("dictType")
	{
		protectedRouter.POST("add", grouter.HandlerFunc(apiController.Create))
		protectedRouter.DELETE("delete", grouter.HandlerFunc(apiController.Delete))
		protectedRouter.PUT("update", grouter.HandlerFunc(apiController.Update))
		protectedRouter.GET("list", grouter.HandlerFunc(apiController.List))
	}
}

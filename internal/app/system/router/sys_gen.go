package router

import (
	"github.com/562589540/jono-gin/api/v1/system"
	"github.com/562589540/jono-gin/ghub/glibrary/grouter"
	"github.com/562589540/jono-gin/internal/app/system/logic/sys_gen"
	"github.com/gin-gonic/gin"
)

func (r *Router) BindMysqlController(_ *gin.RouterGroup, auth *gin.RouterGroup) {
	apiController := system.NewGenApi(sys_gen.New())
	protectedRouter := auth.Group("sys_gen")
	{
		protectedRouter.GET("getCode", grouter.HandlerFunc(apiController.GetCode))
		protectedRouter.GET("list", grouter.HandlerFunc(apiController.List))
		protectedRouter.DELETE("delete", grouter.HandlerFunc(apiController.Delete))
		protectedRouter.GET("tableList", grouter.HandlerFunc(apiController.TableList))
		protectedRouter.GET("tableInfo", grouter.HandlerFunc(apiController.TableInfo))
		protectedRouter.PUT("importDate", grouter.HandlerFunc(apiController.ImportDate))
		protectedRouter.PUT("genCode", grouter.HandlerFunc(apiController.GenCode))
	}
}

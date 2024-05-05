package router

import (
	"github.com/562589540/jono-gin/api/v1/system"
	"github.com/562589540/jono-gin/ghub/glibrary/grouter"
	"github.com/562589540/jono-gin/internal/app/system/logic/attachment"
	"github.com/gin-gonic/gin"
)

func (r *Router) BindAttachmentController(_ *gin.RouterGroup, auth *gin.RouterGroup) {
	apiController := system.NewAttachmentApi(attachment.New())
	protectedRouter := auth.Group("attachment")
	{
		protectedRouter.DELETE("delete", grouter.HandlerFunc(apiController.Delete))
		protectedRouter.GET("list", grouter.HandlerFunc(apiController.List))
	}
}

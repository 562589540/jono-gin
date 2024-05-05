package router

import (
	"github.com/562589540/jono-gin/api/v1/system"
	"github.com/562589540/jono-gin/ghub"
	"github.com/562589540/jono-gin/ghub/glibrary/grouter"
	"github.com/562589540/jono-gin/internal/app/system/logic/uploader"
	"github.com/562589540/jono-gin/internal/app/system/middleware"
	"github.com/gin-gonic/gin"
	"path/filepath"
)

func (r *Router) BindUploaderController(_ *gin.RouterGroup, auth *gin.RouterGroup) {
	path := filepath.Join(ghub.Cfg.Path.ResourcePath, ghub.Cfg.Path.UploadsPath)
	apiController := system.NewUploaderApi(uploader.New(path, ghub.Cfg.System.NodeNumber), path)
	auth.Use(middleware.UploaderAuth())
	protectedRouter := auth.Group("uploader")
	{
		protectedRouter.GET("chunk", grouter.HandlerFunc(apiController.CheckChunkInfo))
		protectedRouter.POST("chunk", grouter.HandlerFunc(apiController.PutChunk))
		protectedRouter.POST("mergeFile", grouter.HandlerFunc(apiController.MergeFile))
	}
}

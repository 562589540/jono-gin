package router

import (
	"github.com/562589540/jono-gin/api/v1/system"
	"github.com/562589540/jono-gin/ghub"
	"github.com/562589540/jono-gin/ghub/glibrary/grouter"
	"github.com/562589540/jono-gin/internal/app/system/logic/job"
	"github.com/gin-gonic/gin"
)

func (r *Router) BindJobController(_ *gin.RouterGroup, auth *gin.RouterGroup) {
	apiController := system.NewJobApi(job.New(ghub.Task))
	protectedRouter := auth.Group("job")
	{
		protectedRouter.POST("add", grouter.HandlerFunc(apiController.Create))
		protectedRouter.DELETE("delete", grouter.HandlerFunc(apiController.Delete))
		protectedRouter.DELETE("deleteJobLog", grouter.HandlerFunc(apiController.DeleteJobLog))
		protectedRouter.DELETE("deleteJobLogAll", grouter.HandlerFunc(apiController.DeleteJobLogAll))
		protectedRouter.PUT("update", grouter.HandlerFunc(apiController.Update))
		protectedRouter.GET("list", grouter.HandlerFunc(apiController.List))
		protectedRouter.GET("getJobLog", grouter.HandlerFunc(apiController.JobLog))
		protectedRouter.POST("once", grouter.HandlerFunc(apiController.Once))
	}
}

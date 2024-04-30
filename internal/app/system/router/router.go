package router

import (
	"github.com/562589540/jono-gin/ghub/gcmd"
	"github.com/562589540/jono-gin/ghub/glibrary/grouter"
	"github.com/562589540/jono-gin/ghub/gutils"
	"github.com/562589540/jono-gin/internal/app/system/middleware"
	"github.com/gin-gonic/gin"
)

func init() {
	R.BindController()
}

var R = new(Router)

type Router struct{}

// BindController 路由自动绑定
func (r *Router) BindController() {
	//注册路由
	gcmd.RegisterRouter(func(public *gin.RouterGroup, auth *gin.RouterGroup) {
		//公开api
		publicRouter := public.Group("system")
		//鉴权token api
		protectedRouter := auth.Group("system")
		//system独立鉴权中间件
		protectedRouter.Use(middleware.AuthRule())
		//记录用户操作中间件
		protectedRouter.Use(middleware.PostRequestLogger())
		//反射自动绑定路由 Bind{xxx}Controller
		if err := grouter.RouterAutoBind(r, publicRouter, protectedRouter); err != nil {
			gutils.ErrorPanic(err)
			return
		}
	})
}

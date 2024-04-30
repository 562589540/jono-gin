package router

import (
	"github.com/562589540/jono-gin/ghub/gcmd"
	"github.com/562589540/jono-gin/ghub/glibrary/grouter"
	"github.com/562589540/jono-gin/ghub/gutils"
	"github.com/gin-gonic/gin"
)

func init() {
	R.BindController()
}

var R = new(Router)

type Router struct{}

func (r *Router) BindController() {
	gcmd.RegisterRouter(func(public *gin.RouterGroup, auth *gin.RouterGroup) {
		if err := grouter.RouterAutoBind(r, public, auth); err != nil {
			gutils.ErrorPanic(err)
			return
		}
	})
}

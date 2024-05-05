package router

import (
	"github.com/562589540/jono-gin/api/v1/system"
	"github.com/562589540/jono-gin/ghub/glibrary/grouter"
	"github.com/562589540/jono-gin/internal/app/system/logic/admin"
	"github.com/562589540/jono-gin/internal/app/system/logic/captcha"
	"github.com/562589540/jono-gin/internal/app/system/logic/login"
	"github.com/562589540/jono-gin/internal/app/system/logic/login_log"
	"github.com/562589540/jono-gin/internal/app/system/logic/token"
	"github.com/562589540/jono-gin/internal/app/system/logic/user_online"
	"github.com/gin-gonic/gin"
)

func (r *Router) BindLoginController(public *gin.RouterGroup, _ *gin.RouterGroup) {
	apiController := system.NewLoginApi(captcha.New(), login.New(), login_log.New(), admin.New(), token.New(), user_online.New())
	publicRouter := public.Group("login")
	{
		publicRouter.GET("captcha", grouter.HandlerFunc(apiController.Captcha))
		//publicRouter.POST("checkCaptcha", grouter.HandlerFunc(apiController.CheckCaptcha))
		publicRouter.POST("login", grouter.HandlerFunc(apiController.Login))
		publicRouter.POST("refreshToken", grouter.HandlerFunc(apiController.RefreshToken))
	}
}

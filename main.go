package main

import (
	"github.com/562589540/jono-gin/ghub/gcmd"
	"github.com/562589540/jono-gin/internal/app/common/middleware"
	_ "github.com/562589540/jono-gin/internal/app/common/router"
	"github.com/562589540/jono-gin/internal/app/common/validate"
	_ "github.com/562589540/jono-gin/internal/app/system/router"
	_ "github.com/562589540/jono-gin/internal/tasks"
	"github.com/gin-gonic/gin"
)

// @title jono-gin
// @version 0.0.1
// @description 基于gin的框架服务
func main() {
	gcmd.New().
		Use(middleware.Cors()).
		Setup("/api/v1/open", "/api/v1", func(r *gin.RouterGroup) {
			//定义非公开路由全局鉴权token中间件
			r.Use(middleware.Auth())
			//注册自定义验证器
			validate.RegisterCustomValidators()
			//注册中间件验证方法
			middleware.RegisterMiddlewareTools()
		}).
		Launch()
}

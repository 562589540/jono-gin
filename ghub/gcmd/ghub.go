package gcmd

import (
	"context"
	"errors"
	"fmt"
	"github.com/562589540/jono-gin/ghub"
	"github.com/gin-gonic/gin"
	"net/http"
	"os/signal"
	"syscall"
	"time"
)

type IRegisterRouter = func(publicRouter *gin.RouterGroup, protectedRouter *gin.RouterGroup)

var (
	//储存路由注册方法
	gfnRouters []IRegisterRouter
)

// RegisterRouter 注册路由
func RegisterRouter(fn IRegisterRouter) {
	if fn == nil {
		return
	}
	gfnRouters = append(gfnRouters, fn)
}

type GinHub struct {
	server          *http.Server
	engine          *gin.Engine
	publicRouter    *gin.RouterGroup
	protectedRouter *gin.RouterGroup
	Ctx             context.Context
}

func New() *GinHub {
	//启动初始化
	return &GinHub{
		engine: gin.Default(),
	}
}

// Use 安装全局中间件
func (g *GinHub) Use(middleware ...gin.HandlerFunc) *GinHub {
	g.engine.Use(middleware...)
	return g
}

// Setup 安装服务
func (g *GinHub) Setup(publicUrl, protectedUrl string, fn func(*gin.RouterGroup)) *GinHub {
	//集成Swagger
	//g.engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	//设置静态目录
	g.engine.Static(ghub.Cfg.Path.Static, ghub.Cfg.Path.ResourcePath)
	//定义公开api
	g.publicRouter = g.engine.Group(publicUrl)
	//定义鉴权api
	g.protectedRouter = g.engine.Group(protectedUrl)

	//运行外部做一些初始化工作并且绑定鉴权中间件
	fn(g.protectedRouter)

	//注册路由
	for _, fnRegisterRouter := range gfnRouters {
		fnRegisterRouter(g.publicRouter, g.protectedRouter)
	}
	return g
}

// Launch 启动服务
func (g *GinHub) Launch() {
	defer func() {
		//服务结束代表服务终止 进去清理工作
		Clean()
	}()

	var cancel context.CancelFunc
	g.Ctx, cancel = signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	g.server = &http.Server{
		Addr:    fmt.Sprintf(":%d", ghub.Cfg.Server.Port),
		Handler: g.engine,
	}

	go func() {
		ghub.Log.Infof("========================gin服务监听端口: %d=====================\n", ghub.Cfg.Server.Port)
		if err := g.server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			ghub.Log.Error("gin服务启动失败:%s", err.Error())
			return
		}
	}()

	<-g.Ctx.Done()

	ctxShutDown, cancelShutDown := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelShutDown()

	if err := g.server.Shutdown(ctxShutDown); err != nil {
		ghub.Log.Error("gin服务关闭失败:%s", err.Error())
		return
	}
	ghub.Log.Info("gin服务关闭成功")
}

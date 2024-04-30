package gcmd

import (
	"fmt"
	"github.com/562589540/jono-gin/ghub"
	"github.com/562589540/jono-gin/ghub/gbootstrap"
	"github.com/562589540/jono-gin/ghub/glibrary/geventbus"
	"github.com/562589540/jono-gin/ghub/glibrary/gjob"
	"github.com/562589540/jono-gin/ghub/glibrary/gscheduler"
	"github.com/562589540/jono-gin/ghub/gutils"
	"github.com/562589540/jono-gin/internal/app/system/dal"
	"github.com/562589540/jono-gin/internal/app/system/logic/casbin"
	"github.com/562589540/jono-gin/temp"
	"os"
)

func init() {
	if err := Setup(); err != nil {
		Clean()
		gutils.ErrorPanic(err)
	}
	InitGen(ghub.Db)
}

func Setup() error {
	//初始化配置
	gbootstrap.InitConfig()

	//检查静态目录是否存在
	if _, err := os.Stat(gbootstrap.Cfg.Path.ResourcePath); os.IsNotExist(err) {
		return fmt.Errorf("resource directory does not exist: %s", gbootstrap.Cfg.Path.ResourcePath)
	}

	//初始化log模块
	ghub.Log = gbootstrap.InitLogger()

	//初始化gorm模块
	if db, err := gbootstrap.InitDb(); err != nil {
		return fmt.Errorf("failed to initialize database: %v", err)
	} else {
		ghub.Db = db
		//初始化dal
		dal.SetDefault(db)
		//初始化casbin
		casbin.New()
	}

	//初始化redis模块
	if redis, err := gbootstrap.InitRedis(); err != nil {
		return fmt.Errorf("failed to initialize Redis: %v", err)
	} else {
		ghub.RedisClient = redis
	}

	//初始化全局事件中心
	ghub.EventBus = geventbus.GetInstance()

	//初始化异步工作池
	ghub.Pool = gjob.GetInstance(10)

	//初始化异步任务
	ghub.Task = gscheduler.GetInstance()

	//开发用的测试模版
	temp.CModel()

	return nil
}

func Clean() {
	fmt.Println("------------- 执行关闭清理工作 -------------")
	//清理异步工作池
	if ghub.Pool != nil {
		ghub.Pool.Close()
	}
	//清理定时任务
	if ghub.Task != nil {
		ghub.Task.StopAll()
	}
}

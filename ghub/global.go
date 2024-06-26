package ghub

import (
	"github.com/562589540/jono-gin/ghub/gbootstrap"
	"github.com/562589540/jono-gin/ghub/glibrary/geventbus"
	"github.com/562589540/jono-gin/ghub/glibrary/gjob"
	"github.com/562589540/jono-gin/ghub/glibrary/gscheduler"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"time"
)

var (
	StartTime   = time.Now()
	Log         *zap.SugaredLogger
	Db          *gorm.DB
	RedisClient *gbootstrap.RedisClient
	EventBus    *geventbus.EventBus
	Pool        *gjob.WorkerPool
	Task        *gscheduler.TaskRunner
	Cfg         *gbootstrap.Configuration
)

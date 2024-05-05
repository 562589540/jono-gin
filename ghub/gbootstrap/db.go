package gbootstrap

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"time"
)

func InitDb() (*gorm.DB, error) {
	logMode := logger.Info

	if !cfg.Mode.Develop {
		logMode = logger.Error
	}

	db, err := gorm.Open(mysql.Open(cfg.DB.DSN), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "sys_", //自定义前缀
			SingularTable: true,   //是否使用复数表名
		},
		Logger: logger.Default.LogMode(logMode),
	})

	if err != nil {
		return nil, err
	}

	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(cfg.DB.MaxIdleCons)
	sqlDB.SetMaxOpenConns(cfg.DB.MaxOpenCons)
	//设置连接可复用的最大时间
	sqlDB.SetConnMaxLifetime(time.Hour)
	sqlDB.SetConnMaxIdleTime(time.Hour)
	return db, nil
}

package model

import (
	"github.com/562589540/jono-gin/ghub/glibrary/gscheduler"
	"time"
)

type SysJob struct {
	ID             uint                   `gorm:"primarykey"`
	CronExpression string                 `gorm:"size:20;not null;comment:cron执行表达式"`
	JobName        string                 `gorm:"size:20;not null;comment:任务名称"`
	InvokeTarget   string                 `gorm:"size:50;not null;comment:任务方法"`
	JobParams      string                 `gorm:"size:255;comment:任务参数"`
	JobGroup       string                 `gorm:"size:50;comment:任务组名"` //默认 系统
	MisfirePolicy  gscheduler.ExecuteType `gorm:"comment:计划执行策略"`       //重复执行 执行一次
	Status         int                    `gorm:"comment:状态"`           //正常 暂停
	CreatedBy      uint                   `gorm:"comment:创建者"`
	UpdatedBy      uint                   `gorm:"comment:更新者"`
	Remark         string                 `gorm:"size:100;comment:备注"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

func (g *SysJob) TableName() string {
	return "sys_job"
}

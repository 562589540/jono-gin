package model

import (
	"time"
)

type TaskLog struct {
	ID        uint   `gorm:"primarykey"`
	JobId     uint   `gorm:"not null;comment:定时任务ID"`
	JobFunc   string `gorm:"size:50;comment:任务方法"`
	ErrorStr  string `gorm:"size:255;comment:错误日记"`
	Status    int    `gorm:"comment:状态"`
	CreatedAt time.Time
}

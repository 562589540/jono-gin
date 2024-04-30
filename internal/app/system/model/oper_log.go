package model

import (
	"time"
)

type OperLog struct {
	ID        uint   `gorm:"primarykey"`
	UserName  string `gorm:"size:64;not null;index;comment:用户名" json:"user_name"`
	Address   string `gorm:"size:64;comment:地址" json:"address"`
	Browser   string `gorm:"size:64;comment:浏览器" json:"browser"`
	Ip        string `gorm:"size:64;comment:ip" json:"ip"`
	Module    string `gorm:"size:64;comment:模块" json:"module"`
	Summary   string `gorm:"comment:操作内容" json:"summary"`
	System    string `gorm:"comment:系统" json:"system"`
	Status    int    `gorm:"comment:状态" json:"status"`
	CreatedAt time.Time
}

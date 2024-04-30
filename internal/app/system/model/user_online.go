package model

import (
	"time"
)

type UserOnline struct {
	ID        uint      `gorm:"primarykey" json:"id"`
	Uid       uint      `gorm:"not null;index;comment:用户id" json:"uid"`
	UserName  string    `gorm:"size:64;not null;index;comment:用户名" json:"user_name"`
	Address   string    `gorm:"size:50;comment:'登陆地址'" json:"address"`
	Browser   string    `gorm:"size:50;comment:'浏览器类型'" json:"browser"`
	Ip        string    `gorm:"size:50;comment:'登陆ip'" json:"ip"`
	System    string    `gorm:"size:50;comment:'操作系统'" json:"system"`
	LoginTime time.Time `gorm:"comment:'登陆时间'" json:"loginTime"`
}

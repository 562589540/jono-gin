package model

import (
	"github.com/562589540/jono-gin/internal/constants/enum"
	"gorm.io/gorm"
	"time"
)

type UserBase struct {
	gorm.Model
	UserName      string      `gorm:"size:64;not null;uniqueIndex;comment:用户名" json:"user_name"`
	Password      string      `gorm:"size:128;not null;comment:密码" json:"password"`
	NickName      string      `gorm:"size:64;index;comment:昵称" json:"nick_name"`
	Avatar        string      `gorm:"size:255;default:'';comment:头像" json:"avatar"`
	Mobile        string      `gorm:"size:11;default:'';comment:手机号" json:"mobile"`
	Email         string      `gorm:"size:128;default:'';comment:邮箱" json:"email"`
	Status        bool        `gorm:"comment:用户状态;0:禁用,1:正常" json:"status"`
	Sex           enum.Gender `gorm:"comment:性别;0:保密,1:男,2:女" json:"sex"`
	Remark        string      `gorm:"default:'';comment:备注" json:"remark"`
	LastLoginIp   string      `gorm:"default:'';comment:最后登陆IP" json:"last_login_ip"`
	LastLoginTime *time.Time  `gorm:"comment:最后登陆时间" json:"last_login_time"`
	Birthday      int         `gorm:"default:0;comment:生日" json:"birthday"`
	Address       string      `gorm:"default:'';comment:地址" json:"address"`
}

//func (u *UserBase) Encrypt() error {
//	hash, err := gutils.Encrypt(u.Password)
//	if err == nil {
//		u.Password = hash
//	}
//	return err
//}
//
//func (u *UserBase) BeforeCreate(tx *gorm.DB) error {
//	return u.Encrypt()
//}

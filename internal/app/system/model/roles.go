package model

import (
	"github.com/562589540/jono-gin/internal/constants"
	"time"
)

type Roles struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Name      string `gorm:"size:50;not null;unique;comment:'角色名称'" json:"name"`
	Code      string `gorm:"size:50;not null;unique;comment:'角色标识'" json:"code"` //没啥用
	Order     int    `gorm:"comment:'排序'" json:"order"`
	Status    bool   `gorm:"comment:'状态'" json:"status"`
	Remark    string `gorm:"size:255;comment:'备注'" json:"remark"`
	Menus     []Menu `gorm:"many2many:role_menus;constraint:OnDelete:CASCADE;'" json:"-"` //与菜单的中间表
}

func (m *Roles) GetStatus() *int {
	if m.Status {
		return &constants.StatusTrue
	}
	return &constants.StatusFalse
}

func (m *Roles) AllowScan() {}

package model

import (
	"github.com/562589540/jono-gin/internal/constants"
	"time"
)

type Dept struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Name      string  `gorm:"size:64;not null;comment:部门" json:"name"`
	Principal string  `gorm:"size:128;default:'';comment:部门负责人" json:"principal"`
	Email     string  `gorm:"size:128;default:'';comment:邮箱" json:"email"`
	Mobile    string  `gorm:"size:11;default:'';comment:手机号" json:"mobile;"`
	ParentID  *uint   `gorm:"index;comment:上级部门" json:"parentId"`
	Sort      int     `gorm:"default:0;comment:排序" json:"sort"`
	Status    bool    `gorm:"comment:状态" json:"status"`
	Remark    string  `gorm:"default:'';comment:备注" json:"remark;"`
	Children  []Dept  `gorm:"foreignKey:ParentID"`
	Admins    []Admin `gorm:"foreignKey:DeptID;references:ID"`
}

func (m *Dept) GetParentID() uint {
	if m.ParentID == nil {
		return 0
	}
	return *m.ParentID
}

func (m *Dept) GetStatus() *int {
	if m.Status {
		return &constants.StatusTrue
	}
	return &constants.StatusFalse
}

package model

import (
	"github.com/562589540/jono-gin/internal/app/common/model"
	"github.com/562589540/jono-gin/internal/constants"
	"github.com/562589540/jono-gin/internal/constants/enum"
)

type Admin struct {
	model.UserBase
	RoleSign []Roles `gorm:"many2many:admin_roles;comment:关联角色" json:"role_sign"`
	DeptID   *uint   `gorm:"index;comment:所属部门ID" json:"deptId"` // 外键字段
	Dept     Dept    `gorm:"foreignKey:DeptID"`                  // 关联部门
}

func (m *Admin) AllowScan() {}

func (m *Admin) GetSex() int {
	if m.Sex == enum.Male {
		return 0
	}
	if m.Sex == enum.Female {
		return 1
	}
	return 2
}

func (m *Admin) GetStatus() *int {
	if m.Status {
		return &constants.StatusTrue
	}
	return &constants.StatusFalse
}

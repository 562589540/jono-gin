package model

import (
	"fmt"
	"github.com/562589540/jono-gin/internal/constants/enum"
	"github.com/google/uuid"
	"time"
)

type Menu struct {
	ID              uint `gorm:"primarykey"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
	MenuType        enum.MenuGenre `gorm:"not null;comment:菜单类型" json:"menuType"`
	Title           string         `gorm:"size:255;not null;comment:菜单名称" json:"title"`
	Name            string         `gorm:"size:255;not null;unique;comment:路由名称" json:"name"`
	Path            string         `gorm:"size:255;not null;comment:路由路径" json:"path"`
	Api             string         `gorm:"size:255;default:'';comment:接口规则" json:"api"`
	ParentID        *uint          `gorm:"index;comment:上级菜单" json:"parentId"`
	Children        []Menu         `gorm:"foreignKey:ParentID"  json:"-"`
	Rank            int            `gorm:"default:0;comment:菜单排序" json:"rank"`
	Redirect        string         `gorm:"size:255;default:'';comment:路由重定向" json:"redirect"`
	Icon            string         `gorm:"size:255;default:'';comment:菜单图标" json:"icon"`
	EnterTransition string         `gorm:"size:255;default:'';comment:进场动画" json:"enterTransition"`
	LeaveTransition string         `gorm:"size:255;default:'';comment:离场动画" json:"leaveTransition"`
	ActivePath      string         `gorm:"size:255;default:'';comment:菜单激活" json:"activePath"`
	ExtraIcon       string         `gorm:"size:255;default:'';comment:右侧图标" json:"extraIcon"`
	FrameSrc        string         `gorm:"size:255;default:'';comment:链接地址" json:"frameSrc"`
	FrameLoading    bool           `gorm:"default:false;comment:加载动画" json:"frameLoading"`
	KeepAlive       bool           `gorm:"default:false;comment:缓存页面" json:"keepAlive"`
	HiddenTag       bool           `gorm:"default:false;comment:标签页" json:"hiddenTag"`
	FixedTag        bool           `gorm:"default:false;comment:固定标签页" json:"fixedTag"`
	ShowLink        bool           `gorm:"default:false;comment:菜单" json:"showLink"`
	ShowParent      bool           `gorm:"default:false;comment:父级菜单" json:"showParent"`
	Auths           string         `gorm:"size:255;default:'';comment:按钮权限标识" json:"auths"`
	Roles           []Roles        `gorm:"many2many:role_menus;constraint:OnDelete:CASCADE;" json:"-"` //获取权限
	Component       string         `gorm:"size:255;default:'';comment:组件路径" json:"component"`
}

func (m *Menu) GetParentID() uint {
	if m.ParentID == nil {
		return 0
	}
	return *m.ParentID
}

func (m *Menu) GetMenuType() *int {
	t := int(m.MenuType)
	return &t
}
func (m *Menu) CheckParentID() {
	if m.ParentID != nil && *m.ParentID == 0 {
		m.ParentID = nil
	}
}

func (m *Menu) CheckSaveParam() error {
	if m.ParentID == nil {
		if m.MenuType != enum.Menu {
			return fmt.Errorf("root目录只能是菜单类型")
		}
	}
	if m.MenuType == enum.Button {
		m.Name = uuid.New().String()
		m.Path = ""
	} else {
		if m.Path == "" || m.Name == "" {
			return fmt.Errorf("非按钮类型 必须有path与name字段")
		}
	}
	return nil
}

func (m *Menu) AllowScan() {}

package dto

import (
	"github.com/562589540/jono-gin/internal/constants"
)

// AddMenuReq 菜单新增
type AddMenuReq struct {
	Menu
}

// UpdateMenuReq 菜单更新
type UpdateMenuReq struct {
	ID uint `json:"id" form:"id" uri:"id"  binding:"required" m:"查询ID不能为空"`
	Menu
}

// Menu 完整菜单
type Menu struct {
	MenuBase
	MenuMeta
	Rank int `form:"rank" json:"rank"` // 菜单排序
}

// MenuBase 基础菜单
type MenuBase struct {
	ID        uint                `form:"id" json:"id,omitempty"`
	ParentID  *uint               `form:"parentId" json:"parentId"`
	MenuType  constants.MenuGenre `form:"menuType" json:"menuType" binding:"min=0,max=3" m:"菜单类型错误"`
	Path      string              `form:"path" json:"path,omitempty"`           // 路由地址
	Component string              `form:"component" json:"component,omitempty"` // 按需加载需要展示的页面
	Name      string              `form:"name" json:"name,omitempty"`           // 路由名字（必须保持唯一）
	Redirect  string              `form:"redirect" json:"redirect,omitempty"`   //路由重定向地址
}

// MenuMeta 路由元信息
type MenuMeta struct {
	Title           string   `form:"title" json:"title,omitempty"`                                      // 菜单名称
	Icon            string   `form:"icon" json:"icon,omitempty"`                                        // 菜单图标
	Api             string   `form:"api" json:"api,omitempty" binding:"omitempty,ruleapi" m:"接口规则格式错误"` // 后端api
	ShowLink        bool     `form:"showLink" json:"showLink"`                                          // 是否在菜单中显示
	ShowParent      bool     `form:"showParent" json:"showParent"`                                      // 是否显示父级菜单 不太需要 不排序就需要
	KeepAlive       bool     `form:"keepAlive" json:"keepAlive"`                                        // 是否缓存该路由页面（开启后，会保存该页面的整体状态，刷新后会清空状态）
	FrameSrc        string   `form:"frameSrc" json:"frameSrc,omitempty"`                                // 需要内嵌的iframe链接地址
	FrameLoading    bool     `form:"frameLoading" json:"frameLoading"`                                  // 内嵌的iframe页面是否开启首次加载动画
	HiddenTag       bool     `form:"hiddenTag" json:"hiddenTag"`                                        // 当前菜单名称或自定义信息禁止添加到标签页
	EnterTransition string   `form:"enterTransition" json:"enterTransition,omitempty"`                  // 当前页面进场动画
	LeaveTransition string   `form:"leaveTransition" json:"leaveTransition,omitempty"`                  // 当前页面离场动画
	Auths           string   `form:"auths" json:"auths,omitempty"`                                      // 按钮级别权限设置
	FixedTag        bool     `form:"fixedTag" json:"fixedTag,omitempty"`
	ActivePath      string   `form:"activePath" json:"activePath"`
	ExtraIcon       string   `form:"extraIcon"  json:"extraIcon"`
	Roles           []string `gorm:"-" json:"roles,omitempty"`
	Rank            int      `form:"rank" json:"rank,omitempty"` // 菜单排序
	//DynamicLevel    int      `form:"dynamicLevel" json:"dynamicLevel,omitempty"`                        // 显示在标签页的最大数量，需满足后面的条件：不显示在菜单中的路由并且是通过query或params传参模式打开的页面
}

type RoleMenu struct {
	ID       uint   `form:"id" json:"id,omitempty"`
	ParentID uint   `form:"parentId" json:"parentId"`
	MenuType int    `form:"menuType" json:"menuType"`
	Title    string `form:"title" json:"title,omitempty"`
}

// MenuList 菜单列表
type MenuList struct {
	ID        uint                `json:"-"`
	MenuType  constants.MenuGenre `json:"-"`
	Path      string              `json:"path,omitempty"`      // 路由地址
	Component string              `json:"component,omitempty"` // 按需加载需要展示的页面 后端不需要
	Name      string              `json:"name,omitempty"`      // 路由名字（必须保持唯一）
	Redirect  string              `json:"redirect,omitempty"`  //路由重定向地址
	Meta      *MenuMeta           `json:"meta,omitempty"`      // 路由元信息
	Children  []*MenuList         `json:"children,omitempty"`  // 子路由
	Rank      int                 `json:"-"`                   // 菜单排序不传参内部使用字段
}

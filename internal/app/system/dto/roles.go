package dto

import (
	"github.com/562589540/jono-gin/ghub/glibrary/gdto"
	"github.com/562589540/jono-gin/internal/app/system/model"
)

type RolesAddReq struct {
	RolesBase
}

type RolesUpdateReq struct {
	RolesBase
}

type RolesBase struct {
	ID     uint   `form:"id" json:"id,omitempty"`
	Code   string `form:"code" json:"code,omitempty" binding:"required" m:"角色名称为必填项"`
	Name   string `form:"name" json:"name,omitempty" binding:"required" m:"角色标识为必填项"`
	Remark string `form:"remark" json:"remark,omitempty"`
	Status string `form:"status" json:"status,omitempty"`
}

func (m *RolesBase) ToModel(mModel *model.Roles, Generate bool) *model.Roles {
	mModel.Code = m.Code
	mModel.Name = m.Name
	mModel.Remark = m.Remark
	if !Generate && m.ID != 0 {
		mModel.ID = m.ID
	}
	if !Generate && m.Status != "" {
		mModel.Status = m.Status == "1"
	}
	if Generate {
		mModel.Status = true
	}
	return mModel
}

type RolesPowerReq struct {
	ID  uint   `json:"id" binding:"required" m:"角色ID不能为空"`
	IDS []uint `json:"ids" binding:"required" m:"菜单列表不能为空"`
}

type RolesSearchReq struct {
	gdto.PaginateReq        //分页
	Code             string `form:"code" json:"code,omitempty"`
	Name             string `form:"name" json:"name,omitempty"`
	Status           string `form:"status" json:"status,omitempty"`
}

type Role struct {
	ID         uint   `json:"id,omitempty"`
	Code       string `json:"code,omitempty"`
	Name       string `json:"name,omitempty"`
	Status     *int   `json:"status,omitempty"`
	CreateTime int64  `json:"createTime,omitempty"`
	UpdateTime int64  `json:"updateTime,omitempty"`
	Remark     string `json:"remark,omitempty"`
}

func (m Role) FromModel(mModel *model.Roles) Role {
	m.ID = mModel.ID
	m.Code = mModel.Code
	m.Name = mModel.Name
	m.Status = mModel.GetStatus()
	m.CreateTime = mModel.CreatedAt.UnixNano() / 1e6
	m.UpdateTime = mModel.UpdatedAt.UnixNano() / 1e6
	m.Remark = mModel.Remark
	return m
}

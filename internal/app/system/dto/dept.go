package dto

import (
	"github.com/562589540/jono-gin/ghub/glibrary/gdto"
	"github.com/562589540/jono-gin/internal/app/system/model"
)

type DeptUpdateReq struct {
	Dept
	ID uint `form:"id" json:"id,omitempty"`
}

func (m *DeptUpdateReq) ToModel(mModel *model.Dept) *model.Dept {
	mModel.ID = m.ID
	return m.Dept.ToModel(mModel)
}

type DeptAddReq struct {
	Dept
}

type DeptSearchReq struct {
	gdto.PaginateReq //分页
}

type Dept struct {
	ID         uint   `form:"id" json:"id"`
	ParentID   uint   `form:"parentId" json:"parentId"`
	Name       string `form:"name" json:"name" binding:"required" m:"部门名称不能为空"`
	Principal  string `form:"principal" json:"principal"`
	Email      string `form:"email" json:"email" binding:"omitempty,email" m:"邮箱格式错误"`
	Mobile     string `form:"mobile" json:"mobile" binding:"omitempty,len=11" m:"手机号格式错误"`
	Sort       int    `form:"sort" json:"sort"`
	Status     *int   `form:"status" json:"status"`
	Remark     string `form:"remark" json:"remark"`
	CreateTime int64  `form:"createTime" json:"createTime"`
}

func (m Dept) ToModel(mModel *model.Dept) *model.Dept {
	var defaultStatus int
	if m.Status != nil {
		defaultStatus = *m.Status
	}
	mModel.Name = m.Name
	mModel.Principal = m.Principal
	mModel.Email = m.Email
	mModel.Mobile = m.Mobile
	mModel.Sort = m.Sort
	mModel.Status = defaultStatus == 1
	mModel.Remark = m.Remark
	if m.ParentID == 0 {
		mModel.ParentID = nil
	} else {
		mModel.ParentID = &m.ParentID
	}

	return mModel
}

func (m Dept) FromModel(mModel *model.Dept) Dept {
	m.ID = mModel.ID
	m.ParentID = mModel.GetParentID()
	m.Name = mModel.Name
	m.Principal = mModel.Principal
	m.Email = mModel.Email
	m.Mobile = mModel.Mobile
	m.Sort = mModel.Sort
	m.Status = mModel.GetStatus()
	m.Remark = mModel.Remark
	m.CreateTime = mModel.CreatedAt.UnixNano() / 1e6
	return m
}

type AdminDept struct {
	ID   uint   `form:"id" json:"id"`
	Name string `form:"name" json:"name"`
}

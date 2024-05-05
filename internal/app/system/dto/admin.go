package dto

import (
	"github.com/562589540/jono-gin/ghub/glibrary/gdto"
	"github.com/562589540/jono-gin/ghub/gutils"
	"github.com/562589540/jono-gin/internal/app/system/model"
	"github.com/562589540/jono-gin/internal/constants/enum"
)

type AdminAddReq struct {
	AdminUpdateReq
	AdminUpdatePassReq
}

func (m *AdminAddReq) ToModel(mModel *model.Admin) (*model.Admin, error) {
	password, err := m.SetPassword()
	if err != nil {
		return nil, err
	}
	mModel = m.AdminUpdateReq.ToModel(mModel)
	mModel.Password = password
	return mModel, nil
}

type AdminSearchReq struct {
	gdto.PaginateReq        //分页
	DeptId           uint   `json:"deptId" form:"deptId"`
	Mobile           string `json:"mobile" form:"mobile"`
	Status           string `json:"status" form:"status"`
	UserName         string `json:"username" form:"username"`
}

type AdminUpdateReq struct {
	AdminUpdateStatusReq
	AdminUpdateAvatar
	ID         uint      `form:"id" json:"id,omitempty"`
	UserName   string    `json:"username,omitempty" binding:"required" m:"账号不能为空"`
	NickName   string    `json:"nickname,omitempty" binding:"required" m:"昵称不能为空"`
	Mobile     string    `json:"mobile,omitempty" binding:"omitempty,len=11" m:"手机号格式错误"`
	DeptId     uint      `json:"parentId,omitempty"` //部门???
	Remark     string    `json:"remark,omitempty"`
	Sex        int       `json:"sex,omitempty"`
	CreateTime int64     `json:"createTime,omitempty"`
	Dept       AdminDept `json:"dept,omitempty"`
	Email      string    `json:"email,omitempty"`
	IP         string    `json:"-"`
}

func (m AdminUpdateReq) ToModel(mModel *model.Admin) *model.Admin {
	mModel.UserName = m.UserName
	mModel.NickName = m.NickName
	mModel.Remark = m.Remark
	mModel.Mobile = m.Mobile
	mModel.Email = m.Email
	mModel.Sex = m.GetSex()
	mModel.DeptID = m.GetDeptId()
	mModel.Status = m.GetStatus()
	return mModel
}

func (m AdminUpdateReq) GetSex() enum.Gender {
	if m.Sex == 0 {
		return enum.Male
	}
	if m.Sex == 1 {
		return enum.Female
	}
	return enum.Other
}

func (m AdminUpdateReq) GetDeptId() *uint {
	if m.DeptId == 0 {
		return nil
	}
	return &m.DeptId
}

type AdminUpdateRoleReq struct {
	ID      uint   `form:"id" json:"id,omitempty"`
	RoleIds []uint `json:"roleIds,omitempty"`
}

type AdminUpdateAvatar struct {
	ID     uint   `form:"id" json:"id,omitempty"`
	Avatar string `json:"avatar,omitempty"`
}

type AdminUpdatePassReq struct {
	ID       uint   `form:"id" json:"id,omitempty"`
	Password string `json:"password" binding:"password" m:"密码格式应为8-18位数字、字母、符号的任意两种组合"`
}

func (m AdminUpdatePassReq) SetPassword() (string, error) {
	hash, err := gutils.Encrypt(m.Password)
	if err != nil {
		return m.Password, err
	}
	return hash, nil
}

type AdminUpdateStatusReq struct {
	ID     uint `form:"id" json:"id,omitempty"`
	Status *int `json:"status,omitempty"`
}

func (m AdminUpdateStatusReq) GetStatus() bool {
	if m.Status == nil {
		return false
	}
	return *m.Status == 1
}

type Admin struct {
	ID         uint      `json:"id"`
	UserName   string    `json:"username"`
	NickName   string    `json:"nickname"`
	Mobile     string    `json:"mobile"`
	DeptId     uint      `json:"parentId"`
	Remark     string    `json:"remark"`
	Sex        int       `json:"sex"`
	Status     *int      `json:"status"`
	Avatar     string    `json:"avatar"`
	CreateTime int64     `json:"createTime"`
	Dept       AdminDept `json:"dept"`
	Email      string    `json:"email"`
	IP         string    `json:"-"`
}

func (m Admin) FromModel(mModel *model.Admin) Admin {
	m.ID = mModel.ID
	m.Avatar = mModel.Avatar
	m.CreateTime = mModel.CreatedAt.UnixNano() / 1e6
	m.Email = mModel.Email
	m.Dept = AdminDept{
		ID:   mModel.Dept.ID,
		Name: mModel.Dept.Name,
	}
	m.Mobile = mModel.Mobile
	m.NickName = mModel.NickName
	m.Remark = mModel.Remark
	m.Sex = mModel.GetSex()
	m.Status = mModel.GetStatus()
	m.UserName = mModel.UserName
	return m
}

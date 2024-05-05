package dto

import (
	"github.com/562589540/jono-gin/ghub/glibrary/gdto"
	"time"
)

type DictTypeSearchReq struct {
	gdto.PaginateReq        //分页
	DictName         string `json:"dictName" form:"dictName"` // 字典名称
	DictType         string `json:"dictType" form:"dictType"` // 字典类型
	Status           string `json:"status" form:"status"`     // 状态（0正常 1停用）
}

type DictTypeUpdateReq struct {
	ID uint `json:"id" binding:"required"` // 字典主键
	DictTypeAddReq
}

type DictTypeAddReq struct {
	DictName string `json:"dictName" binding:"required" m:"字典名称不能为空"` // 字典名称
	DictType string `json:"dictType" binding:"required" m:"字典类型不能为空"` // 字典类型
	Status   int32  `json:"status"`                                   // 状态（0正常 1停用）
	Remark   string `json:"remark"`
}

type DictType struct {
	DictID    uint      `json:"dictId"`   // 字典主键
	DictName  string    `json:"dictName"` // 字典名称
	DictType  string    `json:"dictType"` // 字典类型
	Status    int32     `json:"status"`   // 状态（0正常 1停用）
	CreateBy  int32     `json:"createBy"` // 创建者
	UpdateBy  int32     `json:"updateBy"` // 更新者
	Remark    string    `json:"remark"`
	CreatedAt time.Time `json:"createdAt"` // 创建日期
	UpdatedAt time.Time `json:"updatedAt"` // 修改日期
}

type DictBatchGetReq struct {
	DictTypes []string `json:"dictTypes"`
}

type DictGetReq struct {
	DictType     string `json:"dictType" form:"dictType"`         // 字典类型
	DefaultValue string `json:"defaultValue" form:"defaultValue"` // 默认值
}

type DictGetRes struct {
	Info   *DictDataInfo    `json:"info"`
	Values []*DictDataValue `json:"values"`
}

type DictDataInfo struct {
	Name   string `json:"name"`
	Remark string `json:"remark"`
}

type DictDataValue struct {
	IsDefault int    `json:"isDefault"`
	Key       string `json:"key"`
	Remark    string `json:"remark"`
	Value     string `json:"value"`
}

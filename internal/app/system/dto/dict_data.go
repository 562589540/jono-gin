package dto

import (
	"github.com/562589540/jono-gin/ghub/glibrary/gdto"
)

type DictDataSearchReq struct {
	gdto.PaginateReq      //分页
	DictCode         uint `json:"dictCode" form:"dictCode"` // 字典编码,
	//DictSort         int    `json:"dictSort" form:"dictSort"`   // 字典排序,
	DictLabel string `json:"dictLabel" form:"dictLabel"` // 字典标签,
	//DictValue        string `json:"dictValue" form:"dictValue"` // 字典键值,
	DictType string `json:"dictType" form:"dictType"` // 字典类型,
	//CssClass         string `json:"cssClass" form:"cssClass"`   // 样式属性（其他样式扩展）,
	//ListClass        string `json:"listClass" form:"listClass"` // 表格回显样式,
	//IsDefault        int    `json:"isDefault" form:"isDefault"` // 是否默认（1是 0否）,
	Status string `json:"status" form:"status"` // 状态（0正常 1停用）,
	//CreateBy         uint   `json:"createBy" form:"createBy"`   // 创建者,
	//UpdateBy         uint   `json:"updateBy" form:"updateBy"`   // 更新者,
	//Remark           string `json:"remark" form:"remark"`       // 备注
}

type DictDataUpdateReq struct {
	ID uint `json:"id" binding:"required"`
	DictDataAddReq
}

type DictDataAddReq struct {
	DictSort  int    `json:"dictSort"`                                  // 字典排序,
	DictLabel string `json:"dictLabel" binding:"required" m:"字典标签不能为空"` // 字典标签,
	DictValue string `json:"dictValue" binding:"required" m:"字典键值不能为空"` // 字典键值,
	DictType  string `json:"dictType"  binding:"required" m:"字典类型不能为空"` // 字典类型,
	CssClass  string `json:"cssClass"`                                  // 样式属性（其他样式扩展）,
	ListClass string `json:"listClass"`                                 // 表格回显样式,
	IsDefault int    `json:"isDefault"`                                 // 是否默认（1是 0否）,
	Status    int    `json:"status"`                                    // 状态（0正常 1停用）,
	//CreateBy  uint   `json:"createBy"`                                  // 创建者,
	//UpdateBy  uint   `json:"updateBy"`                                  // 更新者,
	Remark string `json:"remark"` // 备注
}

type DictData struct {
	DictCode  uint   `json:"dictCode"`  // 字典编码,
	DictSort  int    `json:"dictSort"`  // 字典排序,
	DictLabel string `json:"dictLabel"` // 字典标签,
	DictValue string `json:"dictValue"` // 字典键值,
	DictType  string `json:"dictType"`  // 字典类型,
	CssClass  string `json:"cssClass"`  // 样式属性（其他样式扩展）,
	ListClass string `json:"listClass"` // 表格回显样式,
	IsDefault int    `json:"isDefault"` // 是否默认（1是 0否）,
	Status    int    `json:"status"`    // 状态（0正常 1停用）,
	CreateBy  uint   `json:"createBy"`  // 创建者,
	UpdateBy  uint   `json:"updateBy"`  // 更新者,
	Remark    string `json:"remark"`    // 备注
}

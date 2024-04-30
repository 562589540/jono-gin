package model

import (
	"time"
)

// DictData 字典数据表
type DictData struct {
	DictCode  int64     `gorm:"column:dict_code;type:bigint;primaryKey;autoIncrement:true;comment:字典编码" json:"dict_code"` // 字典编码
	DictSort  int32     `gorm:"column:dict_sort;type:int;comment:字典排序" json:"dict_sort"`                                  // 字典排序
	DictLabel string    `gorm:"column:dict_label;type:varchar(100);comment:字典标签" json:"dict_label"`                       // 字典标签
	DictValue string    `gorm:"column:dict_value;type:varchar(100);comment:字典键值" json:"dict_value"`                       // 字典键值
	DictType  string    `gorm:"column:dict_type;type:varchar(100);comment:字典类型" json:"dict_type"`                         // 字典类型
	CSSClass  string    `gorm:"column:css_class;type:varchar(100);comment:样式属性（其他样式扩展）" json:"css_class"`                 // 样式属性（其他样式扩展）
	ListClass string    `gorm:"column:list_class;type:varchar(100);comment:表格回显样式" json:"list_class"`                     // 表格回显样式
	IsDefault bool      `gorm:"column:is_default;type:tinyint(1);comment:是否默认（1是 0否）" json:"is_default"`                  // 是否默认（1是 0否）
	Status    bool      `gorm:"column:status;type:tinyint(1);comment:状态（0正常 1停用）" json:"status"`                          // 状态（0正常 1停用）
	CreateBy  int64     `gorm:"column:create_by;type:bigint unsigned;comment:创建者" json:"create_by"`                       // 创建者
	UpdateBy  int64     `gorm:"column:update_by;type:bigint unsigned;comment:更新者" json:"update_by"`                       // 更新者
	Remark    string    `gorm:"column:remark;type:varchar(500);comment:备注" json:"remark"`                                 // 备注
	CreatedAt time.Time `gorm:"column:created_at;type:datetime;comment:创建时间" json:"created_at"`                           // 创建时间
	UpdatedAt time.Time `gorm:"column:updated_at;type:datetime;comment:修改时间" json:"updated_at"`                           // 修改时间
}

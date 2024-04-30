package model

import (
	"time"
)

// DictType 字典类型表
type DictType struct {
	DictID    uint       `gorm:"column:dict_id;type:bigint unsigned;primaryKey;autoIncrement:true;comment:字典主键" json:"dict_id"`     // 字典主键
	DictName  string     `gorm:"column:dict_name;type:varchar(100);comment:字典名称" json:"dict_name"`                                  // 字典名称
	DictType  string     `gorm:"column:dict_type;type:varchar(100);uniqueIndex:dict_type,priority:1;comment:字典类型" json:"dict_type"` // 字典类型
	Status    int32      `gorm:"column:status;type:tinyint unsigned;comment:状态（0正常 1停用）" json:"status"`                             // 状态（0正常 1停用）
	CreateBy  int32      `gorm:"column:create_by;type:int unsigned;comment:创建者" json:"create_by"`                                   // 创建者
	UpdateBy  int32      `gorm:"column:update_by;type:int unsigned;comment:更新者" json:"update_by"`                                   // 更新者
	Remark    string     `gorm:"column:remark;type:varchar(500);comment:备注" json:"remark"`
	CreatedAt time.Time  `gorm:"column:created_at;type:datetime;comment:创建日期" json:"created_at"` // 创建日期
	UpdatedAt time.Time  `gorm:"column:updated_at;type:datetime;comment:修改日期" json:"updated_at"` // 修改日期
	DictData  []DictData `gorm:"foreignKey:DictType;references:DictType"`
}

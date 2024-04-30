package model

import (
	"encoding/json"
	"github.com/562589540/jono-gin/ghub/glibrary/gtemplate"
	"gorm.io/datatypes"
	"time"
)

type GenDate struct {
	FieldsInfo []*gtemplate.TableFields `json:"fieldsInfo" binding:"required"`
	BaseInfo   *gtemplate.BaseInfo      `json:"baseInfo" binding:"required"`
	GenInfo    *gtemplate.GenInfo       `json:"genInfo" binding:"required"`
}

type TableInfo struct {
	TableName          string `gorm:"column:TABLE_NAME"`
	TableSchema        string `gorm:"column:TABLE_SCHEMA"`
	TableType          string `gorm:"column:TABLE_TYPE"`
	Engine             string `gorm:"column:ENGINE"`
	Version            int    `gorm:"column:VERSION"`
	RowFormat          string `gorm:"column:ROW_FORMAT"`
	TableRows          int    `gorm:"column:TABLE_ROWS"`
	AvgRowLength       int    `gorm:"column:AVG_ROW_LENGTH"`
	DataLength         int    `gorm:"column:DATA_LENGTH"`
	MaxDataLength      int    `gorm:"column:MAX_DATA_LENGTH"`
	IndexLength        int    `gorm:"column:INDEX_LENGTH"`
	DataFree           int    `gorm:"column:DATA_FREE"`
	AutoIncrement      int    `gorm:"column:AUTO_INCREMENT"`
	CreateTime         string `gorm:"column:CREATE_TIME"`
	UpdateTime         string `gorm:"column:UPDATE_TIME"`
	CheckTime          string `gorm:"column:CHECK_TIME"`
	TableCollation     string `gorm:"column:TABLE_COLLATION"`
	Checksum           *int   `gorm:"column:CHECKSUM"`
	CreateTableOptions string `gorm:"column:CREATE_OPTIONS"`
	TableComment       string `gorm:"column:TABLE_COMMENT"`
}

type TableColumn struct {
	ColumnName    string  `gorm:"column:COLUMN_NAME" json:"columnName"`
	ColumnType    string  `gorm:"column:COLUMN_TYPE" json:"columnType"`
	IsNullable    string  `gorm:"column:IS_NULLABLE" json:"isNullable"`
	ColumnKey     string  `gorm:"column:COLUMN_KEY" json:"columnKey"`
	ColumnDefault *string `gorm:"column:COLUMN_DEFAULT" json:"columnDefault"`
	Extra         string  `gorm:"column:EXTRA" json:"extra"`
	ColumnComment string  `gorm:"column:COLUMN_COMMENT" json:"columnComment"`
}

type SysGen struct {
	ID            uint           `gorm:"primarykey"`
	TableNamed    string         `gorm:"size:64;not null;uniqueIndex"`
	TableComment  string         `gorm:"size:50;"`
	FieldsInfo    datatypes.JSON `gorm:"type:json"`
	BaseInfo      datatypes.JSON `gorm:"type:json"`
	GenInfo       datatypes.JSON `gorm:"type:json"`
	GoApiCode     string         `gorm:"type:longtext"`
	GoDtoCode     string         `gorm:"type:longtext"`
	GoLogicCode   string         `gorm:"type:longtext"`
	GoModelCode   string         `gorm:"type:longtext"`
	GoRouterCode  string         `gorm:"type:longtext"`
	GoServiceCode string         `gorm:"type:longtext"`
	VueApiCode    string         `gorm:"type:longtext"`
	VueFormCode   string         `gorm:"type:longtext"`
	VueHookCode   string         `gorm:"type:longtext"`
	VueIndexCode  string         `gorm:"type:longtext"`
	VueRuleCode   string         `gorm:"type:longtext"`
	VueTypesCode  string         `gorm:"type:longtext"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

func (g *SysGen) TableName() string {
	return "sys_gen"
}

// SerializeFieldsInfo 序列化 FieldsInfo 数据为 JSON 格式
func (g *SysGen) SerializeFieldsInfo(data []*gtemplate.TableFields) error {
	b, err := json.Marshal(data)
	if err != nil {
		return err
	}
	g.FieldsInfo = b
	return nil
}

// DeserializeFieldsInfo 反序列化 JSON 数据为 FieldsInfo 结构
func (g *SysGen) DeserializeFieldsInfo() ([]*gtemplate.TableFields, error) {
	var data []*gtemplate.TableFields
	err := json.Unmarshal(g.FieldsInfo, &data)
	return data, err
}

// SerializeBaseInfo 序列化 BaseInfo 数据为 JSON 格式
func (g *SysGen) SerializeBaseInfo(data *gtemplate.BaseInfo) error {
	b, err := json.Marshal(data)
	if err != nil {
		return err
	}
	g.BaseInfo = b
	return nil
}

// DeserializeBaseInfo 反序列化 JSON 数据为 BaseInfo 结构
func (g *SysGen) DeserializeBaseInfo() (*gtemplate.BaseInfo, error) {
	var data *gtemplate.BaseInfo
	err := json.Unmarshal(g.BaseInfo, &data)
	return data, err
}

// SerializeGenInfo 序列化 GenInfo 数据为 JSON 格式
func (g *SysGen) SerializeGenInfo(data *gtemplate.GenInfo) error {
	b, err := json.Marshal(data)
	if err != nil {
		return err
	}
	g.GenInfo = b
	return nil
}

// DeserializeGenInfo 反序列化 JSON 数据为 GenInfo 结构
func (g *SysGen) DeserializeGenInfo() (*gtemplate.GenInfo, error) {
	var data *gtemplate.GenInfo
	err := json.Unmarshal(g.GenInfo, &data)
	return data, err
}

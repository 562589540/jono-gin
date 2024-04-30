package gtemplate

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

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

func GetMysqlInfo() []TableInfo {
	dsn := "root:112233@tcp(localhost:3306)/information_schema?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	var tables []TableInfo
	db.Table("tables").Where("table_schema = ?", "gin").Find(&tables)

	for _, table := range tables {
		fmt.Println(table.TableName, table.TableRows, table.TableComment, table.CreateTime, table.UpdateTime)
	}
	return tables
}

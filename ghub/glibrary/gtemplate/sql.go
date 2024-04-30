package gtemplate

import (
	"gorm.io/gorm"
)

type ColumnInfo struct {
	Field   string `gorm:"column:Field"`
	Type    string `gorm:"column:Type"`
	Null    string `gorm:"column:Null"`
	Key     string `gorm:"column:Key"`
	Default string `gorm:"column:Default"`
	Extra   string `gorm:"column:Extra"`
	Comment string `gorm:"column:Comment"`
}

func GetTableInfo(db *gorm.DB, tableName string) {
	//var columns []ColumnInfo
	//if err := db.Raw("SHOW FULL COLUMNS FROM " + tableName).Scan(&columns).Error; err != nil {
	//	log.Fatal("Failed to query column information:", err)
	//}
	//
	//for _, col := range columns {
	//	fmt.Printf("Field: %s, Type: %s, Null: %s, Key: %s, Default: %s, Extra: %s, Comment: %s\n",
	//		col.Field, col.Type, col.Null, col.Key, col.Default, col.Extra, col.Comment)
	//}
	//fmt.Println("开始执行代码生成......")
	//tp := NewGTemplate("测试代码生成", strings.TrimPrefix(tableName, "sys_"), "demoGen")
	//tp.columns = columns
	//tp.PackColumns()
	//if err := tp.GenerateVueHookTsx(); err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Println("hook.ts 自动创建成功")
	//if err := tp.GenerateVueIndex(); err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Println("index.vue 自动创建成功")
	//if err := tp.GenerateVueForm(); err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Println("form.vue 自动创建成功")
	//if err := tp.GenerateVueTypesTs(); err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Println("types.ts 自动创建成功")
	//if err := tp.GenerateVueApiTs(); err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Println("api.ts 自动创建成功")
	//if err := tp.GenerateVueRuleTs(); err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Println("rule.ts 自动创建成功")
}

package test

import (
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"testing"
)

func TestQ(t *testing.T) {
	dsn := "root:112233@tcp(127.0.0.1:3306)/gin?charset=utf8mb4&parseTime=True&loc=Local"
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "sys_", //自定义前缀
			SingularTable: true,   //是否使用复数表名
		},
	})

	g := gen.NewGenerator(gen.Config{
		FieldWithIndexTag: true,
		FieldWithTypeTag:  true,
		ModelPkgPath:      "demo_model",
		OutPath:           "/Users/zhaojian/框架设计/go/gindemo/internal/app/system/demo_model",
		Mode:              gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface, // generate mode
	})
	g.UseDB(db)
	g.GenerateModel("sys_dict_type")
	g.GenerateModel("sys_dict_data")
	g.Execute()
}

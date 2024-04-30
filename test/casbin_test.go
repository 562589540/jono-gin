package test

import (
	"fmt"
	"github.com/562589540/jono-gin/internal/app/system/model"
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"testing"
)

func TestCa(t *testing.T) {
	dsn := "root:112233@tcp(127.0.0.1:3306)/gin?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "sys_", //自定义前缀
			SingularTable: true,   //是否使用复数表名
		},
	})
	if err != nil {
		panic("failed to connect database")
	}

	// 初始化适配器，确保传入了自定义表的 CasbinRule 引用
	_, err = gormadapter.NewAdapterByDBWithCustomTable(db, &model.CasbinRule{}, "sys_casbin_rule")
	if err != nil {
		panic("failed to initialize the adapter")
	}
}

func TestCasbin(t *testing.T) {

	db, err := gorm.Open(mysql.Open("root:112233@tcp(127.0.0.1:3306)/gin?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "sys_", //自定义前缀
			SingularTable: true,   //是否使用复数表名
		},
	})
	// 提供完整的文件路径
	modelPath := "/Users/zhaojian/框架设计/go/gindemo/resource/casbin/rbac_model.conf"
	a, _ := gormadapter.NewAdapterByDBWithCustomTable(db, &model.CasbinRule{}, "sys_casbin_rule")
	e, _ := casbin.NewEnforcer(modelPath, a)
	_ = e.LoadPolicy()

	//if err != nil {
	//	t.Fatalf("Failed to create enforcer: %v", err)
	//}

	// 确保 e 不是 nil，这是避免 nil pointer dereference 的关键
	if e == nil {
		t.Fatal("Enforcer is nil. Initialization failed.")
	}

	policy, err := e.AddPolicy("zhangsan", "data2", "read")
	if err != nil {
		return
	}

	fmt.Println(policy)
	fmt.Println(err)

	sub := "alice"  // the user that wants to access a resource.
	obj := "data12" // the resource that is going to be accessed.
	act := "read"   // the operation that the user performs on the resource.

	ok, err := e.Enforce(sub, obj, act)

	if err != nil {
		// handle err
		t.Error(err)
	}

	if ok == true {
		// permit alice to read data1
		t.Log("	通过")
	} else {
		// deny the request, show an error
		t.Log("拦截")
	}
}

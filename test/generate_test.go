package test

import (
	"fmt"
	"github.com/562589540/jono-gin/ghub/glibrary/gtemplate"
	"github.com/562589540/jono-gin/internal/app/system/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"testing"
)

func TestName(t *testing.T) {
	dsn := "root:112233@tcp(127.0.0.1:3306)/gin?charset=utf8mb4&parseTime=True&loc=Local"
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "sys_", //自定义前缀
			SingularTable: true,   //是否使用复数表名
		},
	})
	err := db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='后台管理员表'").AutoMigrate(&model.Admin{})
	if err != nil {
		fmt.Println(err)
		return
	}
	//db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='部门表'").AutoMigrate(&model.Dept{})
	//db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='登陆日志表'").AutoMigrate(&model.LoginLog{})
	//db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='菜单表'").AutoMigrate(&model.Menu{})
	//db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='操作日记表'").AutoMigrate(&model.OperLog{})
	//db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='角色表'").AutoMigrate(&model.Roles{})
	//db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='在线用户表'").AutoMigrate(&model.UserOnline{})

	//gtemplate.GetTableInfo(db, "sys_roles")
	gtemplate.GetMysqlInfo()
}

package gcmd

import (
	"github.com/562589540/jono-gin/internal/app/system/model"
	"gorm.io/gen"
	"gorm.io/gorm"
)

// Dynamic SQL
type Querier interface {
	// SELECT * FROM @@table WHERE id=@id
	GetByID(id uint) (gen.T, error)
	// DELETE FROM @@table WHERE id=@id
	DeleteByID(id uint) (gen.RowsAffected, error)
	// DELETE FROM @@table WHERE id IN (@ids)
	DeleteByIDs(ids []uint) (gen.RowsAffected, error)
}

func InitGen(db *gorm.DB) {
	g := gen.NewGenerator(gen.Config{
		//FieldWithIndexTag: true,
		//FieldWithTypeTag:  true,
		OutPath: "./internal/app/system/dal",
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface, // generate mode
	})

	g.UseDB(db)

	g.ApplyBasic(model.Admin{}, model.Dept{}, model.LoginLog{}, model.Menu{}, model.Roles{},
		model.UserOnline{}, model.OperLog{}, model.SysGen{}, model.DictType{}, model.DictData{})

	//给指定模型添加方法
	g.ApplyInterface(func(Querier) {}, model.Admin{}, model.Dept{}, model.LoginLog{}, model.Menu{},
		model.Roles{}, model.UserOnline{}, model.OperLog{}, model.SysGen{}, model.DictType{}, model.DictData{})

	g.Execute()
}

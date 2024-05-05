package gcmd

import (
	"github.com/562589540/jono-gin/internal/app/system/model"
	"gorm.io/gen"
	"gorm.io/gorm"
)

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
		OutPath: "./internal/app/common/dal",                                        //应该全部放在一起
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface, // generate mode
	})

	g.UseDB(db)

	models := make([]interface{}, 0)
	models = append(models, model.Admin{}, model.Dept{}, model.LoginLog{}, model.Menu{}, model.Roles{},
		model.UserOnline{}, model.OperLog{}, model.SysGen{}, model.DictType{}, model.DictData{}, model.SysJob{},
		model.TaskLog{}, model.Attachment{}, model.Chunk{})

	g.ApplyBasic(models...)

	//给指定模型添加方法
	g.ApplyInterface(func(Querier) {}, models...)

	g.Execute()
}

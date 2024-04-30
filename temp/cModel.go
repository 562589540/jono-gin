package temp

import (
	"fmt"
	"github.com/562589540/jono-gin/ghub"
	"github.com/562589540/jono-gin/internal/app/system/model"
)

func CModel() {
	//ghub.Db.Model(&model.MenuBase{}).AddForeignKey("role_id", "roles(id)", "CASCADE", "CASCADE")

	err := ghub.Db.AutoMigrate(&model.DictType{})
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}

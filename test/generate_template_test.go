package test

import (
	"github.com/562589540/jono-gin/resource/template"
	"testing"
)

func TestGenerateTemp(t *testing.T) {
	app := "system"
	utn := "OperLog"
	stm := "operLog"
	template.GenerateFile(app, "model", utn, stm)
	template.GenerateFile(app, "dto", utn, stm)
	template.GenerateFile(app, "service", utn, stm)
	template.GenerateFile(app, "api", utn, stm)
	template.GenerateFile(app, "router", utn, stm)
	template.GenerateFile(app, "logic", utn, stm)
}

//func TestGetCityByIpTemp(t *testing.T) {
//	fmt.Println(gutils.GetCityByIp("106.116.65.246"))
//}

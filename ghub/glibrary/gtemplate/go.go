package gtemplate

import (
	"bytes"
	"github.com/562589540/jono-gin/ghub/glibrary/gstr"
	"path/filepath"
	"text/template"
	"time"
)

type GoTemplate struct {
	NameEnPas   string
	NameEn      string
	AppName     string //应用归类system
	PackageName string //包 service model 小驼峰
	Time        string
}

// GenerateGoStr 创建代码文件从模板
func (m *GTemplate) GenerateGoStr(templateType string) (string, error) {
	// 打开模板文件
	templatePath := filepath.Join(templateBasePath, templateType+".tmpl")
	tmpl, err := template.ParseFiles(templatePath)
	if err != nil {
		return "", err
	}
	var buf bytes.Buffer
	if err = tmpl.Execute(&buf, &GoTemplate{
		NameEnPas:   gstr.ToPascalCase(m.genInfo.BusinessName),
		NameEn:      gstr.ToCamelCase(m.genInfo.BusinessName),
		AppName:     m.genInfo.PackPath,
		PackageName: templateType,
		Time:        time.Now().Format("2006-01-02 15:04:05"),
	}); err != nil {
		return "", err
	}
	return buf.String(), nil
}

func (m *GTemplate) GenerateGoCode(templateType string, code string, cover bool) error {
	var outDir string
	if templateType == "api" {
		outDir = filepath.Join(goRoot+"/api/v1", m.genInfo.PackPath)
	} else if templateType == "logic" {
		outDir = filepath.Join(goRoot+"/internal/app", m.genInfo.PackPath, templateType, gstr.ToSnakeCase(m.genInfo.BusinessName))
	} else {
		outDir = filepath.Join(goRoot+"/internal/app", m.genInfo.PackPath, templateType)
	}
	outDir = filepath.Join(outDir, gstr.ToSnakeCase(m.genInfo.BusinessName)+".go")
	return m.Gen(outDir, code, cover)
}

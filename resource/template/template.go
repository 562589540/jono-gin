package template

import (
	"fmt"
	"github.com/562589540/jono-gin/ghub/glibrary/gstr"
	"os"
	"path/filepath"
	"text/template"
)

const templateBasePath = "/Users/zhaojian/框架设计/go/gindemo/resource/template"
const appPath = "/Users/zhaojian/框架设计/go/gindemo"

// GenerateFile 创建代码文件从模板
func GenerateFile(app, templateType, uppTypeName, smallTypeName string) {
	var outputDir string
	pName := gstr.ToSnakeCase(uppTypeName)

	if templateType == "api" {
		outputDir = filepath.Join(appPath+"/api/v1", app)
	} else if templateType == "logic" {
		outputDir = filepath.Join(appPath+"/internal/app", app, templateType, pName)
	} else {
		outputDir = filepath.Join(appPath+"/internal/app", app, templateType)
	}

	// 确保输出目录存在
	if _, err := os.Stat(outputDir); os.IsNotExist(err) {
		if err := os.MkdirAll(outputDir, 0755); err != nil {
			panic(err) // 在真实环境中应处理错误，而不是直接 panic
		}
	}

	// 拼接文件路径
	filePath := filepath.Join(outputDir, pName+".go")

	// 检查文件是否已存在，避免覆盖
	if _, err := os.Stat(filePath); err == nil {
		return // 如果文件已存在，则退出函数
	}

	// 打开模板文件
	templatePath := filepath.Join(templateBasePath, templateType+".tpl")
	tmpl, err := template.ParseFiles(templatePath)
	if err != nil {
		panic(err) // 在真实环境中应处理错误，而不是直接 panic
	}

	// 创建输出文件
	file, err := os.Create(filePath)
	if err != nil {
		panic(err) // 在真实环境中应处理错误，而不是直接 panic
	}
	defer file.Close()

	// 定义模板数据
	data := struct {
		UppTypeName   string
		SmallTypeName string
		AppName       string
		PackageName   string
	}{
		UppTypeName:   uppTypeName,
		SmallTypeName: smallTypeName,
		AppName:       app,
		PackageName:   templateType,
	}

	fmt.Println("Executing template with data:", data)
	err = tmpl.Execute(file, data)
	if err != nil {
		fmt.Println("Template execution error:", err)
	} else {
		fmt.Println("Template executed successfully")
	}
}

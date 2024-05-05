package internal

import (
	"bytes"
	"github.com/562589540/jono-gin/ghub/glibrary/gfile"
	"github.com/562589540/jono-gin/ghub/glibrary/gstr"
	"github.com/562589540/jono-gin/ghub/glibrary/gtemplate/pkg"
	"os"
	"text/template"
	"time"
)

const (
	goRoot            = "/Users/zhaojian/框架设计/go/gindemo"
	vueRoot           = "/Users/zhaojian/框架设计/vue/vue后台/pureadmin/vue-pure-admin/"
	vueViewPath       = "src/views/"
	vueApiPath        = "src/api/"
	templateBasePath  = "/Users/zhaojian/框架设计/go/gindemo/resource/template/go"
	formTemplatePath  = "/Users/zhaojian/框架设计/go/gindemo/resource/template/vue/form.tmpl"
	indexTemplatePath = "/Users/zhaojian/框架设计/go/gindemo/resource/template/vue/index.tmpl"
	typesTemplatePath = "/Users/zhaojian/框架设计/go/gindemo/resource/template/vue/types.tmpl"
	ruleTemplatePath  = "/Users/zhaojian/框架设计/go/gindemo/resource/template/vue/rule.tmpl"
	apiTemplatePath   = "/Users/zhaojian/框架设计/go/gindemo/resource/template/vue/api.tmpl"
	hookTemplatePath  = "/Users/zhaojian/框架设计/go/gindemo/resource/template/vue/hook.tmpl"
)

type goTemplateData struct {
	NameEnPas   string
	NameEn      string
	NameEnSn    string
	AppName     string //应用归类system
	PackageName string //包 service model 小驼峰
	Time        string
}

type vueTemplateData struct {
	NameZh    string
	NameEn    string
	NameEnPas string
	Wrap      string
	Directory string
	Fields    []*pkg.TableFields
	Time      string
	Author    string
}

type GTemplate struct {
	FieldsInfo []*pkg.TableFields
	BaseInfo   *pkg.BaseInfo
	GenInfo    *pkg.GenInfo
}

func (m *GTemplate) generateGoTmpl(name string, path string, FuncMap template.FuncMap) (*template.Template, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	// 创建模板实例并设置自定义分隔符
	return template.New(name).Funcs(FuncMap).Parse(string(content))
}

func (m *GTemplate) generateVueTmpl(name string, path string, FuncMap template.FuncMap) (*template.Template, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	// 创建模板实例并设置自定义分隔符
	return template.New(name).Delims("[[", "]]").Funcs(FuncMap).Parse(string(content))
}

// go模版生成器
func (m *GTemplate) executeGoStr(tmpl *template.Template, templateType string) (string, error) {
	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, &goTemplateData{
		NameEnSn:    gstr.ToSnakeCase(m.GenInfo.BusinessName),
		NameEnPas:   gstr.ToPascalCase(m.GenInfo.BusinessName),
		NameEn:      gstr.ToCamelCase(m.GenInfo.BusinessName),
		AppName:     m.GenInfo.PackPath,
		PackageName: templateType,
		Time:        time.Now().Format("2006-01-02 15:04:05"),
	}); err != nil {
		return "", err
	}
	return buf.String(), nil
}

// vue模版生成器
func (m *GTemplate) executeStr(tmpl *template.Template) (string, error) {
	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, &vueTemplateData{
		NameZh:    m.GenInfo.FunctionName,                    //中文
		NameEn:    gstr.ToCamelCase(m.GenInfo.BusinessName),  //英文小驼峰
		NameEnPas: gstr.ToPascalCase(m.GenInfo.BusinessName), //英文大驼峰
		Fields:    m.FieldsInfo,                              //字段
		Wrap:      m.GenInfo.PackPath,                        //包位置  //应用归类system
		Directory: m.GenInfo.Directory,                       //所属目录
		Author:    m.BaseInfo.Author,
		Time:      time.Now().Format("2006-01-02 15:04:05"),
	}); err != nil {
		return "", err
	}
	return buf.String(), nil
}

// 写入代码
func (m *GTemplate) gen(outDir string, code string, cover bool) error {
	//查看文件是否存在
	if !gfile.FileIsExisted(outDir) {
		//文件不存在创建并且写入代码
		return gfile.CreateFileAndWrite(outDir, func(file *os.File) error {
			_, err := file.WriteString(code)
			if err != nil {
				return err
			}
			return nil
		})
	}
	//文件存在要求覆盖代码
	if cover {
		err := gfile.WriteToFile(outDir, code)
		if err != nil {
			return err
		}
	}
	return nil
}

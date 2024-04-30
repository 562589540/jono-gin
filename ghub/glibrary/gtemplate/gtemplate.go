package gtemplate

import (
	"bytes"
	"github.com/562589540/jono-gin/ghub/glibrary/gfile"
	"github.com/562589540/jono-gin/ghub/glibrary/gstr"
	_ "github.com/go-sql-driver/mysql"
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

var (
	FieldFilter = map[string]bool{
		"created_at": true,
		"updated_at": true,
		"deleted_at": true,
	}
)

type IGenCode interface {
	GenCodeStr() (string, error)
	GenCode(code string) error
}

type FormData struct {
	NameZh    string
	NameEn    string
	NameEnPas string
	Wrap      string
	Directory string
	Fields    []*TableFields
	Time      string
	Author    string
}

type GTemplate struct {
	fieldsInfo []*TableFields
	baseInfo   *BaseInfo
	genInfo    *GenInfo
}

func NewGTemplate(baseInfo *BaseInfo, genInfo *GenInfo, fieldsInfo []*TableFields) *GTemplate {
	return &GTemplate{
		baseInfo:   baseInfo,
		genInfo:    genInfo,
		fieldsInfo: fieldsInfo,
	}
}

func (m *GTemplate) generateTmpl(name string, path string, FuncMap template.FuncMap) (*template.Template, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	// 创建模板实例并设置自定义分隔符
	return template.New(name).Delims("[[", "]]").Funcs(FuncMap).Parse(string(content))
}

// OpenTpl 打开模板文件
func (m *GTemplate) openTpl(templatePath string) (*template.Template, error) {
	return template.ParseFiles(templatePath)
}

func (m *GTemplate) executeStr(tmpl *template.Template) (string, error) {
	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, &FormData{
		NameZh:    m.genInfo.FunctionName,                    //中文
		NameEn:    gstr.ToCamelCase(m.genInfo.BusinessName),  //英文小驼峰
		NameEnPas: gstr.ToPascalCase(m.genInfo.BusinessName), //英文大驼峰
		Fields:    m.fieldsInfo,                              //字段
		Wrap:      m.genInfo.PackPath,                        //包位置  //应用归类system
		Directory: m.genInfo.Directory,                       //所属目录
		Author:    m.baseInfo.Author,
		Time:      time.Now().Format("2006-01-02 15:04:05"),
	}); err != nil {
		return "", err
	}
	return buf.String(), nil
}

func (m *GTemplate) Gen(outDir string, code string, cover bool) error {
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

// GenCode 写入代码
func (m *GTemplate) GenCode(codes GenCodes) error {
	err := m.GenerateVueApiCode(codes.VueApiCode)
	if err != nil {
		return err
	}
	err = m.GenerateVueHookCode(codes.VueHookCode)
	if err != nil {
		return err
	}
	err = m.GenerateVueFormCode(codes.VueFormCode)
	if err != nil {
		return err
	}
	err = m.GenerateVueIndexCode(codes.VueIndexCode)
	if err != nil {
		return err
	}
	err = m.GenerateVueRuleCode(codes.VueRuleCode)
	if err != nil {
		return err
	}
	err = m.GenerateVueTypesCode(codes.VueTypesCode)
	if err != nil {
		return err
	}
	err = m.GenerateGoCode("api", codes.GoApiCode, true)
	if err != nil {
		return err
	}
	err = m.GenerateGoCode("model", codes.GoModelCode, false)
	if err != nil {
		return err
	}
	err = m.GenerateGoCode("dto", codes.GoDtoCode, false)
	if err != nil {
		return err
	}
	err = m.GenerateGoCode("service", codes.GoServiceCode, false)
	if err != nil {
		return err
	}
	err = m.GenerateGoCode("router", codes.GoRouterCode, false)
	if err != nil {
		return err
	}
	err = m.GenerateGoCode("logic", codes.GoLogicCode, false)
	if err != nil {
		return err
	}
	return nil
}

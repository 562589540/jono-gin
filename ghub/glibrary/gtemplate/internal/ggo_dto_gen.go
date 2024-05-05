package internal

import (
	"fmt"
	"github.com/562589540/jono-gin/ghub/glibrary/gstr"
	"github.com/562589540/jono-gin/ghub/glibrary/gtemplate/pkg"
	"path/filepath"
	"strings"
	"text/template"
)

type GGoDTOGen struct {
	*GTemplate
}

func (m *GGoDTOGen) templateType() string {
	return "dto"
}

func (m *GGoDTOGen) GenCodeStr() (string, error) {
	templatePath := filepath.Join(templateBasePath, m.templateType()+".tmpl")
	tmpl, err := m.generateGoTmpl(m.templateType(), templatePath, template.FuncMap{
		"customSearch": m.customSearch,
		"customAdd":    m.customAdd,
		"customALl":    m.customALl,
	})
	if err != nil {
		return "", err
	}
	return m.executeGoStr(tmpl, m.templateType())
}

func (m *GGoDTOGen) GenCode(codes pkg.GenCodes) error {
	var outDir string
	outDir = filepath.Join(goRoot+"/internal/app", m.GenInfo.PackPath, m.templateType())
	outDir = filepath.Join(outDir, gstr.ToSnakeCase(m.GenInfo.BusinessName)+".go")
	return m.gen(outDir, codes.GoDtoCode, false)
}

func (m *GGoDTOGen) UpdateGeneratedCode(code *pkg.GenCodes, str string) {
	code.GoDtoCode = str
}

// 搜索dto生成
func (m *GGoDTOGen) customSearch() string {
	var sb strings.Builder
	for i, fields := range m.FieldsInfo {
		if fields.Query && fields.Field != "id" {
			sb.WriteString("\t")
			formatted := fmt.Sprintf("%s  %s  `json:\"%s\" form:\"%s\"` // %s",
				fields.GoName, fields.GoType, fields.JsonName, fields.JsonName, fields.FieldDes)
			sb.WriteString(formatted)
			if i != len(m.FieldsInfo)-1 {
				sb.WriteString(`,` + "\n")
			}
		}
	}
	return sb.String()
}

// 新增dto生成
func (m *GGoDTOGen) customAdd() string {
	var sb strings.Builder
	for i, fields := range m.FieldsInfo {
		//允许编辑
		if fields.Edit && fields.Field != "id" {
			sb.WriteString("\t")
			required := ""
			msg := ""
			if fields.Required {
				required = ` binding:"required"`
				msg = ` m:"` + fields.FieldDes + `不能为空"`
			}
			formatted := fmt.Sprintf("%s  %s  `json:\"%s\"%s%s` // %s",
				fields.GoName, fields.GoType, fields.JsonName, required, msg, fields.FieldDes)
			sb.WriteString(formatted)
			if i != len(m.FieldsInfo)-1 {
				sb.WriteString(`,` + "\n")
			}
		}
	}
	return sb.String()
}

// 详情dto生成
func (m *GGoDTOGen) customALl() string {
	var sb strings.Builder
	for i, fields := range m.FieldsInfo {
		if fields.Details {
			sb.WriteString("\t")
			formatted := fmt.Sprintf("%s  %s  `json:\"%s\"` // %s",
				fields.GoName, fields.GoType, fields.JsonName, fields.FieldDes)
			sb.WriteString(formatted)
			if i != len(m.FieldsInfo)-1 {
				sb.WriteString(`,` + "\n")
			}
		}
	}
	return sb.String()
}

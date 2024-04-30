package gtemplate

import (
	"github.com/562589540/jono-gin/ghub/glibrary/gstr"
	"path/filepath"
	"strings"
	"text/template"
)

// GenerateVueTypesTsStr 创建vue types.ts str
func (m *GTemplate) GenerateVueTypesTsStr() (string, error) {
	tmpl, err := m.generateTmpl("types", typesTemplatePath, template.FuncMap{
		"customVar": m.customTsVarGen,
	})
	if err != nil {
		return "", err
	}
	return m.executeStr(tmpl)
}

func (m *GTemplate) customTsVarGen(fields []*TableFields) string {
	var sb strings.Builder
	for _, field := range fields {
		if field.Field == "id" {
			continue
		}
		sb.WriteString("\t" + `/** ` + field.FieldDes + ` */` + "\n")
		sb.WriteString("\t" + field.JsonName + `: ` + field.TsType.String() + `;` + "\n")
	}
	return sb.String()
}
func (m *GTemplate) GenerateVueTypesCode(code string) error {
	//vueRoot vueApiPath/所属菜单/功能名称/utils/
	outDir := filepath.Join(vueRoot, vueViewPath, m.genInfo.Directory,
		gstr.ToCamelCase(m.genInfo.BusinessName), "utils", "types.ts")
	return m.Gen(outDir, code, false)
}

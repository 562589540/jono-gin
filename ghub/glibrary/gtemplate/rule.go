package gtemplate

import (
	"github.com/562589540/jono-gin/ghub/glibrary/gstr"
	"path/filepath"
	"strings"
	"text/template"
)

// GenerateVueRuleTsStr 创建vue rule.ts str
func (m *GTemplate) GenerateVueRuleTsStr() (string, error) {
	tmpl, err := m.generateTmpl("rule", ruleTemplatePath, template.FuncMap{
		"customRule": m.customTsRuleGen,
	})
	if err != nil {
		return "", err
	}
	return m.executeStr(tmpl)
}

func (m *GTemplate) customTsRuleGen(fields []*TableFields) string {
	var sb strings.Builder
	for _, field := range fields {
		if field.Required {
			sb.WriteString("\t" + field.JsonName + `: [{ required: true, message: "` + field.FieldDes + `为必填项", trigger: "blur" }],` + "\n")
		}
	}
	return sb.String()
}

func (m *GTemplate) GenerateVueRuleCode(code string) error {
	//vueRoot vueApiPath/所属菜单/功能名称/utils/
	outDir := filepath.Join(vueRoot, vueViewPath, m.genInfo.Directory,
		gstr.ToCamelCase(m.genInfo.BusinessName), "utils", "rule.ts")
	return m.Gen(outDir, code, false)
}

package internal

import (
	"github.com/562589540/jono-gin/ghub/glibrary/gstr"
	"github.com/562589540/jono-gin/ghub/glibrary/gtemplate/pkg"
	"path/filepath"
	"strings"
	"text/template"
)

type GVueRuleGen struct {
	*GTemplate
}

func (m *GVueRuleGen) GenCodeStr() (string, error) {
	tmpl, err := m.generateVueTmpl("rule", ruleTemplatePath, template.FuncMap{
		"customRule": m.customTsRuleGen,
	})
	if err != nil {
		return "", err
	}
	return m.executeStr(tmpl)
}

func (m *GVueRuleGen) GenCode(codes pkg.GenCodes) error {
	//vueRoot vueApiPath/所属菜单/功能名称/utils/
	outDir := filepath.Join(vueRoot, vueViewPath, m.GenInfo.Directory,
		gstr.ToCamelCase(m.GenInfo.BusinessName), "utils", "rule.ts")
	return m.gen(outDir, codes.VueRuleCode, true)
}

func (m *GVueRuleGen) UpdateGeneratedCode(code *pkg.GenCodes, str string) {
	code.VueRuleCode = str
}

// 生成验证规则
func (m *GVueRuleGen) customTsRuleGen(fields []*pkg.TableFields) string {
	var sb strings.Builder
	for _, field := range fields {
		if field.Required {
			sb.WriteString("\t" + field.JsonName + `: [{ required: true, message: "` + field.FieldDes + `为必填项", trigger: "blur" }],` + "\n")
		}
	}
	return sb.String()
}

package internal

import (
	"github.com/562589540/jono-gin/ghub/glibrary/gstr"
	"github.com/562589540/jono-gin/ghub/glibrary/gtemplate/enum"
	"github.com/562589540/jono-gin/ghub/glibrary/gtemplate/pkg"
	"path/filepath"
	"strings"
	"text/template"
)

type GVueTypesGen struct {
	*GTemplate
}

func (m *GVueTypesGen) GenCodeStr() (string, error) {
	tmpl, err := m.generateVueTmpl("types", typesTemplatePath, template.FuncMap{
		"customVar": m.customTsVarGen,
	})
	if err != nil {
		return "", err
	}
	return m.executeStr(tmpl)
}

func (m *GVueTypesGen) GenCode(codes pkg.GenCodes) error {
	//vueRoot vueApiPath/所属菜单/功能名称/utils/
	outDir := filepath.Join(vueRoot, vueViewPath, m.GenInfo.Directory,
		gstr.ToCamelCase(m.GenInfo.BusinessName), "utils", "types.ts")
	return m.gen(outDir, codes.VueTypesCode, true)
}

func (m *GVueTypesGen) UpdateGeneratedCode(code *pkg.GenCodes, str string) {
	code.VueTypesCode = str
}

func (m *GVueTypesGen) customTsVarGen(fields []*pkg.TableFields) string {
	var sb strings.Builder
	for _, field := range fields {
		if field.Field == "id" {
			continue
		}
		sb.WriteString("\t" + `/** ` + field.FieldDes + ` */` + "\n")
		//数字类型
		if field.TsType == enum.TsNumberType {
			sb.WriteString("\t" + field.JsonName + `: string|` + field.TsType.String() + `;` + "\n")
		} else {
			sb.WriteString("\t" + field.JsonName + `: ` + field.TsType.String() + `;` + "\n")
		}

	}
	return sb.String()
}

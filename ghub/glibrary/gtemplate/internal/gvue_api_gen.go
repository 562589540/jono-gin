package internal

import (
	"github.com/562589540/jono-gin/ghub/glibrary/gstr"
	"github.com/562589540/jono-gin/ghub/glibrary/gtemplate/pkg"
	"path/filepath"
	"text/template"
)

type GVueApiGen struct {
	*GTemplate
}

func (m *GVueApiGen) GenCodeStr() (string, error) {
	tmpl, err := m.generateVueTmpl("api", apiTemplatePath, template.FuncMap{})
	if err != nil {
		return "", err
	}
	return m.executeStr(tmpl)
}

func (m *GVueApiGen) GenCode(codes pkg.GenCodes) error {
	//vueRoot vueApiPath/所属菜单/功能名称
	outDir := filepath.Join(vueRoot, vueApiPath, m.GenInfo.Directory, gstr.ToCamelCase(m.GenInfo.BusinessName)+".ts")
	return m.gen(outDir, codes.VueApiCode, true)
}

func (m *GVueApiGen) UpdateGeneratedCode(code *pkg.GenCodes, str string) {
	code.VueApiCode = str
}

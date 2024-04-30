package gtemplate

import (
	"github.com/562589540/jono-gin/ghub/glibrary/gstr"
	"path/filepath"
	"text/template"
)

// GenerateVueApiTsStr 创建vue api.ts str
func (m *GTemplate) GenerateVueApiTsStr() (string, error) {
	tmpl, err := m.generateTmpl("api", apiTemplatePath, template.FuncMap{})
	if err != nil {
		return "", err
	}
	return m.executeStr(tmpl)
}

func (m *GTemplate) GenerateVueApiCode(code string) error {
	//vueRoot vueApiPath/所属菜单/功能名称
	outDir := filepath.Join(vueRoot, vueApiPath, m.genInfo.Directory, gstr.ToCamelCase(m.genInfo.BusinessName)+".ts")
	return m.Gen(outDir, code, false)
}

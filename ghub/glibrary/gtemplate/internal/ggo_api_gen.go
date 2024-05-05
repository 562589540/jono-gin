package internal

import (
	"github.com/562589540/jono-gin/ghub/glibrary/gstr"
	"github.com/562589540/jono-gin/ghub/glibrary/gtemplate/pkg"
	"path/filepath"
	"text/template"
)

type GGoApiGen struct {
	*GTemplate
}

func (m *GGoApiGen) templateType() string {
	return "api"
}

func (m *GGoApiGen) GenCodeStr() (string, error) {
	templatePath := filepath.Join(templateBasePath, m.templateType()+".tmpl")
	tmpl, err := template.ParseFiles(templatePath)
	if err != nil {
		return "", err
	}
	return m.executeGoStr(tmpl, m.templateType())
}

func (m *GGoApiGen) GenCode(codes pkg.GenCodes) error {
	var outDir string
	outDir = filepath.Join(goRoot+"/api/v1", m.GenInfo.PackPath)
	outDir = filepath.Join(outDir, gstr.ToSnakeCase(m.GenInfo.BusinessName)+".go")
	return m.gen(outDir, codes.GoApiCode, false)
}

func (m *GGoApiGen) UpdateGeneratedCode(code *pkg.GenCodes, str string) {
	code.GoApiCode = str
}

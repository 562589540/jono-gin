package internal

import (
	"github.com/562589540/jono-gin/ghub/glibrary/gstr"
	"github.com/562589540/jono-gin/ghub/glibrary/gtemplate/pkg"
	"path/filepath"
	"text/template"
)

type GGoModelGen struct {
	*GTemplate
}

func (m *GGoModelGen) templateType() string {
	return "model"
}
func (m *GGoModelGen) GenCodeStr() (string, error) {
	templatePath := filepath.Join(templateBasePath, m.templateType()+".tmpl")
	tmpl, err := template.ParseFiles(templatePath)
	if err != nil {
		return "", err
	}
	return m.executeGoStr(tmpl, m.templateType())
}

func (m *GGoModelGen) GenCode(codes pkg.GenCodes) error {
	var outDir string
	outDir = filepath.Join(goRoot+"/internal/app", m.GenInfo.PackPath, m.templateType())
	outDir = filepath.Join(outDir, gstr.ToSnakeCase(m.GenInfo.BusinessName)+".go")
	return m.gen(outDir, codes.GoModelCode, false)
}

func (m *GGoModelGen) UpdateGeneratedCode(code *pkg.GenCodes, str string) {
	code.GoModelCode = str
}

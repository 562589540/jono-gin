package internal

import (
	"github.com/562589540/jono-gin/ghub/glibrary/gstr"
	"github.com/562589540/jono-gin/ghub/glibrary/gtemplate/pkg"
	"path/filepath"
	"text/template"
)

type GGoServiceGen struct {
	*GTemplate
}

func (m *GGoServiceGen) templateType() string {
	return "service"
}
func (m *GGoServiceGen) GenCodeStr() (string, error) {
	templatePath := filepath.Join(templateBasePath, m.templateType()+".tmpl")
	tmpl, err := template.ParseFiles(templatePath)
	if err != nil {
		return "", err
	}
	return m.executeGoStr(tmpl, m.templateType())
}

func (m *GGoServiceGen) GenCode(codes pkg.GenCodes) error {
	var outDir string
	outDir = filepath.Join(goRoot+"/internal/app", m.GenInfo.PackPath, m.templateType())
	outDir = filepath.Join(outDir, gstr.ToSnakeCase(m.GenInfo.BusinessName)+".go")
	return m.gen(outDir, codes.GoServiceCode, false)
}

func (m *GGoServiceGen) UpdateGeneratedCode(code *pkg.GenCodes, str string) {
	code.GoServiceCode = str
}

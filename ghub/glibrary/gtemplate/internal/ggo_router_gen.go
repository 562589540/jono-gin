package internal

import (
	"github.com/562589540/jono-gin/ghub/glibrary/gstr"
	"github.com/562589540/jono-gin/ghub/glibrary/gtemplate/pkg"
	"path/filepath"
	"text/template"
)

type GGoRouterGen struct {
	*GTemplate
}

func (m *GGoRouterGen) templateType() string {
	return "router"
}
func (m *GGoRouterGen) GenCodeStr() (string, error) {
	templatePath := filepath.Join(templateBasePath, m.templateType()+".tmpl")
	tmpl, err := template.ParseFiles(templatePath)
	if err != nil {
		return "", err
	}
	return m.executeGoStr(tmpl, m.templateType())
}

func (m *GGoRouterGen) GenCode(codes pkg.GenCodes) error {
	var outDir string
	outDir = filepath.Join(goRoot+"/internal/app", m.GenInfo.PackPath, m.templateType())
	outDir = filepath.Join(outDir, gstr.ToSnakeCase(m.GenInfo.BusinessName)+".go")
	return m.gen(outDir, codes.GoRouterCode, false)
}

func (m *GGoRouterGen) UpdateGeneratedCode(code *pkg.GenCodes, str string) {
	code.GoRouterCode = str
}

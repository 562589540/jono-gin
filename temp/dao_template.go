package temp

import (
	"os"
	"text/template"
)

type DaoTemplateData struct {
	PackageName string
	TypeName    string
}

const daoTemplate = `package {{.PackageName}}

import (
	"github.com/562589540/jono-gin/internal/app/common/dao"
)

var {{.TypeName}}Dao *{{.TypeName}}Dao

type {{.TypeName}}Dao struct {
	dao.BaseDao
}

func New{{.TypeName}}Dao() *{{.TypeName}}Dao {
	if {{.TypeName}}Dao == nil {
		{{.TypeName}}Dao = &{{.TypeName}}Dao{dao.NewBaseDao()}
	}
	return {{.TypeName}}Dao
}`

func GenerateDaoFile(outputDir string) {
	// Prepare template data
	data := DaoTemplateData{
		PackageName: "dao",
		TypeName:    "RolesBase",
	}

	// Ensure output directory exists
	if _, err := os.Stat(outputDir); os.IsNotExist(err) {
		os.MkdirAll(outputDir, os.ModePerm)
	}

	// Create file
	filePath := outputDir + "/" + data.TypeName + "_dao.go"
	file, err := os.Create(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Create a new template and parse the letter into it
	t := template.Must(template.New("dao").Parse(daoTemplate))

	// Execute the template to file
	err = t.Execute(file, data)
	if err != nil {
		panic(err)
	}
}

// Example usage
// func main() {
//     GenerateDaoFile("./output")
// }

package temp

import (
	"os"
	"text/template"
)

type ServiceTemplateData struct {
	PackageName string
	TypeName    string
	DaoPackage  string
}

const serviceTemplate = `package {{.PackageName}}

import (
	"github.com/562589540/jono-gin/internal/app/common/service"
	"github.com/562589540/jono-gin/internal/app/system/dao"
)

var {{.TypeName}}Service *{{.TypeName}}Service

type {{.TypeName}}Service struct {
	service.BaseService
	Dao *dao.{{.TypeName}}Dao
}

func New{{.TypeName}}Service() *{{.TypeName}}Service {
	if {{.TypeName}}Service == nil {
		{{.TypeName}}Service = &{{.TypeName}}Service{
			Dao: dao.New{{.TypeName}}Dao(),
		}
	}
	return {{.TypeName}}Service
}`

func GenerateServiceFile(outputDir string) {
	// Prepare template data
	data := ServiceTemplateData{
		PackageName: "service",
		TypeName:    "RolesBase",
		DaoPackage:  "dao",
	}

	// Ensure output directory exists
	if _, err := os.Stat(outputDir); os.IsNotExist(err) {
		os.MkdirAll(outputDir, os.ModePerm)
	}

	// Create file
	filePath := outputDir + "/" + data.TypeName + "Service.go"
	file, err := os.Create(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Create a new template and parse the letter into it
	t := template.Must(template.New("service").Parse(serviceTemplate))

	// Execute the template to file
	err = t.Execute(file, data)
	if err != nil {
		panic(err)
	}
}

// Example usage
// func main() {
//     GenerateServiceFile("./output")
// }

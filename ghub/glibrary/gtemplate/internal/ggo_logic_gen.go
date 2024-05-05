package internal

import (
	"fmt"
	"github.com/562589540/jono-gin/ghub/glibrary/gstr"
	"github.com/562589540/jono-gin/ghub/glibrary/gtemplate/enum"
	"github.com/562589540/jono-gin/ghub/glibrary/gtemplate/pkg"
	"path/filepath"
	"strings"
	"text/template"
)

type GGoLogicGen struct {
	*GTemplate
}

func (m *GGoLogicGen) templateType() string {
	return "logic"
}

func (m *GGoLogicGen) GenCodeStr() (string, error) {
	templatePath := filepath.Join(templateBasePath, m.templateType()+".tmpl")
	tmpl, err := m.generateGoTmpl(m.templateType(), templatePath, template.FuncMap{
		"customSearch": m.customSearch,
		"customSort":   m.customSort,
	})
	if err != nil {
		return "", err
	}
	return m.executeGoStr(tmpl, m.templateType())
}

func (m *GGoLogicGen) GenCode(codes pkg.GenCodes) error {
	var outDir string
	outDir = filepath.Join(goRoot+"/internal/app", m.GenInfo.PackPath, m.templateType(), gstr.ToSnakeCase(m.GenInfo.BusinessName))
	outDir = filepath.Join(outDir, gstr.ToSnakeCase(m.GenInfo.BusinessName)+".go")
	return m.gen(outDir, codes.GoLogicCode, false)
}

func (m *GGoLogicGen) UpdateGeneratedCode(code *pkg.GenCodes, str string) {
	code.GoLogicCode = str
}

// 生成查询构造器
func (m *GGoLogicGen) customSearch() string {
	var sb strings.Builder
	for i, fields := range m.FieldsInfo {
		if fields.Query {
			sb.WriteString("\t")
			df := m.changeField(fields.Field) //转换大驼峰匹配dal字段
			gf := fields.GoName               //go变量
			q := ""
			sIf := `search.` + gf + ` != ""`
			//TODO::还有其他可能
			if fields.GoType != enum.GoStringType {
				sIf = `search.` + gf + ` != 0`
			}
			if fields.QueryType == enum.Equal { //等于
				q = `Eq(search.` + gf + `)`
			} else if fields.QueryType == enum.NotEqual { //不等一
				q = `Neq(search.` + gf + `)`
			} else if fields.QueryType == enum.GreaterThan { //大于
				q = `Gt(search.` + gf + `)`
			} else if fields.QueryType == enum.GreaterOrEqual { //大于等于
				q = `Gte(search.` + gf + `)`
			} else if fields.QueryType == enum.LessThan { //小于
				q = `Lt(search.` + gf + `)`
			} else if fields.QueryType == enum.LessOrEqual { //小于等于
				q = `Lte(search.` + gf + `)`
			} else if fields.QueryType == enum.Between { //区间
				sIf = `search.` + gf + ` != nil && len(search.` + gf + `) == 2`
				q = `Between(search.` + gf + `[0], search.` + gf + `[1])`
			} else if fields.QueryType == enum.Like {
				q = `Like(fmt.Sprintf("%%%s%%", search.` + gf + `))`
			}

			sb.WriteString(`if ` + sIf + ` {
		q = q.Where(dao.` + df + `.` + q + `)
	}`)
			if i != len(m.FieldsInfo)-1 {
				sb.WriteString("\n")
			}
		}
	}
	return sb.String()
}

func (m *GGoLogicGen) customSort() string {
	ot := "Asc()"
	if m.BaseInfo.SortWay == "desc" {
		ot = "Desc()"
	}
	return fmt.Sprintf("\tcount, err := q.Order(dao.%s.%s).ScanByPage(&list, search.GetOffset(), search.GetLimit())",
		m.changeField(m.BaseInfo.SortField), ot)
}

func (m *GGoLogicGen) changeField(field string) string {
	if field == "id" {
		return "ID"
	}
	return gstr.SnakeToPascal(field)
}

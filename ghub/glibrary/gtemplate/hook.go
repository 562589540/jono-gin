package gtemplate

import (
	"github.com/562589540/jono-gin/ghub/glibrary/gstr"
	"path/filepath"
	"strings"
	"text/template"
)

// GenerateVueHookTsxStr 创建vue hook.tsx str
func (m *GTemplate) GenerateVueHookTsxStr() (string, error) {
	tmpl, err := m.generateTmpl("hook", hookTemplatePath, template.FuncMap{
		"tableColumnList": m.tableColumnList,
		"switchChange":    m.switchChange,
		"customVar":       m.customVarAllStrGen,
		"openDialogProps": m.openDialogProps,
	})
	if err != nil {
		return "", err
	}
	return m.executeStr(tmpl)
}

// 暂时不考虑 主页的开关选项
func (m *GTemplate) switchChange() string {
	return `//主页快速修改快关占位`
}

// 主页表单列
func (m *GTemplate) tableColumnList(fields []*TableFields) string {
	var sb strings.Builder
	for _, field := range fields {
		if field.List {
			sb.WriteString("\t")
			//这里应该是显示类型
			switch field.ShowType {
			case ShowDate:
				sb.WriteString(`{
     label: "` + field.FieldDes + `",
     prop: "` + field.JsonName + `",
     minWidth: 160,
     formatter: ({ ` + field.JsonName + ` }) =>
       dayjs(` + field.JsonName + `).format("YYYY-MM-DD HH:mm:ss")
   },`)
			default:
				sb.WriteString(`{
     label: "` + field.FieldDes + `",
     prop: "` + field.JsonName + `",
     minWidth: 90
   },`)
			}
			sb.WriteString("\n")
		}
	}
	return sb.String()
}

// 查询需要的变量
func (m *GTemplate) customVarAllStrGen(fields []*TableFields) string {
	var sb strings.Builder
	for i, field := range fields {
		if field.Query {
			sb.WriteString("\t")
			sb.WriteString(field.JsonName + `:""`)
			if i != len(fields)-1 {
				sb.WriteString(`,` + "\n")
			}
		}
	}
	return sb.String()
}

// 打开传值
func (m *GTemplate) openDialogProps(fields []*TableFields) string {
	var sb strings.Builder
	for i, field := range fields {
		//允许编辑
		if field.Edit && field.Field != "id" {
			sb.WriteString("\t\t\t\t\t")
			if field.TsType == TsStringType {
				sb.WriteString(field.JsonName + `: row?.` + field.JsonName + ` ?? 0`)
			} else {
				sb.WriteString(field.JsonName + `: row?.` + field.JsonName + ` ?? ""`)
			}
			if i != len(fields)-1 {
				sb.WriteString(`,` + "\n")
			}
		}
	}
	return sb.String()
}

func (m *GTemplate) GenerateVueHookCode(code string) error {
	//vueRoot vueApiPath/所属菜单/功能名称/utils/
	outDir := filepath.Join(vueRoot, vueViewPath, m.genInfo.Directory,
		gstr.ToCamelCase(m.genInfo.BusinessName), "utils", "hook.tsx")
	return m.Gen(outDir, code, false)
}

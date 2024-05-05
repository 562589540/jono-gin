package internal

import (
	"github.com/562589540/jono-gin/ghub/glibrary/gstr"
	"github.com/562589540/jono-gin/ghub/glibrary/gtemplate/enum"
	"github.com/562589540/jono-gin/ghub/glibrary/gtemplate/pkg"
	"path/filepath"
	"strings"
	"text/template"
)

type GVueHookGen struct {
	*GTemplate
}

func (m *GVueHookGen) GenCodeStr() (string, error) {
	tmpl, err := m.generateVueTmpl("hook", hookTemplatePath, template.FuncMap{
		"tableColumnList": m.tableColumnList,
		"switchChange":    m.switchChange,
		"customVar":       m.customVarAllStrGen,
		"customSearch":    m.customSearch,
		"openDialogProps": m.openDialogProps,
	})
	if err != nil {
		return "", err
	}
	return m.executeStr(tmpl)
}

func (m *GVueHookGen) GenCode(codes pkg.GenCodes) error {
	//vueRoot vueApiPath/所属菜单/功能名称/utils/
	outDir := filepath.Join(vueRoot, vueViewPath, m.GenInfo.Directory,
		gstr.ToCamelCase(m.GenInfo.BusinessName), "utils", "hook.tsx")
	return m.gen(outDir, codes.VueHookCode, true)
}

func (m *GVueHookGen) UpdateGeneratedCode(code *pkg.GenCodes, str string) {
	code.VueHookCode = str
}

// 暂时不考虑 主页的开关选项
func (m *GVueHookGen) switchChange() string {
	return `//主页快速修改快关占位`
}

// 主页表单列
func (m *GVueHookGen) tableColumnList(fields []*pkg.TableFields) string {
	var sb strings.Builder
	for _, field := range fields {
		if field.List {
			sb.WriteString("\t")
			//这里应该是显示类型
			switch field.ShowType {
			case enum.ShowRadio:
				{
					sb.WriteString(`{
      label: "` + field.FieldDes + `",
      prop: "` + field.JsonName + `",
      minWidth: 100,
      cellRenderer: ({ row, props }) => (
        <el-tag size={props.size} style={tagStyle.value(row.` + field.JsonName + `)}>
          {row.` + field.JsonName + ` === 1 ? "启用" : "停用"}
        </el-tag>
      )
    },`)
				}
			case enum.ShowDate:
				{
					sb.WriteString(`{
     label: "` + field.FieldDes + `",
     prop: "` + field.JsonName + `",
     minWidth: 160,
     formatter: ({ ` + field.JsonName + ` }) =>
       dayjs(` + field.JsonName + `).format("YYYY-MM-DD HH:mm:ss")
   },`)
				}
			default:
				{
					sb.WriteString(`{
     label: "` + field.FieldDes + `",
     prop: "` + field.JsonName + `",
     minWidth: 90
   },`)
				}
			}
			sb.WriteString("\n")
		}
	}
	return sb.String()
}

// 查询需要的变量
func (m *GVueHookGen) customVarAllStrGen(fields []*pkg.TableFields) string {
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

// 打开表单vue传值
func (m *GVueHookGen) openDialogProps(fields []*pkg.TableFields) string {
	var sb strings.Builder
	for i, field := range fields {
		//允许编辑
		if field.Edit && field.Field != "id" {
			sb.WriteString("\t\t\t\t\t")
			if field.TsType == enum.TsStringType {
				sb.WriteString(field.JsonName + `: row?.` + field.JsonName + ` ?? ""`)
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

// 表单元素
func (m *GVueHookGen) customSearch(fields []*pkg.TableFields) string {
	var sb strings.Builder
	for i, field := range fields {
		//搜索
		if field.Query && field.Field != "id" {
			sb.WriteString("\t\t")
			switch field.ShowType {
			case enum.ShowRadio:
				m.formSwitch(&sb, field) //单选下拉框
			case enum.ShowDate:
				m.formBetweenTime(&sb, field) //时间区间
			default:
				m.formInput(&sb, field) //普通输入框
			}
			if i != len(fields)-1 {
				sb.WriteString(",\n")
			}
		}
	}
	return sb.String()
}

func (m *GVueHookGen) formInput(sb *strings.Builder, field *pkg.TableFields) {
	sb.WriteString(`{
    label: "` + field.FieldDes + `",
    prop: "` + field.JsonName + `",
    valueType: "copy",
  }`)
}

func (m *GVueHookGen) formSwitch(sb *strings.Builder, field *pkg.TableFields) {
	sb.WriteString(`{
    label: "` + field.FieldDes + `",
    prop: "` + field.JsonName + `",
    valueType: "select",
    options: [
      {
        label: "启用",
        value: "1",
        color: "blue"
      },
      {
        label: "停用",
        value: "0",
        color: "red"
      }
    ]
  }`)
}

func (m *GVueHookGen) formBetweenTime(sb *strings.Builder, field *pkg.TableFields) {
	sb.WriteString(`  {
    label: "` + field.FieldDes + `",
    prop: "` + field.JsonName + `",
    valueType: "date-picker",
    fieldProps: {
      type: "datetimerange",
      startPlaceholder: "请选择",
      endPlaceholder: "请选择"
    }
  }`)
}

func (m *GVueHookGen) formTime(sb *strings.Builder, field *pkg.TableFields) {

}

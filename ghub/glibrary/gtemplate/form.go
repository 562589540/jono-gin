package gtemplate

import (
	"github.com/562589540/jono-gin/ghub/glibrary/gstr"
	"path/filepath"
	"strings"
	"text/template"
)

// GenerateVueFormStr 创建vue Form 页面的字符串
func (m *GTemplate) GenerateVueFormStr() (string, error) {
	tmpl, err := m.generateTmpl("form", formTemplatePath, template.FuncMap{
		"customFormItem": m.customFormItemGen,
		"customVar":      m.customVarGen,
	})
	if err != nil {
		return "", err
	}
	return m.executeStr(tmpl)
}

// 变量
func (m *GTemplate) customVarGen(fields []*TableFields) string {
	var sb strings.Builder
	for i, field := range fields {
		//允许编辑
		if field.Edit && field.Field != "id" {
			sb.WriteString("\t\t")
			if field.TsType == TsNumberType {
				sb.WriteString(field.JsonName + `:0`)
			} else {
				sb.WriteString(field.JsonName + `:""`)
			}
			if i != len(fields)-1 {
				sb.WriteString(`,` + "\n")
			}
		}
	}
	return sb.String()
}

func (m *GTemplate) customFormItemGen(fields []*TableFields) string {
	var sb strings.Builder
	for i, field := range fields {
		//允许编辑
		if field.Edit && field.Field != "id" {
			sb.WriteString("\t\t")
			//是否占一行
			if field.FillUp {
				sb.WriteString(`<re-col>`)
			} else {
				sb.WriteString(`<re-col :value="12" :xs="24" :sm="24">`)
			}
			sb.WriteString("\n\t\t\t")
			switch field.ShowType {
			case ShowTextarea:
				m.formTextarea(&sb, field)
			case ShowRadio:
				m.formSwitch(&sb, field)
			default:
				m.formInput(&sb, field)
			}
			sb.WriteString("\n\t\t" + `</re-col>` + "\n")
			if i != len(fields)-1 {
				sb.WriteString("\n")
			}
		}

	}
	return sb.String()
}

func (m *GTemplate) formInput(sb *strings.Builder, field *TableFields) {
	sb.WriteString(`<el-form-item label="` + field.FieldDes + `" prop="` + field.JsonName + `">
     	<el-input
       	v-model="newFormInline.` + field.JsonName + `"
       	clearable
       	placeholder="请输入` + field.FieldDes + `"
     	/>
   	</el-form-item>`)
}

func (m *GTemplate) formTextarea(sb *strings.Builder, field *TableFields) {
	sb.WriteString(`<el-form-item label="` + field.FieldDes + `">
     	<el-input
       	v-model="newFormInline.` + field.JsonName + `"
       	placeholder="请输入` + field.FieldDes + `信息"
       	type="textarea"
     	/>
   	</el-form-item>`)
}

func (m *GTemplate) formSwitch(sb *strings.Builder, field *TableFields) {
	sb.WriteString(`<el-form-item label="` + field.FieldDes + `">
   	<el-switch
           v-model="newFormInline.` + field.JsonName + `"
           inline-prompt
           :active-value="1"
           :inactive-value="0"
           active-text="启用"
           inactive-text="停用"
         />
       </el-form-item>`)
}

func (m *GTemplate) GenerateVueFormCode(code string) error {
	//vueRoot vueApiPath/所属菜单/功能名称/form.vue
	outDir := filepath.Join(vueRoot, vueViewPath, m.genInfo.Directory,
		gstr.ToCamelCase(m.genInfo.BusinessName), "form.vue")
	return m.Gen(outDir, code, false)
}

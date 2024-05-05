package internal

import (
	"github.com/562589540/jono-gin/ghub/glibrary/gstr"
	"github.com/562589540/jono-gin/ghub/glibrary/gtemplate/enum"
	"github.com/562589540/jono-gin/ghub/glibrary/gtemplate/pkg"
	"path/filepath"
	"strings"
	"text/template"
)

type GVueFormGen struct {
	*GTemplate
}

func (m *GVueFormGen) GenCodeStr() (string, error) {
	tmpl, err := m.generateVueTmpl("form", formTemplatePath, template.FuncMap{
		"customFormItem": m.customFormItemGen,
		"customVar":      m.customVarGen,
	})
	if err != nil {
		return "", err
	}
	return m.executeStr(tmpl)
}

func (m *GVueFormGen) GenCode(codes pkg.GenCodes) error {
	//vueRoot vueApiPath/所属菜单/功能名称/form.vue
	outDir := filepath.Join(vueRoot, vueViewPath, m.GenInfo.Directory,
		gstr.ToCamelCase(m.GenInfo.BusinessName), "form.vue")
	return m.gen(outDir, codes.VueFormCode, true)
}

func (m *GVueFormGen) UpdateGeneratedCode(code *pkg.GenCodes, str string) {
	code.VueFormCode = str
}

// 变量
func (m *GVueFormGen) customVarGen(fields []*pkg.TableFields) string {
	var sb strings.Builder
	for i, field := range fields {
		//允许编辑
		if field.Edit && field.Field != "id" {
			sb.WriteString("\t\t")
			if field.TsType == enum.TsNumberType {
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

// 表单元素
func (m *GVueFormGen) customFormItemGen(fields []*pkg.TableFields) string {
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
			case enum.ShowNum:
				m.formNum(&sb, field)
			case enum.ShowTextarea:
				m.formTextarea(&sb, field)
			case enum.ShowRadio:
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

// 表单input类型
func (m *GVueFormGen) formInput(sb *strings.Builder, field *pkg.TableFields) {
	sb.WriteString(`<el-form-item label="` + field.FieldDes + `" prop="` + field.JsonName + `">
     	<el-input
       	v-model="newFormInline.` + field.JsonName + `"
       	clearable
       	placeholder="请输入` + field.FieldDes + `"
     	/>
   	</el-form-item>`)
}

// 表单数字类型
func (m *GVueFormGen) formNum(sb *strings.Builder, field *pkg.TableFields) {
	sb.WriteString(`<el-form-item label="` + field.FieldDes + `">
          <el-input-number
            v-model="newFormInline.` + field.JsonName + `"
            class="!w-full"
            :min="0"
            :max="9999"
            controls-position="right"
          />
        </el-form-item>`)
}

// 表单Textarea类型
func (m *GVueFormGen) formTextarea(sb *strings.Builder, field *pkg.TableFields) {
	sb.WriteString(`<el-form-item label="` + field.FieldDes + `">
     	<el-input
       	v-model="newFormInline.` + field.JsonName + `"
       	placeholder="请输入` + field.FieldDes + `信息"
       	type="textarea"
     	/>
   	</el-form-item>`)
}

// 表单Switch类型
func (m *GVueFormGen) formSwitch(sb *strings.Builder, field *pkg.TableFields) {
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

package gtemplate

import (
	"github.com/562589540/jono-gin/ghub/glibrary/gstr"
	"path/filepath"
	"strings"
	"text/template"
)

// GenerateVueIndexStr 创建vue Index 页面的字符串
func (m *GTemplate) GenerateVueIndexStr() (string, error) {
	tmpl, err := m.generateTmpl("index", indexTemplatePath, template.FuncMap{
		"customIndexFormItem": m.customIndexFormItem,
	})
	if err != nil {
		return "", err
	}
	return m.executeStr(tmpl)
}

// 主页搜索框
func (m *GTemplate) customIndexFormItem(fields []*TableFields) string {
	var sb strings.Builder
	for i, field := range fields {
		if field.Query {
			sb.WriteString("\t\t")
			switch field.ShowType {
			case ShowRadio:
				m.indexFormRadio(&sb, field)
			default:
				m.indexFormInput(&sb, field)
			}
			if i != len(fields)-1 {
				sb.WriteString("\n")
			}
		}
	}
	return sb.String()
}

func (m *GTemplate) indexFormInput(sb *strings.Builder, field *TableFields) {
	sb.WriteString(`<el-form-item label="` + field.FieldDes + `：" prop="` + field.JsonName + `">
        <el-input
          v-model="form.` + field.JsonName + `"
          placeholder="请输入` + field.FieldDes + `"
          clearable
          class="!w-[180px]"
        />
      </el-form-item>`)
}

func (m *GTemplate) indexFormRadio(sb *strings.Builder, field *TableFields) {
	sb.WriteString(`<el-form-item label="` + field.FieldDes + `：" prop="` + field.JsonName + `">
        <el-select
          v-model="form.` + field.JsonName + `"
          placeholder="请选择` + field.FieldDes + `"
          clearable
          class="!w-[180px]"
        >
          <el-option label="已启用" value="1" />
          <el-option label="已停用" value="0" />
        </el-select>
      </el-form-item>`)
}

func (m *GTemplate) GenerateVueIndexCode(code string) error {
	//vueRoot vueApiPath/所属菜单/功能名称/index.vue
	outDir := filepath.Join(vueRoot, vueViewPath, m.genInfo.Directory,
		gstr.ToCamelCase(m.genInfo.BusinessName), "index.vue")
	return m.Gen(outDir, code, false)
}

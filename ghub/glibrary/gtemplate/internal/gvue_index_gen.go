package internal

import (
	"github.com/562589540/jono-gin/ghub/glibrary/gstr"
	"github.com/562589540/jono-gin/ghub/glibrary/gtemplate/enum"
	"github.com/562589540/jono-gin/ghub/glibrary/gtemplate/pkg"
	"path/filepath"
	"strings"
	"text/template"
)

type GVueIndexGen struct {
	*GTemplate
}

func (m *GVueIndexGen) GenCodeStr() (string, error) {
	tmpl, err := m.generateVueTmpl("index", indexTemplatePath, template.FuncMap{
		"customIndexFormItem": m.customIndexFormItem,
	})
	if err != nil {
		return "", err
	}
	return m.executeStr(tmpl)
}

func (m *GVueIndexGen) GenCode(codes pkg.GenCodes) error {
	//vueRoot vueApiPath/所属菜单/功能名称/index.vue
	outDir := filepath.Join(vueRoot, vueViewPath, m.GenInfo.Directory,
		gstr.ToCamelCase(m.GenInfo.BusinessName), "index.vue")
	return m.gen(outDir, codes.VueIndexCode, true)
}

func (m *GVueIndexGen) UpdateGeneratedCode(code *pkg.GenCodes, str string) {
	code.VueIndexCode = str
}

// 主页搜索框
func (m *GVueIndexGen) customIndexFormItem(fields []*pkg.TableFields) string {
	var sb strings.Builder
	for i, field := range fields {
		if field.Query {
			sb.WriteString("\t\t")
			switch field.ShowType {
			case enum.ShowRadio:
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

func (m *GVueIndexGen) indexFormInput(sb *strings.Builder, field *pkg.TableFields) {
	sb.WriteString(`<el-form-item label="` + field.FieldDes + `：" prop="` + field.JsonName + `">
        <el-input
          v-model="form.` + field.JsonName + `"
          placeholder="请输入` + field.FieldDes + `"
          clearable
          class="!w-[180px]"
        />
      </el-form-item>`)
}

func (m *GVueIndexGen) indexFormRadio(sb *strings.Builder, field *pkg.TableFields) {
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

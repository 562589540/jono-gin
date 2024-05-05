package gtemplate

import (
	"github.com/562589540/jono-gin/ghub/glibrary/gtemplate/internal"
	"github.com/562589540/jono-gin/ghub/glibrary/gtemplate/pkg"
)

func NewGVueIndexGen(tp *internal.GTemplate) pkg.IGenCode {
	return &internal.GVueIndexGen{GTemplate: tp}
}
func NewGVueFormGen(tp *internal.GTemplate) pkg.IGenCode {
	return &internal.GVueFormGen{GTemplate: tp}
}
func NewGVueHookGen(tp *internal.GTemplate) pkg.IGenCode {
	return &internal.GVueHookGen{GTemplate: tp}
}
func NewGVueRuleGen(tp *internal.GTemplate) pkg.IGenCode {
	return &internal.GVueRuleGen{GTemplate: tp}
}
func NewGVueTypesGen(tp *internal.GTemplate) pkg.IGenCode {
	return &internal.GVueTypesGen{GTemplate: tp}
}
func NewGVueApiGen(tp *internal.GTemplate) pkg.IGenCode {
	return &internal.GVueApiGen{GTemplate: tp}
}
func NewGGoModelGen(tp *internal.GTemplate) pkg.IGenCode {
	return &internal.GGoModelGen{GTemplate: tp}
}
func NewGGoLogicGen(tp *internal.GTemplate) pkg.IGenCode {
	return &internal.GGoLogicGen{GTemplate: tp}
}
func NewGGoServiceGen(tp *internal.GTemplate) pkg.IGenCode {
	return &internal.GGoServiceGen{GTemplate: tp}
}
func NewGGoRouterGen(tp *internal.GTemplate) pkg.IGenCode {
	return &internal.GGoRouterGen{GTemplate: tp}
}
func NewGGoDTOGen(tp *internal.GTemplate) pkg.IGenCode {
	return &internal.GGoDTOGen{GTemplate: tp}
}
func NewGGoApiGen(tp *internal.GTemplate) pkg.IGenCode {
	return &internal.GGoApiGen{GTemplate: tp}
}
func NewGTemplate(baseInfo *pkg.BaseInfo, genInfo *pkg.GenInfo, fieldsInfo []*pkg.TableFields) *internal.GTemplate {
	return &internal.GTemplate{
		BaseInfo:   baseInfo,
		GenInfo:    genInfo,
		FieldsInfo: fieldsInfo,
	}
}

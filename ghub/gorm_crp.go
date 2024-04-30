package ghub

import "github.com/562589540/jono-gin/ghub/glibrary/gorm_crp"

func GormCrp[T any]() gorm_crp.QueryBuilder[T] {
	return gorm_crp.New[T]()
}

// ConditionHandler 条件查询
func ConditionHandler() *gorm_crp.ConditionHandler {
	return &gorm_crp.ConditionHandler{}
}

// ModifierHandler 复杂条件查询
func ModifierHandler() *gorm_crp.ModifierHandler {
	return &gorm_crp.ModifierHandler{}
}

// LikeHandler 模糊查询
func LikeHandler() *gorm_crp.LikeHandler {
	return &gorm_crp.LikeHandler{}
}

// OrderHandler 排序
func OrderHandler() *gorm_crp.OrderHandler {
	return &gorm_crp.OrderHandler{}
}

// PreloadHandler 关联
func PreloadHandler() *gorm_crp.PreloadHandler {
	return &gorm_crp.PreloadHandler{}
}

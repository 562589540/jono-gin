package gorm_crp

import (
	"github.com/562589540/jono-gin/ghub/glibrary/gstruct/g"
	"gorm.io/gorm"
)

// ConditionHandler 条件查询处理器
type ConditionHandler struct{}

func (h *ConditionHandler) Handle(query *gorm.DB, options g.FindOptions) (*gorm.DB, error) {
	if len(options.Condition) > 0 {
		query = query.Where(options.Condition)
	}
	return query, nil
}

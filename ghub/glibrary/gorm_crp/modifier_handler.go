package gorm_crp

import (
	"github.com/562589540/jono-gin/ghub/glibrary/gstruct/g"
	"gorm.io/gorm"
)

// LikeHandler 复杂条件查询处理器
type ModifierHandler struct{}

func (h *ModifierHandler) Handle(query *gorm.DB, options g.FindOptions) (*gorm.DB, error) {
	for _, customQuery := range options.CustomQueries {
		query = customQuery.Apply(query)
	}
	return query, nil
}

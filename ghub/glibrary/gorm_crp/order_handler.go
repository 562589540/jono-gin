package gorm_crp

import (
	"github.com/562589540/jono-gin/ghub/glibrary/gstruct/g"
	"gorm.io/gorm"
)

// OrderHandler 排序处理器
type OrderHandler struct{}

func (h *OrderHandler) Handle(query *gorm.DB, options g.FindOptions) (*gorm.DB, error) {
	if options.SortBy != "" {
		orderDirection := "asc"
		if options.IsSortDesc {
			orderDirection = "desc"
		}
		query = query.Order(options.SortBy + " " + orderDirection)
	}
	return query, nil
}

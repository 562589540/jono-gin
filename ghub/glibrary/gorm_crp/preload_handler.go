package gorm_crp

import (
	"github.com/562589540/jono-gin/ghub/glibrary/gstruct/g"
	"gorm.io/gorm"
)

// PreloadHandler 关联预加载处理器
type PreloadHandler struct{}

func (h *PreloadHandler) Handle(query *gorm.DB, options g.FindOptions) (*gorm.DB, error) {
	if len(options.Preload) > 0 {
		for _, preloadRelation := range options.Preload {
			query = query.Preload(preloadRelation)
		}
	}
	return query, nil
}

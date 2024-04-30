package gorm_crp

import (
	"github.com/562589540/jono-gin/ghub/glibrary/gstruct/g"
	"gorm.io/gorm"
)

// LikeHandler 模糊查询处理器
type LikeHandler struct{}

func (h *LikeHandler) Handle(query *gorm.DB, options g.FindOptions) (*gorm.DB, error) {
	if len(options.Like) > 0 {
		for key, value := range options.Like {
			query = query.Where(key+" LIKE ?", "%"+value.(string)+"%")
		}
	}
	return query, nil
}

package gorm_crp

import (
	"github.com/562589540/jono-gin/ghub/glibrary/gstruct/g"
	"gorm.io/gorm"
)

type QueryHandler interface {
	Handle(query *gorm.DB, options g.FindOptions) (*gorm.DB, error)
}

package gorm_crp

import (
	"github.com/562589540/jono-gin/ghub/glibrary/gstruct/g"
	"gorm.io/gorm"
)

type QueryBuilder[T any] struct {
	handlers []QueryHandler
}

func New[T any]() QueryBuilder[T] {
	return QueryBuilder[T]{
		handlers: make([]QueryHandler, 0),
	}
}

func (b *QueryBuilder[T]) AddHandler(handler QueryHandler) {
	b.handlers = append(b.handlers, handler)
}

func (b *QueryBuilder[T]) buildQuery(baseQuery *gorm.DB, options g.FindOptions) (*gorm.DB, error) {
	var err error
	query := baseQuery
	for _, handler := range b.handlers {
		query, err = handler.Handle(query, options)
		if err != nil {
			return nil, err
		}
	}
	return query, nil
}

func (b *QueryBuilder[T]) ExecuteQuery(query *gorm.DB, options *g.FindOptions) (entities []T, totalRows int64, err error) {
	entities = make([]T, 0)
	//构造查询条件
	query, err = b.buildQuery(query, *options)
	if err != nil {
		return
	}
	//分页和总数
	if options.Paginate != nil {
		if err = query.Count(&totalRows).Error; err != nil {
			return
		}
		query = query.Offset((options.Paginate.Page - 1) * options.Paginate.PageSize).Limit(options.Paginate.PageSize)
	}
	//扫描
	if options.ScanStruct != nil {
		err = query.Scan(&options.ScanStruct).Error
	} else {
		err = query.Find(&entities).Error
	}
	return
}

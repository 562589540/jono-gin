package g

import "gorm.io/gorm"

// CustomQuery 复杂查询构造器
type CustomQuery struct {
	Query string        // 查询语句，例如 "age > ? AND score < ?"
	Args  []interface{} // 查询参数
}

func (cq CustomQuery) Apply(query *gorm.DB) *gorm.DB {
	return query.Where(cq.Query, cq.Args...)
}

type Paginate struct {
	Page     int
	PageSize int
}

type FindOptions struct {
	Preload       []string               //关联查询表
	CustomQueries []CustomQuery          //复杂查询构造器
	Condition     map[string]interface{} //查询条件
	Like          map[string]interface{} // 模糊查询条件
	IsSortDesc    bool                   //是否逆序
	ScanStruct    any                    //查询结构扫描到的结构体
	Paginate      *Paginate
	SortBy        string //排序字段
}

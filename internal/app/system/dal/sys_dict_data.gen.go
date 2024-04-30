// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package dal

import (
	"context"
	"strings"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"github.com/562589540/jono-gin/internal/app/system/model"
)

func newDictData(db *gorm.DB, opts ...gen.DOOption) dictData {
	_dictData := dictData{}

	_dictData.dictDataDo.UseDB(db, opts...)
	_dictData.dictDataDo.UseModel(&model.DictData{})

	tableName := _dictData.dictDataDo.TableName()
	_dictData.ALL = field.NewAsterisk(tableName)
	_dictData.DictCode = field.NewInt64(tableName, "dict_code")
	_dictData.DictSort = field.NewInt32(tableName, "dict_sort")
	_dictData.DictLabel = field.NewString(tableName, "dict_label")
	_dictData.DictValue = field.NewString(tableName, "dict_value")
	_dictData.DictType = field.NewString(tableName, "dict_type")
	_dictData.CSSClass = field.NewString(tableName, "css_class")
	_dictData.ListClass = field.NewString(tableName, "list_class")
	_dictData.IsDefault = field.NewBool(tableName, "is_default")
	_dictData.Status = field.NewBool(tableName, "status")
	_dictData.CreateBy = field.NewInt64(tableName, "create_by")
	_dictData.UpdateBy = field.NewInt64(tableName, "update_by")
	_dictData.Remark = field.NewString(tableName, "remark")
	_dictData.CreatedAt = field.NewTime(tableName, "created_at")
	_dictData.UpdatedAt = field.NewTime(tableName, "updated_at")

	_dictData.fillFieldMap()

	return _dictData
}

type dictData struct {
	dictDataDo

	ALL       field.Asterisk
	DictCode  field.Int64
	DictSort  field.Int32
	DictLabel field.String
	DictValue field.String
	DictType  field.String
	CSSClass  field.String
	ListClass field.String
	IsDefault field.Bool
	Status    field.Bool
	CreateBy  field.Int64
	UpdateBy  field.Int64
	Remark    field.String
	CreatedAt field.Time
	UpdatedAt field.Time

	fieldMap map[string]field.Expr
}

func (d dictData) Table(newTableName string) *dictData {
	d.dictDataDo.UseTable(newTableName)
	return d.updateTableName(newTableName)
}

func (d dictData) As(alias string) *dictData {
	d.dictDataDo.DO = *(d.dictDataDo.As(alias).(*gen.DO))
	return d.updateTableName(alias)
}

func (d *dictData) updateTableName(table string) *dictData {
	d.ALL = field.NewAsterisk(table)
	d.DictCode = field.NewInt64(table, "dict_code")
	d.DictSort = field.NewInt32(table, "dict_sort")
	d.DictLabel = field.NewString(table, "dict_label")
	d.DictValue = field.NewString(table, "dict_value")
	d.DictType = field.NewString(table, "dict_type")
	d.CSSClass = field.NewString(table, "css_class")
	d.ListClass = field.NewString(table, "list_class")
	d.IsDefault = field.NewBool(table, "is_default")
	d.Status = field.NewBool(table, "status")
	d.CreateBy = field.NewInt64(table, "create_by")
	d.UpdateBy = field.NewInt64(table, "update_by")
	d.Remark = field.NewString(table, "remark")
	d.CreatedAt = field.NewTime(table, "created_at")
	d.UpdatedAt = field.NewTime(table, "updated_at")

	d.fillFieldMap()

	return d
}

func (d *dictData) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := d.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (d *dictData) fillFieldMap() {
	d.fieldMap = make(map[string]field.Expr, 14)
	d.fieldMap["dict_code"] = d.DictCode
	d.fieldMap["dict_sort"] = d.DictSort
	d.fieldMap["dict_label"] = d.DictLabel
	d.fieldMap["dict_value"] = d.DictValue
	d.fieldMap["dict_type"] = d.DictType
	d.fieldMap["css_class"] = d.CSSClass
	d.fieldMap["list_class"] = d.ListClass
	d.fieldMap["is_default"] = d.IsDefault
	d.fieldMap["status"] = d.Status
	d.fieldMap["create_by"] = d.CreateBy
	d.fieldMap["update_by"] = d.UpdateBy
	d.fieldMap["remark"] = d.Remark
	d.fieldMap["created_at"] = d.CreatedAt
	d.fieldMap["updated_at"] = d.UpdatedAt
}

func (d dictData) clone(db *gorm.DB) dictData {
	d.dictDataDo.ReplaceConnPool(db.Statement.ConnPool)
	return d
}

func (d dictData) replaceDB(db *gorm.DB) dictData {
	d.dictDataDo.ReplaceDB(db)
	return d
}

type dictDataDo struct{ gen.DO }

type IDictDataDo interface {
	gen.SubQuery
	Debug() IDictDataDo
	WithContext(ctx context.Context) IDictDataDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IDictDataDo
	WriteDB() IDictDataDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IDictDataDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IDictDataDo
	Not(conds ...gen.Condition) IDictDataDo
	Or(conds ...gen.Condition) IDictDataDo
	Select(conds ...field.Expr) IDictDataDo
	Where(conds ...gen.Condition) IDictDataDo
	Order(conds ...field.Expr) IDictDataDo
	Distinct(cols ...field.Expr) IDictDataDo
	Omit(cols ...field.Expr) IDictDataDo
	Join(table schema.Tabler, on ...field.Expr) IDictDataDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IDictDataDo
	RightJoin(table schema.Tabler, on ...field.Expr) IDictDataDo
	Group(cols ...field.Expr) IDictDataDo
	Having(conds ...gen.Condition) IDictDataDo
	Limit(limit int) IDictDataDo
	Offset(offset int) IDictDataDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IDictDataDo
	Unscoped() IDictDataDo
	Create(values ...*model.DictData) error
	CreateInBatches(values []*model.DictData, batchSize int) error
	Save(values ...*model.DictData) error
	First() (*model.DictData, error)
	Take() (*model.DictData, error)
	Last() (*model.DictData, error)
	Find() ([]*model.DictData, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.DictData, err error)
	FindInBatches(result *[]*model.DictData, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.DictData) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IDictDataDo
	Assign(attrs ...field.AssignExpr) IDictDataDo
	Joins(fields ...field.RelationField) IDictDataDo
	Preload(fields ...field.RelationField) IDictDataDo
	FirstOrInit() (*model.DictData, error)
	FirstOrCreate() (*model.DictData, error)
	FindByPage(offset int, limit int) (result []*model.DictData, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IDictDataDo
	UnderlyingDB() *gorm.DB
	schema.Tabler

	GetByID(id uint) (result model.DictData, err error)
	DeleteByID(id uint) (rowsAffected int64, err error)
	DeleteByIDs(ids []uint) (rowsAffected int64, err error)
}

// SELECT * FROM @@table WHERE id=@id
func (d dictDataDo) GetByID(id uint) (result model.DictData, err error) {
	var params []interface{}

	var generateSQL strings.Builder
	params = append(params, id)
	generateSQL.WriteString("SELECT * FROM sys_dict_data WHERE id=? ")

	var executeSQL *gorm.DB
	executeSQL = d.UnderlyingDB().Raw(generateSQL.String(), params...).Take(&result) // ignore_security_alert
	err = executeSQL.Error

	return
}

// DELETE FROM @@table WHERE id=@id
func (d dictDataDo) DeleteByID(id uint) (rowsAffected int64, err error) {
	var params []interface{}

	var generateSQL strings.Builder
	params = append(params, id)
	generateSQL.WriteString("DELETE FROM sys_dict_data WHERE id=? ")

	var executeSQL *gorm.DB
	executeSQL = d.UnderlyingDB().Exec(generateSQL.String(), params...) // ignore_security_alert
	rowsAffected = executeSQL.RowsAffected
	err = executeSQL.Error

	return
}

// DELETE FROM @@table WHERE id IN (@ids)
func (d dictDataDo) DeleteByIDs(ids []uint) (rowsAffected int64, err error) {
	var params []interface{}

	var generateSQL strings.Builder
	params = append(params, ids)
	generateSQL.WriteString("DELETE FROM sys_dict_data WHERE id IN (?) ")

	var executeSQL *gorm.DB
	executeSQL = d.UnderlyingDB().Exec(generateSQL.String(), params...) // ignore_security_alert
	rowsAffected = executeSQL.RowsAffected
	err = executeSQL.Error

	return
}

func (d dictDataDo) Debug() IDictDataDo {
	return d.withDO(d.DO.Debug())
}

func (d dictDataDo) WithContext(ctx context.Context) IDictDataDo {
	return d.withDO(d.DO.WithContext(ctx))
}

func (d dictDataDo) ReadDB() IDictDataDo {
	return d.Clauses(dbresolver.Read)
}

func (d dictDataDo) WriteDB() IDictDataDo {
	return d.Clauses(dbresolver.Write)
}

func (d dictDataDo) Session(config *gorm.Session) IDictDataDo {
	return d.withDO(d.DO.Session(config))
}

func (d dictDataDo) Clauses(conds ...clause.Expression) IDictDataDo {
	return d.withDO(d.DO.Clauses(conds...))
}

func (d dictDataDo) Returning(value interface{}, columns ...string) IDictDataDo {
	return d.withDO(d.DO.Returning(value, columns...))
}

func (d dictDataDo) Not(conds ...gen.Condition) IDictDataDo {
	return d.withDO(d.DO.Not(conds...))
}

func (d dictDataDo) Or(conds ...gen.Condition) IDictDataDo {
	return d.withDO(d.DO.Or(conds...))
}

func (d dictDataDo) Select(conds ...field.Expr) IDictDataDo {
	return d.withDO(d.DO.Select(conds...))
}

func (d dictDataDo) Where(conds ...gen.Condition) IDictDataDo {
	return d.withDO(d.DO.Where(conds...))
}

func (d dictDataDo) Order(conds ...field.Expr) IDictDataDo {
	return d.withDO(d.DO.Order(conds...))
}

func (d dictDataDo) Distinct(cols ...field.Expr) IDictDataDo {
	return d.withDO(d.DO.Distinct(cols...))
}

func (d dictDataDo) Omit(cols ...field.Expr) IDictDataDo {
	return d.withDO(d.DO.Omit(cols...))
}

func (d dictDataDo) Join(table schema.Tabler, on ...field.Expr) IDictDataDo {
	return d.withDO(d.DO.Join(table, on...))
}

func (d dictDataDo) LeftJoin(table schema.Tabler, on ...field.Expr) IDictDataDo {
	return d.withDO(d.DO.LeftJoin(table, on...))
}

func (d dictDataDo) RightJoin(table schema.Tabler, on ...field.Expr) IDictDataDo {
	return d.withDO(d.DO.RightJoin(table, on...))
}

func (d dictDataDo) Group(cols ...field.Expr) IDictDataDo {
	return d.withDO(d.DO.Group(cols...))
}

func (d dictDataDo) Having(conds ...gen.Condition) IDictDataDo {
	return d.withDO(d.DO.Having(conds...))
}

func (d dictDataDo) Limit(limit int) IDictDataDo {
	return d.withDO(d.DO.Limit(limit))
}

func (d dictDataDo) Offset(offset int) IDictDataDo {
	return d.withDO(d.DO.Offset(offset))
}

func (d dictDataDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IDictDataDo {
	return d.withDO(d.DO.Scopes(funcs...))
}

func (d dictDataDo) Unscoped() IDictDataDo {
	return d.withDO(d.DO.Unscoped())
}

func (d dictDataDo) Create(values ...*model.DictData) error {
	if len(values) == 0 {
		return nil
	}
	return d.DO.Create(values)
}

func (d dictDataDo) CreateInBatches(values []*model.DictData, batchSize int) error {
	return d.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (d dictDataDo) Save(values ...*model.DictData) error {
	if len(values) == 0 {
		return nil
	}
	return d.DO.Save(values)
}

func (d dictDataDo) First() (*model.DictData, error) {
	if result, err := d.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.DictData), nil
	}
}

func (d dictDataDo) Take() (*model.DictData, error) {
	if result, err := d.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.DictData), nil
	}
}

func (d dictDataDo) Last() (*model.DictData, error) {
	if result, err := d.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.DictData), nil
	}
}

func (d dictDataDo) Find() ([]*model.DictData, error) {
	result, err := d.DO.Find()
	return result.([]*model.DictData), err
}

func (d dictDataDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.DictData, err error) {
	buf := make([]*model.DictData, 0, batchSize)
	err = d.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (d dictDataDo) FindInBatches(result *[]*model.DictData, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return d.DO.FindInBatches(result, batchSize, fc)
}

func (d dictDataDo) Attrs(attrs ...field.AssignExpr) IDictDataDo {
	return d.withDO(d.DO.Attrs(attrs...))
}

func (d dictDataDo) Assign(attrs ...field.AssignExpr) IDictDataDo {
	return d.withDO(d.DO.Assign(attrs...))
}

func (d dictDataDo) Joins(fields ...field.RelationField) IDictDataDo {
	for _, _f := range fields {
		d = *d.withDO(d.DO.Joins(_f))
	}
	return &d
}

func (d dictDataDo) Preload(fields ...field.RelationField) IDictDataDo {
	for _, _f := range fields {
		d = *d.withDO(d.DO.Preload(_f))
	}
	return &d
}

func (d dictDataDo) FirstOrInit() (*model.DictData, error) {
	if result, err := d.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.DictData), nil
	}
}

func (d dictDataDo) FirstOrCreate() (*model.DictData, error) {
	if result, err := d.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.DictData), nil
	}
}

func (d dictDataDo) FindByPage(offset int, limit int) (result []*model.DictData, count int64, err error) {
	result, err = d.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = d.Offset(-1).Limit(-1).Count()
	return
}

func (d dictDataDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = d.Count()
	if err != nil {
		return
	}

	err = d.Offset(offset).Limit(limit).Scan(result)
	return
}

func (d dictDataDo) Scan(result interface{}) (err error) {
	return d.DO.Scan(result)
}

func (d dictDataDo) Delete(models ...*model.DictData) (result gen.ResultInfo, err error) {
	return d.DO.Delete(models)
}

func (d *dictDataDo) withDO(do gen.Dao) *dictDataDo {
	d.DO = *do.(*gen.DO)
	return d
}

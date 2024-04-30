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

func newDictType(db *gorm.DB, opts ...gen.DOOption) dictType {
	_dictType := dictType{}

	_dictType.dictTypeDo.UseDB(db, opts...)
	_dictType.dictTypeDo.UseModel(&model.DictType{})

	tableName := _dictType.dictTypeDo.TableName()
	_dictType.ALL = field.NewAsterisk(tableName)
	_dictType.DictID = field.NewUint(tableName, "dict_id")
	_dictType.DictName = field.NewString(tableName, "dict_name")
	_dictType.DictType = field.NewString(tableName, "dict_type")
	_dictType.Status = field.NewInt32(tableName, "status")
	_dictType.CreateBy = field.NewInt32(tableName, "create_by")
	_dictType.UpdateBy = field.NewInt32(tableName, "update_by")
	_dictType.Remark = field.NewString(tableName, "remark")
	_dictType.CreatedAt = field.NewTime(tableName, "created_at")
	_dictType.UpdatedAt = field.NewTime(tableName, "updated_at")
	_dictType.DictData = dictTypeHasManyDictData{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("DictData", "model.DictData"),
	}

	_dictType.fillFieldMap()

	return _dictType
}

type dictType struct {
	dictTypeDo

	ALL       field.Asterisk
	DictID    field.Uint
	DictName  field.String
	DictType  field.String
	Status    field.Int32
	CreateBy  field.Int32
	UpdateBy  field.Int32
	Remark    field.String
	CreatedAt field.Time
	UpdatedAt field.Time
	DictData  dictTypeHasManyDictData

	fieldMap map[string]field.Expr
}

func (d dictType) Table(newTableName string) *dictType {
	d.dictTypeDo.UseTable(newTableName)
	return d.updateTableName(newTableName)
}

func (d dictType) As(alias string) *dictType {
	d.dictTypeDo.DO = *(d.dictTypeDo.As(alias).(*gen.DO))
	return d.updateTableName(alias)
}

func (d *dictType) updateTableName(table string) *dictType {
	d.ALL = field.NewAsterisk(table)
	d.DictID = field.NewUint(table, "dict_id")
	d.DictName = field.NewString(table, "dict_name")
	d.DictType = field.NewString(table, "dict_type")
	d.Status = field.NewInt32(table, "status")
	d.CreateBy = field.NewInt32(table, "create_by")
	d.UpdateBy = field.NewInt32(table, "update_by")
	d.Remark = field.NewString(table, "remark")
	d.CreatedAt = field.NewTime(table, "created_at")
	d.UpdatedAt = field.NewTime(table, "updated_at")

	d.fillFieldMap()

	return d
}

func (d *dictType) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := d.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (d *dictType) fillFieldMap() {
	d.fieldMap = make(map[string]field.Expr, 10)
	d.fieldMap["dict_id"] = d.DictID
	d.fieldMap["dict_name"] = d.DictName
	d.fieldMap["dict_type"] = d.DictType
	d.fieldMap["status"] = d.Status
	d.fieldMap["create_by"] = d.CreateBy
	d.fieldMap["update_by"] = d.UpdateBy
	d.fieldMap["remark"] = d.Remark
	d.fieldMap["created_at"] = d.CreatedAt
	d.fieldMap["updated_at"] = d.UpdatedAt

}

func (d dictType) clone(db *gorm.DB) dictType {
	d.dictTypeDo.ReplaceConnPool(db.Statement.ConnPool)
	return d
}

func (d dictType) replaceDB(db *gorm.DB) dictType {
	d.dictTypeDo.ReplaceDB(db)
	return d
}

type dictTypeHasManyDictData struct {
	db *gorm.DB

	field.RelationField
}

func (a dictTypeHasManyDictData) Where(conds ...field.Expr) *dictTypeHasManyDictData {
	if len(conds) == 0 {
		return &a
	}

	exprs := make([]clause.Expression, 0, len(conds))
	for _, cond := range conds {
		exprs = append(exprs, cond.BeCond().(clause.Expression))
	}
	a.db = a.db.Clauses(clause.Where{Exprs: exprs})
	return &a
}

func (a dictTypeHasManyDictData) WithContext(ctx context.Context) *dictTypeHasManyDictData {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a dictTypeHasManyDictData) Session(session *gorm.Session) *dictTypeHasManyDictData {
	a.db = a.db.Session(session)
	return &a
}

func (a dictTypeHasManyDictData) Model(m *model.DictType) *dictTypeHasManyDictDataTx {
	return &dictTypeHasManyDictDataTx{a.db.Model(m).Association(a.Name())}
}

type dictTypeHasManyDictDataTx struct{ tx *gorm.Association }

func (a dictTypeHasManyDictDataTx) Find() (result []*model.DictData, err error) {
	return result, a.tx.Find(&result)
}

func (a dictTypeHasManyDictDataTx) Append(values ...*model.DictData) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a dictTypeHasManyDictDataTx) Replace(values ...*model.DictData) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a dictTypeHasManyDictDataTx) Delete(values ...*model.DictData) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a dictTypeHasManyDictDataTx) Clear() error {
	return a.tx.Clear()
}

func (a dictTypeHasManyDictDataTx) Count() int64 {
	return a.tx.Count()
}

type dictTypeDo struct{ gen.DO }

type IDictTypeDo interface {
	gen.SubQuery
	Debug() IDictTypeDo
	WithContext(ctx context.Context) IDictTypeDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IDictTypeDo
	WriteDB() IDictTypeDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IDictTypeDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IDictTypeDo
	Not(conds ...gen.Condition) IDictTypeDo
	Or(conds ...gen.Condition) IDictTypeDo
	Select(conds ...field.Expr) IDictTypeDo
	Where(conds ...gen.Condition) IDictTypeDo
	Order(conds ...field.Expr) IDictTypeDo
	Distinct(cols ...field.Expr) IDictTypeDo
	Omit(cols ...field.Expr) IDictTypeDo
	Join(table schema.Tabler, on ...field.Expr) IDictTypeDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IDictTypeDo
	RightJoin(table schema.Tabler, on ...field.Expr) IDictTypeDo
	Group(cols ...field.Expr) IDictTypeDo
	Having(conds ...gen.Condition) IDictTypeDo
	Limit(limit int) IDictTypeDo
	Offset(offset int) IDictTypeDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IDictTypeDo
	Unscoped() IDictTypeDo
	Create(values ...*model.DictType) error
	CreateInBatches(values []*model.DictType, batchSize int) error
	Save(values ...*model.DictType) error
	First() (*model.DictType, error)
	Take() (*model.DictType, error)
	Last() (*model.DictType, error)
	Find() ([]*model.DictType, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.DictType, err error)
	FindInBatches(result *[]*model.DictType, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.DictType) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IDictTypeDo
	Assign(attrs ...field.AssignExpr) IDictTypeDo
	Joins(fields ...field.RelationField) IDictTypeDo
	Preload(fields ...field.RelationField) IDictTypeDo
	FirstOrInit() (*model.DictType, error)
	FirstOrCreate() (*model.DictType, error)
	FindByPage(offset int, limit int) (result []*model.DictType, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IDictTypeDo
	UnderlyingDB() *gorm.DB
	schema.Tabler

	GetByID(id uint) (result model.DictType, err error)
	DeleteByID(id uint) (rowsAffected int64, err error)
	DeleteByIDs(ids []uint) (rowsAffected int64, err error)
}

// SELECT * FROM @@table WHERE id=@id
func (d dictTypeDo) GetByID(id uint) (result model.DictType, err error) {
	var params []interface{}

	var generateSQL strings.Builder
	params = append(params, id)
	generateSQL.WriteString("SELECT * FROM sys_dict_type WHERE id=? ")

	var executeSQL *gorm.DB
	executeSQL = d.UnderlyingDB().Raw(generateSQL.String(), params...).Take(&result) // ignore_security_alert
	err = executeSQL.Error

	return
}

// DELETE FROM @@table WHERE id=@id
func (d dictTypeDo) DeleteByID(id uint) (rowsAffected int64, err error) {
	var params []interface{}

	var generateSQL strings.Builder
	params = append(params, id)
	generateSQL.WriteString("DELETE FROM sys_dict_type WHERE id=? ")

	var executeSQL *gorm.DB
	executeSQL = d.UnderlyingDB().Exec(generateSQL.String(), params...) // ignore_security_alert
	rowsAffected = executeSQL.RowsAffected
	err = executeSQL.Error

	return
}

// DELETE FROM @@table WHERE id IN (@ids)
func (d dictTypeDo) DeleteByIDs(ids []uint) (rowsAffected int64, err error) {
	var params []interface{}

	var generateSQL strings.Builder
	params = append(params, ids)
	generateSQL.WriteString("DELETE FROM sys_dict_type WHERE id IN (?) ")

	var executeSQL *gorm.DB
	executeSQL = d.UnderlyingDB().Exec(generateSQL.String(), params...) // ignore_security_alert
	rowsAffected = executeSQL.RowsAffected
	err = executeSQL.Error

	return
}

func (d dictTypeDo) Debug() IDictTypeDo {
	return d.withDO(d.DO.Debug())
}

func (d dictTypeDo) WithContext(ctx context.Context) IDictTypeDo {
	return d.withDO(d.DO.WithContext(ctx))
}

func (d dictTypeDo) ReadDB() IDictTypeDo {
	return d.Clauses(dbresolver.Read)
}

func (d dictTypeDo) WriteDB() IDictTypeDo {
	return d.Clauses(dbresolver.Write)
}

func (d dictTypeDo) Session(config *gorm.Session) IDictTypeDo {
	return d.withDO(d.DO.Session(config))
}

func (d dictTypeDo) Clauses(conds ...clause.Expression) IDictTypeDo {
	return d.withDO(d.DO.Clauses(conds...))
}

func (d dictTypeDo) Returning(value interface{}, columns ...string) IDictTypeDo {
	return d.withDO(d.DO.Returning(value, columns...))
}

func (d dictTypeDo) Not(conds ...gen.Condition) IDictTypeDo {
	return d.withDO(d.DO.Not(conds...))
}

func (d dictTypeDo) Or(conds ...gen.Condition) IDictTypeDo {
	return d.withDO(d.DO.Or(conds...))
}

func (d dictTypeDo) Select(conds ...field.Expr) IDictTypeDo {
	return d.withDO(d.DO.Select(conds...))
}

func (d dictTypeDo) Where(conds ...gen.Condition) IDictTypeDo {
	return d.withDO(d.DO.Where(conds...))
}

func (d dictTypeDo) Order(conds ...field.Expr) IDictTypeDo {
	return d.withDO(d.DO.Order(conds...))
}

func (d dictTypeDo) Distinct(cols ...field.Expr) IDictTypeDo {
	return d.withDO(d.DO.Distinct(cols...))
}

func (d dictTypeDo) Omit(cols ...field.Expr) IDictTypeDo {
	return d.withDO(d.DO.Omit(cols...))
}

func (d dictTypeDo) Join(table schema.Tabler, on ...field.Expr) IDictTypeDo {
	return d.withDO(d.DO.Join(table, on...))
}

func (d dictTypeDo) LeftJoin(table schema.Tabler, on ...field.Expr) IDictTypeDo {
	return d.withDO(d.DO.LeftJoin(table, on...))
}

func (d dictTypeDo) RightJoin(table schema.Tabler, on ...field.Expr) IDictTypeDo {
	return d.withDO(d.DO.RightJoin(table, on...))
}

func (d dictTypeDo) Group(cols ...field.Expr) IDictTypeDo {
	return d.withDO(d.DO.Group(cols...))
}

func (d dictTypeDo) Having(conds ...gen.Condition) IDictTypeDo {
	return d.withDO(d.DO.Having(conds...))
}

func (d dictTypeDo) Limit(limit int) IDictTypeDo {
	return d.withDO(d.DO.Limit(limit))
}

func (d dictTypeDo) Offset(offset int) IDictTypeDo {
	return d.withDO(d.DO.Offset(offset))
}

func (d dictTypeDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IDictTypeDo {
	return d.withDO(d.DO.Scopes(funcs...))
}

func (d dictTypeDo) Unscoped() IDictTypeDo {
	return d.withDO(d.DO.Unscoped())
}

func (d dictTypeDo) Create(values ...*model.DictType) error {
	if len(values) == 0 {
		return nil
	}
	return d.DO.Create(values)
}

func (d dictTypeDo) CreateInBatches(values []*model.DictType, batchSize int) error {
	return d.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (d dictTypeDo) Save(values ...*model.DictType) error {
	if len(values) == 0 {
		return nil
	}
	return d.DO.Save(values)
}

func (d dictTypeDo) First() (*model.DictType, error) {
	if result, err := d.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.DictType), nil
	}
}

func (d dictTypeDo) Take() (*model.DictType, error) {
	if result, err := d.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.DictType), nil
	}
}

func (d dictTypeDo) Last() (*model.DictType, error) {
	if result, err := d.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.DictType), nil
	}
}

func (d dictTypeDo) Find() ([]*model.DictType, error) {
	result, err := d.DO.Find()
	return result.([]*model.DictType), err
}

func (d dictTypeDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.DictType, err error) {
	buf := make([]*model.DictType, 0, batchSize)
	err = d.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (d dictTypeDo) FindInBatches(result *[]*model.DictType, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return d.DO.FindInBatches(result, batchSize, fc)
}

func (d dictTypeDo) Attrs(attrs ...field.AssignExpr) IDictTypeDo {
	return d.withDO(d.DO.Attrs(attrs...))
}

func (d dictTypeDo) Assign(attrs ...field.AssignExpr) IDictTypeDo {
	return d.withDO(d.DO.Assign(attrs...))
}

func (d dictTypeDo) Joins(fields ...field.RelationField) IDictTypeDo {
	for _, _f := range fields {
		d = *d.withDO(d.DO.Joins(_f))
	}
	return &d
}

func (d dictTypeDo) Preload(fields ...field.RelationField) IDictTypeDo {
	for _, _f := range fields {
		d = *d.withDO(d.DO.Preload(_f))
	}
	return &d
}

func (d dictTypeDo) FirstOrInit() (*model.DictType, error) {
	if result, err := d.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.DictType), nil
	}
}

func (d dictTypeDo) FirstOrCreate() (*model.DictType, error) {
	if result, err := d.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.DictType), nil
	}
}

func (d dictTypeDo) FindByPage(offset int, limit int) (result []*model.DictType, count int64, err error) {
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

func (d dictTypeDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = d.Count()
	if err != nil {
		return
	}

	err = d.Offset(offset).Limit(limit).Scan(result)
	return
}

func (d dictTypeDo) Scan(result interface{}) (err error) {
	return d.DO.Scan(result)
}

func (d dictTypeDo) Delete(models ...*model.DictType) (result gen.ResultInfo, err error) {
	return d.DO.Delete(models)
}

func (d *dictTypeDo) withDO(do gen.Dao) *dictTypeDo {
	d.DO = *do.(*gen.DO)
	return d
}

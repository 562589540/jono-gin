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

func newOperLog(db *gorm.DB, opts ...gen.DOOption) operLog {
	_operLog := operLog{}

	_operLog.operLogDo.UseDB(db, opts...)
	_operLog.operLogDo.UseModel(&model.OperLog{})

	tableName := _operLog.operLogDo.TableName()
	_operLog.ALL = field.NewAsterisk(tableName)
	_operLog.ID = field.NewUint(tableName, "id")
	_operLog.UserName = field.NewString(tableName, "user_name")
	_operLog.Address = field.NewString(tableName, "address")
	_operLog.Browser = field.NewString(tableName, "browser")
	_operLog.Ip = field.NewString(tableName, "ip")
	_operLog.Module = field.NewString(tableName, "module")
	_operLog.Summary = field.NewString(tableName, "summary")
	_operLog.System = field.NewString(tableName, "system")
	_operLog.Status = field.NewInt(tableName, "status")
	_operLog.CreatedAt = field.NewTime(tableName, "created_at")

	_operLog.fillFieldMap()

	return _operLog
}

type operLog struct {
	operLogDo

	ALL       field.Asterisk
	ID        field.Uint
	UserName  field.String
	Address   field.String
	Browser   field.String
	Ip        field.String
	Module    field.String
	Summary   field.String
	System    field.String
	Status    field.Int
	CreatedAt field.Time

	fieldMap map[string]field.Expr
}

func (o operLog) Table(newTableName string) *operLog {
	o.operLogDo.UseTable(newTableName)
	return o.updateTableName(newTableName)
}

func (o operLog) As(alias string) *operLog {
	o.operLogDo.DO = *(o.operLogDo.As(alias).(*gen.DO))
	return o.updateTableName(alias)
}

func (o *operLog) updateTableName(table string) *operLog {
	o.ALL = field.NewAsterisk(table)
	o.ID = field.NewUint(table, "id")
	o.UserName = field.NewString(table, "user_name")
	o.Address = field.NewString(table, "address")
	o.Browser = field.NewString(table, "browser")
	o.Ip = field.NewString(table, "ip")
	o.Module = field.NewString(table, "module")
	o.Summary = field.NewString(table, "summary")
	o.System = field.NewString(table, "system")
	o.Status = field.NewInt(table, "status")
	o.CreatedAt = field.NewTime(table, "created_at")

	o.fillFieldMap()

	return o
}

func (o *operLog) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := o.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (o *operLog) fillFieldMap() {
	o.fieldMap = make(map[string]field.Expr, 10)
	o.fieldMap["id"] = o.ID
	o.fieldMap["user_name"] = o.UserName
	o.fieldMap["address"] = o.Address
	o.fieldMap["browser"] = o.Browser
	o.fieldMap["ip"] = o.Ip
	o.fieldMap["module"] = o.Module
	o.fieldMap["summary"] = o.Summary
	o.fieldMap["system"] = o.System
	o.fieldMap["status"] = o.Status
	o.fieldMap["created_at"] = o.CreatedAt
}

func (o operLog) clone(db *gorm.DB) operLog {
	o.operLogDo.ReplaceConnPool(db.Statement.ConnPool)
	return o
}

func (o operLog) replaceDB(db *gorm.DB) operLog {
	o.operLogDo.ReplaceDB(db)
	return o
}

type operLogDo struct{ gen.DO }

type IOperLogDo interface {
	gen.SubQuery
	Debug() IOperLogDo
	WithContext(ctx context.Context) IOperLogDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IOperLogDo
	WriteDB() IOperLogDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IOperLogDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IOperLogDo
	Not(conds ...gen.Condition) IOperLogDo
	Or(conds ...gen.Condition) IOperLogDo
	Select(conds ...field.Expr) IOperLogDo
	Where(conds ...gen.Condition) IOperLogDo
	Order(conds ...field.Expr) IOperLogDo
	Distinct(cols ...field.Expr) IOperLogDo
	Omit(cols ...field.Expr) IOperLogDo
	Join(table schema.Tabler, on ...field.Expr) IOperLogDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IOperLogDo
	RightJoin(table schema.Tabler, on ...field.Expr) IOperLogDo
	Group(cols ...field.Expr) IOperLogDo
	Having(conds ...gen.Condition) IOperLogDo
	Limit(limit int) IOperLogDo
	Offset(offset int) IOperLogDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IOperLogDo
	Unscoped() IOperLogDo
	Create(values ...*model.OperLog) error
	CreateInBatches(values []*model.OperLog, batchSize int) error
	Save(values ...*model.OperLog) error
	First() (*model.OperLog, error)
	Take() (*model.OperLog, error)
	Last() (*model.OperLog, error)
	Find() ([]*model.OperLog, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.OperLog, err error)
	FindInBatches(result *[]*model.OperLog, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.OperLog) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IOperLogDo
	Assign(attrs ...field.AssignExpr) IOperLogDo
	Joins(fields ...field.RelationField) IOperLogDo
	Preload(fields ...field.RelationField) IOperLogDo
	FirstOrInit() (*model.OperLog, error)
	FirstOrCreate() (*model.OperLog, error)
	FindByPage(offset int, limit int) (result []*model.OperLog, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IOperLogDo
	UnderlyingDB() *gorm.DB
	schema.Tabler

	GetByID(id uint) (result model.OperLog, err error)
	DeleteByID(id uint) (rowsAffected int64, err error)
	DeleteByIDs(ids []uint) (rowsAffected int64, err error)
}

// SELECT * FROM @@table WHERE id=@id
func (o operLogDo) GetByID(id uint) (result model.OperLog, err error) {
	var params []interface{}

	var generateSQL strings.Builder
	params = append(params, id)
	generateSQL.WriteString("SELECT * FROM sys_oper_log WHERE id=? ")

	var executeSQL *gorm.DB
	executeSQL = o.UnderlyingDB().Raw(generateSQL.String(), params...).Take(&result) // ignore_security_alert
	err = executeSQL.Error

	return
}

// DELETE FROM @@table WHERE id=@id
func (o operLogDo) DeleteByID(id uint) (rowsAffected int64, err error) {
	var params []interface{}

	var generateSQL strings.Builder
	params = append(params, id)
	generateSQL.WriteString("DELETE FROM sys_oper_log WHERE id=? ")

	var executeSQL *gorm.DB
	executeSQL = o.UnderlyingDB().Exec(generateSQL.String(), params...) // ignore_security_alert
	rowsAffected = executeSQL.RowsAffected
	err = executeSQL.Error

	return
}

// DELETE FROM @@table WHERE id IN (@ids)
func (o operLogDo) DeleteByIDs(ids []uint) (rowsAffected int64, err error) {
	var params []interface{}

	var generateSQL strings.Builder
	params = append(params, ids)
	generateSQL.WriteString("DELETE FROM sys_oper_log WHERE id IN (?) ")

	var executeSQL *gorm.DB
	executeSQL = o.UnderlyingDB().Exec(generateSQL.String(), params...) // ignore_security_alert
	rowsAffected = executeSQL.RowsAffected
	err = executeSQL.Error

	return
}

func (o operLogDo) Debug() IOperLogDo {
	return o.withDO(o.DO.Debug())
}

func (o operLogDo) WithContext(ctx context.Context) IOperLogDo {
	return o.withDO(o.DO.WithContext(ctx))
}

func (o operLogDo) ReadDB() IOperLogDo {
	return o.Clauses(dbresolver.Read)
}

func (o operLogDo) WriteDB() IOperLogDo {
	return o.Clauses(dbresolver.Write)
}

func (o operLogDo) Session(config *gorm.Session) IOperLogDo {
	return o.withDO(o.DO.Session(config))
}

func (o operLogDo) Clauses(conds ...clause.Expression) IOperLogDo {
	return o.withDO(o.DO.Clauses(conds...))
}

func (o operLogDo) Returning(value interface{}, columns ...string) IOperLogDo {
	return o.withDO(o.DO.Returning(value, columns...))
}

func (o operLogDo) Not(conds ...gen.Condition) IOperLogDo {
	return o.withDO(o.DO.Not(conds...))
}

func (o operLogDo) Or(conds ...gen.Condition) IOperLogDo {
	return o.withDO(o.DO.Or(conds...))
}

func (o operLogDo) Select(conds ...field.Expr) IOperLogDo {
	return o.withDO(o.DO.Select(conds...))
}

func (o operLogDo) Where(conds ...gen.Condition) IOperLogDo {
	return o.withDO(o.DO.Where(conds...))
}

func (o operLogDo) Order(conds ...field.Expr) IOperLogDo {
	return o.withDO(o.DO.Order(conds...))
}

func (o operLogDo) Distinct(cols ...field.Expr) IOperLogDo {
	return o.withDO(o.DO.Distinct(cols...))
}

func (o operLogDo) Omit(cols ...field.Expr) IOperLogDo {
	return o.withDO(o.DO.Omit(cols...))
}

func (o operLogDo) Join(table schema.Tabler, on ...field.Expr) IOperLogDo {
	return o.withDO(o.DO.Join(table, on...))
}

func (o operLogDo) LeftJoin(table schema.Tabler, on ...field.Expr) IOperLogDo {
	return o.withDO(o.DO.LeftJoin(table, on...))
}

func (o operLogDo) RightJoin(table schema.Tabler, on ...field.Expr) IOperLogDo {
	return o.withDO(o.DO.RightJoin(table, on...))
}

func (o operLogDo) Group(cols ...field.Expr) IOperLogDo {
	return o.withDO(o.DO.Group(cols...))
}

func (o operLogDo) Having(conds ...gen.Condition) IOperLogDo {
	return o.withDO(o.DO.Having(conds...))
}

func (o operLogDo) Limit(limit int) IOperLogDo {
	return o.withDO(o.DO.Limit(limit))
}

func (o operLogDo) Offset(offset int) IOperLogDo {
	return o.withDO(o.DO.Offset(offset))
}

func (o operLogDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IOperLogDo {
	return o.withDO(o.DO.Scopes(funcs...))
}

func (o operLogDo) Unscoped() IOperLogDo {
	return o.withDO(o.DO.Unscoped())
}

func (o operLogDo) Create(values ...*model.OperLog) error {
	if len(values) == 0 {
		return nil
	}
	return o.DO.Create(values)
}

func (o operLogDo) CreateInBatches(values []*model.OperLog, batchSize int) error {
	return o.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (o operLogDo) Save(values ...*model.OperLog) error {
	if len(values) == 0 {
		return nil
	}
	return o.DO.Save(values)
}

func (o operLogDo) First() (*model.OperLog, error) {
	if result, err := o.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.OperLog), nil
	}
}

func (o operLogDo) Take() (*model.OperLog, error) {
	if result, err := o.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.OperLog), nil
	}
}

func (o operLogDo) Last() (*model.OperLog, error) {
	if result, err := o.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.OperLog), nil
	}
}

func (o operLogDo) Find() ([]*model.OperLog, error) {
	result, err := o.DO.Find()
	return result.([]*model.OperLog), err
}

func (o operLogDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.OperLog, err error) {
	buf := make([]*model.OperLog, 0, batchSize)
	err = o.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (o operLogDo) FindInBatches(result *[]*model.OperLog, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return o.DO.FindInBatches(result, batchSize, fc)
}

func (o operLogDo) Attrs(attrs ...field.AssignExpr) IOperLogDo {
	return o.withDO(o.DO.Attrs(attrs...))
}

func (o operLogDo) Assign(attrs ...field.AssignExpr) IOperLogDo {
	return o.withDO(o.DO.Assign(attrs...))
}

func (o operLogDo) Joins(fields ...field.RelationField) IOperLogDo {
	for _, _f := range fields {
		o = *o.withDO(o.DO.Joins(_f))
	}
	return &o
}

func (o operLogDo) Preload(fields ...field.RelationField) IOperLogDo {
	for _, _f := range fields {
		o = *o.withDO(o.DO.Preload(_f))
	}
	return &o
}

func (o operLogDo) FirstOrInit() (*model.OperLog, error) {
	if result, err := o.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.OperLog), nil
	}
}

func (o operLogDo) FirstOrCreate() (*model.OperLog, error) {
	if result, err := o.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.OperLog), nil
	}
}

func (o operLogDo) FindByPage(offset int, limit int) (result []*model.OperLog, count int64, err error) {
	result, err = o.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = o.Offset(-1).Limit(-1).Count()
	return
}

func (o operLogDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = o.Count()
	if err != nil {
		return
	}

	err = o.Offset(offset).Limit(limit).Scan(result)
	return
}

func (o operLogDo) Scan(result interface{}) (err error) {
	return o.DO.Scan(result)
}

func (o operLogDo) Delete(models ...*model.OperLog) (result gen.ResultInfo, err error) {
	return o.DO.Delete(models)
}

func (o *operLogDo) withDO(do gen.Dao) *operLogDo {
	o.DO = *do.(*gen.DO)
	return o
}

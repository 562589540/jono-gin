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

func newTaskLog(db *gorm.DB, opts ...gen.DOOption) taskLog {
	_taskLog := taskLog{}

	_taskLog.taskLogDo.UseDB(db, opts...)
	_taskLog.taskLogDo.UseModel(&model.TaskLog{})

	tableName := _taskLog.taskLogDo.TableName()
	_taskLog.ALL = field.NewAsterisk(tableName)
	_taskLog.ID = field.NewUint(tableName, "id")
	_taskLog.JobId = field.NewUint(tableName, "job_id")
	_taskLog.JobFunc = field.NewString(tableName, "job_func")
	_taskLog.ErrorStr = field.NewString(tableName, "error_str")
	_taskLog.Status = field.NewInt(tableName, "status")
	_taskLog.CreatedAt = field.NewTime(tableName, "created_at")

	_taskLog.fillFieldMap()

	return _taskLog
}

type taskLog struct {
	taskLogDo

	ALL       field.Asterisk
	ID        field.Uint
	JobId     field.Uint
	JobFunc   field.String
	ErrorStr  field.String
	Status    field.Int
	CreatedAt field.Time

	fieldMap map[string]field.Expr
}

func (t taskLog) Table(newTableName string) *taskLog {
	t.taskLogDo.UseTable(newTableName)
	return t.updateTableName(newTableName)
}

func (t taskLog) As(alias string) *taskLog {
	t.taskLogDo.DO = *(t.taskLogDo.As(alias).(*gen.DO))
	return t.updateTableName(alias)
}

func (t *taskLog) updateTableName(table string) *taskLog {
	t.ALL = field.NewAsterisk(table)
	t.ID = field.NewUint(table, "id")
	t.JobId = field.NewUint(table, "job_id")
	t.JobFunc = field.NewString(table, "job_func")
	t.ErrorStr = field.NewString(table, "error_str")
	t.Status = field.NewInt(table, "status")
	t.CreatedAt = field.NewTime(table, "created_at")

	t.fillFieldMap()

	return t
}

func (t *taskLog) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := t.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (t *taskLog) fillFieldMap() {
	t.fieldMap = make(map[string]field.Expr, 6)
	t.fieldMap["id"] = t.ID
	t.fieldMap["job_id"] = t.JobId
	t.fieldMap["job_func"] = t.JobFunc
	t.fieldMap["error_str"] = t.ErrorStr
	t.fieldMap["status"] = t.Status
	t.fieldMap["created_at"] = t.CreatedAt
}

func (t taskLog) clone(db *gorm.DB) taskLog {
	t.taskLogDo.ReplaceConnPool(db.Statement.ConnPool)
	return t
}

func (t taskLog) replaceDB(db *gorm.DB) taskLog {
	t.taskLogDo.ReplaceDB(db)
	return t
}

type taskLogDo struct{ gen.DO }

type ITaskLogDo interface {
	gen.SubQuery
	Debug() ITaskLogDo
	WithContext(ctx context.Context) ITaskLogDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ITaskLogDo
	WriteDB() ITaskLogDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ITaskLogDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ITaskLogDo
	Not(conds ...gen.Condition) ITaskLogDo
	Or(conds ...gen.Condition) ITaskLogDo
	Select(conds ...field.Expr) ITaskLogDo
	Where(conds ...gen.Condition) ITaskLogDo
	Order(conds ...field.Expr) ITaskLogDo
	Distinct(cols ...field.Expr) ITaskLogDo
	Omit(cols ...field.Expr) ITaskLogDo
	Join(table schema.Tabler, on ...field.Expr) ITaskLogDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ITaskLogDo
	RightJoin(table schema.Tabler, on ...field.Expr) ITaskLogDo
	Group(cols ...field.Expr) ITaskLogDo
	Having(conds ...gen.Condition) ITaskLogDo
	Limit(limit int) ITaskLogDo
	Offset(offset int) ITaskLogDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ITaskLogDo
	Unscoped() ITaskLogDo
	Create(values ...*model.TaskLog) error
	CreateInBatches(values []*model.TaskLog, batchSize int) error
	Save(values ...*model.TaskLog) error
	First() (*model.TaskLog, error)
	Take() (*model.TaskLog, error)
	Last() (*model.TaskLog, error)
	Find() ([]*model.TaskLog, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.TaskLog, err error)
	FindInBatches(result *[]*model.TaskLog, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.TaskLog) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ITaskLogDo
	Assign(attrs ...field.AssignExpr) ITaskLogDo
	Joins(fields ...field.RelationField) ITaskLogDo
	Preload(fields ...field.RelationField) ITaskLogDo
	FirstOrInit() (*model.TaskLog, error)
	FirstOrCreate() (*model.TaskLog, error)
	FindByPage(offset int, limit int) (result []*model.TaskLog, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ITaskLogDo
	UnderlyingDB() *gorm.DB
	schema.Tabler

	GetByID(id uint) (result model.TaskLog, err error)
	DeleteByID(id uint) (rowsAffected int64, err error)
	DeleteByIDs(ids []uint) (rowsAffected int64, err error)
}

// SELECT * FROM @@table WHERE id=@id
func (t taskLogDo) GetByID(id uint) (result model.TaskLog, err error) {
	var params []interface{}

	var generateSQL strings.Builder
	params = append(params, id)
	generateSQL.WriteString("SELECT * FROM sys_task_log WHERE id=? ")

	var executeSQL *gorm.DB
	executeSQL = t.UnderlyingDB().Raw(generateSQL.String(), params...).Take(&result) // ignore_security_alert
	err = executeSQL.Error

	return
}

// DELETE FROM @@table WHERE id=@id
func (t taskLogDo) DeleteByID(id uint) (rowsAffected int64, err error) {
	var params []interface{}

	var generateSQL strings.Builder
	params = append(params, id)
	generateSQL.WriteString("DELETE FROM sys_task_log WHERE id=? ")

	var executeSQL *gorm.DB
	executeSQL = t.UnderlyingDB().Exec(generateSQL.String(), params...) // ignore_security_alert
	rowsAffected = executeSQL.RowsAffected
	err = executeSQL.Error

	return
}

// DELETE FROM @@table WHERE id IN (@ids)
func (t taskLogDo) DeleteByIDs(ids []uint) (rowsAffected int64, err error) {
	var params []interface{}

	var generateSQL strings.Builder
	params = append(params, ids)
	generateSQL.WriteString("DELETE FROM sys_task_log WHERE id IN (?) ")

	var executeSQL *gorm.DB
	executeSQL = t.UnderlyingDB().Exec(generateSQL.String(), params...) // ignore_security_alert
	rowsAffected = executeSQL.RowsAffected
	err = executeSQL.Error

	return
}

func (t taskLogDo) Debug() ITaskLogDo {
	return t.withDO(t.DO.Debug())
}

func (t taskLogDo) WithContext(ctx context.Context) ITaskLogDo {
	return t.withDO(t.DO.WithContext(ctx))
}

func (t taskLogDo) ReadDB() ITaskLogDo {
	return t.Clauses(dbresolver.Read)
}

func (t taskLogDo) WriteDB() ITaskLogDo {
	return t.Clauses(dbresolver.Write)
}

func (t taskLogDo) Session(config *gorm.Session) ITaskLogDo {
	return t.withDO(t.DO.Session(config))
}

func (t taskLogDo) Clauses(conds ...clause.Expression) ITaskLogDo {
	return t.withDO(t.DO.Clauses(conds...))
}

func (t taskLogDo) Returning(value interface{}, columns ...string) ITaskLogDo {
	return t.withDO(t.DO.Returning(value, columns...))
}

func (t taskLogDo) Not(conds ...gen.Condition) ITaskLogDo {
	return t.withDO(t.DO.Not(conds...))
}

func (t taskLogDo) Or(conds ...gen.Condition) ITaskLogDo {
	return t.withDO(t.DO.Or(conds...))
}

func (t taskLogDo) Select(conds ...field.Expr) ITaskLogDo {
	return t.withDO(t.DO.Select(conds...))
}

func (t taskLogDo) Where(conds ...gen.Condition) ITaskLogDo {
	return t.withDO(t.DO.Where(conds...))
}

func (t taskLogDo) Order(conds ...field.Expr) ITaskLogDo {
	return t.withDO(t.DO.Order(conds...))
}

func (t taskLogDo) Distinct(cols ...field.Expr) ITaskLogDo {
	return t.withDO(t.DO.Distinct(cols...))
}

func (t taskLogDo) Omit(cols ...field.Expr) ITaskLogDo {
	return t.withDO(t.DO.Omit(cols...))
}

func (t taskLogDo) Join(table schema.Tabler, on ...field.Expr) ITaskLogDo {
	return t.withDO(t.DO.Join(table, on...))
}

func (t taskLogDo) LeftJoin(table schema.Tabler, on ...field.Expr) ITaskLogDo {
	return t.withDO(t.DO.LeftJoin(table, on...))
}

func (t taskLogDo) RightJoin(table schema.Tabler, on ...field.Expr) ITaskLogDo {
	return t.withDO(t.DO.RightJoin(table, on...))
}

func (t taskLogDo) Group(cols ...field.Expr) ITaskLogDo {
	return t.withDO(t.DO.Group(cols...))
}

func (t taskLogDo) Having(conds ...gen.Condition) ITaskLogDo {
	return t.withDO(t.DO.Having(conds...))
}

func (t taskLogDo) Limit(limit int) ITaskLogDo {
	return t.withDO(t.DO.Limit(limit))
}

func (t taskLogDo) Offset(offset int) ITaskLogDo {
	return t.withDO(t.DO.Offset(offset))
}

func (t taskLogDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ITaskLogDo {
	return t.withDO(t.DO.Scopes(funcs...))
}

func (t taskLogDo) Unscoped() ITaskLogDo {
	return t.withDO(t.DO.Unscoped())
}

func (t taskLogDo) Create(values ...*model.TaskLog) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Create(values)
}

func (t taskLogDo) CreateInBatches(values []*model.TaskLog, batchSize int) error {
	return t.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (t taskLogDo) Save(values ...*model.TaskLog) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Save(values)
}

func (t taskLogDo) First() (*model.TaskLog, error) {
	if result, err := t.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.TaskLog), nil
	}
}

func (t taskLogDo) Take() (*model.TaskLog, error) {
	if result, err := t.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.TaskLog), nil
	}
}

func (t taskLogDo) Last() (*model.TaskLog, error) {
	if result, err := t.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.TaskLog), nil
	}
}

func (t taskLogDo) Find() ([]*model.TaskLog, error) {
	result, err := t.DO.Find()
	return result.([]*model.TaskLog), err
}

func (t taskLogDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.TaskLog, err error) {
	buf := make([]*model.TaskLog, 0, batchSize)
	err = t.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (t taskLogDo) FindInBatches(result *[]*model.TaskLog, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return t.DO.FindInBatches(result, batchSize, fc)
}

func (t taskLogDo) Attrs(attrs ...field.AssignExpr) ITaskLogDo {
	return t.withDO(t.DO.Attrs(attrs...))
}

func (t taskLogDo) Assign(attrs ...field.AssignExpr) ITaskLogDo {
	return t.withDO(t.DO.Assign(attrs...))
}

func (t taskLogDo) Joins(fields ...field.RelationField) ITaskLogDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Joins(_f))
	}
	return &t
}

func (t taskLogDo) Preload(fields ...field.RelationField) ITaskLogDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Preload(_f))
	}
	return &t
}

func (t taskLogDo) FirstOrInit() (*model.TaskLog, error) {
	if result, err := t.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.TaskLog), nil
	}
}

func (t taskLogDo) FirstOrCreate() (*model.TaskLog, error) {
	if result, err := t.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.TaskLog), nil
	}
}

func (t taskLogDo) FindByPage(offset int, limit int) (result []*model.TaskLog, count int64, err error) {
	result, err = t.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = t.Offset(-1).Limit(-1).Count()
	return
}

func (t taskLogDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = t.Count()
	if err != nil {
		return
	}

	err = t.Offset(offset).Limit(limit).Scan(result)
	return
}

func (t taskLogDo) Scan(result interface{}) (err error) {
	return t.DO.Scan(result)
}

func (t taskLogDo) Delete(models ...*model.TaskLog) (result gen.ResultInfo, err error) {
	return t.DO.Delete(models)
}

func (t *taskLogDo) withDO(do gen.Dao) *taskLogDo {
	t.DO = *do.(*gen.DO)
	return t
}

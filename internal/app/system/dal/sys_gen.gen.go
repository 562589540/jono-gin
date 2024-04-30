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

func newSysGen(db *gorm.DB, opts ...gen.DOOption) sysGen {
	_sysGen := sysGen{}

	_sysGen.sysGenDo.UseDB(db, opts...)
	_sysGen.sysGenDo.UseModel(&model.SysGen{})

	tableName := _sysGen.sysGenDo.TableName()
	_sysGen.ALL = field.NewAsterisk(tableName)
	_sysGen.ID = field.NewUint(tableName, "id")
	_sysGen.TableNamed = field.NewString(tableName, "table_named")
	_sysGen.TableComment = field.NewString(tableName, "table_comment")
	_sysGen.FieldsInfo = field.NewField(tableName, "fields_info")
	_sysGen.BaseInfo = field.NewField(tableName, "base_info")
	_sysGen.GenInfo = field.NewField(tableName, "gen_info")
	_sysGen.GoApiCode = field.NewString(tableName, "go_api_code")
	_sysGen.GoDtoCode = field.NewString(tableName, "go_dto_code")
	_sysGen.GoLogicCode = field.NewString(tableName, "go_logic_code")
	_sysGen.GoModelCode = field.NewString(tableName, "go_model_code")
	_sysGen.GoRouterCode = field.NewString(tableName, "go_router_code")
	_sysGen.GoServiceCode = field.NewString(tableName, "go_service_code")
	_sysGen.VueApiCode = field.NewString(tableName, "vue_api_code")
	_sysGen.VueFormCode = field.NewString(tableName, "vue_form_code")
	_sysGen.VueHookCode = field.NewString(tableName, "vue_hook_code")
	_sysGen.VueIndexCode = field.NewString(tableName, "vue_index_code")
	_sysGen.VueRuleCode = field.NewString(tableName, "vue_rule_code")
	_sysGen.VueTypesCode = field.NewString(tableName, "vue_types_code")
	_sysGen.CreatedAt = field.NewTime(tableName, "created_at")
	_sysGen.UpdatedAt = field.NewTime(tableName, "updated_at")

	_sysGen.fillFieldMap()

	return _sysGen
}

type sysGen struct {
	sysGenDo

	ALL           field.Asterisk
	ID            field.Uint
	TableNamed    field.String
	TableComment  field.String
	FieldsInfo    field.Field
	BaseInfo      field.Field
	GenInfo       field.Field
	GoApiCode     field.String
	GoDtoCode     field.String
	GoLogicCode   field.String
	GoModelCode   field.String
	GoRouterCode  field.String
	GoServiceCode field.String
	VueApiCode    field.String
	VueFormCode   field.String
	VueHookCode   field.String
	VueIndexCode  field.String
	VueRuleCode   field.String
	VueTypesCode  field.String
	CreatedAt     field.Time
	UpdatedAt     field.Time

	fieldMap map[string]field.Expr
}

func (s sysGen) Table(newTableName string) *sysGen {
	s.sysGenDo.UseTable(newTableName)
	return s.updateTableName(newTableName)
}

func (s sysGen) As(alias string) *sysGen {
	s.sysGenDo.DO = *(s.sysGenDo.As(alias).(*gen.DO))
	return s.updateTableName(alias)
}

func (s *sysGen) updateTableName(table string) *sysGen {
	s.ALL = field.NewAsterisk(table)
	s.ID = field.NewUint(table, "id")
	s.TableNamed = field.NewString(table, "table_named")
	s.TableComment = field.NewString(table, "table_comment")
	s.FieldsInfo = field.NewField(table, "fields_info")
	s.BaseInfo = field.NewField(table, "base_info")
	s.GenInfo = field.NewField(table, "gen_info")
	s.GoApiCode = field.NewString(table, "go_api_code")
	s.GoDtoCode = field.NewString(table, "go_dto_code")
	s.GoLogicCode = field.NewString(table, "go_logic_code")
	s.GoModelCode = field.NewString(table, "go_model_code")
	s.GoRouterCode = field.NewString(table, "go_router_code")
	s.GoServiceCode = field.NewString(table, "go_service_code")
	s.VueApiCode = field.NewString(table, "vue_api_code")
	s.VueFormCode = field.NewString(table, "vue_form_code")
	s.VueHookCode = field.NewString(table, "vue_hook_code")
	s.VueIndexCode = field.NewString(table, "vue_index_code")
	s.VueRuleCode = field.NewString(table, "vue_rule_code")
	s.VueTypesCode = field.NewString(table, "vue_types_code")
	s.CreatedAt = field.NewTime(table, "created_at")
	s.UpdatedAt = field.NewTime(table, "updated_at")

	s.fillFieldMap()

	return s
}

func (s *sysGen) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *sysGen) fillFieldMap() {
	s.fieldMap = make(map[string]field.Expr, 20)
	s.fieldMap["id"] = s.ID
	s.fieldMap["table_named"] = s.TableNamed
	s.fieldMap["table_comment"] = s.TableComment
	s.fieldMap["fields_info"] = s.FieldsInfo
	s.fieldMap["base_info"] = s.BaseInfo
	s.fieldMap["gen_info"] = s.GenInfo
	s.fieldMap["go_api_code"] = s.GoApiCode
	s.fieldMap["go_dto_code"] = s.GoDtoCode
	s.fieldMap["go_logic_code"] = s.GoLogicCode
	s.fieldMap["go_model_code"] = s.GoModelCode
	s.fieldMap["go_router_code"] = s.GoRouterCode
	s.fieldMap["go_service_code"] = s.GoServiceCode
	s.fieldMap["vue_api_code"] = s.VueApiCode
	s.fieldMap["vue_form_code"] = s.VueFormCode
	s.fieldMap["vue_hook_code"] = s.VueHookCode
	s.fieldMap["vue_index_code"] = s.VueIndexCode
	s.fieldMap["vue_rule_code"] = s.VueRuleCode
	s.fieldMap["vue_types_code"] = s.VueTypesCode
	s.fieldMap["created_at"] = s.CreatedAt
	s.fieldMap["updated_at"] = s.UpdatedAt
}

func (s sysGen) clone(db *gorm.DB) sysGen {
	s.sysGenDo.ReplaceConnPool(db.Statement.ConnPool)
	return s
}

func (s sysGen) replaceDB(db *gorm.DB) sysGen {
	s.sysGenDo.ReplaceDB(db)
	return s
}

type sysGenDo struct{ gen.DO }

type ISysGenDo interface {
	gen.SubQuery
	Debug() ISysGenDo
	WithContext(ctx context.Context) ISysGenDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ISysGenDo
	WriteDB() ISysGenDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ISysGenDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ISysGenDo
	Not(conds ...gen.Condition) ISysGenDo
	Or(conds ...gen.Condition) ISysGenDo
	Select(conds ...field.Expr) ISysGenDo
	Where(conds ...gen.Condition) ISysGenDo
	Order(conds ...field.Expr) ISysGenDo
	Distinct(cols ...field.Expr) ISysGenDo
	Omit(cols ...field.Expr) ISysGenDo
	Join(table schema.Tabler, on ...field.Expr) ISysGenDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ISysGenDo
	RightJoin(table schema.Tabler, on ...field.Expr) ISysGenDo
	Group(cols ...field.Expr) ISysGenDo
	Having(conds ...gen.Condition) ISysGenDo
	Limit(limit int) ISysGenDo
	Offset(offset int) ISysGenDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ISysGenDo
	Unscoped() ISysGenDo
	Create(values ...*model.SysGen) error
	CreateInBatches(values []*model.SysGen, batchSize int) error
	Save(values ...*model.SysGen) error
	First() (*model.SysGen, error)
	Take() (*model.SysGen, error)
	Last() (*model.SysGen, error)
	Find() ([]*model.SysGen, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.SysGen, err error)
	FindInBatches(result *[]*model.SysGen, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.SysGen) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ISysGenDo
	Assign(attrs ...field.AssignExpr) ISysGenDo
	Joins(fields ...field.RelationField) ISysGenDo
	Preload(fields ...field.RelationField) ISysGenDo
	FirstOrInit() (*model.SysGen, error)
	FirstOrCreate() (*model.SysGen, error)
	FindByPage(offset int, limit int) (result []*model.SysGen, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ISysGenDo
	UnderlyingDB() *gorm.DB
	schema.Tabler

	GetByID(id uint) (result model.SysGen, err error)
	DeleteByID(id uint) (rowsAffected int64, err error)
	DeleteByIDs(ids []uint) (rowsAffected int64, err error)
}

// SELECT * FROM @@table WHERE id=@id
func (s sysGenDo) GetByID(id uint) (result model.SysGen, err error) {
	var params []interface{}

	var generateSQL strings.Builder
	params = append(params, id)
	generateSQL.WriteString("SELECT * FROM sys_gen WHERE id=? ")

	var executeSQL *gorm.DB
	executeSQL = s.UnderlyingDB().Raw(generateSQL.String(), params...).Take(&result) // ignore_security_alert
	err = executeSQL.Error

	return
}

// DELETE FROM @@table WHERE id=@id
func (s sysGenDo) DeleteByID(id uint) (rowsAffected int64, err error) {
	var params []interface{}

	var generateSQL strings.Builder
	params = append(params, id)
	generateSQL.WriteString("DELETE FROM sys_gen WHERE id=? ")

	var executeSQL *gorm.DB
	executeSQL = s.UnderlyingDB().Exec(generateSQL.String(), params...) // ignore_security_alert
	rowsAffected = executeSQL.RowsAffected
	err = executeSQL.Error

	return
}

// DELETE FROM @@table WHERE id IN (@ids)
func (s sysGenDo) DeleteByIDs(ids []uint) (rowsAffected int64, err error) {
	var params []interface{}

	var generateSQL strings.Builder
	params = append(params, ids)
	generateSQL.WriteString("DELETE FROM sys_gen WHERE id IN (?) ")

	var executeSQL *gorm.DB
	executeSQL = s.UnderlyingDB().Exec(generateSQL.String(), params...) // ignore_security_alert
	rowsAffected = executeSQL.RowsAffected
	err = executeSQL.Error

	return
}

func (s sysGenDo) Debug() ISysGenDo {
	return s.withDO(s.DO.Debug())
}

func (s sysGenDo) WithContext(ctx context.Context) ISysGenDo {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s sysGenDo) ReadDB() ISysGenDo {
	return s.Clauses(dbresolver.Read)
}

func (s sysGenDo) WriteDB() ISysGenDo {
	return s.Clauses(dbresolver.Write)
}

func (s sysGenDo) Session(config *gorm.Session) ISysGenDo {
	return s.withDO(s.DO.Session(config))
}

func (s sysGenDo) Clauses(conds ...clause.Expression) ISysGenDo {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s sysGenDo) Returning(value interface{}, columns ...string) ISysGenDo {
	return s.withDO(s.DO.Returning(value, columns...))
}

func (s sysGenDo) Not(conds ...gen.Condition) ISysGenDo {
	return s.withDO(s.DO.Not(conds...))
}

func (s sysGenDo) Or(conds ...gen.Condition) ISysGenDo {
	return s.withDO(s.DO.Or(conds...))
}

func (s sysGenDo) Select(conds ...field.Expr) ISysGenDo {
	return s.withDO(s.DO.Select(conds...))
}

func (s sysGenDo) Where(conds ...gen.Condition) ISysGenDo {
	return s.withDO(s.DO.Where(conds...))
}

func (s sysGenDo) Order(conds ...field.Expr) ISysGenDo {
	return s.withDO(s.DO.Order(conds...))
}

func (s sysGenDo) Distinct(cols ...field.Expr) ISysGenDo {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s sysGenDo) Omit(cols ...field.Expr) ISysGenDo {
	return s.withDO(s.DO.Omit(cols...))
}

func (s sysGenDo) Join(table schema.Tabler, on ...field.Expr) ISysGenDo {
	return s.withDO(s.DO.Join(table, on...))
}

func (s sysGenDo) LeftJoin(table schema.Tabler, on ...field.Expr) ISysGenDo {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s sysGenDo) RightJoin(table schema.Tabler, on ...field.Expr) ISysGenDo {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s sysGenDo) Group(cols ...field.Expr) ISysGenDo {
	return s.withDO(s.DO.Group(cols...))
}

func (s sysGenDo) Having(conds ...gen.Condition) ISysGenDo {
	return s.withDO(s.DO.Having(conds...))
}

func (s sysGenDo) Limit(limit int) ISysGenDo {
	return s.withDO(s.DO.Limit(limit))
}

func (s sysGenDo) Offset(offset int) ISysGenDo {
	return s.withDO(s.DO.Offset(offset))
}

func (s sysGenDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ISysGenDo {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s sysGenDo) Unscoped() ISysGenDo {
	return s.withDO(s.DO.Unscoped())
}

func (s sysGenDo) Create(values ...*model.SysGen) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s sysGenDo) CreateInBatches(values []*model.SysGen, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s sysGenDo) Save(values ...*model.SysGen) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s sysGenDo) First() (*model.SysGen, error) {
	if result, err := s.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysGen), nil
	}
}

func (s sysGenDo) Take() (*model.SysGen, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysGen), nil
	}
}

func (s sysGenDo) Last() (*model.SysGen, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysGen), nil
	}
}

func (s sysGenDo) Find() ([]*model.SysGen, error) {
	result, err := s.DO.Find()
	return result.([]*model.SysGen), err
}

func (s sysGenDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.SysGen, err error) {
	buf := make([]*model.SysGen, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s sysGenDo) FindInBatches(result *[]*model.SysGen, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s sysGenDo) Attrs(attrs ...field.AssignExpr) ISysGenDo {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s sysGenDo) Assign(attrs ...field.AssignExpr) ISysGenDo {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s sysGenDo) Joins(fields ...field.RelationField) ISysGenDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Joins(_f))
	}
	return &s
}

func (s sysGenDo) Preload(fields ...field.RelationField) ISysGenDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Preload(_f))
	}
	return &s
}

func (s sysGenDo) FirstOrInit() (*model.SysGen, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysGen), nil
	}
}

func (s sysGenDo) FirstOrCreate() (*model.SysGen, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysGen), nil
	}
}

func (s sysGenDo) FindByPage(offset int, limit int) (result []*model.SysGen, count int64, err error) {
	result, err = s.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = s.Offset(-1).Limit(-1).Count()
	return
}

func (s sysGenDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s sysGenDo) Scan(result interface{}) (err error) {
	return s.DO.Scan(result)
}

func (s sysGenDo) Delete(models ...*model.SysGen) (result gen.ResultInfo, err error) {
	return s.DO.Delete(models)
}

func (s *sysGenDo) withDO(do gen.Dao) *sysGenDo {
	s.DO = *do.(*gen.DO)
	return s
}

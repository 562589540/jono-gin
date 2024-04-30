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

func newUserOnline(db *gorm.DB, opts ...gen.DOOption) userOnline {
	_userOnline := userOnline{}

	_userOnline.userOnlineDo.UseDB(db, opts...)
	_userOnline.userOnlineDo.UseModel(&model.UserOnline{})

	tableName := _userOnline.userOnlineDo.TableName()
	_userOnline.ALL = field.NewAsterisk(tableName)
	_userOnline.ID = field.NewUint(tableName, "id")
	_userOnline.Uid = field.NewUint(tableName, "uid")
	_userOnline.UserName = field.NewString(tableName, "user_name")
	_userOnline.Address = field.NewString(tableName, "address")
	_userOnline.Browser = field.NewString(tableName, "browser")
	_userOnline.Ip = field.NewString(tableName, "ip")
	_userOnline.System = field.NewString(tableName, "system")
	_userOnline.LoginTime = field.NewTime(tableName, "login_time")

	_userOnline.fillFieldMap()

	return _userOnline
}

type userOnline struct {
	userOnlineDo

	ALL       field.Asterisk
	ID        field.Uint
	Uid       field.Uint
	UserName  field.String
	Address   field.String
	Browser   field.String
	Ip        field.String
	System    field.String
	LoginTime field.Time

	fieldMap map[string]field.Expr
}

func (u userOnline) Table(newTableName string) *userOnline {
	u.userOnlineDo.UseTable(newTableName)
	return u.updateTableName(newTableName)
}

func (u userOnline) As(alias string) *userOnline {
	u.userOnlineDo.DO = *(u.userOnlineDo.As(alias).(*gen.DO))
	return u.updateTableName(alias)
}

func (u *userOnline) updateTableName(table string) *userOnline {
	u.ALL = field.NewAsterisk(table)
	u.ID = field.NewUint(table, "id")
	u.Uid = field.NewUint(table, "uid")
	u.UserName = field.NewString(table, "user_name")
	u.Address = field.NewString(table, "address")
	u.Browser = field.NewString(table, "browser")
	u.Ip = field.NewString(table, "ip")
	u.System = field.NewString(table, "system")
	u.LoginTime = field.NewTime(table, "login_time")

	u.fillFieldMap()

	return u
}

func (u *userOnline) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := u.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (u *userOnline) fillFieldMap() {
	u.fieldMap = make(map[string]field.Expr, 8)
	u.fieldMap["id"] = u.ID
	u.fieldMap["uid"] = u.Uid
	u.fieldMap["user_name"] = u.UserName
	u.fieldMap["address"] = u.Address
	u.fieldMap["browser"] = u.Browser
	u.fieldMap["ip"] = u.Ip
	u.fieldMap["system"] = u.System
	u.fieldMap["login_time"] = u.LoginTime
}

func (u userOnline) clone(db *gorm.DB) userOnline {
	u.userOnlineDo.ReplaceConnPool(db.Statement.ConnPool)
	return u
}

func (u userOnline) replaceDB(db *gorm.DB) userOnline {
	u.userOnlineDo.ReplaceDB(db)
	return u
}

type userOnlineDo struct{ gen.DO }

type IUserOnlineDo interface {
	gen.SubQuery
	Debug() IUserOnlineDo
	WithContext(ctx context.Context) IUserOnlineDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IUserOnlineDo
	WriteDB() IUserOnlineDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IUserOnlineDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IUserOnlineDo
	Not(conds ...gen.Condition) IUserOnlineDo
	Or(conds ...gen.Condition) IUserOnlineDo
	Select(conds ...field.Expr) IUserOnlineDo
	Where(conds ...gen.Condition) IUserOnlineDo
	Order(conds ...field.Expr) IUserOnlineDo
	Distinct(cols ...field.Expr) IUserOnlineDo
	Omit(cols ...field.Expr) IUserOnlineDo
	Join(table schema.Tabler, on ...field.Expr) IUserOnlineDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IUserOnlineDo
	RightJoin(table schema.Tabler, on ...field.Expr) IUserOnlineDo
	Group(cols ...field.Expr) IUserOnlineDo
	Having(conds ...gen.Condition) IUserOnlineDo
	Limit(limit int) IUserOnlineDo
	Offset(offset int) IUserOnlineDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IUserOnlineDo
	Unscoped() IUserOnlineDo
	Create(values ...*model.UserOnline) error
	CreateInBatches(values []*model.UserOnline, batchSize int) error
	Save(values ...*model.UserOnline) error
	First() (*model.UserOnline, error)
	Take() (*model.UserOnline, error)
	Last() (*model.UserOnline, error)
	Find() ([]*model.UserOnline, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.UserOnline, err error)
	FindInBatches(result *[]*model.UserOnline, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.UserOnline) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IUserOnlineDo
	Assign(attrs ...field.AssignExpr) IUserOnlineDo
	Joins(fields ...field.RelationField) IUserOnlineDo
	Preload(fields ...field.RelationField) IUserOnlineDo
	FirstOrInit() (*model.UserOnline, error)
	FirstOrCreate() (*model.UserOnline, error)
	FindByPage(offset int, limit int) (result []*model.UserOnline, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IUserOnlineDo
	UnderlyingDB() *gorm.DB
	schema.Tabler

	GetByID(id uint) (result model.UserOnline, err error)
	DeleteByID(id uint) (rowsAffected int64, err error)
	DeleteByIDs(ids []uint) (rowsAffected int64, err error)
}

// SELECT * FROM @@table WHERE id=@id
func (u userOnlineDo) GetByID(id uint) (result model.UserOnline, err error) {
	var params []interface{}

	var generateSQL strings.Builder
	params = append(params, id)
	generateSQL.WriteString("SELECT * FROM sys_user_online WHERE id=? ")

	var executeSQL *gorm.DB
	executeSQL = u.UnderlyingDB().Raw(generateSQL.String(), params...).Take(&result) // ignore_security_alert
	err = executeSQL.Error

	return
}

// DELETE FROM @@table WHERE id=@id
func (u userOnlineDo) DeleteByID(id uint) (rowsAffected int64, err error) {
	var params []interface{}

	var generateSQL strings.Builder
	params = append(params, id)
	generateSQL.WriteString("DELETE FROM sys_user_online WHERE id=? ")

	var executeSQL *gorm.DB
	executeSQL = u.UnderlyingDB().Exec(generateSQL.String(), params...) // ignore_security_alert
	rowsAffected = executeSQL.RowsAffected
	err = executeSQL.Error

	return
}

// DELETE FROM @@table WHERE id IN (@ids)
func (u userOnlineDo) DeleteByIDs(ids []uint) (rowsAffected int64, err error) {
	var params []interface{}

	var generateSQL strings.Builder
	params = append(params, ids)
	generateSQL.WriteString("DELETE FROM sys_user_online WHERE id IN (?) ")

	var executeSQL *gorm.DB
	executeSQL = u.UnderlyingDB().Exec(generateSQL.String(), params...) // ignore_security_alert
	rowsAffected = executeSQL.RowsAffected
	err = executeSQL.Error

	return
}

func (u userOnlineDo) Debug() IUserOnlineDo {
	return u.withDO(u.DO.Debug())
}

func (u userOnlineDo) WithContext(ctx context.Context) IUserOnlineDo {
	return u.withDO(u.DO.WithContext(ctx))
}

func (u userOnlineDo) ReadDB() IUserOnlineDo {
	return u.Clauses(dbresolver.Read)
}

func (u userOnlineDo) WriteDB() IUserOnlineDo {
	return u.Clauses(dbresolver.Write)
}

func (u userOnlineDo) Session(config *gorm.Session) IUserOnlineDo {
	return u.withDO(u.DO.Session(config))
}

func (u userOnlineDo) Clauses(conds ...clause.Expression) IUserOnlineDo {
	return u.withDO(u.DO.Clauses(conds...))
}

func (u userOnlineDo) Returning(value interface{}, columns ...string) IUserOnlineDo {
	return u.withDO(u.DO.Returning(value, columns...))
}

func (u userOnlineDo) Not(conds ...gen.Condition) IUserOnlineDo {
	return u.withDO(u.DO.Not(conds...))
}

func (u userOnlineDo) Or(conds ...gen.Condition) IUserOnlineDo {
	return u.withDO(u.DO.Or(conds...))
}

func (u userOnlineDo) Select(conds ...field.Expr) IUserOnlineDo {
	return u.withDO(u.DO.Select(conds...))
}

func (u userOnlineDo) Where(conds ...gen.Condition) IUserOnlineDo {
	return u.withDO(u.DO.Where(conds...))
}

func (u userOnlineDo) Order(conds ...field.Expr) IUserOnlineDo {
	return u.withDO(u.DO.Order(conds...))
}

func (u userOnlineDo) Distinct(cols ...field.Expr) IUserOnlineDo {
	return u.withDO(u.DO.Distinct(cols...))
}

func (u userOnlineDo) Omit(cols ...field.Expr) IUserOnlineDo {
	return u.withDO(u.DO.Omit(cols...))
}

func (u userOnlineDo) Join(table schema.Tabler, on ...field.Expr) IUserOnlineDo {
	return u.withDO(u.DO.Join(table, on...))
}

func (u userOnlineDo) LeftJoin(table schema.Tabler, on ...field.Expr) IUserOnlineDo {
	return u.withDO(u.DO.LeftJoin(table, on...))
}

func (u userOnlineDo) RightJoin(table schema.Tabler, on ...field.Expr) IUserOnlineDo {
	return u.withDO(u.DO.RightJoin(table, on...))
}

func (u userOnlineDo) Group(cols ...field.Expr) IUserOnlineDo {
	return u.withDO(u.DO.Group(cols...))
}

func (u userOnlineDo) Having(conds ...gen.Condition) IUserOnlineDo {
	return u.withDO(u.DO.Having(conds...))
}

func (u userOnlineDo) Limit(limit int) IUserOnlineDo {
	return u.withDO(u.DO.Limit(limit))
}

func (u userOnlineDo) Offset(offset int) IUserOnlineDo {
	return u.withDO(u.DO.Offset(offset))
}

func (u userOnlineDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IUserOnlineDo {
	return u.withDO(u.DO.Scopes(funcs...))
}

func (u userOnlineDo) Unscoped() IUserOnlineDo {
	return u.withDO(u.DO.Unscoped())
}

func (u userOnlineDo) Create(values ...*model.UserOnline) error {
	if len(values) == 0 {
		return nil
	}
	return u.DO.Create(values)
}

func (u userOnlineDo) CreateInBatches(values []*model.UserOnline, batchSize int) error {
	return u.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (u userOnlineDo) Save(values ...*model.UserOnline) error {
	if len(values) == 0 {
		return nil
	}
	return u.DO.Save(values)
}

func (u userOnlineDo) First() (*model.UserOnline, error) {
	if result, err := u.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserOnline), nil
	}
}

func (u userOnlineDo) Take() (*model.UserOnline, error) {
	if result, err := u.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserOnline), nil
	}
}

func (u userOnlineDo) Last() (*model.UserOnline, error) {
	if result, err := u.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserOnline), nil
	}
}

func (u userOnlineDo) Find() ([]*model.UserOnline, error) {
	result, err := u.DO.Find()
	return result.([]*model.UserOnline), err
}

func (u userOnlineDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.UserOnline, err error) {
	buf := make([]*model.UserOnline, 0, batchSize)
	err = u.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (u userOnlineDo) FindInBatches(result *[]*model.UserOnline, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return u.DO.FindInBatches(result, batchSize, fc)
}

func (u userOnlineDo) Attrs(attrs ...field.AssignExpr) IUserOnlineDo {
	return u.withDO(u.DO.Attrs(attrs...))
}

func (u userOnlineDo) Assign(attrs ...field.AssignExpr) IUserOnlineDo {
	return u.withDO(u.DO.Assign(attrs...))
}

func (u userOnlineDo) Joins(fields ...field.RelationField) IUserOnlineDo {
	for _, _f := range fields {
		u = *u.withDO(u.DO.Joins(_f))
	}
	return &u
}

func (u userOnlineDo) Preload(fields ...field.RelationField) IUserOnlineDo {
	for _, _f := range fields {
		u = *u.withDO(u.DO.Preload(_f))
	}
	return &u
}

func (u userOnlineDo) FirstOrInit() (*model.UserOnline, error) {
	if result, err := u.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserOnline), nil
	}
}

func (u userOnlineDo) FirstOrCreate() (*model.UserOnline, error) {
	if result, err := u.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserOnline), nil
	}
}

func (u userOnlineDo) FindByPage(offset int, limit int) (result []*model.UserOnline, count int64, err error) {
	result, err = u.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = u.Offset(-1).Limit(-1).Count()
	return
}

func (u userOnlineDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = u.Count()
	if err != nil {
		return
	}

	err = u.Offset(offset).Limit(limit).Scan(result)
	return
}

func (u userOnlineDo) Scan(result interface{}) (err error) {
	return u.DO.Scan(result)
}

func (u userOnlineDo) Delete(models ...*model.UserOnline) (result gen.ResultInfo, err error) {
	return u.DO.Delete(models)
}

func (u *userOnlineDo) withDO(do gen.Dao) *userOnlineDo {
	u.DO = *do.(*gen.DO)
	return u
}
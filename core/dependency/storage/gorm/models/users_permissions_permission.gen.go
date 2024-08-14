// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package models

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"river0825/cleanarchitecture/core/dependency/storage/gorm/table"
)

func newUsersPermissionsPermission(db *gorm.DB, opts ...gen.DOOption) usersPermissionsPermission {
	_usersPermissionsPermission := usersPermissionsPermission{}

	_usersPermissionsPermission.usersPermissionsPermissionDo.UseDB(db, opts...)
	_usersPermissionsPermission.usersPermissionsPermissionDo.UseModel(&table.UsersPermissionsPermission{})

	tableName := _usersPermissionsPermission.usersPermissionsPermissionDo.TableName()
	_usersPermissionsPermission.ALL = field.NewAsterisk(tableName)
	_usersPermissionsPermission.ID = field.NewInt32(tableName, "id")
	_usersPermissionsPermission.Type = field.NewString(tableName, "type")
	_usersPermissionsPermission.Controller = field.NewString(tableName, "controller")
	_usersPermissionsPermission.Action = field.NewString(tableName, "action")
	_usersPermissionsPermission.Enabled = field.NewBool(tableName, "enabled")
	_usersPermissionsPermission.Policy = field.NewString(tableName, "policy")
	_usersPermissionsPermission.Role = field.NewInt32(tableName, "role")
	_usersPermissionsPermission.CreatedBy = field.NewInt32(tableName, "created_by")
	_usersPermissionsPermission.UpdatedBy = field.NewInt32(tableName, "updated_by")

	_usersPermissionsPermission.fillFieldMap()

	return _usersPermissionsPermission
}

type usersPermissionsPermission struct {
	usersPermissionsPermissionDo

	ALL        field.Asterisk
	ID         field.Int32
	Type       field.String
	Controller field.String
	Action     field.String
	Enabled    field.Bool
	Policy     field.String
	Role       field.Int32
	CreatedBy  field.Int32
	UpdatedBy  field.Int32

	fieldMap map[string]field.Expr
}

func (u usersPermissionsPermission) Table(newTableName string) *usersPermissionsPermission {
	u.usersPermissionsPermissionDo.UseTable(newTableName)
	return u.updateTableName(newTableName)
}

func (u usersPermissionsPermission) As(alias string) *usersPermissionsPermission {
	u.usersPermissionsPermissionDo.DO = *(u.usersPermissionsPermissionDo.As(alias).(*gen.DO))
	return u.updateTableName(alias)
}

func (u *usersPermissionsPermission) updateTableName(table string) *usersPermissionsPermission {
	u.ALL = field.NewAsterisk(table)
	u.ID = field.NewInt32(table, "id")
	u.Type = field.NewString(table, "type")
	u.Controller = field.NewString(table, "controller")
	u.Action = field.NewString(table, "action")
	u.Enabled = field.NewBool(table, "enabled")
	u.Policy = field.NewString(table, "policy")
	u.Role = field.NewInt32(table, "role")
	u.CreatedBy = field.NewInt32(table, "created_by")
	u.UpdatedBy = field.NewInt32(table, "updated_by")

	u.fillFieldMap()

	return u
}

func (u *usersPermissionsPermission) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := u.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (u *usersPermissionsPermission) fillFieldMap() {
	u.fieldMap = make(map[string]field.Expr, 9)
	u.fieldMap["id"] = u.ID
	u.fieldMap["type"] = u.Type
	u.fieldMap["controller"] = u.Controller
	u.fieldMap["action"] = u.Action
	u.fieldMap["enabled"] = u.Enabled
	u.fieldMap["policy"] = u.Policy
	u.fieldMap["role"] = u.Role
	u.fieldMap["created_by"] = u.CreatedBy
	u.fieldMap["updated_by"] = u.UpdatedBy
}

func (u usersPermissionsPermission) clone(db *gorm.DB) usersPermissionsPermission {
	u.usersPermissionsPermissionDo.ReplaceConnPool(db.Statement.ConnPool)
	return u
}

func (u usersPermissionsPermission) replaceDB(db *gorm.DB) usersPermissionsPermission {
	u.usersPermissionsPermissionDo.ReplaceDB(db)
	return u
}

type usersPermissionsPermissionDo struct{ gen.DO }

type IUsersPermissionsPermissionDo interface {
	gen.SubQuery
	Debug() IUsersPermissionsPermissionDo
	WithContext(ctx context.Context) IUsersPermissionsPermissionDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IUsersPermissionsPermissionDo
	WriteDB() IUsersPermissionsPermissionDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IUsersPermissionsPermissionDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IUsersPermissionsPermissionDo
	Not(conds ...gen.Condition) IUsersPermissionsPermissionDo
	Or(conds ...gen.Condition) IUsersPermissionsPermissionDo
	Select(conds ...field.Expr) IUsersPermissionsPermissionDo
	Where(conds ...gen.Condition) IUsersPermissionsPermissionDo
	Order(conds ...field.Expr) IUsersPermissionsPermissionDo
	Distinct(cols ...field.Expr) IUsersPermissionsPermissionDo
	Omit(cols ...field.Expr) IUsersPermissionsPermissionDo
	Join(table schema.Tabler, on ...field.Expr) IUsersPermissionsPermissionDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IUsersPermissionsPermissionDo
	RightJoin(table schema.Tabler, on ...field.Expr) IUsersPermissionsPermissionDo
	Group(cols ...field.Expr) IUsersPermissionsPermissionDo
	Having(conds ...gen.Condition) IUsersPermissionsPermissionDo
	Limit(limit int) IUsersPermissionsPermissionDo
	Offset(offset int) IUsersPermissionsPermissionDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IUsersPermissionsPermissionDo
	Unscoped() IUsersPermissionsPermissionDo
	Create(values ...*table.UsersPermissionsPermission) error
	CreateInBatches(values []*table.UsersPermissionsPermission, batchSize int) error
	Save(values ...*table.UsersPermissionsPermission) error
	First() (*table.UsersPermissionsPermission, error)
	Take() (*table.UsersPermissionsPermission, error)
	Last() (*table.UsersPermissionsPermission, error)
	Find() ([]*table.UsersPermissionsPermission, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*table.UsersPermissionsPermission, err error)
	FindInBatches(result *[]*table.UsersPermissionsPermission, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*table.UsersPermissionsPermission) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IUsersPermissionsPermissionDo
	Assign(attrs ...field.AssignExpr) IUsersPermissionsPermissionDo
	Joins(fields ...field.RelationField) IUsersPermissionsPermissionDo
	Preload(fields ...field.RelationField) IUsersPermissionsPermissionDo
	FirstOrInit() (*table.UsersPermissionsPermission, error)
	FirstOrCreate() (*table.UsersPermissionsPermission, error)
	FindByPage(offset int, limit int) (result []*table.UsersPermissionsPermission, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IUsersPermissionsPermissionDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (u usersPermissionsPermissionDo) Debug() IUsersPermissionsPermissionDo {
	return u.withDO(u.DO.Debug())
}

func (u usersPermissionsPermissionDo) WithContext(ctx context.Context) IUsersPermissionsPermissionDo {
	return u.withDO(u.DO.WithContext(ctx))
}

func (u usersPermissionsPermissionDo) ReadDB() IUsersPermissionsPermissionDo {
	return u.Clauses(dbresolver.Read)
}

func (u usersPermissionsPermissionDo) WriteDB() IUsersPermissionsPermissionDo {
	return u.Clauses(dbresolver.Write)
}

func (u usersPermissionsPermissionDo) Session(config *gorm.Session) IUsersPermissionsPermissionDo {
	return u.withDO(u.DO.Session(config))
}

func (u usersPermissionsPermissionDo) Clauses(conds ...clause.Expression) IUsersPermissionsPermissionDo {
	return u.withDO(u.DO.Clauses(conds...))
}

func (u usersPermissionsPermissionDo) Returning(value interface{}, columns ...string) IUsersPermissionsPermissionDo {
	return u.withDO(u.DO.Returning(value, columns...))
}

func (u usersPermissionsPermissionDo) Not(conds ...gen.Condition) IUsersPermissionsPermissionDo {
	return u.withDO(u.DO.Not(conds...))
}

func (u usersPermissionsPermissionDo) Or(conds ...gen.Condition) IUsersPermissionsPermissionDo {
	return u.withDO(u.DO.Or(conds...))
}

func (u usersPermissionsPermissionDo) Select(conds ...field.Expr) IUsersPermissionsPermissionDo {
	return u.withDO(u.DO.Select(conds...))
}

func (u usersPermissionsPermissionDo) Where(conds ...gen.Condition) IUsersPermissionsPermissionDo {
	return u.withDO(u.DO.Where(conds...))
}

func (u usersPermissionsPermissionDo) Order(conds ...field.Expr) IUsersPermissionsPermissionDo {
	return u.withDO(u.DO.Order(conds...))
}

func (u usersPermissionsPermissionDo) Distinct(cols ...field.Expr) IUsersPermissionsPermissionDo {
	return u.withDO(u.DO.Distinct(cols...))
}

func (u usersPermissionsPermissionDo) Omit(cols ...field.Expr) IUsersPermissionsPermissionDo {
	return u.withDO(u.DO.Omit(cols...))
}

func (u usersPermissionsPermissionDo) Join(table schema.Tabler, on ...field.Expr) IUsersPermissionsPermissionDo {
	return u.withDO(u.DO.Join(table, on...))
}

func (u usersPermissionsPermissionDo) LeftJoin(table schema.Tabler, on ...field.Expr) IUsersPermissionsPermissionDo {
	return u.withDO(u.DO.LeftJoin(table, on...))
}

func (u usersPermissionsPermissionDo) RightJoin(table schema.Tabler, on ...field.Expr) IUsersPermissionsPermissionDo {
	return u.withDO(u.DO.RightJoin(table, on...))
}

func (u usersPermissionsPermissionDo) Group(cols ...field.Expr) IUsersPermissionsPermissionDo {
	return u.withDO(u.DO.Group(cols...))
}

func (u usersPermissionsPermissionDo) Having(conds ...gen.Condition) IUsersPermissionsPermissionDo {
	return u.withDO(u.DO.Having(conds...))
}

func (u usersPermissionsPermissionDo) Limit(limit int) IUsersPermissionsPermissionDo {
	return u.withDO(u.DO.Limit(limit))
}

func (u usersPermissionsPermissionDo) Offset(offset int) IUsersPermissionsPermissionDo {
	return u.withDO(u.DO.Offset(offset))
}

func (u usersPermissionsPermissionDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IUsersPermissionsPermissionDo {
	return u.withDO(u.DO.Scopes(funcs...))
}

func (u usersPermissionsPermissionDo) Unscoped() IUsersPermissionsPermissionDo {
	return u.withDO(u.DO.Unscoped())
}

func (u usersPermissionsPermissionDo) Create(values ...*table.UsersPermissionsPermission) error {
	if len(values) == 0 {
		return nil
	}
	return u.DO.Create(values)
}

func (u usersPermissionsPermissionDo) CreateInBatches(values []*table.UsersPermissionsPermission, batchSize int) error {
	return u.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (u usersPermissionsPermissionDo) Save(values ...*table.UsersPermissionsPermission) error {
	if len(values) == 0 {
		return nil
	}
	return u.DO.Save(values)
}

func (u usersPermissionsPermissionDo) First() (*table.UsersPermissionsPermission, error) {
	if result, err := u.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*table.UsersPermissionsPermission), nil
	}
}

func (u usersPermissionsPermissionDo) Take() (*table.UsersPermissionsPermission, error) {
	if result, err := u.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*table.UsersPermissionsPermission), nil
	}
}

func (u usersPermissionsPermissionDo) Last() (*table.UsersPermissionsPermission, error) {
	if result, err := u.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*table.UsersPermissionsPermission), nil
	}
}

func (u usersPermissionsPermissionDo) Find() ([]*table.UsersPermissionsPermission, error) {
	result, err := u.DO.Find()
	return result.([]*table.UsersPermissionsPermission), err
}

func (u usersPermissionsPermissionDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*table.UsersPermissionsPermission, err error) {
	buf := make([]*table.UsersPermissionsPermission, 0, batchSize)
	err = u.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (u usersPermissionsPermissionDo) FindInBatches(result *[]*table.UsersPermissionsPermission, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return u.DO.FindInBatches(result, batchSize, fc)
}

func (u usersPermissionsPermissionDo) Attrs(attrs ...field.AssignExpr) IUsersPermissionsPermissionDo {
	return u.withDO(u.DO.Attrs(attrs...))
}

func (u usersPermissionsPermissionDo) Assign(attrs ...field.AssignExpr) IUsersPermissionsPermissionDo {
	return u.withDO(u.DO.Assign(attrs...))
}

func (u usersPermissionsPermissionDo) Joins(fields ...field.RelationField) IUsersPermissionsPermissionDo {
	for _, _f := range fields {
		u = *u.withDO(u.DO.Joins(_f))
	}
	return &u
}

func (u usersPermissionsPermissionDo) Preload(fields ...field.RelationField) IUsersPermissionsPermissionDo {
	for _, _f := range fields {
		u = *u.withDO(u.DO.Preload(_f))
	}
	return &u
}

func (u usersPermissionsPermissionDo) FirstOrInit() (*table.UsersPermissionsPermission, error) {
	if result, err := u.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*table.UsersPermissionsPermission), nil
	}
}

func (u usersPermissionsPermissionDo) FirstOrCreate() (*table.UsersPermissionsPermission, error) {
	if result, err := u.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*table.UsersPermissionsPermission), nil
	}
}

func (u usersPermissionsPermissionDo) FindByPage(offset int, limit int) (result []*table.UsersPermissionsPermission, count int64, err error) {
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

func (u usersPermissionsPermissionDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = u.Count()
	if err != nil {
		return
	}

	err = u.Offset(offset).Limit(limit).Scan(result)
	return
}

func (u usersPermissionsPermissionDo) Scan(result interface{}) (err error) {
	return u.DO.Scan(result)
}

func (u usersPermissionsPermissionDo) Delete(models ...*table.UsersPermissionsPermission) (result gen.ResultInfo, err error) {
	return u.DO.Delete(models)
}

func (u *usersPermissionsPermissionDo) withDO(do gen.Dao) *usersPermissionsPermissionDo {
	u.DO = *do.(*gen.DO)
	return u
}
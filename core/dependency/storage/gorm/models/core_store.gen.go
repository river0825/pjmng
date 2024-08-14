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

func newCoreStore(db *gorm.DB, opts ...gen.DOOption) coreStore {
	_coreStore := coreStore{}

	_coreStore.coreStoreDo.UseDB(db, opts...)
	_coreStore.coreStoreDo.UseModel(&table.CoreStore{})

	tableName := _coreStore.coreStoreDo.TableName()
	_coreStore.ALL = field.NewAsterisk(tableName)
	_coreStore.ID = field.NewInt32(tableName, "id")
	_coreStore.Key = field.NewString(tableName, "key")
	_coreStore.Value = field.NewString(tableName, "value")
	_coreStore.Type = field.NewString(tableName, "type")
	_coreStore.Environment = field.NewString(tableName, "environment")
	_coreStore.Tag = field.NewString(tableName, "tag")

	_coreStore.fillFieldMap()

	return _coreStore
}

type coreStore struct {
	coreStoreDo

	ALL         field.Asterisk
	ID          field.Int32
	Key         field.String
	Value       field.String
	Type        field.String
	Environment field.String
	Tag         field.String

	fieldMap map[string]field.Expr
}

func (c coreStore) Table(newTableName string) *coreStore {
	c.coreStoreDo.UseTable(newTableName)
	return c.updateTableName(newTableName)
}

func (c coreStore) As(alias string) *coreStore {
	c.coreStoreDo.DO = *(c.coreStoreDo.As(alias).(*gen.DO))
	return c.updateTableName(alias)
}

func (c *coreStore) updateTableName(table string) *coreStore {
	c.ALL = field.NewAsterisk(table)
	c.ID = field.NewInt32(table, "id")
	c.Key = field.NewString(table, "key")
	c.Value = field.NewString(table, "value")
	c.Type = field.NewString(table, "type")
	c.Environment = field.NewString(table, "environment")
	c.Tag = field.NewString(table, "tag")

	c.fillFieldMap()

	return c
}

func (c *coreStore) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := c.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (c *coreStore) fillFieldMap() {
	c.fieldMap = make(map[string]field.Expr, 6)
	c.fieldMap["id"] = c.ID
	c.fieldMap["key"] = c.Key
	c.fieldMap["value"] = c.Value
	c.fieldMap["type"] = c.Type
	c.fieldMap["environment"] = c.Environment
	c.fieldMap["tag"] = c.Tag
}

func (c coreStore) clone(db *gorm.DB) coreStore {
	c.coreStoreDo.ReplaceConnPool(db.Statement.ConnPool)
	return c
}

func (c coreStore) replaceDB(db *gorm.DB) coreStore {
	c.coreStoreDo.ReplaceDB(db)
	return c
}

type coreStoreDo struct{ gen.DO }

type ICoreStoreDo interface {
	gen.SubQuery
	Debug() ICoreStoreDo
	WithContext(ctx context.Context) ICoreStoreDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ICoreStoreDo
	WriteDB() ICoreStoreDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ICoreStoreDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ICoreStoreDo
	Not(conds ...gen.Condition) ICoreStoreDo
	Or(conds ...gen.Condition) ICoreStoreDo
	Select(conds ...field.Expr) ICoreStoreDo
	Where(conds ...gen.Condition) ICoreStoreDo
	Order(conds ...field.Expr) ICoreStoreDo
	Distinct(cols ...field.Expr) ICoreStoreDo
	Omit(cols ...field.Expr) ICoreStoreDo
	Join(table schema.Tabler, on ...field.Expr) ICoreStoreDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ICoreStoreDo
	RightJoin(table schema.Tabler, on ...field.Expr) ICoreStoreDo
	Group(cols ...field.Expr) ICoreStoreDo
	Having(conds ...gen.Condition) ICoreStoreDo
	Limit(limit int) ICoreStoreDo
	Offset(offset int) ICoreStoreDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ICoreStoreDo
	Unscoped() ICoreStoreDo
	Create(values ...*table.CoreStore) error
	CreateInBatches(values []*table.CoreStore, batchSize int) error
	Save(values ...*table.CoreStore) error
	First() (*table.CoreStore, error)
	Take() (*table.CoreStore, error)
	Last() (*table.CoreStore, error)
	Find() ([]*table.CoreStore, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*table.CoreStore, err error)
	FindInBatches(result *[]*table.CoreStore, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*table.CoreStore) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ICoreStoreDo
	Assign(attrs ...field.AssignExpr) ICoreStoreDo
	Joins(fields ...field.RelationField) ICoreStoreDo
	Preload(fields ...field.RelationField) ICoreStoreDo
	FirstOrInit() (*table.CoreStore, error)
	FirstOrCreate() (*table.CoreStore, error)
	FindByPage(offset int, limit int) (result []*table.CoreStore, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ICoreStoreDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (c coreStoreDo) Debug() ICoreStoreDo {
	return c.withDO(c.DO.Debug())
}

func (c coreStoreDo) WithContext(ctx context.Context) ICoreStoreDo {
	return c.withDO(c.DO.WithContext(ctx))
}

func (c coreStoreDo) ReadDB() ICoreStoreDo {
	return c.Clauses(dbresolver.Read)
}

func (c coreStoreDo) WriteDB() ICoreStoreDo {
	return c.Clauses(dbresolver.Write)
}

func (c coreStoreDo) Session(config *gorm.Session) ICoreStoreDo {
	return c.withDO(c.DO.Session(config))
}

func (c coreStoreDo) Clauses(conds ...clause.Expression) ICoreStoreDo {
	return c.withDO(c.DO.Clauses(conds...))
}

func (c coreStoreDo) Returning(value interface{}, columns ...string) ICoreStoreDo {
	return c.withDO(c.DO.Returning(value, columns...))
}

func (c coreStoreDo) Not(conds ...gen.Condition) ICoreStoreDo {
	return c.withDO(c.DO.Not(conds...))
}

func (c coreStoreDo) Or(conds ...gen.Condition) ICoreStoreDo {
	return c.withDO(c.DO.Or(conds...))
}

func (c coreStoreDo) Select(conds ...field.Expr) ICoreStoreDo {
	return c.withDO(c.DO.Select(conds...))
}

func (c coreStoreDo) Where(conds ...gen.Condition) ICoreStoreDo {
	return c.withDO(c.DO.Where(conds...))
}

func (c coreStoreDo) Order(conds ...field.Expr) ICoreStoreDo {
	return c.withDO(c.DO.Order(conds...))
}

func (c coreStoreDo) Distinct(cols ...field.Expr) ICoreStoreDo {
	return c.withDO(c.DO.Distinct(cols...))
}

func (c coreStoreDo) Omit(cols ...field.Expr) ICoreStoreDo {
	return c.withDO(c.DO.Omit(cols...))
}

func (c coreStoreDo) Join(table schema.Tabler, on ...field.Expr) ICoreStoreDo {
	return c.withDO(c.DO.Join(table, on...))
}

func (c coreStoreDo) LeftJoin(table schema.Tabler, on ...field.Expr) ICoreStoreDo {
	return c.withDO(c.DO.LeftJoin(table, on...))
}

func (c coreStoreDo) RightJoin(table schema.Tabler, on ...field.Expr) ICoreStoreDo {
	return c.withDO(c.DO.RightJoin(table, on...))
}

func (c coreStoreDo) Group(cols ...field.Expr) ICoreStoreDo {
	return c.withDO(c.DO.Group(cols...))
}

func (c coreStoreDo) Having(conds ...gen.Condition) ICoreStoreDo {
	return c.withDO(c.DO.Having(conds...))
}

func (c coreStoreDo) Limit(limit int) ICoreStoreDo {
	return c.withDO(c.DO.Limit(limit))
}

func (c coreStoreDo) Offset(offset int) ICoreStoreDo {
	return c.withDO(c.DO.Offset(offset))
}

func (c coreStoreDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ICoreStoreDo {
	return c.withDO(c.DO.Scopes(funcs...))
}

func (c coreStoreDo) Unscoped() ICoreStoreDo {
	return c.withDO(c.DO.Unscoped())
}

func (c coreStoreDo) Create(values ...*table.CoreStore) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Create(values)
}

func (c coreStoreDo) CreateInBatches(values []*table.CoreStore, batchSize int) error {
	return c.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (c coreStoreDo) Save(values ...*table.CoreStore) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Save(values)
}

func (c coreStoreDo) First() (*table.CoreStore, error) {
	if result, err := c.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*table.CoreStore), nil
	}
}

func (c coreStoreDo) Take() (*table.CoreStore, error) {
	if result, err := c.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*table.CoreStore), nil
	}
}

func (c coreStoreDo) Last() (*table.CoreStore, error) {
	if result, err := c.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*table.CoreStore), nil
	}
}

func (c coreStoreDo) Find() ([]*table.CoreStore, error) {
	result, err := c.DO.Find()
	return result.([]*table.CoreStore), err
}

func (c coreStoreDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*table.CoreStore, err error) {
	buf := make([]*table.CoreStore, 0, batchSize)
	err = c.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (c coreStoreDo) FindInBatches(result *[]*table.CoreStore, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return c.DO.FindInBatches(result, batchSize, fc)
}

func (c coreStoreDo) Attrs(attrs ...field.AssignExpr) ICoreStoreDo {
	return c.withDO(c.DO.Attrs(attrs...))
}

func (c coreStoreDo) Assign(attrs ...field.AssignExpr) ICoreStoreDo {
	return c.withDO(c.DO.Assign(attrs...))
}

func (c coreStoreDo) Joins(fields ...field.RelationField) ICoreStoreDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Joins(_f))
	}
	return &c
}

func (c coreStoreDo) Preload(fields ...field.RelationField) ICoreStoreDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Preload(_f))
	}
	return &c
}

func (c coreStoreDo) FirstOrInit() (*table.CoreStore, error) {
	if result, err := c.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*table.CoreStore), nil
	}
}

func (c coreStoreDo) FirstOrCreate() (*table.CoreStore, error) {
	if result, err := c.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*table.CoreStore), nil
	}
}

func (c coreStoreDo) FindByPage(offset int, limit int) (result []*table.CoreStore, count int64, err error) {
	result, err = c.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = c.Offset(-1).Limit(-1).Count()
	return
}

func (c coreStoreDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = c.Count()
	if err != nil {
		return
	}

	err = c.Offset(offset).Limit(limit).Scan(result)
	return
}

func (c coreStoreDo) Scan(result interface{}) (err error) {
	return c.DO.Scan(result)
}

func (c coreStoreDo) Delete(models ...*table.CoreStore) (result gen.ResultInfo, err error) {
	return c.DO.Delete(models)
}

func (c *coreStoreDo) withDO(do gen.Dao) *coreStoreDo {
	c.DO = *do.(*gen.DO)
	return c
}
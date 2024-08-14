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

func newStrapiWebhook(db *gorm.DB, opts ...gen.DOOption) strapiWebhook {
	_strapiWebhook := strapiWebhook{}

	_strapiWebhook.strapiWebhookDo.UseDB(db, opts...)
	_strapiWebhook.strapiWebhookDo.UseModel(&table.StrapiWebhook{})

	tableName := _strapiWebhook.strapiWebhookDo.TableName()
	_strapiWebhook.ALL = field.NewAsterisk(tableName)
	_strapiWebhook.ID = field.NewInt32(tableName, "id")
	_strapiWebhook.Name = field.NewString(tableName, "name")
	_strapiWebhook.URL = field.NewString(tableName, "url")
	_strapiWebhook.Headers = field.NewField(tableName, "headers")
	_strapiWebhook.Events = field.NewField(tableName, "events")
	_strapiWebhook.Enabled = field.NewBool(tableName, "enabled")

	_strapiWebhook.fillFieldMap()

	return _strapiWebhook
}

type strapiWebhook struct {
	strapiWebhookDo

	ALL     field.Asterisk
	ID      field.Int32
	Name    field.String
	URL     field.String
	Headers field.Field
	Events  field.Field
	Enabled field.Bool

	fieldMap map[string]field.Expr
}

func (s strapiWebhook) Table(newTableName string) *strapiWebhook {
	s.strapiWebhookDo.UseTable(newTableName)
	return s.updateTableName(newTableName)
}

func (s strapiWebhook) As(alias string) *strapiWebhook {
	s.strapiWebhookDo.DO = *(s.strapiWebhookDo.As(alias).(*gen.DO))
	return s.updateTableName(alias)
}

func (s *strapiWebhook) updateTableName(table string) *strapiWebhook {
	s.ALL = field.NewAsterisk(table)
	s.ID = field.NewInt32(table, "id")
	s.Name = field.NewString(table, "name")
	s.URL = field.NewString(table, "url")
	s.Headers = field.NewField(table, "headers")
	s.Events = field.NewField(table, "events")
	s.Enabled = field.NewBool(table, "enabled")

	s.fillFieldMap()

	return s
}

func (s *strapiWebhook) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *strapiWebhook) fillFieldMap() {
	s.fieldMap = make(map[string]field.Expr, 6)
	s.fieldMap["id"] = s.ID
	s.fieldMap["name"] = s.Name
	s.fieldMap["url"] = s.URL
	s.fieldMap["headers"] = s.Headers
	s.fieldMap["events"] = s.Events
	s.fieldMap["enabled"] = s.Enabled
}

func (s strapiWebhook) clone(db *gorm.DB) strapiWebhook {
	s.strapiWebhookDo.ReplaceConnPool(db.Statement.ConnPool)
	return s
}

func (s strapiWebhook) replaceDB(db *gorm.DB) strapiWebhook {
	s.strapiWebhookDo.ReplaceDB(db)
	return s
}

type strapiWebhookDo struct{ gen.DO }

type IStrapiWebhookDo interface {
	gen.SubQuery
	Debug() IStrapiWebhookDo
	WithContext(ctx context.Context) IStrapiWebhookDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IStrapiWebhookDo
	WriteDB() IStrapiWebhookDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IStrapiWebhookDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IStrapiWebhookDo
	Not(conds ...gen.Condition) IStrapiWebhookDo
	Or(conds ...gen.Condition) IStrapiWebhookDo
	Select(conds ...field.Expr) IStrapiWebhookDo
	Where(conds ...gen.Condition) IStrapiWebhookDo
	Order(conds ...field.Expr) IStrapiWebhookDo
	Distinct(cols ...field.Expr) IStrapiWebhookDo
	Omit(cols ...field.Expr) IStrapiWebhookDo
	Join(table schema.Tabler, on ...field.Expr) IStrapiWebhookDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IStrapiWebhookDo
	RightJoin(table schema.Tabler, on ...field.Expr) IStrapiWebhookDo
	Group(cols ...field.Expr) IStrapiWebhookDo
	Having(conds ...gen.Condition) IStrapiWebhookDo
	Limit(limit int) IStrapiWebhookDo
	Offset(offset int) IStrapiWebhookDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IStrapiWebhookDo
	Unscoped() IStrapiWebhookDo
	Create(values ...*table.StrapiWebhook) error
	CreateInBatches(values []*table.StrapiWebhook, batchSize int) error
	Save(values ...*table.StrapiWebhook) error
	First() (*table.StrapiWebhook, error)
	Take() (*table.StrapiWebhook, error)
	Last() (*table.StrapiWebhook, error)
	Find() ([]*table.StrapiWebhook, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*table.StrapiWebhook, err error)
	FindInBatches(result *[]*table.StrapiWebhook, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*table.StrapiWebhook) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IStrapiWebhookDo
	Assign(attrs ...field.AssignExpr) IStrapiWebhookDo
	Joins(fields ...field.RelationField) IStrapiWebhookDo
	Preload(fields ...field.RelationField) IStrapiWebhookDo
	FirstOrInit() (*table.StrapiWebhook, error)
	FirstOrCreate() (*table.StrapiWebhook, error)
	FindByPage(offset int, limit int) (result []*table.StrapiWebhook, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IStrapiWebhookDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (s strapiWebhookDo) Debug() IStrapiWebhookDo {
	return s.withDO(s.DO.Debug())
}

func (s strapiWebhookDo) WithContext(ctx context.Context) IStrapiWebhookDo {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s strapiWebhookDo) ReadDB() IStrapiWebhookDo {
	return s.Clauses(dbresolver.Read)
}

func (s strapiWebhookDo) WriteDB() IStrapiWebhookDo {
	return s.Clauses(dbresolver.Write)
}

func (s strapiWebhookDo) Session(config *gorm.Session) IStrapiWebhookDo {
	return s.withDO(s.DO.Session(config))
}

func (s strapiWebhookDo) Clauses(conds ...clause.Expression) IStrapiWebhookDo {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s strapiWebhookDo) Returning(value interface{}, columns ...string) IStrapiWebhookDo {
	return s.withDO(s.DO.Returning(value, columns...))
}

func (s strapiWebhookDo) Not(conds ...gen.Condition) IStrapiWebhookDo {
	return s.withDO(s.DO.Not(conds...))
}

func (s strapiWebhookDo) Or(conds ...gen.Condition) IStrapiWebhookDo {
	return s.withDO(s.DO.Or(conds...))
}

func (s strapiWebhookDo) Select(conds ...field.Expr) IStrapiWebhookDo {
	return s.withDO(s.DO.Select(conds...))
}

func (s strapiWebhookDo) Where(conds ...gen.Condition) IStrapiWebhookDo {
	return s.withDO(s.DO.Where(conds...))
}

func (s strapiWebhookDo) Order(conds ...field.Expr) IStrapiWebhookDo {
	return s.withDO(s.DO.Order(conds...))
}

func (s strapiWebhookDo) Distinct(cols ...field.Expr) IStrapiWebhookDo {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s strapiWebhookDo) Omit(cols ...field.Expr) IStrapiWebhookDo {
	return s.withDO(s.DO.Omit(cols...))
}

func (s strapiWebhookDo) Join(table schema.Tabler, on ...field.Expr) IStrapiWebhookDo {
	return s.withDO(s.DO.Join(table, on...))
}

func (s strapiWebhookDo) LeftJoin(table schema.Tabler, on ...field.Expr) IStrapiWebhookDo {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s strapiWebhookDo) RightJoin(table schema.Tabler, on ...field.Expr) IStrapiWebhookDo {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s strapiWebhookDo) Group(cols ...field.Expr) IStrapiWebhookDo {
	return s.withDO(s.DO.Group(cols...))
}

func (s strapiWebhookDo) Having(conds ...gen.Condition) IStrapiWebhookDo {
	return s.withDO(s.DO.Having(conds...))
}

func (s strapiWebhookDo) Limit(limit int) IStrapiWebhookDo {
	return s.withDO(s.DO.Limit(limit))
}

func (s strapiWebhookDo) Offset(offset int) IStrapiWebhookDo {
	return s.withDO(s.DO.Offset(offset))
}

func (s strapiWebhookDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IStrapiWebhookDo {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s strapiWebhookDo) Unscoped() IStrapiWebhookDo {
	return s.withDO(s.DO.Unscoped())
}

func (s strapiWebhookDo) Create(values ...*table.StrapiWebhook) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s strapiWebhookDo) CreateInBatches(values []*table.StrapiWebhook, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s strapiWebhookDo) Save(values ...*table.StrapiWebhook) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s strapiWebhookDo) First() (*table.StrapiWebhook, error) {
	if result, err := s.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*table.StrapiWebhook), nil
	}
}

func (s strapiWebhookDo) Take() (*table.StrapiWebhook, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*table.StrapiWebhook), nil
	}
}

func (s strapiWebhookDo) Last() (*table.StrapiWebhook, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*table.StrapiWebhook), nil
	}
}

func (s strapiWebhookDo) Find() ([]*table.StrapiWebhook, error) {
	result, err := s.DO.Find()
	return result.([]*table.StrapiWebhook), err
}

func (s strapiWebhookDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*table.StrapiWebhook, err error) {
	buf := make([]*table.StrapiWebhook, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s strapiWebhookDo) FindInBatches(result *[]*table.StrapiWebhook, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s strapiWebhookDo) Attrs(attrs ...field.AssignExpr) IStrapiWebhookDo {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s strapiWebhookDo) Assign(attrs ...field.AssignExpr) IStrapiWebhookDo {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s strapiWebhookDo) Joins(fields ...field.RelationField) IStrapiWebhookDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Joins(_f))
	}
	return &s
}

func (s strapiWebhookDo) Preload(fields ...field.RelationField) IStrapiWebhookDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Preload(_f))
	}
	return &s
}

func (s strapiWebhookDo) FirstOrInit() (*table.StrapiWebhook, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*table.StrapiWebhook), nil
	}
}

func (s strapiWebhookDo) FirstOrCreate() (*table.StrapiWebhook, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*table.StrapiWebhook), nil
	}
}

func (s strapiWebhookDo) FindByPage(offset int, limit int) (result []*table.StrapiWebhook, count int64, err error) {
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

func (s strapiWebhookDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s strapiWebhookDo) Scan(result interface{}) (err error) {
	return s.DO.Scan(result)
}

func (s strapiWebhookDo) Delete(models ...*table.StrapiWebhook) (result gen.ResultInfo, err error) {
	return s.DO.Delete(models)
}

func (s *strapiWebhookDo) withDO(do gen.Dao) *strapiWebhookDo {
	s.DO = *do.(*gen.DO)
	return s
}

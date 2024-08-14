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

func newMaterialTag(db *gorm.DB, opts ...gen.DOOption) materialTag {
	_materialTag := materialTag{}

	_materialTag.materialTagDo.UseDB(db, opts...)
	_materialTag.materialTagDo.UseModel(&table.MaterialTag{})

	tableName := _materialTag.materialTagDo.TableName()
	_materialTag.ALL = field.NewAsterisk(tableName)
	_materialTag.ID = field.NewInt32(tableName, "id")
	_materialTag.User = field.NewInt32(tableName, "user")
	_materialTag.TagList = field.NewField(tableName, "tag_list")
	_materialTag.PublishedAt = field.NewTime(tableName, "published_at")
	_materialTag.CreatedBy = field.NewInt32(tableName, "created_by")
	_materialTag.UpdatedBy = field.NewInt32(tableName, "updated_by")
	_materialTag.CreatedAt = field.NewTime(tableName, "created_at")
	_materialTag.UpdatedAt = field.NewTime(tableName, "updated_at")
	_materialTag.Tags = field.NewString(tableName, "tags")

	_materialTag.fillFieldMap()

	return _materialTag
}

type materialTag struct {
	materialTagDo

	ALL         field.Asterisk
	ID          field.Int32
	User        field.Int32
	TagList     field.Field
	PublishedAt field.Time
	CreatedBy   field.Int32
	UpdatedBy   field.Int32
	CreatedAt   field.Time
	UpdatedAt   field.Time
	Tags        field.String

	fieldMap map[string]field.Expr
}

func (m materialTag) Table(newTableName string) *materialTag {
	m.materialTagDo.UseTable(newTableName)
	return m.updateTableName(newTableName)
}

func (m materialTag) As(alias string) *materialTag {
	m.materialTagDo.DO = *(m.materialTagDo.As(alias).(*gen.DO))
	return m.updateTableName(alias)
}

func (m *materialTag) updateTableName(table string) *materialTag {
	m.ALL = field.NewAsterisk(table)
	m.ID = field.NewInt32(table, "id")
	m.User = field.NewInt32(table, "user")
	m.TagList = field.NewField(table, "tag_list")
	m.PublishedAt = field.NewTime(table, "published_at")
	m.CreatedBy = field.NewInt32(table, "created_by")
	m.UpdatedBy = field.NewInt32(table, "updated_by")
	m.CreatedAt = field.NewTime(table, "created_at")
	m.UpdatedAt = field.NewTime(table, "updated_at")
	m.Tags = field.NewString(table, "tags")

	m.fillFieldMap()

	return m
}

func (m *materialTag) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := m.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (m *materialTag) fillFieldMap() {
	m.fieldMap = make(map[string]field.Expr, 9)
	m.fieldMap["id"] = m.ID
	m.fieldMap["user"] = m.User
	m.fieldMap["tag_list"] = m.TagList
	m.fieldMap["published_at"] = m.PublishedAt
	m.fieldMap["created_by"] = m.CreatedBy
	m.fieldMap["updated_by"] = m.UpdatedBy
	m.fieldMap["created_at"] = m.CreatedAt
	m.fieldMap["updated_at"] = m.UpdatedAt
	m.fieldMap["tags"] = m.Tags
}

func (m materialTag) clone(db *gorm.DB) materialTag {
	m.materialTagDo.ReplaceConnPool(db.Statement.ConnPool)
	return m
}

func (m materialTag) replaceDB(db *gorm.DB) materialTag {
	m.materialTagDo.ReplaceDB(db)
	return m
}

type materialTagDo struct{ gen.DO }

type IMaterialTagDo interface {
	gen.SubQuery
	Debug() IMaterialTagDo
	WithContext(ctx context.Context) IMaterialTagDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IMaterialTagDo
	WriteDB() IMaterialTagDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IMaterialTagDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IMaterialTagDo
	Not(conds ...gen.Condition) IMaterialTagDo
	Or(conds ...gen.Condition) IMaterialTagDo
	Select(conds ...field.Expr) IMaterialTagDo
	Where(conds ...gen.Condition) IMaterialTagDo
	Order(conds ...field.Expr) IMaterialTagDo
	Distinct(cols ...field.Expr) IMaterialTagDo
	Omit(cols ...field.Expr) IMaterialTagDo
	Join(table schema.Tabler, on ...field.Expr) IMaterialTagDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IMaterialTagDo
	RightJoin(table schema.Tabler, on ...field.Expr) IMaterialTagDo
	Group(cols ...field.Expr) IMaterialTagDo
	Having(conds ...gen.Condition) IMaterialTagDo
	Limit(limit int) IMaterialTagDo
	Offset(offset int) IMaterialTagDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IMaterialTagDo
	Unscoped() IMaterialTagDo
	Create(values ...*table.MaterialTag) error
	CreateInBatches(values []*table.MaterialTag, batchSize int) error
	Save(values ...*table.MaterialTag) error
	First() (*table.MaterialTag, error)
	Take() (*table.MaterialTag, error)
	Last() (*table.MaterialTag, error)
	Find() ([]*table.MaterialTag, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*table.MaterialTag, err error)
	FindInBatches(result *[]*table.MaterialTag, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*table.MaterialTag) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IMaterialTagDo
	Assign(attrs ...field.AssignExpr) IMaterialTagDo
	Joins(fields ...field.RelationField) IMaterialTagDo
	Preload(fields ...field.RelationField) IMaterialTagDo
	FirstOrInit() (*table.MaterialTag, error)
	FirstOrCreate() (*table.MaterialTag, error)
	FindByPage(offset int, limit int) (result []*table.MaterialTag, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IMaterialTagDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (m materialTagDo) Debug() IMaterialTagDo {
	return m.withDO(m.DO.Debug())
}

func (m materialTagDo) WithContext(ctx context.Context) IMaterialTagDo {
	return m.withDO(m.DO.WithContext(ctx))
}

func (m materialTagDo) ReadDB() IMaterialTagDo {
	return m.Clauses(dbresolver.Read)
}

func (m materialTagDo) WriteDB() IMaterialTagDo {
	return m.Clauses(dbresolver.Write)
}

func (m materialTagDo) Session(config *gorm.Session) IMaterialTagDo {
	return m.withDO(m.DO.Session(config))
}

func (m materialTagDo) Clauses(conds ...clause.Expression) IMaterialTagDo {
	return m.withDO(m.DO.Clauses(conds...))
}

func (m materialTagDo) Returning(value interface{}, columns ...string) IMaterialTagDo {
	return m.withDO(m.DO.Returning(value, columns...))
}

func (m materialTagDo) Not(conds ...gen.Condition) IMaterialTagDo {
	return m.withDO(m.DO.Not(conds...))
}

func (m materialTagDo) Or(conds ...gen.Condition) IMaterialTagDo {
	return m.withDO(m.DO.Or(conds...))
}

func (m materialTagDo) Select(conds ...field.Expr) IMaterialTagDo {
	return m.withDO(m.DO.Select(conds...))
}

func (m materialTagDo) Where(conds ...gen.Condition) IMaterialTagDo {
	return m.withDO(m.DO.Where(conds...))
}

func (m materialTagDo) Order(conds ...field.Expr) IMaterialTagDo {
	return m.withDO(m.DO.Order(conds...))
}

func (m materialTagDo) Distinct(cols ...field.Expr) IMaterialTagDo {
	return m.withDO(m.DO.Distinct(cols...))
}

func (m materialTagDo) Omit(cols ...field.Expr) IMaterialTagDo {
	return m.withDO(m.DO.Omit(cols...))
}

func (m materialTagDo) Join(table schema.Tabler, on ...field.Expr) IMaterialTagDo {
	return m.withDO(m.DO.Join(table, on...))
}

func (m materialTagDo) LeftJoin(table schema.Tabler, on ...field.Expr) IMaterialTagDo {
	return m.withDO(m.DO.LeftJoin(table, on...))
}

func (m materialTagDo) RightJoin(table schema.Tabler, on ...field.Expr) IMaterialTagDo {
	return m.withDO(m.DO.RightJoin(table, on...))
}

func (m materialTagDo) Group(cols ...field.Expr) IMaterialTagDo {
	return m.withDO(m.DO.Group(cols...))
}

func (m materialTagDo) Having(conds ...gen.Condition) IMaterialTagDo {
	return m.withDO(m.DO.Having(conds...))
}

func (m materialTagDo) Limit(limit int) IMaterialTagDo {
	return m.withDO(m.DO.Limit(limit))
}

func (m materialTagDo) Offset(offset int) IMaterialTagDo {
	return m.withDO(m.DO.Offset(offset))
}

func (m materialTagDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IMaterialTagDo {
	return m.withDO(m.DO.Scopes(funcs...))
}

func (m materialTagDo) Unscoped() IMaterialTagDo {
	return m.withDO(m.DO.Unscoped())
}

func (m materialTagDo) Create(values ...*table.MaterialTag) error {
	if len(values) == 0 {
		return nil
	}
	return m.DO.Create(values)
}

func (m materialTagDo) CreateInBatches(values []*table.MaterialTag, batchSize int) error {
	return m.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (m materialTagDo) Save(values ...*table.MaterialTag) error {
	if len(values) == 0 {
		return nil
	}
	return m.DO.Save(values)
}

func (m materialTagDo) First() (*table.MaterialTag, error) {
	if result, err := m.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*table.MaterialTag), nil
	}
}

func (m materialTagDo) Take() (*table.MaterialTag, error) {
	if result, err := m.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*table.MaterialTag), nil
	}
}

func (m materialTagDo) Last() (*table.MaterialTag, error) {
	if result, err := m.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*table.MaterialTag), nil
	}
}

func (m materialTagDo) Find() ([]*table.MaterialTag, error) {
	result, err := m.DO.Find()
	return result.([]*table.MaterialTag), err
}

func (m materialTagDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*table.MaterialTag, err error) {
	buf := make([]*table.MaterialTag, 0, batchSize)
	err = m.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (m materialTagDo) FindInBatches(result *[]*table.MaterialTag, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return m.DO.FindInBatches(result, batchSize, fc)
}

func (m materialTagDo) Attrs(attrs ...field.AssignExpr) IMaterialTagDo {
	return m.withDO(m.DO.Attrs(attrs...))
}

func (m materialTagDo) Assign(attrs ...field.AssignExpr) IMaterialTagDo {
	return m.withDO(m.DO.Assign(attrs...))
}

func (m materialTagDo) Joins(fields ...field.RelationField) IMaterialTagDo {
	for _, _f := range fields {
		m = *m.withDO(m.DO.Joins(_f))
	}
	return &m
}

func (m materialTagDo) Preload(fields ...field.RelationField) IMaterialTagDo {
	for _, _f := range fields {
		m = *m.withDO(m.DO.Preload(_f))
	}
	return &m
}

func (m materialTagDo) FirstOrInit() (*table.MaterialTag, error) {
	if result, err := m.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*table.MaterialTag), nil
	}
}

func (m materialTagDo) FirstOrCreate() (*table.MaterialTag, error) {
	if result, err := m.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*table.MaterialTag), nil
	}
}

func (m materialTagDo) FindByPage(offset int, limit int) (result []*table.MaterialTag, count int64, err error) {
	result, err = m.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = m.Offset(-1).Limit(-1).Count()
	return
}

func (m materialTagDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = m.Count()
	if err != nil {
		return
	}

	err = m.Offset(offset).Limit(limit).Scan(result)
	return
}

func (m materialTagDo) Scan(result interface{}) (err error) {
	return m.DO.Scan(result)
}

func (m materialTagDo) Delete(models ...*table.MaterialTag) (result gen.ResultInfo, err error) {
	return m.DO.Delete(models)
}

func (m *materialTagDo) withDO(do gen.Dao) *materialTagDo {
	m.DO = *do.(*gen.DO)
	return m
}
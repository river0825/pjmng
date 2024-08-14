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

func newMaterialContentSchema(db *gorm.DB, opts ...gen.DOOption) materialContentSchema {
	_materialContentSchema := materialContentSchema{}

	_materialContentSchema.materialContentSchemaDo.UseDB(db, opts...)
	_materialContentSchema.materialContentSchemaDo.UseModel(&table.MaterialContentSchema{})

	tableName := _materialContentSchema.materialContentSchemaDo.TableName()
	_materialContentSchema.ALL = field.NewAsterisk(tableName)
	_materialContentSchema.ID = field.NewInt32(tableName, "id")
	_materialContentSchema.JSONSchema = field.NewField(tableName, "json_schema")
	_materialContentSchema.PublishedAt = field.NewTime(tableName, "published_at")
	_materialContentSchema.CreatedBy = field.NewInt32(tableName, "created_by")
	_materialContentSchema.UpdatedBy = field.NewInt32(tableName, "updated_by")
	_materialContentSchema.CreatedAt = field.NewTime(tableName, "created_at")
	_materialContentSchema.UpdatedAt = field.NewTime(tableName, "updated_at")

	_materialContentSchema.fillFieldMap()

	return _materialContentSchema
}

type materialContentSchema struct {
	materialContentSchemaDo

	ALL         field.Asterisk
	ID          field.Int32
	JSONSchema  field.Field
	PublishedAt field.Time
	CreatedBy   field.Int32
	UpdatedBy   field.Int32
	CreatedAt   field.Time
	UpdatedAt   field.Time

	fieldMap map[string]field.Expr
}

func (m materialContentSchema) Table(newTableName string) *materialContentSchema {
	m.materialContentSchemaDo.UseTable(newTableName)
	return m.updateTableName(newTableName)
}

func (m materialContentSchema) As(alias string) *materialContentSchema {
	m.materialContentSchemaDo.DO = *(m.materialContentSchemaDo.As(alias).(*gen.DO))
	return m.updateTableName(alias)
}

func (m *materialContentSchema) updateTableName(table string) *materialContentSchema {
	m.ALL = field.NewAsterisk(table)
	m.ID = field.NewInt32(table, "id")
	m.JSONSchema = field.NewField(table, "json_schema")
	m.PublishedAt = field.NewTime(table, "published_at")
	m.CreatedBy = field.NewInt32(table, "created_by")
	m.UpdatedBy = field.NewInt32(table, "updated_by")
	m.CreatedAt = field.NewTime(table, "created_at")
	m.UpdatedAt = field.NewTime(table, "updated_at")

	m.fillFieldMap()

	return m
}

func (m *materialContentSchema) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := m.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (m *materialContentSchema) fillFieldMap() {
	m.fieldMap = make(map[string]field.Expr, 7)
	m.fieldMap["id"] = m.ID
	m.fieldMap["json_schema"] = m.JSONSchema
	m.fieldMap["published_at"] = m.PublishedAt
	m.fieldMap["created_by"] = m.CreatedBy
	m.fieldMap["updated_by"] = m.UpdatedBy
	m.fieldMap["created_at"] = m.CreatedAt
	m.fieldMap["updated_at"] = m.UpdatedAt
}

func (m materialContentSchema) clone(db *gorm.DB) materialContentSchema {
	m.materialContentSchemaDo.ReplaceConnPool(db.Statement.ConnPool)
	return m
}

func (m materialContentSchema) replaceDB(db *gorm.DB) materialContentSchema {
	m.materialContentSchemaDo.ReplaceDB(db)
	return m
}

type materialContentSchemaDo struct{ gen.DO }

type IMaterialContentSchemaDo interface {
	gen.SubQuery
	Debug() IMaterialContentSchemaDo
	WithContext(ctx context.Context) IMaterialContentSchemaDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IMaterialContentSchemaDo
	WriteDB() IMaterialContentSchemaDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IMaterialContentSchemaDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IMaterialContentSchemaDo
	Not(conds ...gen.Condition) IMaterialContentSchemaDo
	Or(conds ...gen.Condition) IMaterialContentSchemaDo
	Select(conds ...field.Expr) IMaterialContentSchemaDo
	Where(conds ...gen.Condition) IMaterialContentSchemaDo
	Order(conds ...field.Expr) IMaterialContentSchemaDo
	Distinct(cols ...field.Expr) IMaterialContentSchemaDo
	Omit(cols ...field.Expr) IMaterialContentSchemaDo
	Join(table schema.Tabler, on ...field.Expr) IMaterialContentSchemaDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IMaterialContentSchemaDo
	RightJoin(table schema.Tabler, on ...field.Expr) IMaterialContentSchemaDo
	Group(cols ...field.Expr) IMaterialContentSchemaDo
	Having(conds ...gen.Condition) IMaterialContentSchemaDo
	Limit(limit int) IMaterialContentSchemaDo
	Offset(offset int) IMaterialContentSchemaDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IMaterialContentSchemaDo
	Unscoped() IMaterialContentSchemaDo
	Create(values ...*table.MaterialContentSchema) error
	CreateInBatches(values []*table.MaterialContentSchema, batchSize int) error
	Save(values ...*table.MaterialContentSchema) error
	First() (*table.MaterialContentSchema, error)
	Take() (*table.MaterialContentSchema, error)
	Last() (*table.MaterialContentSchema, error)
	Find() ([]*table.MaterialContentSchema, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*table.MaterialContentSchema, err error)
	FindInBatches(result *[]*table.MaterialContentSchema, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*table.MaterialContentSchema) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IMaterialContentSchemaDo
	Assign(attrs ...field.AssignExpr) IMaterialContentSchemaDo
	Joins(fields ...field.RelationField) IMaterialContentSchemaDo
	Preload(fields ...field.RelationField) IMaterialContentSchemaDo
	FirstOrInit() (*table.MaterialContentSchema, error)
	FirstOrCreate() (*table.MaterialContentSchema, error)
	FindByPage(offset int, limit int) (result []*table.MaterialContentSchema, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IMaterialContentSchemaDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (m materialContentSchemaDo) Debug() IMaterialContentSchemaDo {
	return m.withDO(m.DO.Debug())
}

func (m materialContentSchemaDo) WithContext(ctx context.Context) IMaterialContentSchemaDo {
	return m.withDO(m.DO.WithContext(ctx))
}

func (m materialContentSchemaDo) ReadDB() IMaterialContentSchemaDo {
	return m.Clauses(dbresolver.Read)
}

func (m materialContentSchemaDo) WriteDB() IMaterialContentSchemaDo {
	return m.Clauses(dbresolver.Write)
}

func (m materialContentSchemaDo) Session(config *gorm.Session) IMaterialContentSchemaDo {
	return m.withDO(m.DO.Session(config))
}

func (m materialContentSchemaDo) Clauses(conds ...clause.Expression) IMaterialContentSchemaDo {
	return m.withDO(m.DO.Clauses(conds...))
}

func (m materialContentSchemaDo) Returning(value interface{}, columns ...string) IMaterialContentSchemaDo {
	return m.withDO(m.DO.Returning(value, columns...))
}

func (m materialContentSchemaDo) Not(conds ...gen.Condition) IMaterialContentSchemaDo {
	return m.withDO(m.DO.Not(conds...))
}

func (m materialContentSchemaDo) Or(conds ...gen.Condition) IMaterialContentSchemaDo {
	return m.withDO(m.DO.Or(conds...))
}

func (m materialContentSchemaDo) Select(conds ...field.Expr) IMaterialContentSchemaDo {
	return m.withDO(m.DO.Select(conds...))
}

func (m materialContentSchemaDo) Where(conds ...gen.Condition) IMaterialContentSchemaDo {
	return m.withDO(m.DO.Where(conds...))
}

func (m materialContentSchemaDo) Order(conds ...field.Expr) IMaterialContentSchemaDo {
	return m.withDO(m.DO.Order(conds...))
}

func (m materialContentSchemaDo) Distinct(cols ...field.Expr) IMaterialContentSchemaDo {
	return m.withDO(m.DO.Distinct(cols...))
}

func (m materialContentSchemaDo) Omit(cols ...field.Expr) IMaterialContentSchemaDo {
	return m.withDO(m.DO.Omit(cols...))
}

func (m materialContentSchemaDo) Join(table schema.Tabler, on ...field.Expr) IMaterialContentSchemaDo {
	return m.withDO(m.DO.Join(table, on...))
}

func (m materialContentSchemaDo) LeftJoin(table schema.Tabler, on ...field.Expr) IMaterialContentSchemaDo {
	return m.withDO(m.DO.LeftJoin(table, on...))
}

func (m materialContentSchemaDo) RightJoin(table schema.Tabler, on ...field.Expr) IMaterialContentSchemaDo {
	return m.withDO(m.DO.RightJoin(table, on...))
}

func (m materialContentSchemaDo) Group(cols ...field.Expr) IMaterialContentSchemaDo {
	return m.withDO(m.DO.Group(cols...))
}

func (m materialContentSchemaDo) Having(conds ...gen.Condition) IMaterialContentSchemaDo {
	return m.withDO(m.DO.Having(conds...))
}

func (m materialContentSchemaDo) Limit(limit int) IMaterialContentSchemaDo {
	return m.withDO(m.DO.Limit(limit))
}

func (m materialContentSchemaDo) Offset(offset int) IMaterialContentSchemaDo {
	return m.withDO(m.DO.Offset(offset))
}

func (m materialContentSchemaDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IMaterialContentSchemaDo {
	return m.withDO(m.DO.Scopes(funcs...))
}

func (m materialContentSchemaDo) Unscoped() IMaterialContentSchemaDo {
	return m.withDO(m.DO.Unscoped())
}

func (m materialContentSchemaDo) Create(values ...*table.MaterialContentSchema) error {
	if len(values) == 0 {
		return nil
	}
	return m.DO.Create(values)
}

func (m materialContentSchemaDo) CreateInBatches(values []*table.MaterialContentSchema, batchSize int) error {
	return m.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (m materialContentSchemaDo) Save(values ...*table.MaterialContentSchema) error {
	if len(values) == 0 {
		return nil
	}
	return m.DO.Save(values)
}

func (m materialContentSchemaDo) First() (*table.MaterialContentSchema, error) {
	if result, err := m.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*table.MaterialContentSchema), nil
	}
}

func (m materialContentSchemaDo) Take() (*table.MaterialContentSchema, error) {
	if result, err := m.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*table.MaterialContentSchema), nil
	}
}

func (m materialContentSchemaDo) Last() (*table.MaterialContentSchema, error) {
	if result, err := m.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*table.MaterialContentSchema), nil
	}
}

func (m materialContentSchemaDo) Find() ([]*table.MaterialContentSchema, error) {
	result, err := m.DO.Find()
	return result.([]*table.MaterialContentSchema), err
}

func (m materialContentSchemaDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*table.MaterialContentSchema, err error) {
	buf := make([]*table.MaterialContentSchema, 0, batchSize)
	err = m.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (m materialContentSchemaDo) FindInBatches(result *[]*table.MaterialContentSchema, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return m.DO.FindInBatches(result, batchSize, fc)
}

func (m materialContentSchemaDo) Attrs(attrs ...field.AssignExpr) IMaterialContentSchemaDo {
	return m.withDO(m.DO.Attrs(attrs...))
}

func (m materialContentSchemaDo) Assign(attrs ...field.AssignExpr) IMaterialContentSchemaDo {
	return m.withDO(m.DO.Assign(attrs...))
}

func (m materialContentSchemaDo) Joins(fields ...field.RelationField) IMaterialContentSchemaDo {
	for _, _f := range fields {
		m = *m.withDO(m.DO.Joins(_f))
	}
	return &m
}

func (m materialContentSchemaDo) Preload(fields ...field.RelationField) IMaterialContentSchemaDo {
	for _, _f := range fields {
		m = *m.withDO(m.DO.Preload(_f))
	}
	return &m
}

func (m materialContentSchemaDo) FirstOrInit() (*table.MaterialContentSchema, error) {
	if result, err := m.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*table.MaterialContentSchema), nil
	}
}

func (m materialContentSchemaDo) FirstOrCreate() (*table.MaterialContentSchema, error) {
	if result, err := m.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*table.MaterialContentSchema), nil
	}
}

func (m materialContentSchemaDo) FindByPage(offset int, limit int) (result []*table.MaterialContentSchema, count int64, err error) {
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

func (m materialContentSchemaDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = m.Count()
	if err != nil {
		return
	}

	err = m.Offset(offset).Limit(limit).Scan(result)
	return
}

func (m materialContentSchemaDo) Scan(result interface{}) (err error) {
	return m.DO.Scan(result)
}

func (m materialContentSchemaDo) Delete(models ...*table.MaterialContentSchema) (result gen.ResultInfo, err error) {
	return m.DO.Delete(models)
}

func (m *materialContentSchemaDo) withDO(do gen.Dao) *materialContentSchemaDo {
	m.DO = *do.(*gen.DO)
	return m
}

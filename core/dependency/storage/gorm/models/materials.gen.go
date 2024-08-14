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

func newMaterial(db *gorm.DB, opts ...gen.DOOption) material {
	_material := material{}

	_material.materialDo.UseDB(db, opts...)
	_material.materialDo.UseModel(&table.Material{})

	tableName := _material.materialDo.TableName()
	_material.ALL = field.NewAsterisk(tableName)
	_material.ID = field.NewInt32(tableName, "id")
	_material.Name = field.NewString(tableName, "name")
	_material.Tag = field.NewString(tableName, "tag")
	_material.User = field.NewInt32(tableName, "user")
	_material.Content = field.NewField(tableName, "content")
	_material.PublishedAt = field.NewTime(tableName, "published_at")
	_material.CreatedBy = field.NewInt32(tableName, "created_by")
	_material.UpdatedBy = field.NewInt32(tableName, "updated_by")
	_material.CreatedAt = field.NewTime(tableName, "created_at")
	_material.UpdatedAt = field.NewTime(tableName, "updated_at")
	_material.Element = field.NewInt32(tableName, "element")

	_material.fillFieldMap()

	return _material
}

type material struct {
	materialDo

	ALL         field.Asterisk
	ID          field.Int32
	Name        field.String
	Tag         field.String
	User        field.Int32
	Content     field.Field
	PublishedAt field.Time
	CreatedBy   field.Int32
	UpdatedBy   field.Int32
	CreatedAt   field.Time
	UpdatedAt   field.Time
	Element     field.Int32

	fieldMap map[string]field.Expr
}

func (m material) Table(newTableName string) *material {
	m.materialDo.UseTable(newTableName)
	return m.updateTableName(newTableName)
}

func (m material) As(alias string) *material {
	m.materialDo.DO = *(m.materialDo.As(alias).(*gen.DO))
	return m.updateTableName(alias)
}

func (m *material) updateTableName(table string) *material {
	m.ALL = field.NewAsterisk(table)
	m.ID = field.NewInt32(table, "id")
	m.Name = field.NewString(table, "name")
	m.Tag = field.NewString(table, "tag")
	m.User = field.NewInt32(table, "user")
	m.Content = field.NewField(table, "content")
	m.PublishedAt = field.NewTime(table, "published_at")
	m.CreatedBy = field.NewInt32(table, "created_by")
	m.UpdatedBy = field.NewInt32(table, "updated_by")
	m.CreatedAt = field.NewTime(table, "created_at")
	m.UpdatedAt = field.NewTime(table, "updated_at")
	m.Element = field.NewInt32(table, "element")

	m.fillFieldMap()

	return m
}

func (m *material) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := m.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (m *material) fillFieldMap() {
	m.fieldMap = make(map[string]field.Expr, 11)
	m.fieldMap["id"] = m.ID
	m.fieldMap["name"] = m.Name
	m.fieldMap["tag"] = m.Tag
	m.fieldMap["user"] = m.User
	m.fieldMap["content"] = m.Content
	m.fieldMap["published_at"] = m.PublishedAt
	m.fieldMap["created_by"] = m.CreatedBy
	m.fieldMap["updated_by"] = m.UpdatedBy
	m.fieldMap["created_at"] = m.CreatedAt
	m.fieldMap["updated_at"] = m.UpdatedAt
	m.fieldMap["element"] = m.Element
}

func (m material) clone(db *gorm.DB) material {
	m.materialDo.ReplaceConnPool(db.Statement.ConnPool)
	return m
}

func (m material) replaceDB(db *gorm.DB) material {
	m.materialDo.ReplaceDB(db)
	return m
}

type materialDo struct{ gen.DO }

type IMaterialDo interface {
	gen.SubQuery
	Debug() IMaterialDo
	WithContext(ctx context.Context) IMaterialDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IMaterialDo
	WriteDB() IMaterialDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IMaterialDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IMaterialDo
	Not(conds ...gen.Condition) IMaterialDo
	Or(conds ...gen.Condition) IMaterialDo
	Select(conds ...field.Expr) IMaterialDo
	Where(conds ...gen.Condition) IMaterialDo
	Order(conds ...field.Expr) IMaterialDo
	Distinct(cols ...field.Expr) IMaterialDo
	Omit(cols ...field.Expr) IMaterialDo
	Join(table schema.Tabler, on ...field.Expr) IMaterialDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IMaterialDo
	RightJoin(table schema.Tabler, on ...field.Expr) IMaterialDo
	Group(cols ...field.Expr) IMaterialDo
	Having(conds ...gen.Condition) IMaterialDo
	Limit(limit int) IMaterialDo
	Offset(offset int) IMaterialDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IMaterialDo
	Unscoped() IMaterialDo
	Create(values ...*table.Material) error
	CreateInBatches(values []*table.Material, batchSize int) error
	Save(values ...*table.Material) error
	First() (*table.Material, error)
	Take() (*table.Material, error)
	Last() (*table.Material, error)
	Find() ([]*table.Material, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*table.Material, err error)
	FindInBatches(result *[]*table.Material, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*table.Material) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IMaterialDo
	Assign(attrs ...field.AssignExpr) IMaterialDo
	Joins(fields ...field.RelationField) IMaterialDo
	Preload(fields ...field.RelationField) IMaterialDo
	FirstOrInit() (*table.Material, error)
	FirstOrCreate() (*table.Material, error)
	FindByPage(offset int, limit int) (result []*table.Material, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IMaterialDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (m materialDo) Debug() IMaterialDo {
	return m.withDO(m.DO.Debug())
}

func (m materialDo) WithContext(ctx context.Context) IMaterialDo {
	return m.withDO(m.DO.WithContext(ctx))
}

func (m materialDo) ReadDB() IMaterialDo {
	return m.Clauses(dbresolver.Read)
}

func (m materialDo) WriteDB() IMaterialDo {
	return m.Clauses(dbresolver.Write)
}

func (m materialDo) Session(config *gorm.Session) IMaterialDo {
	return m.withDO(m.DO.Session(config))
}

func (m materialDo) Clauses(conds ...clause.Expression) IMaterialDo {
	return m.withDO(m.DO.Clauses(conds...))
}

func (m materialDo) Returning(value interface{}, columns ...string) IMaterialDo {
	return m.withDO(m.DO.Returning(value, columns...))
}

func (m materialDo) Not(conds ...gen.Condition) IMaterialDo {
	return m.withDO(m.DO.Not(conds...))
}

func (m materialDo) Or(conds ...gen.Condition) IMaterialDo {
	return m.withDO(m.DO.Or(conds...))
}

func (m materialDo) Select(conds ...field.Expr) IMaterialDo {
	return m.withDO(m.DO.Select(conds...))
}

func (m materialDo) Where(conds ...gen.Condition) IMaterialDo {
	return m.withDO(m.DO.Where(conds...))
}

func (m materialDo) Order(conds ...field.Expr) IMaterialDo {
	return m.withDO(m.DO.Order(conds...))
}

func (m materialDo) Distinct(cols ...field.Expr) IMaterialDo {
	return m.withDO(m.DO.Distinct(cols...))
}

func (m materialDo) Omit(cols ...field.Expr) IMaterialDo {
	return m.withDO(m.DO.Omit(cols...))
}

func (m materialDo) Join(table schema.Tabler, on ...field.Expr) IMaterialDo {
	return m.withDO(m.DO.Join(table, on...))
}

func (m materialDo) LeftJoin(table schema.Tabler, on ...field.Expr) IMaterialDo {
	return m.withDO(m.DO.LeftJoin(table, on...))
}

func (m materialDo) RightJoin(table schema.Tabler, on ...field.Expr) IMaterialDo {
	return m.withDO(m.DO.RightJoin(table, on...))
}

func (m materialDo) Group(cols ...field.Expr) IMaterialDo {
	return m.withDO(m.DO.Group(cols...))
}

func (m materialDo) Having(conds ...gen.Condition) IMaterialDo {
	return m.withDO(m.DO.Having(conds...))
}

func (m materialDo) Limit(limit int) IMaterialDo {
	return m.withDO(m.DO.Limit(limit))
}

func (m materialDo) Offset(offset int) IMaterialDo {
	return m.withDO(m.DO.Offset(offset))
}

func (m materialDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IMaterialDo {
	return m.withDO(m.DO.Scopes(funcs...))
}

func (m materialDo) Unscoped() IMaterialDo {
	return m.withDO(m.DO.Unscoped())
}

func (m materialDo) Create(values ...*table.Material) error {
	if len(values) == 0 {
		return nil
	}
	return m.DO.Create(values)
}

func (m materialDo) CreateInBatches(values []*table.Material, batchSize int) error {
	return m.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (m materialDo) Save(values ...*table.Material) error {
	if len(values) == 0 {
		return nil
	}
	return m.DO.Save(values)
}

func (m materialDo) First() (*table.Material, error) {
	if result, err := m.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*table.Material), nil
	}
}

func (m materialDo) Take() (*table.Material, error) {
	if result, err := m.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*table.Material), nil
	}
}

func (m materialDo) Last() (*table.Material, error) {
	if result, err := m.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*table.Material), nil
	}
}

func (m materialDo) Find() ([]*table.Material, error) {
	result, err := m.DO.Find()
	return result.([]*table.Material), err
}

func (m materialDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*table.Material, err error) {
	buf := make([]*table.Material, 0, batchSize)
	err = m.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (m materialDo) FindInBatches(result *[]*table.Material, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return m.DO.FindInBatches(result, batchSize, fc)
}

func (m materialDo) Attrs(attrs ...field.AssignExpr) IMaterialDo {
	return m.withDO(m.DO.Attrs(attrs...))
}

func (m materialDo) Assign(attrs ...field.AssignExpr) IMaterialDo {
	return m.withDO(m.DO.Assign(attrs...))
}

func (m materialDo) Joins(fields ...field.RelationField) IMaterialDo {
	for _, _f := range fields {
		m = *m.withDO(m.DO.Joins(_f))
	}
	return &m
}

func (m materialDo) Preload(fields ...field.RelationField) IMaterialDo {
	for _, _f := range fields {
		m = *m.withDO(m.DO.Preload(_f))
	}
	return &m
}

func (m materialDo) FirstOrInit() (*table.Material, error) {
	if result, err := m.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*table.Material), nil
	}
}

func (m materialDo) FirstOrCreate() (*table.Material, error) {
	if result, err := m.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*table.Material), nil
	}
}

func (m materialDo) FindByPage(offset int, limit int) (result []*table.Material, count int64, err error) {
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

func (m materialDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = m.Count()
	if err != nil {
		return
	}

	err = m.Offset(offset).Limit(limit).Scan(result)
	return
}

func (m materialDo) Scan(result interface{}) (err error) {
	return m.DO.Scan(result)
}

func (m materialDo) Delete(models ...*table.Material) (result gen.ResultInfo, err error) {
	return m.DO.Delete(models)
}

func (m *materialDo) withDO(do gen.Dao) *materialDo {
	m.DO = *do.(*gen.DO)
	return m
}

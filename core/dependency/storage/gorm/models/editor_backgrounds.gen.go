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

func newEditorBackground(db *gorm.DB, opts ...gen.DOOption) editorBackground {
	_editorBackground := editorBackground{}

	_editorBackground.editorBackgroundDo.UseDB(db, opts...)
	_editorBackground.editorBackgroundDo.UseModel(&table.EditorBackground{})

	tableName := _editorBackground.editorBackgroundDo.TableName()
	_editorBackground.ALL = field.NewAsterisk(tableName)
	_editorBackground.ID = field.NewInt32(tableName, "id")
	_editorBackground.PublishedAt = field.NewTime(tableName, "published_at")
	_editorBackground.CreatedBy = field.NewInt32(tableName, "created_by")
	_editorBackground.UpdatedBy = field.NewInt32(tableName, "updated_by")
	_editorBackground.CreatedAt = field.NewTime(tableName, "created_at")
	_editorBackground.UpdatedAt = field.NewTime(tableName, "updated_at")

	_editorBackground.fillFieldMap()

	return _editorBackground
}

type editorBackground struct {
	editorBackgroundDo

	ALL         field.Asterisk
	ID          field.Int32
	PublishedAt field.Time
	CreatedBy   field.Int32
	UpdatedBy   field.Int32
	CreatedAt   field.Time
	UpdatedAt   field.Time

	fieldMap map[string]field.Expr
}

func (e editorBackground) Table(newTableName string) *editorBackground {
	e.editorBackgroundDo.UseTable(newTableName)
	return e.updateTableName(newTableName)
}

func (e editorBackground) As(alias string) *editorBackground {
	e.editorBackgroundDo.DO = *(e.editorBackgroundDo.As(alias).(*gen.DO))
	return e.updateTableName(alias)
}

func (e *editorBackground) updateTableName(table string) *editorBackground {
	e.ALL = field.NewAsterisk(table)
	e.ID = field.NewInt32(table, "id")
	e.PublishedAt = field.NewTime(table, "published_at")
	e.CreatedBy = field.NewInt32(table, "created_by")
	e.UpdatedBy = field.NewInt32(table, "updated_by")
	e.CreatedAt = field.NewTime(table, "created_at")
	e.UpdatedAt = field.NewTime(table, "updated_at")

	e.fillFieldMap()

	return e
}

func (e *editorBackground) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := e.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (e *editorBackground) fillFieldMap() {
	e.fieldMap = make(map[string]field.Expr, 6)
	e.fieldMap["id"] = e.ID
	e.fieldMap["published_at"] = e.PublishedAt
	e.fieldMap["created_by"] = e.CreatedBy
	e.fieldMap["updated_by"] = e.UpdatedBy
	e.fieldMap["created_at"] = e.CreatedAt
	e.fieldMap["updated_at"] = e.UpdatedAt
}

func (e editorBackground) clone(db *gorm.DB) editorBackground {
	e.editorBackgroundDo.ReplaceConnPool(db.Statement.ConnPool)
	return e
}

func (e editorBackground) replaceDB(db *gorm.DB) editorBackground {
	e.editorBackgroundDo.ReplaceDB(db)
	return e
}

type editorBackgroundDo struct{ gen.DO }

type IEditorBackgroundDo interface {
	gen.SubQuery
	Debug() IEditorBackgroundDo
	WithContext(ctx context.Context) IEditorBackgroundDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IEditorBackgroundDo
	WriteDB() IEditorBackgroundDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IEditorBackgroundDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IEditorBackgroundDo
	Not(conds ...gen.Condition) IEditorBackgroundDo
	Or(conds ...gen.Condition) IEditorBackgroundDo
	Select(conds ...field.Expr) IEditorBackgroundDo
	Where(conds ...gen.Condition) IEditorBackgroundDo
	Order(conds ...field.Expr) IEditorBackgroundDo
	Distinct(cols ...field.Expr) IEditorBackgroundDo
	Omit(cols ...field.Expr) IEditorBackgroundDo
	Join(table schema.Tabler, on ...field.Expr) IEditorBackgroundDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IEditorBackgroundDo
	RightJoin(table schema.Tabler, on ...field.Expr) IEditorBackgroundDo
	Group(cols ...field.Expr) IEditorBackgroundDo
	Having(conds ...gen.Condition) IEditorBackgroundDo
	Limit(limit int) IEditorBackgroundDo
	Offset(offset int) IEditorBackgroundDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IEditorBackgroundDo
	Unscoped() IEditorBackgroundDo
	Create(values ...*table.EditorBackground) error
	CreateInBatches(values []*table.EditorBackground, batchSize int) error
	Save(values ...*table.EditorBackground) error
	First() (*table.EditorBackground, error)
	Take() (*table.EditorBackground, error)
	Last() (*table.EditorBackground, error)
	Find() ([]*table.EditorBackground, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*table.EditorBackground, err error)
	FindInBatches(result *[]*table.EditorBackground, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*table.EditorBackground) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IEditorBackgroundDo
	Assign(attrs ...field.AssignExpr) IEditorBackgroundDo
	Joins(fields ...field.RelationField) IEditorBackgroundDo
	Preload(fields ...field.RelationField) IEditorBackgroundDo
	FirstOrInit() (*table.EditorBackground, error)
	FirstOrCreate() (*table.EditorBackground, error)
	FindByPage(offset int, limit int) (result []*table.EditorBackground, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IEditorBackgroundDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (e editorBackgroundDo) Debug() IEditorBackgroundDo {
	return e.withDO(e.DO.Debug())
}

func (e editorBackgroundDo) WithContext(ctx context.Context) IEditorBackgroundDo {
	return e.withDO(e.DO.WithContext(ctx))
}

func (e editorBackgroundDo) ReadDB() IEditorBackgroundDo {
	return e.Clauses(dbresolver.Read)
}

func (e editorBackgroundDo) WriteDB() IEditorBackgroundDo {
	return e.Clauses(dbresolver.Write)
}

func (e editorBackgroundDo) Session(config *gorm.Session) IEditorBackgroundDo {
	return e.withDO(e.DO.Session(config))
}

func (e editorBackgroundDo) Clauses(conds ...clause.Expression) IEditorBackgroundDo {
	return e.withDO(e.DO.Clauses(conds...))
}

func (e editorBackgroundDo) Returning(value interface{}, columns ...string) IEditorBackgroundDo {
	return e.withDO(e.DO.Returning(value, columns...))
}

func (e editorBackgroundDo) Not(conds ...gen.Condition) IEditorBackgroundDo {
	return e.withDO(e.DO.Not(conds...))
}

func (e editorBackgroundDo) Or(conds ...gen.Condition) IEditorBackgroundDo {
	return e.withDO(e.DO.Or(conds...))
}

func (e editorBackgroundDo) Select(conds ...field.Expr) IEditorBackgroundDo {
	return e.withDO(e.DO.Select(conds...))
}

func (e editorBackgroundDo) Where(conds ...gen.Condition) IEditorBackgroundDo {
	return e.withDO(e.DO.Where(conds...))
}

func (e editorBackgroundDo) Order(conds ...field.Expr) IEditorBackgroundDo {
	return e.withDO(e.DO.Order(conds...))
}

func (e editorBackgroundDo) Distinct(cols ...field.Expr) IEditorBackgroundDo {
	return e.withDO(e.DO.Distinct(cols...))
}

func (e editorBackgroundDo) Omit(cols ...field.Expr) IEditorBackgroundDo {
	return e.withDO(e.DO.Omit(cols...))
}

func (e editorBackgroundDo) Join(table schema.Tabler, on ...field.Expr) IEditorBackgroundDo {
	return e.withDO(e.DO.Join(table, on...))
}

func (e editorBackgroundDo) LeftJoin(table schema.Tabler, on ...field.Expr) IEditorBackgroundDo {
	return e.withDO(e.DO.LeftJoin(table, on...))
}

func (e editorBackgroundDo) RightJoin(table schema.Tabler, on ...field.Expr) IEditorBackgroundDo {
	return e.withDO(e.DO.RightJoin(table, on...))
}

func (e editorBackgroundDo) Group(cols ...field.Expr) IEditorBackgroundDo {
	return e.withDO(e.DO.Group(cols...))
}

func (e editorBackgroundDo) Having(conds ...gen.Condition) IEditorBackgroundDo {
	return e.withDO(e.DO.Having(conds...))
}

func (e editorBackgroundDo) Limit(limit int) IEditorBackgroundDo {
	return e.withDO(e.DO.Limit(limit))
}

func (e editorBackgroundDo) Offset(offset int) IEditorBackgroundDo {
	return e.withDO(e.DO.Offset(offset))
}

func (e editorBackgroundDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IEditorBackgroundDo {
	return e.withDO(e.DO.Scopes(funcs...))
}

func (e editorBackgroundDo) Unscoped() IEditorBackgroundDo {
	return e.withDO(e.DO.Unscoped())
}

func (e editorBackgroundDo) Create(values ...*table.EditorBackground) error {
	if len(values) == 0 {
		return nil
	}
	return e.DO.Create(values)
}

func (e editorBackgroundDo) CreateInBatches(values []*table.EditorBackground, batchSize int) error {
	return e.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (e editorBackgroundDo) Save(values ...*table.EditorBackground) error {
	if len(values) == 0 {
		return nil
	}
	return e.DO.Save(values)
}

func (e editorBackgroundDo) First() (*table.EditorBackground, error) {
	if result, err := e.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*table.EditorBackground), nil
	}
}

func (e editorBackgroundDo) Take() (*table.EditorBackground, error) {
	if result, err := e.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*table.EditorBackground), nil
	}
}

func (e editorBackgroundDo) Last() (*table.EditorBackground, error) {
	if result, err := e.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*table.EditorBackground), nil
	}
}

func (e editorBackgroundDo) Find() ([]*table.EditorBackground, error) {
	result, err := e.DO.Find()
	return result.([]*table.EditorBackground), err
}

func (e editorBackgroundDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*table.EditorBackground, err error) {
	buf := make([]*table.EditorBackground, 0, batchSize)
	err = e.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (e editorBackgroundDo) FindInBatches(result *[]*table.EditorBackground, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return e.DO.FindInBatches(result, batchSize, fc)
}

func (e editorBackgroundDo) Attrs(attrs ...field.AssignExpr) IEditorBackgroundDo {
	return e.withDO(e.DO.Attrs(attrs...))
}

func (e editorBackgroundDo) Assign(attrs ...field.AssignExpr) IEditorBackgroundDo {
	return e.withDO(e.DO.Assign(attrs...))
}

func (e editorBackgroundDo) Joins(fields ...field.RelationField) IEditorBackgroundDo {
	for _, _f := range fields {
		e = *e.withDO(e.DO.Joins(_f))
	}
	return &e
}

func (e editorBackgroundDo) Preload(fields ...field.RelationField) IEditorBackgroundDo {
	for _, _f := range fields {
		e = *e.withDO(e.DO.Preload(_f))
	}
	return &e
}

func (e editorBackgroundDo) FirstOrInit() (*table.EditorBackground, error) {
	if result, err := e.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*table.EditorBackground), nil
	}
}

func (e editorBackgroundDo) FirstOrCreate() (*table.EditorBackground, error) {
	if result, err := e.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*table.EditorBackground), nil
	}
}

func (e editorBackgroundDo) FindByPage(offset int, limit int) (result []*table.EditorBackground, count int64, err error) {
	result, err = e.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = e.Offset(-1).Limit(-1).Count()
	return
}

func (e editorBackgroundDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = e.Count()
	if err != nil {
		return
	}

	err = e.Offset(offset).Limit(limit).Scan(result)
	return
}

func (e editorBackgroundDo) Scan(result interface{}) (err error) {
	return e.DO.Scan(result)
}

func (e editorBackgroundDo) Delete(models ...*table.EditorBackground) (result gen.ResultInfo, err error) {
	return e.DO.Delete(models)
}

func (e *editorBackgroundDo) withDO(do gen.Dao) *editorBackgroundDo {
	e.DO = *do.(*gen.DO)
	return e
}
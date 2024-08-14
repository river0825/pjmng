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

func newEditorCategoryAudio(db *gorm.DB, opts ...gen.DOOption) editorCategoryAudio {
	_editorCategoryAudio := editorCategoryAudio{}

	_editorCategoryAudio.editorCategoryAudioDo.UseDB(db, opts...)
	_editorCategoryAudio.editorCategoryAudioDo.UseModel(&table.EditorCategoryAudio{})

	tableName := _editorCategoryAudio.editorCategoryAudioDo.TableName()
	_editorCategoryAudio.ALL = field.NewAsterisk(tableName)
	_editorCategoryAudio.ID = field.NewInt32(tableName, "id")
	_editorCategoryAudio.Category = field.NewString(tableName, "category")
	_editorCategoryAudio.PublishedAt = field.NewTime(tableName, "published_at")
	_editorCategoryAudio.CreatedBy = field.NewInt32(tableName, "created_by")
	_editorCategoryAudio.UpdatedBy = field.NewInt32(tableName, "updated_by")
	_editorCategoryAudio.CreatedAt = field.NewTime(tableName, "created_at")
	_editorCategoryAudio.UpdatedAt = field.NewTime(tableName, "updated_at")

	_editorCategoryAudio.fillFieldMap()

	return _editorCategoryAudio
}

type editorCategoryAudio struct {
	editorCategoryAudioDo

	ALL         field.Asterisk
	ID          field.Int32
	Category    field.String
	PublishedAt field.Time
	CreatedBy   field.Int32
	UpdatedBy   field.Int32
	CreatedAt   field.Time
	UpdatedAt   field.Time

	fieldMap map[string]field.Expr
}

func (e editorCategoryAudio) Table(newTableName string) *editorCategoryAudio {
	e.editorCategoryAudioDo.UseTable(newTableName)
	return e.updateTableName(newTableName)
}

func (e editorCategoryAudio) As(alias string) *editorCategoryAudio {
	e.editorCategoryAudioDo.DO = *(e.editorCategoryAudioDo.As(alias).(*gen.DO))
	return e.updateTableName(alias)
}

func (e *editorCategoryAudio) updateTableName(table string) *editorCategoryAudio {
	e.ALL = field.NewAsterisk(table)
	e.ID = field.NewInt32(table, "id")
	e.Category = field.NewString(table, "category")
	e.PublishedAt = field.NewTime(table, "published_at")
	e.CreatedBy = field.NewInt32(table, "created_by")
	e.UpdatedBy = field.NewInt32(table, "updated_by")
	e.CreatedAt = field.NewTime(table, "created_at")
	e.UpdatedAt = field.NewTime(table, "updated_at")

	e.fillFieldMap()

	return e
}

func (e *editorCategoryAudio) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := e.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (e *editorCategoryAudio) fillFieldMap() {
	e.fieldMap = make(map[string]field.Expr, 7)
	e.fieldMap["id"] = e.ID
	e.fieldMap["category"] = e.Category
	e.fieldMap["published_at"] = e.PublishedAt
	e.fieldMap["created_by"] = e.CreatedBy
	e.fieldMap["updated_by"] = e.UpdatedBy
	e.fieldMap["created_at"] = e.CreatedAt
	e.fieldMap["updated_at"] = e.UpdatedAt
}

func (e editorCategoryAudio) clone(db *gorm.DB) editorCategoryAudio {
	e.editorCategoryAudioDo.ReplaceConnPool(db.Statement.ConnPool)
	return e
}

func (e editorCategoryAudio) replaceDB(db *gorm.DB) editorCategoryAudio {
	e.editorCategoryAudioDo.ReplaceDB(db)
	return e
}

type editorCategoryAudioDo struct{ gen.DO }

type IEditorCategoryAudioDo interface {
	gen.SubQuery
	Debug() IEditorCategoryAudioDo
	WithContext(ctx context.Context) IEditorCategoryAudioDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IEditorCategoryAudioDo
	WriteDB() IEditorCategoryAudioDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IEditorCategoryAudioDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IEditorCategoryAudioDo
	Not(conds ...gen.Condition) IEditorCategoryAudioDo
	Or(conds ...gen.Condition) IEditorCategoryAudioDo
	Select(conds ...field.Expr) IEditorCategoryAudioDo
	Where(conds ...gen.Condition) IEditorCategoryAudioDo
	Order(conds ...field.Expr) IEditorCategoryAudioDo
	Distinct(cols ...field.Expr) IEditorCategoryAudioDo
	Omit(cols ...field.Expr) IEditorCategoryAudioDo
	Join(table schema.Tabler, on ...field.Expr) IEditorCategoryAudioDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IEditorCategoryAudioDo
	RightJoin(table schema.Tabler, on ...field.Expr) IEditorCategoryAudioDo
	Group(cols ...field.Expr) IEditorCategoryAudioDo
	Having(conds ...gen.Condition) IEditorCategoryAudioDo
	Limit(limit int) IEditorCategoryAudioDo
	Offset(offset int) IEditorCategoryAudioDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IEditorCategoryAudioDo
	Unscoped() IEditorCategoryAudioDo
	Create(values ...*table.EditorCategoryAudio) error
	CreateInBatches(values []*table.EditorCategoryAudio, batchSize int) error
	Save(values ...*table.EditorCategoryAudio) error
	First() (*table.EditorCategoryAudio, error)
	Take() (*table.EditorCategoryAudio, error)
	Last() (*table.EditorCategoryAudio, error)
	Find() ([]*table.EditorCategoryAudio, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*table.EditorCategoryAudio, err error)
	FindInBatches(result *[]*table.EditorCategoryAudio, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*table.EditorCategoryAudio) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IEditorCategoryAudioDo
	Assign(attrs ...field.AssignExpr) IEditorCategoryAudioDo
	Joins(fields ...field.RelationField) IEditorCategoryAudioDo
	Preload(fields ...field.RelationField) IEditorCategoryAudioDo
	FirstOrInit() (*table.EditorCategoryAudio, error)
	FirstOrCreate() (*table.EditorCategoryAudio, error)
	FindByPage(offset int, limit int) (result []*table.EditorCategoryAudio, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IEditorCategoryAudioDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (e editorCategoryAudioDo) Debug() IEditorCategoryAudioDo {
	return e.withDO(e.DO.Debug())
}

func (e editorCategoryAudioDo) WithContext(ctx context.Context) IEditorCategoryAudioDo {
	return e.withDO(e.DO.WithContext(ctx))
}

func (e editorCategoryAudioDo) ReadDB() IEditorCategoryAudioDo {
	return e.Clauses(dbresolver.Read)
}

func (e editorCategoryAudioDo) WriteDB() IEditorCategoryAudioDo {
	return e.Clauses(dbresolver.Write)
}

func (e editorCategoryAudioDo) Session(config *gorm.Session) IEditorCategoryAudioDo {
	return e.withDO(e.DO.Session(config))
}

func (e editorCategoryAudioDo) Clauses(conds ...clause.Expression) IEditorCategoryAudioDo {
	return e.withDO(e.DO.Clauses(conds...))
}

func (e editorCategoryAudioDo) Returning(value interface{}, columns ...string) IEditorCategoryAudioDo {
	return e.withDO(e.DO.Returning(value, columns...))
}

func (e editorCategoryAudioDo) Not(conds ...gen.Condition) IEditorCategoryAudioDo {
	return e.withDO(e.DO.Not(conds...))
}

func (e editorCategoryAudioDo) Or(conds ...gen.Condition) IEditorCategoryAudioDo {
	return e.withDO(e.DO.Or(conds...))
}

func (e editorCategoryAudioDo) Select(conds ...field.Expr) IEditorCategoryAudioDo {
	return e.withDO(e.DO.Select(conds...))
}

func (e editorCategoryAudioDo) Where(conds ...gen.Condition) IEditorCategoryAudioDo {
	return e.withDO(e.DO.Where(conds...))
}

func (e editorCategoryAudioDo) Order(conds ...field.Expr) IEditorCategoryAudioDo {
	return e.withDO(e.DO.Order(conds...))
}

func (e editorCategoryAudioDo) Distinct(cols ...field.Expr) IEditorCategoryAudioDo {
	return e.withDO(e.DO.Distinct(cols...))
}

func (e editorCategoryAudioDo) Omit(cols ...field.Expr) IEditorCategoryAudioDo {
	return e.withDO(e.DO.Omit(cols...))
}

func (e editorCategoryAudioDo) Join(table schema.Tabler, on ...field.Expr) IEditorCategoryAudioDo {
	return e.withDO(e.DO.Join(table, on...))
}

func (e editorCategoryAudioDo) LeftJoin(table schema.Tabler, on ...field.Expr) IEditorCategoryAudioDo {
	return e.withDO(e.DO.LeftJoin(table, on...))
}

func (e editorCategoryAudioDo) RightJoin(table schema.Tabler, on ...field.Expr) IEditorCategoryAudioDo {
	return e.withDO(e.DO.RightJoin(table, on...))
}

func (e editorCategoryAudioDo) Group(cols ...field.Expr) IEditorCategoryAudioDo {
	return e.withDO(e.DO.Group(cols...))
}

func (e editorCategoryAudioDo) Having(conds ...gen.Condition) IEditorCategoryAudioDo {
	return e.withDO(e.DO.Having(conds...))
}

func (e editorCategoryAudioDo) Limit(limit int) IEditorCategoryAudioDo {
	return e.withDO(e.DO.Limit(limit))
}

func (e editorCategoryAudioDo) Offset(offset int) IEditorCategoryAudioDo {
	return e.withDO(e.DO.Offset(offset))
}

func (e editorCategoryAudioDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IEditorCategoryAudioDo {
	return e.withDO(e.DO.Scopes(funcs...))
}

func (e editorCategoryAudioDo) Unscoped() IEditorCategoryAudioDo {
	return e.withDO(e.DO.Unscoped())
}

func (e editorCategoryAudioDo) Create(values ...*table.EditorCategoryAudio) error {
	if len(values) == 0 {
		return nil
	}
	return e.DO.Create(values)
}

func (e editorCategoryAudioDo) CreateInBatches(values []*table.EditorCategoryAudio, batchSize int) error {
	return e.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (e editorCategoryAudioDo) Save(values ...*table.EditorCategoryAudio) error {
	if len(values) == 0 {
		return nil
	}
	return e.DO.Save(values)
}

func (e editorCategoryAudioDo) First() (*table.EditorCategoryAudio, error) {
	if result, err := e.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*table.EditorCategoryAudio), nil
	}
}

func (e editorCategoryAudioDo) Take() (*table.EditorCategoryAudio, error) {
	if result, err := e.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*table.EditorCategoryAudio), nil
	}
}

func (e editorCategoryAudioDo) Last() (*table.EditorCategoryAudio, error) {
	if result, err := e.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*table.EditorCategoryAudio), nil
	}
}

func (e editorCategoryAudioDo) Find() ([]*table.EditorCategoryAudio, error) {
	result, err := e.DO.Find()
	return result.([]*table.EditorCategoryAudio), err
}

func (e editorCategoryAudioDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*table.EditorCategoryAudio, err error) {
	buf := make([]*table.EditorCategoryAudio, 0, batchSize)
	err = e.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (e editorCategoryAudioDo) FindInBatches(result *[]*table.EditorCategoryAudio, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return e.DO.FindInBatches(result, batchSize, fc)
}

func (e editorCategoryAudioDo) Attrs(attrs ...field.AssignExpr) IEditorCategoryAudioDo {
	return e.withDO(e.DO.Attrs(attrs...))
}

func (e editorCategoryAudioDo) Assign(attrs ...field.AssignExpr) IEditorCategoryAudioDo {
	return e.withDO(e.DO.Assign(attrs...))
}

func (e editorCategoryAudioDo) Joins(fields ...field.RelationField) IEditorCategoryAudioDo {
	for _, _f := range fields {
		e = *e.withDO(e.DO.Joins(_f))
	}
	return &e
}

func (e editorCategoryAudioDo) Preload(fields ...field.RelationField) IEditorCategoryAudioDo {
	for _, _f := range fields {
		e = *e.withDO(e.DO.Preload(_f))
	}
	return &e
}

func (e editorCategoryAudioDo) FirstOrInit() (*table.EditorCategoryAudio, error) {
	if result, err := e.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*table.EditorCategoryAudio), nil
	}
}

func (e editorCategoryAudioDo) FirstOrCreate() (*table.EditorCategoryAudio, error) {
	if result, err := e.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*table.EditorCategoryAudio), nil
	}
}

func (e editorCategoryAudioDo) FindByPage(offset int, limit int) (result []*table.EditorCategoryAudio, count int64, err error) {
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

func (e editorCategoryAudioDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = e.Count()
	if err != nil {
		return
	}

	err = e.Offset(offset).Limit(limit).Scan(result)
	return
}

func (e editorCategoryAudioDo) Scan(result interface{}) (err error) {
	return e.DO.Scan(result)
}

func (e editorCategoryAudioDo) Delete(models ...*table.EditorCategoryAudio) (result gen.ResultInfo, err error) {
	return e.DO.Delete(models)
}

func (e *editorCategoryAudioDo) withDO(do gen.Dao) *editorCategoryAudioDo {
	e.DO = *do.(*gen.DO)
	return e
}
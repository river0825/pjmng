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

func newLegalDocument(db *gorm.DB, opts ...gen.DOOption) legalDocument {
	_legalDocument := legalDocument{}

	_legalDocument.legalDocumentDo.UseDB(db, opts...)
	_legalDocument.legalDocumentDo.UseModel(&table.LegalDocument{})

	tableName := _legalDocument.legalDocumentDo.TableName()
	_legalDocument.ALL = field.NewAsterisk(tableName)
	_legalDocument.ID = field.NewInt32(tableName, "id")
	_legalDocument.Name = field.NewString(tableName, "name")
	_legalDocument.Version = field.NewString(tableName, "version")
	_legalDocument.Content = field.NewString(tableName, "content")
	_legalDocument.Type = field.NewString(tableName, "type")
	_legalDocument.PublishedAt = field.NewTime(tableName, "published_at")
	_legalDocument.CreatedBy = field.NewInt32(tableName, "created_by")
	_legalDocument.UpdatedBy = field.NewInt32(tableName, "updated_by")
	_legalDocument.CreatedAt = field.NewTime(tableName, "created_at")
	_legalDocument.UpdatedAt = field.NewTime(tableName, "updated_at")

	_legalDocument.fillFieldMap()

	return _legalDocument
}

type legalDocument struct {
	legalDocumentDo

	ALL         field.Asterisk
	ID          field.Int32
	Name        field.String
	Version     field.String
	Content     field.String
	Type        field.String
	PublishedAt field.Time
	CreatedBy   field.Int32
	UpdatedBy   field.Int32
	CreatedAt   field.Time
	UpdatedAt   field.Time

	fieldMap map[string]field.Expr
}

func (l legalDocument) Table(newTableName string) *legalDocument {
	l.legalDocumentDo.UseTable(newTableName)
	return l.updateTableName(newTableName)
}

func (l legalDocument) As(alias string) *legalDocument {
	l.legalDocumentDo.DO = *(l.legalDocumentDo.As(alias).(*gen.DO))
	return l.updateTableName(alias)
}

func (l *legalDocument) updateTableName(table string) *legalDocument {
	l.ALL = field.NewAsterisk(table)
	l.ID = field.NewInt32(table, "id")
	l.Name = field.NewString(table, "name")
	l.Version = field.NewString(table, "version")
	l.Content = field.NewString(table, "content")
	l.Type = field.NewString(table, "type")
	l.PublishedAt = field.NewTime(table, "published_at")
	l.CreatedBy = field.NewInt32(table, "created_by")
	l.UpdatedBy = field.NewInt32(table, "updated_by")
	l.CreatedAt = field.NewTime(table, "created_at")
	l.UpdatedAt = field.NewTime(table, "updated_at")

	l.fillFieldMap()

	return l
}

func (l *legalDocument) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := l.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (l *legalDocument) fillFieldMap() {
	l.fieldMap = make(map[string]field.Expr, 10)
	l.fieldMap["id"] = l.ID
	l.fieldMap["name"] = l.Name
	l.fieldMap["version"] = l.Version
	l.fieldMap["content"] = l.Content
	l.fieldMap["type"] = l.Type
	l.fieldMap["published_at"] = l.PublishedAt
	l.fieldMap["created_by"] = l.CreatedBy
	l.fieldMap["updated_by"] = l.UpdatedBy
	l.fieldMap["created_at"] = l.CreatedAt
	l.fieldMap["updated_at"] = l.UpdatedAt
}

func (l legalDocument) clone(db *gorm.DB) legalDocument {
	l.legalDocumentDo.ReplaceConnPool(db.Statement.ConnPool)
	return l
}

func (l legalDocument) replaceDB(db *gorm.DB) legalDocument {
	l.legalDocumentDo.ReplaceDB(db)
	return l
}

type legalDocumentDo struct{ gen.DO }

type ILegalDocumentDo interface {
	gen.SubQuery
	Debug() ILegalDocumentDo
	WithContext(ctx context.Context) ILegalDocumentDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ILegalDocumentDo
	WriteDB() ILegalDocumentDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ILegalDocumentDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ILegalDocumentDo
	Not(conds ...gen.Condition) ILegalDocumentDo
	Or(conds ...gen.Condition) ILegalDocumentDo
	Select(conds ...field.Expr) ILegalDocumentDo
	Where(conds ...gen.Condition) ILegalDocumentDo
	Order(conds ...field.Expr) ILegalDocumentDo
	Distinct(cols ...field.Expr) ILegalDocumentDo
	Omit(cols ...field.Expr) ILegalDocumentDo
	Join(table schema.Tabler, on ...field.Expr) ILegalDocumentDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ILegalDocumentDo
	RightJoin(table schema.Tabler, on ...field.Expr) ILegalDocumentDo
	Group(cols ...field.Expr) ILegalDocumentDo
	Having(conds ...gen.Condition) ILegalDocumentDo
	Limit(limit int) ILegalDocumentDo
	Offset(offset int) ILegalDocumentDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ILegalDocumentDo
	Unscoped() ILegalDocumentDo
	Create(values ...*table.LegalDocument) error
	CreateInBatches(values []*table.LegalDocument, batchSize int) error
	Save(values ...*table.LegalDocument) error
	First() (*table.LegalDocument, error)
	Take() (*table.LegalDocument, error)
	Last() (*table.LegalDocument, error)
	Find() ([]*table.LegalDocument, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*table.LegalDocument, err error)
	FindInBatches(result *[]*table.LegalDocument, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*table.LegalDocument) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ILegalDocumentDo
	Assign(attrs ...field.AssignExpr) ILegalDocumentDo
	Joins(fields ...field.RelationField) ILegalDocumentDo
	Preload(fields ...field.RelationField) ILegalDocumentDo
	FirstOrInit() (*table.LegalDocument, error)
	FirstOrCreate() (*table.LegalDocument, error)
	FindByPage(offset int, limit int) (result []*table.LegalDocument, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ILegalDocumentDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (l legalDocumentDo) Debug() ILegalDocumentDo {
	return l.withDO(l.DO.Debug())
}

func (l legalDocumentDo) WithContext(ctx context.Context) ILegalDocumentDo {
	return l.withDO(l.DO.WithContext(ctx))
}

func (l legalDocumentDo) ReadDB() ILegalDocumentDo {
	return l.Clauses(dbresolver.Read)
}

func (l legalDocumentDo) WriteDB() ILegalDocumentDo {
	return l.Clauses(dbresolver.Write)
}

func (l legalDocumentDo) Session(config *gorm.Session) ILegalDocumentDo {
	return l.withDO(l.DO.Session(config))
}

func (l legalDocumentDo) Clauses(conds ...clause.Expression) ILegalDocumentDo {
	return l.withDO(l.DO.Clauses(conds...))
}

func (l legalDocumentDo) Returning(value interface{}, columns ...string) ILegalDocumentDo {
	return l.withDO(l.DO.Returning(value, columns...))
}

func (l legalDocumentDo) Not(conds ...gen.Condition) ILegalDocumentDo {
	return l.withDO(l.DO.Not(conds...))
}

func (l legalDocumentDo) Or(conds ...gen.Condition) ILegalDocumentDo {
	return l.withDO(l.DO.Or(conds...))
}

func (l legalDocumentDo) Select(conds ...field.Expr) ILegalDocumentDo {
	return l.withDO(l.DO.Select(conds...))
}

func (l legalDocumentDo) Where(conds ...gen.Condition) ILegalDocumentDo {
	return l.withDO(l.DO.Where(conds...))
}

func (l legalDocumentDo) Order(conds ...field.Expr) ILegalDocumentDo {
	return l.withDO(l.DO.Order(conds...))
}

func (l legalDocumentDo) Distinct(cols ...field.Expr) ILegalDocumentDo {
	return l.withDO(l.DO.Distinct(cols...))
}

func (l legalDocumentDo) Omit(cols ...field.Expr) ILegalDocumentDo {
	return l.withDO(l.DO.Omit(cols...))
}

func (l legalDocumentDo) Join(table schema.Tabler, on ...field.Expr) ILegalDocumentDo {
	return l.withDO(l.DO.Join(table, on...))
}

func (l legalDocumentDo) LeftJoin(table schema.Tabler, on ...field.Expr) ILegalDocumentDo {
	return l.withDO(l.DO.LeftJoin(table, on...))
}

func (l legalDocumentDo) RightJoin(table schema.Tabler, on ...field.Expr) ILegalDocumentDo {
	return l.withDO(l.DO.RightJoin(table, on...))
}

func (l legalDocumentDo) Group(cols ...field.Expr) ILegalDocumentDo {
	return l.withDO(l.DO.Group(cols...))
}

func (l legalDocumentDo) Having(conds ...gen.Condition) ILegalDocumentDo {
	return l.withDO(l.DO.Having(conds...))
}

func (l legalDocumentDo) Limit(limit int) ILegalDocumentDo {
	return l.withDO(l.DO.Limit(limit))
}

func (l legalDocumentDo) Offset(offset int) ILegalDocumentDo {
	return l.withDO(l.DO.Offset(offset))
}

func (l legalDocumentDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ILegalDocumentDo {
	return l.withDO(l.DO.Scopes(funcs...))
}

func (l legalDocumentDo) Unscoped() ILegalDocumentDo {
	return l.withDO(l.DO.Unscoped())
}

func (l legalDocumentDo) Create(values ...*table.LegalDocument) error {
	if len(values) == 0 {
		return nil
	}
	return l.DO.Create(values)
}

func (l legalDocumentDo) CreateInBatches(values []*table.LegalDocument, batchSize int) error {
	return l.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (l legalDocumentDo) Save(values ...*table.LegalDocument) error {
	if len(values) == 0 {
		return nil
	}
	return l.DO.Save(values)
}

func (l legalDocumentDo) First() (*table.LegalDocument, error) {
	if result, err := l.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*table.LegalDocument), nil
	}
}

func (l legalDocumentDo) Take() (*table.LegalDocument, error) {
	if result, err := l.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*table.LegalDocument), nil
	}
}

func (l legalDocumentDo) Last() (*table.LegalDocument, error) {
	if result, err := l.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*table.LegalDocument), nil
	}
}

func (l legalDocumentDo) Find() ([]*table.LegalDocument, error) {
	result, err := l.DO.Find()
	return result.([]*table.LegalDocument), err
}

func (l legalDocumentDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*table.LegalDocument, err error) {
	buf := make([]*table.LegalDocument, 0, batchSize)
	err = l.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (l legalDocumentDo) FindInBatches(result *[]*table.LegalDocument, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return l.DO.FindInBatches(result, batchSize, fc)
}

func (l legalDocumentDo) Attrs(attrs ...field.AssignExpr) ILegalDocumentDo {
	return l.withDO(l.DO.Attrs(attrs...))
}

func (l legalDocumentDo) Assign(attrs ...field.AssignExpr) ILegalDocumentDo {
	return l.withDO(l.DO.Assign(attrs...))
}

func (l legalDocumentDo) Joins(fields ...field.RelationField) ILegalDocumentDo {
	for _, _f := range fields {
		l = *l.withDO(l.DO.Joins(_f))
	}
	return &l
}

func (l legalDocumentDo) Preload(fields ...field.RelationField) ILegalDocumentDo {
	for _, _f := range fields {
		l = *l.withDO(l.DO.Preload(_f))
	}
	return &l
}

func (l legalDocumentDo) FirstOrInit() (*table.LegalDocument, error) {
	if result, err := l.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*table.LegalDocument), nil
	}
}

func (l legalDocumentDo) FirstOrCreate() (*table.LegalDocument, error) {
	if result, err := l.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*table.LegalDocument), nil
	}
}

func (l legalDocumentDo) FindByPage(offset int, limit int) (result []*table.LegalDocument, count int64, err error) {
	result, err = l.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = l.Offset(-1).Limit(-1).Count()
	return
}

func (l legalDocumentDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = l.Count()
	if err != nil {
		return
	}

	err = l.Offset(offset).Limit(limit).Scan(result)
	return
}

func (l legalDocumentDo) Scan(result interface{}) (err error) {
	return l.DO.Scan(result)
}

func (l legalDocumentDo) Delete(models ...*table.LegalDocument) (result gen.ResultInfo, err error) {
	return l.DO.Delete(models)
}

func (l *legalDocumentDo) withDO(do gen.Dao) *legalDocumentDo {
	l.DO = *do.(*gen.DO)
	return l
}

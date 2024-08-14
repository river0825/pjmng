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

func newUploadFileMorph(db *gorm.DB, opts ...gen.DOOption) uploadFileMorph {
	_uploadFileMorph := uploadFileMorph{}

	_uploadFileMorph.uploadFileMorphDo.UseDB(db, opts...)
	_uploadFileMorph.uploadFileMorphDo.UseModel(&table.UploadFileMorph{})

	tableName := _uploadFileMorph.uploadFileMorphDo.TableName()
	_uploadFileMorph.ALL = field.NewAsterisk(tableName)
	_uploadFileMorph.ID = field.NewInt32(tableName, "id")
	_uploadFileMorph.UploadFileID = field.NewInt32(tableName, "upload_file_id")
	_uploadFileMorph.RelatedID = field.NewInt32(tableName, "related_id")
	_uploadFileMorph.RelatedType = field.NewString(tableName, "related_type")
	_uploadFileMorph.Field = field.NewString(tableName, "field")
	_uploadFileMorph.Order_ = field.NewInt32(tableName, "order")

	_uploadFileMorph.fillFieldMap()

	return _uploadFileMorph
}

type uploadFileMorph struct {
	uploadFileMorphDo

	ALL          field.Asterisk
	ID           field.Int32
	UploadFileID field.Int32
	RelatedID    field.Int32
	RelatedType  field.String
	Field        field.String
	Order_       field.Int32

	fieldMap map[string]field.Expr
}

func (u uploadFileMorph) Table(newTableName string) *uploadFileMorph {
	u.uploadFileMorphDo.UseTable(newTableName)
	return u.updateTableName(newTableName)
}

func (u uploadFileMorph) As(alias string) *uploadFileMorph {
	u.uploadFileMorphDo.DO = *(u.uploadFileMorphDo.As(alias).(*gen.DO))
	return u.updateTableName(alias)
}

func (u *uploadFileMorph) updateTableName(table string) *uploadFileMorph {
	u.ALL = field.NewAsterisk(table)
	u.ID = field.NewInt32(table, "id")
	u.UploadFileID = field.NewInt32(table, "upload_file_id")
	u.RelatedID = field.NewInt32(table, "related_id")
	u.RelatedType = field.NewString(table, "related_type")
	u.Field = field.NewString(table, "field")
	u.Order_ = field.NewInt32(table, "order")

	u.fillFieldMap()

	return u
}

func (u *uploadFileMorph) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := u.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (u *uploadFileMorph) fillFieldMap() {
	u.fieldMap = make(map[string]field.Expr, 6)
	u.fieldMap["id"] = u.ID
	u.fieldMap["upload_file_id"] = u.UploadFileID
	u.fieldMap["related_id"] = u.RelatedID
	u.fieldMap["related_type"] = u.RelatedType
	u.fieldMap["field"] = u.Field
	u.fieldMap["order"] = u.Order_
}

func (u uploadFileMorph) clone(db *gorm.DB) uploadFileMorph {
	u.uploadFileMorphDo.ReplaceConnPool(db.Statement.ConnPool)
	return u
}

func (u uploadFileMorph) replaceDB(db *gorm.DB) uploadFileMorph {
	u.uploadFileMorphDo.ReplaceDB(db)
	return u
}

type uploadFileMorphDo struct{ gen.DO }

type IUploadFileMorphDo interface {
	gen.SubQuery
	Debug() IUploadFileMorphDo
	WithContext(ctx context.Context) IUploadFileMorphDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IUploadFileMorphDo
	WriteDB() IUploadFileMorphDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IUploadFileMorphDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IUploadFileMorphDo
	Not(conds ...gen.Condition) IUploadFileMorphDo
	Or(conds ...gen.Condition) IUploadFileMorphDo
	Select(conds ...field.Expr) IUploadFileMorphDo
	Where(conds ...gen.Condition) IUploadFileMorphDo
	Order(conds ...field.Expr) IUploadFileMorphDo
	Distinct(cols ...field.Expr) IUploadFileMorphDo
	Omit(cols ...field.Expr) IUploadFileMorphDo
	Join(table schema.Tabler, on ...field.Expr) IUploadFileMorphDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IUploadFileMorphDo
	RightJoin(table schema.Tabler, on ...field.Expr) IUploadFileMorphDo
	Group(cols ...field.Expr) IUploadFileMorphDo
	Having(conds ...gen.Condition) IUploadFileMorphDo
	Limit(limit int) IUploadFileMorphDo
	Offset(offset int) IUploadFileMorphDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IUploadFileMorphDo
	Unscoped() IUploadFileMorphDo
	Create(values ...*table.UploadFileMorph) error
	CreateInBatches(values []*table.UploadFileMorph, batchSize int) error
	Save(values ...*table.UploadFileMorph) error
	First() (*table.UploadFileMorph, error)
	Take() (*table.UploadFileMorph, error)
	Last() (*table.UploadFileMorph, error)
	Find() ([]*table.UploadFileMorph, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*table.UploadFileMorph, err error)
	FindInBatches(result *[]*table.UploadFileMorph, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*table.UploadFileMorph) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IUploadFileMorphDo
	Assign(attrs ...field.AssignExpr) IUploadFileMorphDo
	Joins(fields ...field.RelationField) IUploadFileMorphDo
	Preload(fields ...field.RelationField) IUploadFileMorphDo
	FirstOrInit() (*table.UploadFileMorph, error)
	FirstOrCreate() (*table.UploadFileMorph, error)
	FindByPage(offset int, limit int) (result []*table.UploadFileMorph, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IUploadFileMorphDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (u uploadFileMorphDo) Debug() IUploadFileMorphDo {
	return u.withDO(u.DO.Debug())
}

func (u uploadFileMorphDo) WithContext(ctx context.Context) IUploadFileMorphDo {
	return u.withDO(u.DO.WithContext(ctx))
}

func (u uploadFileMorphDo) ReadDB() IUploadFileMorphDo {
	return u.Clauses(dbresolver.Read)
}

func (u uploadFileMorphDo) WriteDB() IUploadFileMorphDo {
	return u.Clauses(dbresolver.Write)
}

func (u uploadFileMorphDo) Session(config *gorm.Session) IUploadFileMorphDo {
	return u.withDO(u.DO.Session(config))
}

func (u uploadFileMorphDo) Clauses(conds ...clause.Expression) IUploadFileMorphDo {
	return u.withDO(u.DO.Clauses(conds...))
}

func (u uploadFileMorphDo) Returning(value interface{}, columns ...string) IUploadFileMorphDo {
	return u.withDO(u.DO.Returning(value, columns...))
}

func (u uploadFileMorphDo) Not(conds ...gen.Condition) IUploadFileMorphDo {
	return u.withDO(u.DO.Not(conds...))
}

func (u uploadFileMorphDo) Or(conds ...gen.Condition) IUploadFileMorphDo {
	return u.withDO(u.DO.Or(conds...))
}

func (u uploadFileMorphDo) Select(conds ...field.Expr) IUploadFileMorphDo {
	return u.withDO(u.DO.Select(conds...))
}

func (u uploadFileMorphDo) Where(conds ...gen.Condition) IUploadFileMorphDo {
	return u.withDO(u.DO.Where(conds...))
}

func (u uploadFileMorphDo) Order(conds ...field.Expr) IUploadFileMorphDo {
	return u.withDO(u.DO.Order(conds...))
}

func (u uploadFileMorphDo) Distinct(cols ...field.Expr) IUploadFileMorphDo {
	return u.withDO(u.DO.Distinct(cols...))
}

func (u uploadFileMorphDo) Omit(cols ...field.Expr) IUploadFileMorphDo {
	return u.withDO(u.DO.Omit(cols...))
}

func (u uploadFileMorphDo) Join(table schema.Tabler, on ...field.Expr) IUploadFileMorphDo {
	return u.withDO(u.DO.Join(table, on...))
}

func (u uploadFileMorphDo) LeftJoin(table schema.Tabler, on ...field.Expr) IUploadFileMorphDo {
	return u.withDO(u.DO.LeftJoin(table, on...))
}

func (u uploadFileMorphDo) RightJoin(table schema.Tabler, on ...field.Expr) IUploadFileMorphDo {
	return u.withDO(u.DO.RightJoin(table, on...))
}

func (u uploadFileMorphDo) Group(cols ...field.Expr) IUploadFileMorphDo {
	return u.withDO(u.DO.Group(cols...))
}

func (u uploadFileMorphDo) Having(conds ...gen.Condition) IUploadFileMorphDo {
	return u.withDO(u.DO.Having(conds...))
}

func (u uploadFileMorphDo) Limit(limit int) IUploadFileMorphDo {
	return u.withDO(u.DO.Limit(limit))
}

func (u uploadFileMorphDo) Offset(offset int) IUploadFileMorphDo {
	return u.withDO(u.DO.Offset(offset))
}

func (u uploadFileMorphDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IUploadFileMorphDo {
	return u.withDO(u.DO.Scopes(funcs...))
}

func (u uploadFileMorphDo) Unscoped() IUploadFileMorphDo {
	return u.withDO(u.DO.Unscoped())
}

func (u uploadFileMorphDo) Create(values ...*table.UploadFileMorph) error {
	if len(values) == 0 {
		return nil
	}
	return u.DO.Create(values)
}

func (u uploadFileMorphDo) CreateInBatches(values []*table.UploadFileMorph, batchSize int) error {
	return u.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (u uploadFileMorphDo) Save(values ...*table.UploadFileMorph) error {
	if len(values) == 0 {
		return nil
	}
	return u.DO.Save(values)
}

func (u uploadFileMorphDo) First() (*table.UploadFileMorph, error) {
	if result, err := u.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*table.UploadFileMorph), nil
	}
}

func (u uploadFileMorphDo) Take() (*table.UploadFileMorph, error) {
	if result, err := u.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*table.UploadFileMorph), nil
	}
}

func (u uploadFileMorphDo) Last() (*table.UploadFileMorph, error) {
	if result, err := u.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*table.UploadFileMorph), nil
	}
}

func (u uploadFileMorphDo) Find() ([]*table.UploadFileMorph, error) {
	result, err := u.DO.Find()
	return result.([]*table.UploadFileMorph), err
}

func (u uploadFileMorphDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*table.UploadFileMorph, err error) {
	buf := make([]*table.UploadFileMorph, 0, batchSize)
	err = u.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (u uploadFileMorphDo) FindInBatches(result *[]*table.UploadFileMorph, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return u.DO.FindInBatches(result, batchSize, fc)
}

func (u uploadFileMorphDo) Attrs(attrs ...field.AssignExpr) IUploadFileMorphDo {
	return u.withDO(u.DO.Attrs(attrs...))
}

func (u uploadFileMorphDo) Assign(attrs ...field.AssignExpr) IUploadFileMorphDo {
	return u.withDO(u.DO.Assign(attrs...))
}

func (u uploadFileMorphDo) Joins(fields ...field.RelationField) IUploadFileMorphDo {
	for _, _f := range fields {
		u = *u.withDO(u.DO.Joins(_f))
	}
	return &u
}

func (u uploadFileMorphDo) Preload(fields ...field.RelationField) IUploadFileMorphDo {
	for _, _f := range fields {
		u = *u.withDO(u.DO.Preload(_f))
	}
	return &u
}

func (u uploadFileMorphDo) FirstOrInit() (*table.UploadFileMorph, error) {
	if result, err := u.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*table.UploadFileMorph), nil
	}
}

func (u uploadFileMorphDo) FirstOrCreate() (*table.UploadFileMorph, error) {
	if result, err := u.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*table.UploadFileMorph), nil
	}
}

func (u uploadFileMorphDo) FindByPage(offset int, limit int) (result []*table.UploadFileMorph, count int64, err error) {
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

func (u uploadFileMorphDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = u.Count()
	if err != nil {
		return
	}

	err = u.Offset(offset).Limit(limit).Scan(result)
	return
}

func (u uploadFileMorphDo) Scan(result interface{}) (err error) {
	return u.DO.Scan(result)
}

func (u uploadFileMorphDo) Delete(models ...*table.UploadFileMorph) (result gen.ResultInfo, err error) {
	return u.DO.Delete(models)
}

func (u *uploadFileMorphDo) withDO(do gen.Dao) *uploadFileMorphDo {
	u.DO = *do.(*gen.DO)
	return u
}

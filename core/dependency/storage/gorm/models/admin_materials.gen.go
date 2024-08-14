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

func newAdminMaterial(db *gorm.DB, opts ...gen.DOOption) adminMaterial {
	_adminMaterial := adminMaterial{}

	_adminMaterial.adminMaterialDo.UseDB(db, opts...)
	_adminMaterial.adminMaterialDo.UseModel(&table.AdminMaterial{})

	tableName := _adminMaterial.adminMaterialDo.TableName()
	_adminMaterial.ALL = field.NewAsterisk(tableName)
	_adminMaterial.ID = field.NewInt32(tableName, "id")
	_adminMaterial.Name = field.NewString(tableName, "name")
	_adminMaterial.Tag = field.NewString(tableName, "tag")
	_adminMaterial.User = field.NewInt32(tableName, "user")
	_adminMaterial.Content = field.NewField(tableName, "content")
	_adminMaterial.Lesson = field.NewInt32(tableName, "lesson")
	_adminMaterial.Grammer = field.NewInt32(tableName, "grammer")
	_adminMaterial.PublishedAt = field.NewTime(tableName, "published_at")
	_adminMaterial.CreatedBy = field.NewInt32(tableName, "created_by")
	_adminMaterial.UpdatedBy = field.NewInt32(tableName, "updated_by")
	_adminMaterial.CreatedAt = field.NewTime(tableName, "created_at")
	_adminMaterial.UpdatedAt = field.NewTime(tableName, "updated_at")
	_adminMaterial.Organization = field.NewInt32(tableName, "organization")
	_adminMaterial.Isfree = field.NewBool(tableName, "isfree")
	_adminMaterial.TagSimplify = field.NewString(tableName, "tag_simplify")

	_adminMaterial.fillFieldMap()

	return _adminMaterial
}

type adminMaterial struct {
	adminMaterialDo

	ALL          field.Asterisk
	ID           field.Int32
	Name         field.String
	Tag          field.String
	User         field.Int32
	Content      field.Field
	Lesson       field.Int32
	Grammer      field.Int32
	PublishedAt  field.Time
	CreatedBy    field.Int32
	UpdatedBy    field.Int32
	CreatedAt    field.Time
	UpdatedAt    field.Time
	Organization field.Int32
	Isfree       field.Bool
	TagSimplify  field.String

	fieldMap map[string]field.Expr
}

func (a adminMaterial) Table(newTableName string) *adminMaterial {
	a.adminMaterialDo.UseTable(newTableName)
	return a.updateTableName(newTableName)
}

func (a adminMaterial) As(alias string) *adminMaterial {
	a.adminMaterialDo.DO = *(a.adminMaterialDo.As(alias).(*gen.DO))
	return a.updateTableName(alias)
}

func (a *adminMaterial) updateTableName(table string) *adminMaterial {
	a.ALL = field.NewAsterisk(table)
	a.ID = field.NewInt32(table, "id")
	a.Name = field.NewString(table, "name")
	a.Tag = field.NewString(table, "tag")
	a.User = field.NewInt32(table, "user")
	a.Content = field.NewField(table, "content")
	a.Lesson = field.NewInt32(table, "lesson")
	a.Grammer = field.NewInt32(table, "grammer")
	a.PublishedAt = field.NewTime(table, "published_at")
	a.CreatedBy = field.NewInt32(table, "created_by")
	a.UpdatedBy = field.NewInt32(table, "updated_by")
	a.CreatedAt = field.NewTime(table, "created_at")
	a.UpdatedAt = field.NewTime(table, "updated_at")
	a.Organization = field.NewInt32(table, "organization")
	a.Isfree = field.NewBool(table, "isfree")
	a.TagSimplify = field.NewString(table, "tag_simplify")

	a.fillFieldMap()

	return a
}

func (a *adminMaterial) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := a.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (a *adminMaterial) fillFieldMap() {
	a.fieldMap = make(map[string]field.Expr, 15)
	a.fieldMap["id"] = a.ID
	a.fieldMap["name"] = a.Name
	a.fieldMap["tag"] = a.Tag
	a.fieldMap["user"] = a.User
	a.fieldMap["content"] = a.Content
	a.fieldMap["lesson"] = a.Lesson
	a.fieldMap["grammer"] = a.Grammer
	a.fieldMap["published_at"] = a.PublishedAt
	a.fieldMap["created_by"] = a.CreatedBy
	a.fieldMap["updated_by"] = a.UpdatedBy
	a.fieldMap["created_at"] = a.CreatedAt
	a.fieldMap["updated_at"] = a.UpdatedAt
	a.fieldMap["organization"] = a.Organization
	a.fieldMap["isfree"] = a.Isfree
	a.fieldMap["tag_simplify"] = a.TagSimplify
}

func (a adminMaterial) clone(db *gorm.DB) adminMaterial {
	a.adminMaterialDo.ReplaceConnPool(db.Statement.ConnPool)
	return a
}

func (a adminMaterial) replaceDB(db *gorm.DB) adminMaterial {
	a.adminMaterialDo.ReplaceDB(db)
	return a
}

type adminMaterialDo struct{ gen.DO }

type IAdminMaterialDo interface {
	gen.SubQuery
	Debug() IAdminMaterialDo
	WithContext(ctx context.Context) IAdminMaterialDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IAdminMaterialDo
	WriteDB() IAdminMaterialDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IAdminMaterialDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IAdminMaterialDo
	Not(conds ...gen.Condition) IAdminMaterialDo
	Or(conds ...gen.Condition) IAdminMaterialDo
	Select(conds ...field.Expr) IAdminMaterialDo
	Where(conds ...gen.Condition) IAdminMaterialDo
	Order(conds ...field.Expr) IAdminMaterialDo
	Distinct(cols ...field.Expr) IAdminMaterialDo
	Omit(cols ...field.Expr) IAdminMaterialDo
	Join(table schema.Tabler, on ...field.Expr) IAdminMaterialDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IAdminMaterialDo
	RightJoin(table schema.Tabler, on ...field.Expr) IAdminMaterialDo
	Group(cols ...field.Expr) IAdminMaterialDo
	Having(conds ...gen.Condition) IAdminMaterialDo
	Limit(limit int) IAdminMaterialDo
	Offset(offset int) IAdminMaterialDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IAdminMaterialDo
	Unscoped() IAdminMaterialDo
	Create(values ...*table.AdminMaterial) error
	CreateInBatches(values []*table.AdminMaterial, batchSize int) error
	Save(values ...*table.AdminMaterial) error
	First() (*table.AdminMaterial, error)
	Take() (*table.AdminMaterial, error)
	Last() (*table.AdminMaterial, error)
	Find() ([]*table.AdminMaterial, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*table.AdminMaterial, err error)
	FindInBatches(result *[]*table.AdminMaterial, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*table.AdminMaterial) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IAdminMaterialDo
	Assign(attrs ...field.AssignExpr) IAdminMaterialDo
	Joins(fields ...field.RelationField) IAdminMaterialDo
	Preload(fields ...field.RelationField) IAdminMaterialDo
	FirstOrInit() (*table.AdminMaterial, error)
	FirstOrCreate() (*table.AdminMaterial, error)
	FindByPage(offset int, limit int) (result []*table.AdminMaterial, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IAdminMaterialDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (a adminMaterialDo) Debug() IAdminMaterialDo {
	return a.withDO(a.DO.Debug())
}

func (a adminMaterialDo) WithContext(ctx context.Context) IAdminMaterialDo {
	return a.withDO(a.DO.WithContext(ctx))
}

func (a adminMaterialDo) ReadDB() IAdminMaterialDo {
	return a.Clauses(dbresolver.Read)
}

func (a adminMaterialDo) WriteDB() IAdminMaterialDo {
	return a.Clauses(dbresolver.Write)
}

func (a adminMaterialDo) Session(config *gorm.Session) IAdminMaterialDo {
	return a.withDO(a.DO.Session(config))
}

func (a adminMaterialDo) Clauses(conds ...clause.Expression) IAdminMaterialDo {
	return a.withDO(a.DO.Clauses(conds...))
}

func (a adminMaterialDo) Returning(value interface{}, columns ...string) IAdminMaterialDo {
	return a.withDO(a.DO.Returning(value, columns...))
}

func (a adminMaterialDo) Not(conds ...gen.Condition) IAdminMaterialDo {
	return a.withDO(a.DO.Not(conds...))
}

func (a adminMaterialDo) Or(conds ...gen.Condition) IAdminMaterialDo {
	return a.withDO(a.DO.Or(conds...))
}

func (a adminMaterialDo) Select(conds ...field.Expr) IAdminMaterialDo {
	return a.withDO(a.DO.Select(conds...))
}

func (a adminMaterialDo) Where(conds ...gen.Condition) IAdminMaterialDo {
	return a.withDO(a.DO.Where(conds...))
}

func (a adminMaterialDo) Order(conds ...field.Expr) IAdminMaterialDo {
	return a.withDO(a.DO.Order(conds...))
}

func (a adminMaterialDo) Distinct(cols ...field.Expr) IAdminMaterialDo {
	return a.withDO(a.DO.Distinct(cols...))
}

func (a adminMaterialDo) Omit(cols ...field.Expr) IAdminMaterialDo {
	return a.withDO(a.DO.Omit(cols...))
}

func (a adminMaterialDo) Join(table schema.Tabler, on ...field.Expr) IAdminMaterialDo {
	return a.withDO(a.DO.Join(table, on...))
}

func (a adminMaterialDo) LeftJoin(table schema.Tabler, on ...field.Expr) IAdminMaterialDo {
	return a.withDO(a.DO.LeftJoin(table, on...))
}

func (a adminMaterialDo) RightJoin(table schema.Tabler, on ...field.Expr) IAdminMaterialDo {
	return a.withDO(a.DO.RightJoin(table, on...))
}

func (a adminMaterialDo) Group(cols ...field.Expr) IAdminMaterialDo {
	return a.withDO(a.DO.Group(cols...))
}

func (a adminMaterialDo) Having(conds ...gen.Condition) IAdminMaterialDo {
	return a.withDO(a.DO.Having(conds...))
}

func (a adminMaterialDo) Limit(limit int) IAdminMaterialDo {
	return a.withDO(a.DO.Limit(limit))
}

func (a adminMaterialDo) Offset(offset int) IAdminMaterialDo {
	return a.withDO(a.DO.Offset(offset))
}

func (a adminMaterialDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IAdminMaterialDo {
	return a.withDO(a.DO.Scopes(funcs...))
}

func (a adminMaterialDo) Unscoped() IAdminMaterialDo {
	return a.withDO(a.DO.Unscoped())
}

func (a adminMaterialDo) Create(values ...*table.AdminMaterial) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Create(values)
}

func (a adminMaterialDo) CreateInBatches(values []*table.AdminMaterial, batchSize int) error {
	return a.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (a adminMaterialDo) Save(values ...*table.AdminMaterial) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Save(values)
}

func (a adminMaterialDo) First() (*table.AdminMaterial, error) {
	if result, err := a.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*table.AdminMaterial), nil
	}
}

func (a adminMaterialDo) Take() (*table.AdminMaterial, error) {
	if result, err := a.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*table.AdminMaterial), nil
	}
}

func (a adminMaterialDo) Last() (*table.AdminMaterial, error) {
	if result, err := a.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*table.AdminMaterial), nil
	}
}

func (a adminMaterialDo) Find() ([]*table.AdminMaterial, error) {
	result, err := a.DO.Find()
	return result.([]*table.AdminMaterial), err
}

func (a adminMaterialDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*table.AdminMaterial, err error) {
	buf := make([]*table.AdminMaterial, 0, batchSize)
	err = a.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (a adminMaterialDo) FindInBatches(result *[]*table.AdminMaterial, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return a.DO.FindInBatches(result, batchSize, fc)
}

func (a adminMaterialDo) Attrs(attrs ...field.AssignExpr) IAdminMaterialDo {
	return a.withDO(a.DO.Attrs(attrs...))
}

func (a adminMaterialDo) Assign(attrs ...field.AssignExpr) IAdminMaterialDo {
	return a.withDO(a.DO.Assign(attrs...))
}

func (a adminMaterialDo) Joins(fields ...field.RelationField) IAdminMaterialDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Joins(_f))
	}
	return &a
}

func (a adminMaterialDo) Preload(fields ...field.RelationField) IAdminMaterialDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Preload(_f))
	}
	return &a
}

func (a adminMaterialDo) FirstOrInit() (*table.AdminMaterial, error) {
	if result, err := a.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*table.AdminMaterial), nil
	}
}

func (a adminMaterialDo) FirstOrCreate() (*table.AdminMaterial, error) {
	if result, err := a.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*table.AdminMaterial), nil
	}
}

func (a adminMaterialDo) FindByPage(offset int, limit int) (result []*table.AdminMaterial, count int64, err error) {
	result, err = a.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = a.Offset(-1).Limit(-1).Count()
	return
}

func (a adminMaterialDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = a.Count()
	if err != nil {
		return
	}

	err = a.Offset(offset).Limit(limit).Scan(result)
	return
}

func (a adminMaterialDo) Scan(result interface{}) (err error) {
	return a.DO.Scan(result)
}

func (a adminMaterialDo) Delete(models ...*table.AdminMaterial) (result gen.ResultInfo, err error) {
	return a.DO.Delete(models)
}

func (a *adminMaterialDo) withDO(do gen.Dao) *adminMaterialDo {
	a.DO = *do.(*gen.DO)
	return a
}

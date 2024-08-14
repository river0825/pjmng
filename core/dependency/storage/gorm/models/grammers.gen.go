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

func newGrammer(db *gorm.DB, opts ...gen.DOOption) grammer {
	_grammer := grammer{}

	_grammer.grammerDo.UseDB(db, opts...)
	_grammer.grammerDo.UseModel(&table.Grammer{})

	tableName := _grammer.grammerDo.TableName()
	_grammer.ALL = field.NewAsterisk(tableName)
	_grammer.ID = field.NewInt32(tableName, "id")
	_grammer.Name = field.NewString(tableName, "name")
	_grammer.Tag = field.NewString(tableName, "tag")
	_grammer.AdminMaterial = field.NewInt32(tableName, "admin_material")
	_grammer.PublishedAt = field.NewTime(tableName, "published_at")
	_grammer.CreatedBy = field.NewInt32(tableName, "created_by")
	_grammer.UpdatedBy = field.NewInt32(tableName, "updated_by")
	_grammer.CreatedAt = field.NewTime(tableName, "created_at")
	_grammer.UpdatedAt = field.NewTime(tableName, "updated_at")
	_grammer.Content = field.NewField(tableName, "content")
	_grammer.Isfree = field.NewBool(tableName, "isfree")
	_grammer.Order_ = field.NewInt64(tableName, "order")

	_grammer.fillFieldMap()

	return _grammer
}

type grammer struct {
	grammerDo

	ALL           field.Asterisk
	ID            field.Int32
	Name          field.String
	Tag           field.String
	AdminMaterial field.Int32
	PublishedAt   field.Time
	CreatedBy     field.Int32
	UpdatedBy     field.Int32
	CreatedAt     field.Time
	UpdatedAt     field.Time
	Content       field.Field
	Isfree        field.Bool
	Order_        field.Int64

	fieldMap map[string]field.Expr
}

func (g grammer) Table(newTableName string) *grammer {
	g.grammerDo.UseTable(newTableName)
	return g.updateTableName(newTableName)
}

func (g grammer) As(alias string) *grammer {
	g.grammerDo.DO = *(g.grammerDo.As(alias).(*gen.DO))
	return g.updateTableName(alias)
}

func (g *grammer) updateTableName(table string) *grammer {
	g.ALL = field.NewAsterisk(table)
	g.ID = field.NewInt32(table, "id")
	g.Name = field.NewString(table, "name")
	g.Tag = field.NewString(table, "tag")
	g.AdminMaterial = field.NewInt32(table, "admin_material")
	g.PublishedAt = field.NewTime(table, "published_at")
	g.CreatedBy = field.NewInt32(table, "created_by")
	g.UpdatedBy = field.NewInt32(table, "updated_by")
	g.CreatedAt = field.NewTime(table, "created_at")
	g.UpdatedAt = field.NewTime(table, "updated_at")
	g.Content = field.NewField(table, "content")
	g.Isfree = field.NewBool(table, "isfree")
	g.Order_ = field.NewInt64(table, "order")

	g.fillFieldMap()

	return g
}

func (g *grammer) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := g.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (g *grammer) fillFieldMap() {
	g.fieldMap = make(map[string]field.Expr, 12)
	g.fieldMap["id"] = g.ID
	g.fieldMap["name"] = g.Name
	g.fieldMap["tag"] = g.Tag
	g.fieldMap["admin_material"] = g.AdminMaterial
	g.fieldMap["published_at"] = g.PublishedAt
	g.fieldMap["created_by"] = g.CreatedBy
	g.fieldMap["updated_by"] = g.UpdatedBy
	g.fieldMap["created_at"] = g.CreatedAt
	g.fieldMap["updated_at"] = g.UpdatedAt
	g.fieldMap["content"] = g.Content
	g.fieldMap["isfree"] = g.Isfree
	g.fieldMap["order"] = g.Order_
}

func (g grammer) clone(db *gorm.DB) grammer {
	g.grammerDo.ReplaceConnPool(db.Statement.ConnPool)
	return g
}

func (g grammer) replaceDB(db *gorm.DB) grammer {
	g.grammerDo.ReplaceDB(db)
	return g
}

type grammerDo struct{ gen.DO }

type IGrammerDo interface {
	gen.SubQuery
	Debug() IGrammerDo
	WithContext(ctx context.Context) IGrammerDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IGrammerDo
	WriteDB() IGrammerDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IGrammerDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IGrammerDo
	Not(conds ...gen.Condition) IGrammerDo
	Or(conds ...gen.Condition) IGrammerDo
	Select(conds ...field.Expr) IGrammerDo
	Where(conds ...gen.Condition) IGrammerDo
	Order(conds ...field.Expr) IGrammerDo
	Distinct(cols ...field.Expr) IGrammerDo
	Omit(cols ...field.Expr) IGrammerDo
	Join(table schema.Tabler, on ...field.Expr) IGrammerDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IGrammerDo
	RightJoin(table schema.Tabler, on ...field.Expr) IGrammerDo
	Group(cols ...field.Expr) IGrammerDo
	Having(conds ...gen.Condition) IGrammerDo
	Limit(limit int) IGrammerDo
	Offset(offset int) IGrammerDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IGrammerDo
	Unscoped() IGrammerDo
	Create(values ...*table.Grammer) error
	CreateInBatches(values []*table.Grammer, batchSize int) error
	Save(values ...*table.Grammer) error
	First() (*table.Grammer, error)
	Take() (*table.Grammer, error)
	Last() (*table.Grammer, error)
	Find() ([]*table.Grammer, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*table.Grammer, err error)
	FindInBatches(result *[]*table.Grammer, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*table.Grammer) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IGrammerDo
	Assign(attrs ...field.AssignExpr) IGrammerDo
	Joins(fields ...field.RelationField) IGrammerDo
	Preload(fields ...field.RelationField) IGrammerDo
	FirstOrInit() (*table.Grammer, error)
	FirstOrCreate() (*table.Grammer, error)
	FindByPage(offset int, limit int) (result []*table.Grammer, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IGrammerDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (g grammerDo) Debug() IGrammerDo {
	return g.withDO(g.DO.Debug())
}

func (g grammerDo) WithContext(ctx context.Context) IGrammerDo {
	return g.withDO(g.DO.WithContext(ctx))
}

func (g grammerDo) ReadDB() IGrammerDo {
	return g.Clauses(dbresolver.Read)
}

func (g grammerDo) WriteDB() IGrammerDo {
	return g.Clauses(dbresolver.Write)
}

func (g grammerDo) Session(config *gorm.Session) IGrammerDo {
	return g.withDO(g.DO.Session(config))
}

func (g grammerDo) Clauses(conds ...clause.Expression) IGrammerDo {
	return g.withDO(g.DO.Clauses(conds...))
}

func (g grammerDo) Returning(value interface{}, columns ...string) IGrammerDo {
	return g.withDO(g.DO.Returning(value, columns...))
}

func (g grammerDo) Not(conds ...gen.Condition) IGrammerDo {
	return g.withDO(g.DO.Not(conds...))
}

func (g grammerDo) Or(conds ...gen.Condition) IGrammerDo {
	return g.withDO(g.DO.Or(conds...))
}

func (g grammerDo) Select(conds ...field.Expr) IGrammerDo {
	return g.withDO(g.DO.Select(conds...))
}

func (g grammerDo) Where(conds ...gen.Condition) IGrammerDo {
	return g.withDO(g.DO.Where(conds...))
}

func (g grammerDo) Order(conds ...field.Expr) IGrammerDo {
	return g.withDO(g.DO.Order(conds...))
}

func (g grammerDo) Distinct(cols ...field.Expr) IGrammerDo {
	return g.withDO(g.DO.Distinct(cols...))
}

func (g grammerDo) Omit(cols ...field.Expr) IGrammerDo {
	return g.withDO(g.DO.Omit(cols...))
}

func (g grammerDo) Join(table schema.Tabler, on ...field.Expr) IGrammerDo {
	return g.withDO(g.DO.Join(table, on...))
}

func (g grammerDo) LeftJoin(table schema.Tabler, on ...field.Expr) IGrammerDo {
	return g.withDO(g.DO.LeftJoin(table, on...))
}

func (g grammerDo) RightJoin(table schema.Tabler, on ...field.Expr) IGrammerDo {
	return g.withDO(g.DO.RightJoin(table, on...))
}

func (g grammerDo) Group(cols ...field.Expr) IGrammerDo {
	return g.withDO(g.DO.Group(cols...))
}

func (g grammerDo) Having(conds ...gen.Condition) IGrammerDo {
	return g.withDO(g.DO.Having(conds...))
}

func (g grammerDo) Limit(limit int) IGrammerDo {
	return g.withDO(g.DO.Limit(limit))
}

func (g grammerDo) Offset(offset int) IGrammerDo {
	return g.withDO(g.DO.Offset(offset))
}

func (g grammerDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IGrammerDo {
	return g.withDO(g.DO.Scopes(funcs...))
}

func (g grammerDo) Unscoped() IGrammerDo {
	return g.withDO(g.DO.Unscoped())
}

func (g grammerDo) Create(values ...*table.Grammer) error {
	if len(values) == 0 {
		return nil
	}
	return g.DO.Create(values)
}

func (g grammerDo) CreateInBatches(values []*table.Grammer, batchSize int) error {
	return g.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (g grammerDo) Save(values ...*table.Grammer) error {
	if len(values) == 0 {
		return nil
	}
	return g.DO.Save(values)
}

func (g grammerDo) First() (*table.Grammer, error) {
	if result, err := g.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*table.Grammer), nil
	}
}

func (g grammerDo) Take() (*table.Grammer, error) {
	if result, err := g.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*table.Grammer), nil
	}
}

func (g grammerDo) Last() (*table.Grammer, error) {
	if result, err := g.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*table.Grammer), nil
	}
}

func (g grammerDo) Find() ([]*table.Grammer, error) {
	result, err := g.DO.Find()
	return result.([]*table.Grammer), err
}

func (g grammerDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*table.Grammer, err error) {
	buf := make([]*table.Grammer, 0, batchSize)
	err = g.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (g grammerDo) FindInBatches(result *[]*table.Grammer, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return g.DO.FindInBatches(result, batchSize, fc)
}

func (g grammerDo) Attrs(attrs ...field.AssignExpr) IGrammerDo {
	return g.withDO(g.DO.Attrs(attrs...))
}

func (g grammerDo) Assign(attrs ...field.AssignExpr) IGrammerDo {
	return g.withDO(g.DO.Assign(attrs...))
}

func (g grammerDo) Joins(fields ...field.RelationField) IGrammerDo {
	for _, _f := range fields {
		g = *g.withDO(g.DO.Joins(_f))
	}
	return &g
}

func (g grammerDo) Preload(fields ...field.RelationField) IGrammerDo {
	for _, _f := range fields {
		g = *g.withDO(g.DO.Preload(_f))
	}
	return &g
}

func (g grammerDo) FirstOrInit() (*table.Grammer, error) {
	if result, err := g.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*table.Grammer), nil
	}
}

func (g grammerDo) FirstOrCreate() (*table.Grammer, error) {
	if result, err := g.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*table.Grammer), nil
	}
}

func (g grammerDo) FindByPage(offset int, limit int) (result []*table.Grammer, count int64, err error) {
	result, err = g.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = g.Offset(-1).Limit(-1).Count()
	return
}

func (g grammerDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = g.Count()
	if err != nil {
		return
	}

	err = g.Offset(offset).Limit(limit).Scan(result)
	return
}

func (g grammerDo) Scan(result interface{}) (err error) {
	return g.DO.Scan(result)
}

func (g grammerDo) Delete(models ...*table.Grammer) (result gen.ResultInfo, err error) {
	return g.DO.Delete(models)
}

func (g *grammerDo) withDO(do gen.Dao) *grammerDo {
	g.DO = *do.(*gen.DO)
	return g
}
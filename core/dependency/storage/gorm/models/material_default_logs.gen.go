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

func newMaterialDefaultLog(db *gorm.DB, opts ...gen.DOOption) materialDefaultLog {
	_materialDefaultLog := materialDefaultLog{}

	_materialDefaultLog.materialDefaultLogDo.UseDB(db, opts...)
	_materialDefaultLog.materialDefaultLogDo.UseModel(&table.MaterialDefaultLog{})

	tableName := _materialDefaultLog.materialDefaultLogDo.TableName()
	_materialDefaultLog.ALL = field.NewAsterisk(tableName)
	_materialDefaultLog.ID = field.NewInt32(tableName, "id")
	_materialDefaultLog.UserID = field.NewInt32(tableName, "user_id")
	_materialDefaultLog.CreatedBy = field.NewInt32(tableName, "created_by")
	_materialDefaultLog.UpdatedBy = field.NewInt32(tableName, "updated_by")
	_materialDefaultLog.CreatedAt = field.NewTime(tableName, "created_at")
	_materialDefaultLog.UpdatedAt = field.NewTime(tableName, "updated_at")

	_materialDefaultLog.fillFieldMap()

	return _materialDefaultLog
}

type materialDefaultLog struct {
	materialDefaultLogDo

	ALL       field.Asterisk
	ID        field.Int32
	UserID    field.Int32
	CreatedBy field.Int32
	UpdatedBy field.Int32
	CreatedAt field.Time
	UpdatedAt field.Time

	fieldMap map[string]field.Expr
}

func (m materialDefaultLog) Table(newTableName string) *materialDefaultLog {
	m.materialDefaultLogDo.UseTable(newTableName)
	return m.updateTableName(newTableName)
}

func (m materialDefaultLog) As(alias string) *materialDefaultLog {
	m.materialDefaultLogDo.DO = *(m.materialDefaultLogDo.As(alias).(*gen.DO))
	return m.updateTableName(alias)
}

func (m *materialDefaultLog) updateTableName(table string) *materialDefaultLog {
	m.ALL = field.NewAsterisk(table)
	m.ID = field.NewInt32(table, "id")
	m.UserID = field.NewInt32(table, "user_id")
	m.CreatedBy = field.NewInt32(table, "created_by")
	m.UpdatedBy = field.NewInt32(table, "updated_by")
	m.CreatedAt = field.NewTime(table, "created_at")
	m.UpdatedAt = field.NewTime(table, "updated_at")

	m.fillFieldMap()

	return m
}

func (m *materialDefaultLog) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := m.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (m *materialDefaultLog) fillFieldMap() {
	m.fieldMap = make(map[string]field.Expr, 6)
	m.fieldMap["id"] = m.ID
	m.fieldMap["user_id"] = m.UserID
	m.fieldMap["created_by"] = m.CreatedBy
	m.fieldMap["updated_by"] = m.UpdatedBy
	m.fieldMap["created_at"] = m.CreatedAt
	m.fieldMap["updated_at"] = m.UpdatedAt
}

func (m materialDefaultLog) clone(db *gorm.DB) materialDefaultLog {
	m.materialDefaultLogDo.ReplaceConnPool(db.Statement.ConnPool)
	return m
}

func (m materialDefaultLog) replaceDB(db *gorm.DB) materialDefaultLog {
	m.materialDefaultLogDo.ReplaceDB(db)
	return m
}

type materialDefaultLogDo struct{ gen.DO }

type IMaterialDefaultLogDo interface {
	gen.SubQuery
	Debug() IMaterialDefaultLogDo
	WithContext(ctx context.Context) IMaterialDefaultLogDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IMaterialDefaultLogDo
	WriteDB() IMaterialDefaultLogDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IMaterialDefaultLogDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IMaterialDefaultLogDo
	Not(conds ...gen.Condition) IMaterialDefaultLogDo
	Or(conds ...gen.Condition) IMaterialDefaultLogDo
	Select(conds ...field.Expr) IMaterialDefaultLogDo
	Where(conds ...gen.Condition) IMaterialDefaultLogDo
	Order(conds ...field.Expr) IMaterialDefaultLogDo
	Distinct(cols ...field.Expr) IMaterialDefaultLogDo
	Omit(cols ...field.Expr) IMaterialDefaultLogDo
	Join(table schema.Tabler, on ...field.Expr) IMaterialDefaultLogDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IMaterialDefaultLogDo
	RightJoin(table schema.Tabler, on ...field.Expr) IMaterialDefaultLogDo
	Group(cols ...field.Expr) IMaterialDefaultLogDo
	Having(conds ...gen.Condition) IMaterialDefaultLogDo
	Limit(limit int) IMaterialDefaultLogDo
	Offset(offset int) IMaterialDefaultLogDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IMaterialDefaultLogDo
	Unscoped() IMaterialDefaultLogDo
	Create(values ...*table.MaterialDefaultLog) error
	CreateInBatches(values []*table.MaterialDefaultLog, batchSize int) error
	Save(values ...*table.MaterialDefaultLog) error
	First() (*table.MaterialDefaultLog, error)
	Take() (*table.MaterialDefaultLog, error)
	Last() (*table.MaterialDefaultLog, error)
	Find() ([]*table.MaterialDefaultLog, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*table.MaterialDefaultLog, err error)
	FindInBatches(result *[]*table.MaterialDefaultLog, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*table.MaterialDefaultLog) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IMaterialDefaultLogDo
	Assign(attrs ...field.AssignExpr) IMaterialDefaultLogDo
	Joins(fields ...field.RelationField) IMaterialDefaultLogDo
	Preload(fields ...field.RelationField) IMaterialDefaultLogDo
	FirstOrInit() (*table.MaterialDefaultLog, error)
	FirstOrCreate() (*table.MaterialDefaultLog, error)
	FindByPage(offset int, limit int) (result []*table.MaterialDefaultLog, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IMaterialDefaultLogDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (m materialDefaultLogDo) Debug() IMaterialDefaultLogDo {
	return m.withDO(m.DO.Debug())
}

func (m materialDefaultLogDo) WithContext(ctx context.Context) IMaterialDefaultLogDo {
	return m.withDO(m.DO.WithContext(ctx))
}

func (m materialDefaultLogDo) ReadDB() IMaterialDefaultLogDo {
	return m.Clauses(dbresolver.Read)
}

func (m materialDefaultLogDo) WriteDB() IMaterialDefaultLogDo {
	return m.Clauses(dbresolver.Write)
}

func (m materialDefaultLogDo) Session(config *gorm.Session) IMaterialDefaultLogDo {
	return m.withDO(m.DO.Session(config))
}

func (m materialDefaultLogDo) Clauses(conds ...clause.Expression) IMaterialDefaultLogDo {
	return m.withDO(m.DO.Clauses(conds...))
}

func (m materialDefaultLogDo) Returning(value interface{}, columns ...string) IMaterialDefaultLogDo {
	return m.withDO(m.DO.Returning(value, columns...))
}

func (m materialDefaultLogDo) Not(conds ...gen.Condition) IMaterialDefaultLogDo {
	return m.withDO(m.DO.Not(conds...))
}

func (m materialDefaultLogDo) Or(conds ...gen.Condition) IMaterialDefaultLogDo {
	return m.withDO(m.DO.Or(conds...))
}

func (m materialDefaultLogDo) Select(conds ...field.Expr) IMaterialDefaultLogDo {
	return m.withDO(m.DO.Select(conds...))
}

func (m materialDefaultLogDo) Where(conds ...gen.Condition) IMaterialDefaultLogDo {
	return m.withDO(m.DO.Where(conds...))
}

func (m materialDefaultLogDo) Order(conds ...field.Expr) IMaterialDefaultLogDo {
	return m.withDO(m.DO.Order(conds...))
}

func (m materialDefaultLogDo) Distinct(cols ...field.Expr) IMaterialDefaultLogDo {
	return m.withDO(m.DO.Distinct(cols...))
}

func (m materialDefaultLogDo) Omit(cols ...field.Expr) IMaterialDefaultLogDo {
	return m.withDO(m.DO.Omit(cols...))
}

func (m materialDefaultLogDo) Join(table schema.Tabler, on ...field.Expr) IMaterialDefaultLogDo {
	return m.withDO(m.DO.Join(table, on...))
}

func (m materialDefaultLogDo) LeftJoin(table schema.Tabler, on ...field.Expr) IMaterialDefaultLogDo {
	return m.withDO(m.DO.LeftJoin(table, on...))
}

func (m materialDefaultLogDo) RightJoin(table schema.Tabler, on ...field.Expr) IMaterialDefaultLogDo {
	return m.withDO(m.DO.RightJoin(table, on...))
}

func (m materialDefaultLogDo) Group(cols ...field.Expr) IMaterialDefaultLogDo {
	return m.withDO(m.DO.Group(cols...))
}

func (m materialDefaultLogDo) Having(conds ...gen.Condition) IMaterialDefaultLogDo {
	return m.withDO(m.DO.Having(conds...))
}

func (m materialDefaultLogDo) Limit(limit int) IMaterialDefaultLogDo {
	return m.withDO(m.DO.Limit(limit))
}

func (m materialDefaultLogDo) Offset(offset int) IMaterialDefaultLogDo {
	return m.withDO(m.DO.Offset(offset))
}

func (m materialDefaultLogDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IMaterialDefaultLogDo {
	return m.withDO(m.DO.Scopes(funcs...))
}

func (m materialDefaultLogDo) Unscoped() IMaterialDefaultLogDo {
	return m.withDO(m.DO.Unscoped())
}

func (m materialDefaultLogDo) Create(values ...*table.MaterialDefaultLog) error {
	if len(values) == 0 {
		return nil
	}
	return m.DO.Create(values)
}

func (m materialDefaultLogDo) CreateInBatches(values []*table.MaterialDefaultLog, batchSize int) error {
	return m.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (m materialDefaultLogDo) Save(values ...*table.MaterialDefaultLog) error {
	if len(values) == 0 {
		return nil
	}
	return m.DO.Save(values)
}

func (m materialDefaultLogDo) First() (*table.MaterialDefaultLog, error) {
	if result, err := m.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*table.MaterialDefaultLog), nil
	}
}

func (m materialDefaultLogDo) Take() (*table.MaterialDefaultLog, error) {
	if result, err := m.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*table.MaterialDefaultLog), nil
	}
}

func (m materialDefaultLogDo) Last() (*table.MaterialDefaultLog, error) {
	if result, err := m.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*table.MaterialDefaultLog), nil
	}
}

func (m materialDefaultLogDo) Find() ([]*table.MaterialDefaultLog, error) {
	result, err := m.DO.Find()
	return result.([]*table.MaterialDefaultLog), err
}

func (m materialDefaultLogDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*table.MaterialDefaultLog, err error) {
	buf := make([]*table.MaterialDefaultLog, 0, batchSize)
	err = m.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (m materialDefaultLogDo) FindInBatches(result *[]*table.MaterialDefaultLog, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return m.DO.FindInBatches(result, batchSize, fc)
}

func (m materialDefaultLogDo) Attrs(attrs ...field.AssignExpr) IMaterialDefaultLogDo {
	return m.withDO(m.DO.Attrs(attrs...))
}

func (m materialDefaultLogDo) Assign(attrs ...field.AssignExpr) IMaterialDefaultLogDo {
	return m.withDO(m.DO.Assign(attrs...))
}

func (m materialDefaultLogDo) Joins(fields ...field.RelationField) IMaterialDefaultLogDo {
	for _, _f := range fields {
		m = *m.withDO(m.DO.Joins(_f))
	}
	return &m
}

func (m materialDefaultLogDo) Preload(fields ...field.RelationField) IMaterialDefaultLogDo {
	for _, _f := range fields {
		m = *m.withDO(m.DO.Preload(_f))
	}
	return &m
}

func (m materialDefaultLogDo) FirstOrInit() (*table.MaterialDefaultLog, error) {
	if result, err := m.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*table.MaterialDefaultLog), nil
	}
}

func (m materialDefaultLogDo) FirstOrCreate() (*table.MaterialDefaultLog, error) {
	if result, err := m.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*table.MaterialDefaultLog), nil
	}
}

func (m materialDefaultLogDo) FindByPage(offset int, limit int) (result []*table.MaterialDefaultLog, count int64, err error) {
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

func (m materialDefaultLogDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = m.Count()
	if err != nil {
		return
	}

	err = m.Offset(offset).Limit(limit).Scan(result)
	return
}

func (m materialDefaultLogDo) Scan(result interface{}) (err error) {
	return m.DO.Scan(result)
}

func (m materialDefaultLogDo) Delete(models ...*table.MaterialDefaultLog) (result gen.ResultInfo, err error) {
	return m.DO.Delete(models)
}

func (m *materialDefaultLogDo) withDO(do gen.Dao) *materialDefaultLogDo {
	m.DO = *do.(*gen.DO)
	return m
}
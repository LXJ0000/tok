// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"github.com/LXJ0000/tok/backend/basicService/internal/data/model"
)

func newCollection(db *gorm.DB, opts ...gen.DOOption) collection {
	_collection := collection{}

	_collection.collectionDo.UseDB(db, opts...)
	_collection.collectionDo.UseModel(&model.Collection{})

	tableName := _collection.collectionDo.TableName()
	_collection.ALL = field.NewAsterisk(tableName)
	_collection.ID = field.NewInt64(tableName, "id")
	_collection.UserID = field.NewInt64(tableName, "user_id")
	_collection.Title = field.NewString(tableName, "title")
	_collection.Description = field.NewString(tableName, "description")
	_collection.IsDeleted = field.NewInt32(tableName, "is_deleted")
	_collection.CreateTime = field.NewTime(tableName, "create_time")
	_collection.UpdateTime = field.NewTime(tableName, "update_time")

	_collection.fillFieldMap()

	return _collection
}

type collection struct {
	collectionDo

	ALL         field.Asterisk
	ID          field.Int64  // 收藏夹ID
	UserID      field.Int64  // 用户ID
	Title       field.String // 标题
	Description field.String // 描述
	IsDeleted   field.Int32  // 是否删除
	CreateTime  field.Time   // 创建时间
	UpdateTime  field.Time   // 更新时间

	fieldMap map[string]field.Expr
}

func (c collection) Table(newTableName string) *collection {
	c.collectionDo.UseTable(newTableName)
	return c.updateTableName(newTableName)
}

func (c collection) As(alias string) *collection {
	c.collectionDo.DO = *(c.collectionDo.As(alias).(*gen.DO))
	return c.updateTableName(alias)
}

func (c *collection) updateTableName(table string) *collection {
	c.ALL = field.NewAsterisk(table)
	c.ID = field.NewInt64(table, "id")
	c.UserID = field.NewInt64(table, "user_id")
	c.Title = field.NewString(table, "title")
	c.Description = field.NewString(table, "description")
	c.IsDeleted = field.NewInt32(table, "is_deleted")
	c.CreateTime = field.NewTime(table, "create_time")
	c.UpdateTime = field.NewTime(table, "update_time")

	c.fillFieldMap()

	return c
}

func (c *collection) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := c.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (c *collection) fillFieldMap() {
	c.fieldMap = make(map[string]field.Expr, 7)
	c.fieldMap["id"] = c.ID
	c.fieldMap["user_id"] = c.UserID
	c.fieldMap["title"] = c.Title
	c.fieldMap["description"] = c.Description
	c.fieldMap["is_deleted"] = c.IsDeleted
	c.fieldMap["create_time"] = c.CreateTime
	c.fieldMap["update_time"] = c.UpdateTime
}

func (c collection) clone(db *gorm.DB) collection {
	c.collectionDo.ReplaceConnPool(db.Statement.ConnPool)
	return c
}

func (c collection) replaceDB(db *gorm.DB) collection {
	c.collectionDo.ReplaceDB(db)
	return c
}

type collectionDo struct{ gen.DO }

type ICollectionDo interface {
	gen.SubQuery
	Debug() ICollectionDo
	WithContext(ctx context.Context) ICollectionDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ICollectionDo
	WriteDB() ICollectionDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ICollectionDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ICollectionDo
	Not(conds ...gen.Condition) ICollectionDo
	Or(conds ...gen.Condition) ICollectionDo
	Select(conds ...field.Expr) ICollectionDo
	Where(conds ...gen.Condition) ICollectionDo
	Order(conds ...field.Expr) ICollectionDo
	Distinct(cols ...field.Expr) ICollectionDo
	Omit(cols ...field.Expr) ICollectionDo
	Join(table schema.Tabler, on ...field.Expr) ICollectionDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ICollectionDo
	RightJoin(table schema.Tabler, on ...field.Expr) ICollectionDo
	Group(cols ...field.Expr) ICollectionDo
	Having(conds ...gen.Condition) ICollectionDo
	Limit(limit int) ICollectionDo
	Offset(offset int) ICollectionDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ICollectionDo
	Unscoped() ICollectionDo
	Create(values ...*model.Collection) error
	CreateInBatches(values []*model.Collection, batchSize int) error
	Save(values ...*model.Collection) error
	First() (*model.Collection, error)
	Take() (*model.Collection, error)
	Last() (*model.Collection, error)
	Find() ([]*model.Collection, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Collection, err error)
	FindInBatches(result *[]*model.Collection, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.Collection) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ICollectionDo
	Assign(attrs ...field.AssignExpr) ICollectionDo
	Joins(fields ...field.RelationField) ICollectionDo
	Preload(fields ...field.RelationField) ICollectionDo
	FirstOrInit() (*model.Collection, error)
	FirstOrCreate() (*model.Collection, error)
	FindByPage(offset int, limit int) (result []*model.Collection, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ICollectionDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (c collectionDo) Debug() ICollectionDo {
	return c.withDO(c.DO.Debug())
}

func (c collectionDo) WithContext(ctx context.Context) ICollectionDo {
	return c.withDO(c.DO.WithContext(ctx))
}

func (c collectionDo) ReadDB() ICollectionDo {
	return c.Clauses(dbresolver.Read)
}

func (c collectionDo) WriteDB() ICollectionDo {
	return c.Clauses(dbresolver.Write)
}

func (c collectionDo) Session(config *gorm.Session) ICollectionDo {
	return c.withDO(c.DO.Session(config))
}

func (c collectionDo) Clauses(conds ...clause.Expression) ICollectionDo {
	return c.withDO(c.DO.Clauses(conds...))
}

func (c collectionDo) Returning(value interface{}, columns ...string) ICollectionDo {
	return c.withDO(c.DO.Returning(value, columns...))
}

func (c collectionDo) Not(conds ...gen.Condition) ICollectionDo {
	return c.withDO(c.DO.Not(conds...))
}

func (c collectionDo) Or(conds ...gen.Condition) ICollectionDo {
	return c.withDO(c.DO.Or(conds...))
}

func (c collectionDo) Select(conds ...field.Expr) ICollectionDo {
	return c.withDO(c.DO.Select(conds...))
}

func (c collectionDo) Where(conds ...gen.Condition) ICollectionDo {
	return c.withDO(c.DO.Where(conds...))
}

func (c collectionDo) Order(conds ...field.Expr) ICollectionDo {
	return c.withDO(c.DO.Order(conds...))
}

func (c collectionDo) Distinct(cols ...field.Expr) ICollectionDo {
	return c.withDO(c.DO.Distinct(cols...))
}

func (c collectionDo) Omit(cols ...field.Expr) ICollectionDo {
	return c.withDO(c.DO.Omit(cols...))
}

func (c collectionDo) Join(table schema.Tabler, on ...field.Expr) ICollectionDo {
	return c.withDO(c.DO.Join(table, on...))
}

func (c collectionDo) LeftJoin(table schema.Tabler, on ...field.Expr) ICollectionDo {
	return c.withDO(c.DO.LeftJoin(table, on...))
}

func (c collectionDo) RightJoin(table schema.Tabler, on ...field.Expr) ICollectionDo {
	return c.withDO(c.DO.RightJoin(table, on...))
}

func (c collectionDo) Group(cols ...field.Expr) ICollectionDo {
	return c.withDO(c.DO.Group(cols...))
}

func (c collectionDo) Having(conds ...gen.Condition) ICollectionDo {
	return c.withDO(c.DO.Having(conds...))
}

func (c collectionDo) Limit(limit int) ICollectionDo {
	return c.withDO(c.DO.Limit(limit))
}

func (c collectionDo) Offset(offset int) ICollectionDo {
	return c.withDO(c.DO.Offset(offset))
}

func (c collectionDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ICollectionDo {
	return c.withDO(c.DO.Scopes(funcs...))
}

func (c collectionDo) Unscoped() ICollectionDo {
	return c.withDO(c.DO.Unscoped())
}

func (c collectionDo) Create(values ...*model.Collection) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Create(values)
}

func (c collectionDo) CreateInBatches(values []*model.Collection, batchSize int) error {
	return c.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (c collectionDo) Save(values ...*model.Collection) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Save(values)
}

func (c collectionDo) First() (*model.Collection, error) {
	if result, err := c.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Collection), nil
	}
}

func (c collectionDo) Take() (*model.Collection, error) {
	if result, err := c.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Collection), nil
	}
}

func (c collectionDo) Last() (*model.Collection, error) {
	if result, err := c.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Collection), nil
	}
}

func (c collectionDo) Find() ([]*model.Collection, error) {
	result, err := c.DO.Find()
	return result.([]*model.Collection), err
}

func (c collectionDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Collection, err error) {
	buf := make([]*model.Collection, 0, batchSize)
	err = c.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (c collectionDo) FindInBatches(result *[]*model.Collection, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return c.DO.FindInBatches(result, batchSize, fc)
}

func (c collectionDo) Attrs(attrs ...field.AssignExpr) ICollectionDo {
	return c.withDO(c.DO.Attrs(attrs...))
}

func (c collectionDo) Assign(attrs ...field.AssignExpr) ICollectionDo {
	return c.withDO(c.DO.Assign(attrs...))
}

func (c collectionDo) Joins(fields ...field.RelationField) ICollectionDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Joins(_f))
	}
	return &c
}

func (c collectionDo) Preload(fields ...field.RelationField) ICollectionDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Preload(_f))
	}
	return &c
}

func (c collectionDo) FirstOrInit() (*model.Collection, error) {
	if result, err := c.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Collection), nil
	}
}

func (c collectionDo) FirstOrCreate() (*model.Collection, error) {
	if result, err := c.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Collection), nil
	}
}

func (c collectionDo) FindByPage(offset int, limit int) (result []*model.Collection, count int64, err error) {
	result, err = c.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = c.Offset(-1).Limit(-1).Count()
	return
}

func (c collectionDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = c.Count()
	if err != nil {
		return
	}

	err = c.Offset(offset).Limit(limit).Scan(result)
	return
}

func (c collectionDo) Scan(result interface{}) (err error) {
	return c.DO.Scan(result)
}

func (c collectionDo) Delete(models ...*model.Collection) (result gen.ResultInfo, err error) {
	return c.DO.Delete(models)
}

func (c *collectionDo) withDO(do gen.Dao) *collectionDo {
	c.DO = *do.(*gen.DO)
	return c
}

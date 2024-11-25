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

func newAccount(db *gorm.DB, opts ...gen.DOOption) account {
	_account := account{}

	_account.accountDo.UseDB(db, opts...)
	_account.accountDo.UseModel(&model.Account{})

	tableName := _account.accountDo.TableName()
	_account.ALL = field.NewAsterisk(tableName)
	_account.ID = field.NewInt64(tableName, "id")
	_account.Mobile = field.NewString(tableName, "mobile")
	_account.Email = field.NewString(tableName, "email")
	_account.Password = field.NewString(tableName, "password")
	_account.Salt = field.NewString(tableName, "salt")
	_account.Number = field.NewString(tableName, "number")
	_account.IsDeleted = field.NewBool(tableName, "is_deleted")
	_account.CreateTime = field.NewTime(tableName, "create_time")
	_account.UpdateTime = field.NewTime(tableName, "update_time")

	_account.fillFieldMap()

	return _account
}

type account struct {
	accountDo

	ALL        field.Asterisk
	ID         field.Int64  // 账户ID
	Mobile     field.String // 手机号
	Email      field.String // 电子邮件
	Password   field.String // 密码
	Salt       field.String // 密码盐
	Number     field.String // 号码
	IsDeleted  field.Bool   // 是否删除
	CreateTime field.Time   // 创建时间
	UpdateTime field.Time   // 更新时间

	fieldMap map[string]field.Expr
}

func (a account) Table(newTableName string) *account {
	a.accountDo.UseTable(newTableName)
	return a.updateTableName(newTableName)
}

func (a account) As(alias string) *account {
	a.accountDo.DO = *(a.accountDo.As(alias).(*gen.DO))
	return a.updateTableName(alias)
}

func (a *account) updateTableName(table string) *account {
	a.ALL = field.NewAsterisk(table)
	a.ID = field.NewInt64(table, "id")
	a.Mobile = field.NewString(table, "mobile")
	a.Email = field.NewString(table, "email")
	a.Password = field.NewString(table, "password")
	a.Salt = field.NewString(table, "salt")
	a.Number = field.NewString(table, "number")
	a.IsDeleted = field.NewBool(table, "is_deleted")
	a.CreateTime = field.NewTime(table, "create_time")
	a.UpdateTime = field.NewTime(table, "update_time")

	a.fillFieldMap()

	return a
}

func (a *account) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := a.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (a *account) fillFieldMap() {
	a.fieldMap = make(map[string]field.Expr, 9)
	a.fieldMap["id"] = a.ID
	a.fieldMap["mobile"] = a.Mobile
	a.fieldMap["email"] = a.Email
	a.fieldMap["password"] = a.Password
	a.fieldMap["salt"] = a.Salt
	a.fieldMap["number"] = a.Number
	a.fieldMap["is_deleted"] = a.IsDeleted
	a.fieldMap["create_time"] = a.CreateTime
	a.fieldMap["update_time"] = a.UpdateTime
}

func (a account) clone(db *gorm.DB) account {
	a.accountDo.ReplaceConnPool(db.Statement.ConnPool)
	return a
}

func (a account) replaceDB(db *gorm.DB) account {
	a.accountDo.ReplaceDB(db)
	return a
}

type accountDo struct{ gen.DO }

type IAccountDo interface {
	gen.SubQuery
	Debug() IAccountDo
	WithContext(ctx context.Context) IAccountDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IAccountDo
	WriteDB() IAccountDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IAccountDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IAccountDo
	Not(conds ...gen.Condition) IAccountDo
	Or(conds ...gen.Condition) IAccountDo
	Select(conds ...field.Expr) IAccountDo
	Where(conds ...gen.Condition) IAccountDo
	Order(conds ...field.Expr) IAccountDo
	Distinct(cols ...field.Expr) IAccountDo
	Omit(cols ...field.Expr) IAccountDo
	Join(table schema.Tabler, on ...field.Expr) IAccountDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IAccountDo
	RightJoin(table schema.Tabler, on ...field.Expr) IAccountDo
	Group(cols ...field.Expr) IAccountDo
	Having(conds ...gen.Condition) IAccountDo
	Limit(limit int) IAccountDo
	Offset(offset int) IAccountDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IAccountDo
	Unscoped() IAccountDo
	Create(values ...*model.Account) error
	CreateInBatches(values []*model.Account, batchSize int) error
	Save(values ...*model.Account) error
	First() (*model.Account, error)
	Take() (*model.Account, error)
	Last() (*model.Account, error)
	Find() ([]*model.Account, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Account, err error)
	FindInBatches(result *[]*model.Account, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.Account) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IAccountDo
	Assign(attrs ...field.AssignExpr) IAccountDo
	Joins(fields ...field.RelationField) IAccountDo
	Preload(fields ...field.RelationField) IAccountDo
	FirstOrInit() (*model.Account, error)
	FirstOrCreate() (*model.Account, error)
	FindByPage(offset int, limit int) (result []*model.Account, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IAccountDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (a accountDo) Debug() IAccountDo {
	return a.withDO(a.DO.Debug())
}

func (a accountDo) WithContext(ctx context.Context) IAccountDo {
	return a.withDO(a.DO.WithContext(ctx))
}

func (a accountDo) ReadDB() IAccountDo {
	return a.Clauses(dbresolver.Read)
}

func (a accountDo) WriteDB() IAccountDo {
	return a.Clauses(dbresolver.Write)
}

func (a accountDo) Session(config *gorm.Session) IAccountDo {
	return a.withDO(a.DO.Session(config))
}

func (a accountDo) Clauses(conds ...clause.Expression) IAccountDo {
	return a.withDO(a.DO.Clauses(conds...))
}

func (a accountDo) Returning(value interface{}, columns ...string) IAccountDo {
	return a.withDO(a.DO.Returning(value, columns...))
}

func (a accountDo) Not(conds ...gen.Condition) IAccountDo {
	return a.withDO(a.DO.Not(conds...))
}

func (a accountDo) Or(conds ...gen.Condition) IAccountDo {
	return a.withDO(a.DO.Or(conds...))
}

func (a accountDo) Select(conds ...field.Expr) IAccountDo {
	return a.withDO(a.DO.Select(conds...))
}

func (a accountDo) Where(conds ...gen.Condition) IAccountDo {
	return a.withDO(a.DO.Where(conds...))
}

func (a accountDo) Order(conds ...field.Expr) IAccountDo {
	return a.withDO(a.DO.Order(conds...))
}

func (a accountDo) Distinct(cols ...field.Expr) IAccountDo {
	return a.withDO(a.DO.Distinct(cols...))
}

func (a accountDo) Omit(cols ...field.Expr) IAccountDo {
	return a.withDO(a.DO.Omit(cols...))
}

func (a accountDo) Join(table schema.Tabler, on ...field.Expr) IAccountDo {
	return a.withDO(a.DO.Join(table, on...))
}

func (a accountDo) LeftJoin(table schema.Tabler, on ...field.Expr) IAccountDo {
	return a.withDO(a.DO.LeftJoin(table, on...))
}

func (a accountDo) RightJoin(table schema.Tabler, on ...field.Expr) IAccountDo {
	return a.withDO(a.DO.RightJoin(table, on...))
}

func (a accountDo) Group(cols ...field.Expr) IAccountDo {
	return a.withDO(a.DO.Group(cols...))
}

func (a accountDo) Having(conds ...gen.Condition) IAccountDo {
	return a.withDO(a.DO.Having(conds...))
}

func (a accountDo) Limit(limit int) IAccountDo {
	return a.withDO(a.DO.Limit(limit))
}

func (a accountDo) Offset(offset int) IAccountDo {
	return a.withDO(a.DO.Offset(offset))
}

func (a accountDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IAccountDo {
	return a.withDO(a.DO.Scopes(funcs...))
}

func (a accountDo) Unscoped() IAccountDo {
	return a.withDO(a.DO.Unscoped())
}

func (a accountDo) Create(values ...*model.Account) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Create(values)
}

func (a accountDo) CreateInBatches(values []*model.Account, batchSize int) error {
	return a.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (a accountDo) Save(values ...*model.Account) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Save(values)
}

func (a accountDo) First() (*model.Account, error) {
	if result, err := a.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Account), nil
	}
}

func (a accountDo) Take() (*model.Account, error) {
	if result, err := a.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Account), nil
	}
}

func (a accountDo) Last() (*model.Account, error) {
	if result, err := a.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Account), nil
	}
}

func (a accountDo) Find() ([]*model.Account, error) {
	result, err := a.DO.Find()
	return result.([]*model.Account), err
}

func (a accountDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Account, err error) {
	buf := make([]*model.Account, 0, batchSize)
	err = a.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (a accountDo) FindInBatches(result *[]*model.Account, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return a.DO.FindInBatches(result, batchSize, fc)
}

func (a accountDo) Attrs(attrs ...field.AssignExpr) IAccountDo {
	return a.withDO(a.DO.Attrs(attrs...))
}

func (a accountDo) Assign(attrs ...field.AssignExpr) IAccountDo {
	return a.withDO(a.DO.Assign(attrs...))
}

func (a accountDo) Joins(fields ...field.RelationField) IAccountDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Joins(_f))
	}
	return &a
}

func (a accountDo) Preload(fields ...field.RelationField) IAccountDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Preload(_f))
	}
	return &a
}

func (a accountDo) FirstOrInit() (*model.Account, error) {
	if result, err := a.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Account), nil
	}
}

func (a accountDo) FirstOrCreate() (*model.Account, error) {
	if result, err := a.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Account), nil
	}
}

func (a accountDo) FindByPage(offset int, limit int) (result []*model.Account, count int64, err error) {
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

func (a accountDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = a.Count()
	if err != nil {
		return
	}

	err = a.Offset(offset).Limit(limit).Scan(result)
	return
}

func (a accountDo) Scan(result interface{}) (err error) {
	return a.DO.Scan(result)
}

func (a accountDo) Delete(models ...*model.Account) (result gen.ResultInfo, err error) {
	return a.DO.Delete(models)
}

func (a *accountDo) withDO(do gen.Dao) *accountDo {
	a.DO = *do.(*gen.DO)
	return a
}

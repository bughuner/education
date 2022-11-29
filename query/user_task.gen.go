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

	"education/model"
)

func newUserTask(db *gorm.DB, opts ...gen.DOOption) userTask {
	_userTask := userTask{}

	_userTask.userTaskDo.UseDB(db, opts...)
	_userTask.userTaskDo.UseModel(&model.UserTask{})

	tableName := _userTask.userTaskDo.TableName()
	_userTask.ALL = field.NewAsterisk(tableName)
	_userTask.ID = field.NewString(tableName, "id")
	_userTask.UserID = field.NewString(tableName, "user_id")
	_userTask.TaskID = field.NewString(tableName, "task_id")

	_userTask.fillFieldMap()

	return _userTask
}

type userTask struct {
	userTaskDo userTaskDo

	ALL    field.Asterisk
	ID     field.String // 用户持有任务id
	UserID field.String // 用户id
	TaskID field.String // 任务id

	fieldMap map[string]field.Expr
}

func (u userTask) Table(newTableName string) *userTask {
	u.userTaskDo.UseTable(newTableName)
	return u.updateTableName(newTableName)
}

func (u userTask) As(alias string) *userTask {
	u.userTaskDo.DO = *(u.userTaskDo.As(alias).(*gen.DO))
	return u.updateTableName(alias)
}

func (u *userTask) updateTableName(table string) *userTask {
	u.ALL = field.NewAsterisk(table)
	u.ID = field.NewString(table, "id")
	u.UserID = field.NewString(table, "user_id")
	u.TaskID = field.NewString(table, "task_id")

	u.fillFieldMap()

	return u
}

func (u *userTask) WithContext(ctx context.Context) IUserTaskDo { return u.userTaskDo.WithContext(ctx) }

func (u userTask) TableName() string { return u.userTaskDo.TableName() }

func (u userTask) Alias() string { return u.userTaskDo.Alias() }

func (u *userTask) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := u.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (u *userTask) fillFieldMap() {
	u.fieldMap = make(map[string]field.Expr, 3)
	u.fieldMap["id"] = u.ID
	u.fieldMap["user_id"] = u.UserID
	u.fieldMap["task_id"] = u.TaskID
}

func (u userTask) clone(db *gorm.DB) userTask {
	u.userTaskDo.ReplaceConnPool(db.Statement.ConnPool)
	return u
}

func (u userTask) replaceDB(db *gorm.DB) userTask {
	u.userTaskDo.ReplaceDB(db)
	return u
}

type userTaskDo struct{ gen.DO }

type IUserTaskDo interface {
	gen.SubQuery
	Debug() IUserTaskDo
	WithContext(ctx context.Context) IUserTaskDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IUserTaskDo
	WriteDB() IUserTaskDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IUserTaskDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IUserTaskDo
	Not(conds ...gen.Condition) IUserTaskDo
	Or(conds ...gen.Condition) IUserTaskDo
	Select(conds ...field.Expr) IUserTaskDo
	Where(conds ...gen.Condition) IUserTaskDo
	Order(conds ...field.Expr) IUserTaskDo
	Distinct(cols ...field.Expr) IUserTaskDo
	Omit(cols ...field.Expr) IUserTaskDo
	Join(table schema.Tabler, on ...field.Expr) IUserTaskDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IUserTaskDo
	RightJoin(table schema.Tabler, on ...field.Expr) IUserTaskDo
	Group(cols ...field.Expr) IUserTaskDo
	Having(conds ...gen.Condition) IUserTaskDo
	Limit(limit int) IUserTaskDo
	Offset(offset int) IUserTaskDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IUserTaskDo
	Unscoped() IUserTaskDo
	Create(values ...*model.UserTask) error
	CreateInBatches(values []*model.UserTask, batchSize int) error
	Save(values ...*model.UserTask) error
	First() (*model.UserTask, error)
	Take() (*model.UserTask, error)
	Last() (*model.UserTask, error)
	Find() ([]*model.UserTask, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.UserTask, err error)
	FindInBatches(result *[]*model.UserTask, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.UserTask) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IUserTaskDo
	Assign(attrs ...field.AssignExpr) IUserTaskDo
	Joins(fields ...field.RelationField) IUserTaskDo
	Preload(fields ...field.RelationField) IUserTaskDo
	FirstOrInit() (*model.UserTask, error)
	FirstOrCreate() (*model.UserTask, error)
	FindByPage(offset int, limit int) (result []*model.UserTask, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IUserTaskDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (u userTaskDo) Debug() IUserTaskDo {
	return u.withDO(u.DO.Debug())
}

func (u userTaskDo) WithContext(ctx context.Context) IUserTaskDo {
	return u.withDO(u.DO.WithContext(ctx))
}

func (u userTaskDo) ReadDB() IUserTaskDo {
	return u.Clauses(dbresolver.Read)
}

func (u userTaskDo) WriteDB() IUserTaskDo {
	return u.Clauses(dbresolver.Write)
}

func (u userTaskDo) Session(config *gorm.Session) IUserTaskDo {
	return u.withDO(u.DO.Session(config))
}

func (u userTaskDo) Clauses(conds ...clause.Expression) IUserTaskDo {
	return u.withDO(u.DO.Clauses(conds...))
}

func (u userTaskDo) Returning(value interface{}, columns ...string) IUserTaskDo {
	return u.withDO(u.DO.Returning(value, columns...))
}

func (u userTaskDo) Not(conds ...gen.Condition) IUserTaskDo {
	return u.withDO(u.DO.Not(conds...))
}

func (u userTaskDo) Or(conds ...gen.Condition) IUserTaskDo {
	return u.withDO(u.DO.Or(conds...))
}

func (u userTaskDo) Select(conds ...field.Expr) IUserTaskDo {
	return u.withDO(u.DO.Select(conds...))
}

func (u userTaskDo) Where(conds ...gen.Condition) IUserTaskDo {
	return u.withDO(u.DO.Where(conds...))
}

func (u userTaskDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) IUserTaskDo {
	return u.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (u userTaskDo) Order(conds ...field.Expr) IUserTaskDo {
	return u.withDO(u.DO.Order(conds...))
}

func (u userTaskDo) Distinct(cols ...field.Expr) IUserTaskDo {
	return u.withDO(u.DO.Distinct(cols...))
}

func (u userTaskDo) Omit(cols ...field.Expr) IUserTaskDo {
	return u.withDO(u.DO.Omit(cols...))
}

func (u userTaskDo) Join(table schema.Tabler, on ...field.Expr) IUserTaskDo {
	return u.withDO(u.DO.Join(table, on...))
}

func (u userTaskDo) LeftJoin(table schema.Tabler, on ...field.Expr) IUserTaskDo {
	return u.withDO(u.DO.LeftJoin(table, on...))
}

func (u userTaskDo) RightJoin(table schema.Tabler, on ...field.Expr) IUserTaskDo {
	return u.withDO(u.DO.RightJoin(table, on...))
}

func (u userTaskDo) Group(cols ...field.Expr) IUserTaskDo {
	return u.withDO(u.DO.Group(cols...))
}

func (u userTaskDo) Having(conds ...gen.Condition) IUserTaskDo {
	return u.withDO(u.DO.Having(conds...))
}

func (u userTaskDo) Limit(limit int) IUserTaskDo {
	return u.withDO(u.DO.Limit(limit))
}

func (u userTaskDo) Offset(offset int) IUserTaskDo {
	return u.withDO(u.DO.Offset(offset))
}

func (u userTaskDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IUserTaskDo {
	return u.withDO(u.DO.Scopes(funcs...))
}

func (u userTaskDo) Unscoped() IUserTaskDo {
	return u.withDO(u.DO.Unscoped())
}

func (u userTaskDo) Create(values ...*model.UserTask) error {
	if len(values) == 0 {
		return nil
	}
	return u.DO.Create(values)
}

func (u userTaskDo) CreateInBatches(values []*model.UserTask, batchSize int) error {
	return u.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (u userTaskDo) Save(values ...*model.UserTask) error {
	if len(values) == 0 {
		return nil
	}
	return u.DO.Save(values)
}

func (u userTaskDo) First() (*model.UserTask, error) {
	if result, err := u.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserTask), nil
	}
}

func (u userTaskDo) Take() (*model.UserTask, error) {
	if result, err := u.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserTask), nil
	}
}

func (u userTaskDo) Last() (*model.UserTask, error) {
	if result, err := u.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserTask), nil
	}
}

func (u userTaskDo) Find() ([]*model.UserTask, error) {
	result, err := u.DO.Find()
	return result.([]*model.UserTask), err
}

func (u userTaskDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.UserTask, err error) {
	buf := make([]*model.UserTask, 0, batchSize)
	err = u.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (u userTaskDo) FindInBatches(result *[]*model.UserTask, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return u.DO.FindInBatches(result, batchSize, fc)
}

func (u userTaskDo) Attrs(attrs ...field.AssignExpr) IUserTaskDo {
	return u.withDO(u.DO.Attrs(attrs...))
}

func (u userTaskDo) Assign(attrs ...field.AssignExpr) IUserTaskDo {
	return u.withDO(u.DO.Assign(attrs...))
}

func (u userTaskDo) Joins(fields ...field.RelationField) IUserTaskDo {
	for _, _f := range fields {
		u = *u.withDO(u.DO.Joins(_f))
	}
	return &u
}

func (u userTaskDo) Preload(fields ...field.RelationField) IUserTaskDo {
	for _, _f := range fields {
		u = *u.withDO(u.DO.Preload(_f))
	}
	return &u
}

func (u userTaskDo) FirstOrInit() (*model.UserTask, error) {
	if result, err := u.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserTask), nil
	}
}

func (u userTaskDo) FirstOrCreate() (*model.UserTask, error) {
	if result, err := u.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserTask), nil
	}
}

func (u userTaskDo) FindByPage(offset int, limit int) (result []*model.UserTask, count int64, err error) {
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

func (u userTaskDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = u.Count()
	if err != nil {
		return
	}

	err = u.Offset(offset).Limit(limit).Scan(result)
	return
}

func (u userTaskDo) Scan(result interface{}) (err error) {
	return u.DO.Scan(result)
}

func (u userTaskDo) Delete(models ...*model.UserTask) (result gen.ResultInfo, err error) {
	return u.DO.Delete(models)
}

func (u *userTaskDo) withDO(do gen.Dao) *userTaskDo {
	u.DO = *do.(*gen.DO)
	return u
}

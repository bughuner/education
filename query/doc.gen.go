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

func newDoc(db *gorm.DB, opts ...gen.DOOption) doc {
	_doc := doc{}

	_doc.docDo.UseDB(db, opts...)
	_doc.docDo.UseModel(&model.Doc{})

	tableName := _doc.docDo.TableName()
	_doc.ALL = field.NewAsterisk(tableName)
	_doc.ID = field.NewString(tableName, "id")
	_doc.Link = field.NewString(tableName, "link")
	_doc.Author = field.NewString(tableName, "author")
	_doc.Content = field.NewString(tableName, "content")
	_doc.Status = field.NewInt64(tableName, "status")
	_doc.Type = field.NewString(tableName, "type")
	_doc.Label = field.NewString(tableName, "label")
	_doc.Title = field.NewString(tableName, "title")

	_doc.fillFieldMap()

	return _doc
}

type doc struct {
	docDo docDo

	ALL     field.Asterisk
	ID      field.String // 文档
	Link    field.String // 链接
	Author  field.String // 作者
	Content field.String // 内容
	Status  field.Int64  // 状态 0-待审核 1-审核
	Type    field.String // 类型
	Label   field.String // 标签
	Title   field.String // 标题

	fieldMap map[string]field.Expr
}

func (d doc) Table(newTableName string) *doc {
	d.docDo.UseTable(newTableName)
	return d.updateTableName(newTableName)
}

func (d doc) As(alias string) *doc {
	d.docDo.DO = *(d.docDo.As(alias).(*gen.DO))
	return d.updateTableName(alias)
}

func (d *doc) updateTableName(table string) *doc {
	d.ALL = field.NewAsterisk(table)
	d.ID = field.NewString(table, "id")
	d.Link = field.NewString(table, "link")
	d.Author = field.NewString(table, "author")
	d.Content = field.NewString(table, "content")
	d.Status = field.NewInt64(table, "status")
	d.Type = field.NewString(table, "type")
	d.Label = field.NewString(table, "label")
	d.Title = field.NewString(table, "title")

	d.fillFieldMap()

	return d
}

func (d *doc) WithContext(ctx context.Context) IDocDo { return d.docDo.WithContext(ctx) }

func (d doc) TableName() string { return d.docDo.TableName() }

func (d doc) Alias() string { return d.docDo.Alias() }

func (d *doc) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := d.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (d *doc) fillFieldMap() {
	d.fieldMap = make(map[string]field.Expr, 8)
	d.fieldMap["id"] = d.ID
	d.fieldMap["link"] = d.Link
	d.fieldMap["author"] = d.Author
	d.fieldMap["content"] = d.Content
	d.fieldMap["status"] = d.Status
	d.fieldMap["type"] = d.Type
	d.fieldMap["label"] = d.Label
	d.fieldMap["title"] = d.Title
}

func (d doc) clone(db *gorm.DB) doc {
	d.docDo.ReplaceConnPool(db.Statement.ConnPool)
	return d
}

func (d doc) replaceDB(db *gorm.DB) doc {
	d.docDo.ReplaceDB(db)
	return d
}

type docDo struct{ gen.DO }

type IDocDo interface {
	gen.SubQuery
	Debug() IDocDo
	WithContext(ctx context.Context) IDocDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IDocDo
	WriteDB() IDocDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IDocDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IDocDo
	Not(conds ...gen.Condition) IDocDo
	Or(conds ...gen.Condition) IDocDo
	Select(conds ...field.Expr) IDocDo
	Where(conds ...gen.Condition) IDocDo
	Order(conds ...field.Expr) IDocDo
	Distinct(cols ...field.Expr) IDocDo
	Omit(cols ...field.Expr) IDocDo
	Join(table schema.Tabler, on ...field.Expr) IDocDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IDocDo
	RightJoin(table schema.Tabler, on ...field.Expr) IDocDo
	Group(cols ...field.Expr) IDocDo
	Having(conds ...gen.Condition) IDocDo
	Limit(limit int) IDocDo
	Offset(offset int) IDocDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IDocDo
	Unscoped() IDocDo
	Create(values ...*model.Doc) error
	CreateInBatches(values []*model.Doc, batchSize int) error
	Save(values ...*model.Doc) error
	First() (*model.Doc, error)
	Take() (*model.Doc, error)
	Last() (*model.Doc, error)
	Find() ([]*model.Doc, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Doc, err error)
	FindInBatches(result *[]*model.Doc, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.Doc) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IDocDo
	Assign(attrs ...field.AssignExpr) IDocDo
	Joins(fields ...field.RelationField) IDocDo
	Preload(fields ...field.RelationField) IDocDo
	FirstOrInit() (*model.Doc, error)
	FirstOrCreate() (*model.Doc, error)
	FindByPage(offset int, limit int) (result []*model.Doc, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IDocDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (d docDo) Debug() IDocDo {
	return d.withDO(d.DO.Debug())
}

func (d docDo) WithContext(ctx context.Context) IDocDo {
	return d.withDO(d.DO.WithContext(ctx))
}

func (d docDo) ReadDB() IDocDo {
	return d.Clauses(dbresolver.Read)
}

func (d docDo) WriteDB() IDocDo {
	return d.Clauses(dbresolver.Write)
}

func (d docDo) Session(config *gorm.Session) IDocDo {
	return d.withDO(d.DO.Session(config))
}

func (d docDo) Clauses(conds ...clause.Expression) IDocDo {
	return d.withDO(d.DO.Clauses(conds...))
}

func (d docDo) Returning(value interface{}, columns ...string) IDocDo {
	return d.withDO(d.DO.Returning(value, columns...))
}

func (d docDo) Not(conds ...gen.Condition) IDocDo {
	return d.withDO(d.DO.Not(conds...))
}

func (d docDo) Or(conds ...gen.Condition) IDocDo {
	return d.withDO(d.DO.Or(conds...))
}

func (d docDo) Select(conds ...field.Expr) IDocDo {
	return d.withDO(d.DO.Select(conds...))
}

func (d docDo) Where(conds ...gen.Condition) IDocDo {
	return d.withDO(d.DO.Where(conds...))
}

func (d docDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) IDocDo {
	return d.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (d docDo) Order(conds ...field.Expr) IDocDo {
	return d.withDO(d.DO.Order(conds...))
}

func (d docDo) Distinct(cols ...field.Expr) IDocDo {
	return d.withDO(d.DO.Distinct(cols...))
}

func (d docDo) Omit(cols ...field.Expr) IDocDo {
	return d.withDO(d.DO.Omit(cols...))
}

func (d docDo) Join(table schema.Tabler, on ...field.Expr) IDocDo {
	return d.withDO(d.DO.Join(table, on...))
}

func (d docDo) LeftJoin(table schema.Tabler, on ...field.Expr) IDocDo {
	return d.withDO(d.DO.LeftJoin(table, on...))
}

func (d docDo) RightJoin(table schema.Tabler, on ...field.Expr) IDocDo {
	return d.withDO(d.DO.RightJoin(table, on...))
}

func (d docDo) Group(cols ...field.Expr) IDocDo {
	return d.withDO(d.DO.Group(cols...))
}

func (d docDo) Having(conds ...gen.Condition) IDocDo {
	return d.withDO(d.DO.Having(conds...))
}

func (d docDo) Limit(limit int) IDocDo {
	return d.withDO(d.DO.Limit(limit))
}

func (d docDo) Offset(offset int) IDocDo {
	return d.withDO(d.DO.Offset(offset))
}

func (d docDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IDocDo {
	return d.withDO(d.DO.Scopes(funcs...))
}

func (d docDo) Unscoped() IDocDo {
	return d.withDO(d.DO.Unscoped())
}

func (d docDo) Create(values ...*model.Doc) error {
	if len(values) == 0 {
		return nil
	}
	return d.DO.Create(values)
}

func (d docDo) CreateInBatches(values []*model.Doc, batchSize int) error {
	return d.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (d docDo) Save(values ...*model.Doc) error {
	if len(values) == 0 {
		return nil
	}
	return d.DO.Save(values)
}

func (d docDo) First() (*model.Doc, error) {
	if result, err := d.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Doc), nil
	}
}

func (d docDo) Take() (*model.Doc, error) {
	if result, err := d.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Doc), nil
	}
}

func (d docDo) Last() (*model.Doc, error) {
	if result, err := d.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Doc), nil
	}
}

func (d docDo) Find() ([]*model.Doc, error) {
	result, err := d.DO.Find()
	return result.([]*model.Doc), err
}

func (d docDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Doc, err error) {
	buf := make([]*model.Doc, 0, batchSize)
	err = d.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (d docDo) FindInBatches(result *[]*model.Doc, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return d.DO.FindInBatches(result, batchSize, fc)
}

func (d docDo) Attrs(attrs ...field.AssignExpr) IDocDo {
	return d.withDO(d.DO.Attrs(attrs...))
}

func (d docDo) Assign(attrs ...field.AssignExpr) IDocDo {
	return d.withDO(d.DO.Assign(attrs...))
}

func (d docDo) Joins(fields ...field.RelationField) IDocDo {
	for _, _f := range fields {
		d = *d.withDO(d.DO.Joins(_f))
	}
	return &d
}

func (d docDo) Preload(fields ...field.RelationField) IDocDo {
	for _, _f := range fields {
		d = *d.withDO(d.DO.Preload(_f))
	}
	return &d
}

func (d docDo) FirstOrInit() (*model.Doc, error) {
	if result, err := d.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Doc), nil
	}
}

func (d docDo) FirstOrCreate() (*model.Doc, error) {
	if result, err := d.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Doc), nil
	}
}

func (d docDo) FindByPage(offset int, limit int) (result []*model.Doc, count int64, err error) {
	result, err = d.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = d.Offset(-1).Limit(-1).Count()
	return
}

func (d docDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = d.Count()
	if err != nil {
		return
	}

	err = d.Offset(offset).Limit(limit).Scan(result)
	return
}

func (d docDo) Scan(result interface{}) (err error) {
	return d.DO.Scan(result)
}

func (d docDo) Delete(models ...*model.Doc) (result gen.ResultInfo, err error) {
	return d.DO.Delete(models)
}

func (d *docDo) withDO(do gen.Dao) *docDo {
	d.DO = *do.(*gen.DO)
	return d
}

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

	"github.com/zjzjzjzj1874/best-pracrice-go-zero/__test__/gorm/dal/model"
)

func newTMock(db *gorm.DB) tMock {
	_tMock := tMock{}

	_tMock.tMockDo.UseDB(db)
	_tMock.tMockDo.UseModel(&model.TMock{})

	tableName := _tMock.tMockDo.TableName()
	_tMock.ALL = field.NewField(tableName, "*")
	_tMock.ID = field.NewInt64(tableName, "id")
	_tMock.CreatedAt = field.NewTime(tableName, "created_at")
	_tMock.UpdatedAt = field.NewTime(tableName, "updated_at")
	_tMock.DeletedAt = field.NewField(tableName, "deleted_at")
	_tMock.TestName = field.NewString(tableName, "test_name")
	_tMock.Hobbies = field.NewString(tableName, "hobbies")
	_tMock.SLATimes = field.NewString(tableName, "sla_times")

	_tMock.fillFieldMap()

	return _tMock
}

type tMock struct {
	tMockDo tMockDo

	ALL       field.Field
	ID        field.Int64
	CreatedAt field.Time
	UpdatedAt field.Time
	DeletedAt field.Field
	TestName  field.String
	Hobbies   field.String
	SLATimes  field.String

	fieldMap map[string]field.Expr
}

func (t tMock) Table(newTableName string) *tMock {
	t.tMockDo.UseTable(newTableName)
	return t.updateTableName(newTableName)
}

func (t tMock) As(alias string) *tMock {
	t.tMockDo.DO = *(t.tMockDo.As(alias).(*gen.DO))
	return t.updateTableName(alias)
}

func (t *tMock) updateTableName(table string) *tMock {
	t.ALL = field.NewField(table, "*")
	t.ID = field.NewInt64(table, "id")
	t.CreatedAt = field.NewTime(table, "created_at")
	t.UpdatedAt = field.NewTime(table, "updated_at")
	t.DeletedAt = field.NewField(table, "deleted_at")
	t.TestName = field.NewString(table, "test_name")
	t.Hobbies = field.NewString(table, "hobbies")
	t.SLATimes = field.NewString(table, "sla_times")

	t.fillFieldMap()

	return t
}

func (t *tMock) WithContext(ctx context.Context) *tMockDo { return t.tMockDo.WithContext(ctx) }

func (t tMock) TableName() string { return t.tMockDo.TableName() }

func (t tMock) Alias() string { return t.tMockDo.Alias() }

func (t *tMock) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := t.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (t *tMock) fillFieldMap() {
	t.fieldMap = make(map[string]field.Expr, 7)
	t.fieldMap["id"] = t.ID
	t.fieldMap["created_at"] = t.CreatedAt
	t.fieldMap["updated_at"] = t.UpdatedAt
	t.fieldMap["deleted_at"] = t.DeletedAt
	t.fieldMap["test_name"] = t.TestName
	t.fieldMap["hobbies"] = t.Hobbies
	t.fieldMap["sla_times"] = t.SLATimes
}

func (t tMock) clone(db *gorm.DB) tMock {
	t.tMockDo.ReplaceDB(db)
	return t
}

type tMockDo struct{ gen.DO }

func (t tMockDo) Debug() *tMockDo {
	return t.withDO(t.DO.Debug())
}

func (t tMockDo) WithContext(ctx context.Context) *tMockDo {
	return t.withDO(t.DO.WithContext(ctx))
}

func (t tMockDo) ReadDB() *tMockDo {
	return t.Clauses(dbresolver.Read)
}

func (t tMockDo) WriteDB() *tMockDo {
	return t.Clauses(dbresolver.Write)
}

func (t tMockDo) Clauses(conds ...clause.Expression) *tMockDo {
	return t.withDO(t.DO.Clauses(conds...))
}

func (t tMockDo) Returning(value interface{}, columns ...string) *tMockDo {
	return t.withDO(t.DO.Returning(value, columns...))
}

func (t tMockDo) Not(conds ...gen.Condition) *tMockDo {
	return t.withDO(t.DO.Not(conds...))
}

func (t tMockDo) Or(conds ...gen.Condition) *tMockDo {
	return t.withDO(t.DO.Or(conds...))
}

func (t tMockDo) Select(conds ...field.Expr) *tMockDo {
	return t.withDO(t.DO.Select(conds...))
}

func (t tMockDo) Where(conds ...gen.Condition) *tMockDo {
	return t.withDO(t.DO.Where(conds...))
}

func (t tMockDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *tMockDo {
	return t.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (t tMockDo) Order(conds ...field.Expr) *tMockDo {
	return t.withDO(t.DO.Order(conds...))
}

func (t tMockDo) Distinct(cols ...field.Expr) *tMockDo {
	return t.withDO(t.DO.Distinct(cols...))
}

func (t tMockDo) Omit(cols ...field.Expr) *tMockDo {
	return t.withDO(t.DO.Omit(cols...))
}

func (t tMockDo) Join(table schema.Tabler, on ...field.Expr) *tMockDo {
	return t.withDO(t.DO.Join(table, on...))
}

func (t tMockDo) LeftJoin(table schema.Tabler, on ...field.Expr) *tMockDo {
	return t.withDO(t.DO.LeftJoin(table, on...))
}

func (t tMockDo) RightJoin(table schema.Tabler, on ...field.Expr) *tMockDo {
	return t.withDO(t.DO.RightJoin(table, on...))
}

func (t tMockDo) Group(cols ...field.Expr) *tMockDo {
	return t.withDO(t.DO.Group(cols...))
}

func (t tMockDo) Having(conds ...gen.Condition) *tMockDo {
	return t.withDO(t.DO.Having(conds...))
}

func (t tMockDo) Limit(limit int) *tMockDo {
	return t.withDO(t.DO.Limit(limit))
}

func (t tMockDo) Offset(offset int) *tMockDo {
	return t.withDO(t.DO.Offset(offset))
}

func (t tMockDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *tMockDo {
	return t.withDO(t.DO.Scopes(funcs...))
}

func (t tMockDo) Unscoped() *tMockDo {
	return t.withDO(t.DO.Unscoped())
}

func (t tMockDo) Create(values ...*model.TMock) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Create(values)
}

func (t tMockDo) CreateInBatches(values []*model.TMock, batchSize int) error {
	return t.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (t tMockDo) Save(values ...*model.TMock) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Save(values)
}

func (t tMockDo) First() (*model.TMock, error) {
	if result, err := t.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.TMock), nil
	}
}

func (t tMockDo) Take() (*model.TMock, error) {
	if result, err := t.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.TMock), nil
	}
}

func (t tMockDo) Last() (*model.TMock, error) {
	if result, err := t.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.TMock), nil
	}
}

func (t tMockDo) Find() ([]*model.TMock, error) {
	result, err := t.DO.Find()
	return result.([]*model.TMock), err
}

func (t tMockDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.TMock, err error) {
	buf := make([]*model.TMock, 0, batchSize)
	err = t.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (t tMockDo) FindInBatches(result *[]*model.TMock, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return t.DO.FindInBatches(result, batchSize, fc)
}

func (t tMockDo) Attrs(attrs ...field.AssignExpr) *tMockDo {
	return t.withDO(t.DO.Attrs(attrs...))
}

func (t tMockDo) Assign(attrs ...field.AssignExpr) *tMockDo {
	return t.withDO(t.DO.Assign(attrs...))
}

func (t tMockDo) Joins(fields ...field.RelationField) *tMockDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Joins(_f))
	}
	return &t
}

func (t tMockDo) Preload(fields ...field.RelationField) *tMockDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Preload(_f))
	}
	return &t
}

func (t tMockDo) FirstOrInit() (*model.TMock, error) {
	if result, err := t.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.TMock), nil
	}
}

func (t tMockDo) FirstOrCreate() (*model.TMock, error) {
	if result, err := t.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.TMock), nil
	}
}

func (t tMockDo) FindByPage(offset int, limit int) (result []*model.TMock, count int64, err error) {
	result, err = t.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = t.Offset(-1).Limit(-1).Count()
	return
}

func (t tMockDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = t.Count()
	if err != nil {
		return
	}

	err = t.Offset(offset).Limit(limit).Scan(result)
	return
}

func (t tMockDo) Scan(result interface{}) (err error) {
	return t.DO.Scan(result)
}

func (t *tMockDo) withDO(do gen.Dao) *tMockDo {
	t.DO = *do.(*gen.DO)
	return t
}
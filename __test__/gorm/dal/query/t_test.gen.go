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

func newTTest(db *gorm.DB) tTest {
	_tTest := tTest{}

	_tTest.tTestDo.UseDB(db)
	_tTest.tTestDo.UseModel(&model.TTest{})

	tableName := _tTest.tTestDo.TableName()
	_tTest.ALL = field.NewField(tableName, "*")
	_tTest.ID = field.NewInt64(tableName, "id")
	_tTest.CreatedAt = field.NewTime(tableName, "created_at")
	_tTest.UpdatedAt = field.NewTime(tableName, "updated_at")
	_tTest.DeletedAt = field.NewField(tableName, "deleted_at")
	_tTest.ReceiveTime = field.NewTime(tableName, "receive_time")
	_tTest.UpdateTime = field.NewTime(tableName, "update_time")
	_tTest.CallbackTime = field.NewTime(tableName, "callback_time")
	_tTest.TestTime = field.NewTime(tableName, "test_time")
	_tTest.TestName = field.NewString(tableName, "test_name")
	_tTest.TestTime1 = field.NewTime(tableName, "test_time1")
	_tTest.Emotion = field.NewString(tableName, "emotion")

	_tTest.fillFieldMap()

	return _tTest
}

type tTest struct {
	tTestDo tTestDo

	ALL          field.Field
	ID           field.Int64
	CreatedAt    field.Time
	UpdatedAt    field.Time
	DeletedAt    field.Field
	ReceiveTime  field.Time
	UpdateTime   field.Time
	CallbackTime field.Time
	TestTime     field.Time
	TestName     field.String
	TestTime1    field.Time
	Emotion      field.String

	fieldMap map[string]field.Expr
}

func (t tTest) Table(newTableName string) *tTest {
	t.tTestDo.UseTable(newTableName)
	return t.updateTableName(newTableName)
}

func (t tTest) As(alias string) *tTest {
	t.tTestDo.DO = *(t.tTestDo.As(alias).(*gen.DO))
	return t.updateTableName(alias)
}

func (t *tTest) updateTableName(table string) *tTest {
	t.ALL = field.NewField(table, "*")
	t.ID = field.NewInt64(table, "id")
	t.CreatedAt = field.NewTime(table, "created_at")
	t.UpdatedAt = field.NewTime(table, "updated_at")
	t.DeletedAt = field.NewField(table, "deleted_at")
	t.ReceiveTime = field.NewTime(table, "receive_time")
	t.UpdateTime = field.NewTime(table, "update_time")
	t.CallbackTime = field.NewTime(table, "callback_time")
	t.TestTime = field.NewTime(table, "test_time")
	t.TestName = field.NewString(table, "test_name")
	t.TestTime1 = field.NewTime(table, "test_time1")
	t.Emotion = field.NewString(table, "emotion")

	t.fillFieldMap()

	return t
}

func (t *tTest) WithContext(ctx context.Context) *tTestDo { return t.tTestDo.WithContext(ctx) }

func (t tTest) TableName() string { return t.tTestDo.TableName() }

func (t tTest) Alias() string { return t.tTestDo.Alias() }

func (t *tTest) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := t.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (t *tTest) fillFieldMap() {
	t.fieldMap = make(map[string]field.Expr, 11)
	t.fieldMap["id"] = t.ID
	t.fieldMap["created_at"] = t.CreatedAt
	t.fieldMap["updated_at"] = t.UpdatedAt
	t.fieldMap["deleted_at"] = t.DeletedAt
	t.fieldMap["receive_time"] = t.ReceiveTime
	t.fieldMap["update_time"] = t.UpdateTime
	t.fieldMap["callback_time"] = t.CallbackTime
	t.fieldMap["test_time"] = t.TestTime
	t.fieldMap["test_name"] = t.TestName
	t.fieldMap["test_time1"] = t.TestTime1
	t.fieldMap["emotion"] = t.Emotion
}

func (t tTest) clone(db *gorm.DB) tTest {
	t.tTestDo.ReplaceDB(db)
	return t
}

type tTestDo struct{ gen.DO }

func (t tTestDo) Debug() *tTestDo {
	return t.withDO(t.DO.Debug())
}

func (t tTestDo) WithContext(ctx context.Context) *tTestDo {
	return t.withDO(t.DO.WithContext(ctx))
}

func (t tTestDo) ReadDB() *tTestDo {
	return t.Clauses(dbresolver.Read)
}

func (t tTestDo) WriteDB() *tTestDo {
	return t.Clauses(dbresolver.Write)
}

func (t tTestDo) Clauses(conds ...clause.Expression) *tTestDo {
	return t.withDO(t.DO.Clauses(conds...))
}

func (t tTestDo) Returning(value interface{}, columns ...string) *tTestDo {
	return t.withDO(t.DO.Returning(value, columns...))
}

func (t tTestDo) Not(conds ...gen.Condition) *tTestDo {
	return t.withDO(t.DO.Not(conds...))
}

func (t tTestDo) Or(conds ...gen.Condition) *tTestDo {
	return t.withDO(t.DO.Or(conds...))
}

func (t tTestDo) Select(conds ...field.Expr) *tTestDo {
	return t.withDO(t.DO.Select(conds...))
}

func (t tTestDo) Where(conds ...gen.Condition) *tTestDo {
	return t.withDO(t.DO.Where(conds...))
}

func (t tTestDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *tTestDo {
	return t.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (t tTestDo) Order(conds ...field.Expr) *tTestDo {
	return t.withDO(t.DO.Order(conds...))
}

func (t tTestDo) Distinct(cols ...field.Expr) *tTestDo {
	return t.withDO(t.DO.Distinct(cols...))
}

func (t tTestDo) Omit(cols ...field.Expr) *tTestDo {
	return t.withDO(t.DO.Omit(cols...))
}

func (t tTestDo) Join(table schema.Tabler, on ...field.Expr) *tTestDo {
	return t.withDO(t.DO.Join(table, on...))
}

func (t tTestDo) LeftJoin(table schema.Tabler, on ...field.Expr) *tTestDo {
	return t.withDO(t.DO.LeftJoin(table, on...))
}

func (t tTestDo) RightJoin(table schema.Tabler, on ...field.Expr) *tTestDo {
	return t.withDO(t.DO.RightJoin(table, on...))
}

func (t tTestDo) Group(cols ...field.Expr) *tTestDo {
	return t.withDO(t.DO.Group(cols...))
}

func (t tTestDo) Having(conds ...gen.Condition) *tTestDo {
	return t.withDO(t.DO.Having(conds...))
}

func (t tTestDo) Limit(limit int) *tTestDo {
	return t.withDO(t.DO.Limit(limit))
}

func (t tTestDo) Offset(offset int) *tTestDo {
	return t.withDO(t.DO.Offset(offset))
}

func (t tTestDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *tTestDo {
	return t.withDO(t.DO.Scopes(funcs...))
}

func (t tTestDo) Unscoped() *tTestDo {
	return t.withDO(t.DO.Unscoped())
}

func (t tTestDo) Create(values ...*model.TTest) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Create(values)
}

func (t tTestDo) CreateInBatches(values []*model.TTest, batchSize int) error {
	return t.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (t tTestDo) Save(values ...*model.TTest) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Save(values)
}

func (t tTestDo) First() (*model.TTest, error) {
	if result, err := t.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.TTest), nil
	}
}

func (t tTestDo) Take() (*model.TTest, error) {
	if result, err := t.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.TTest), nil
	}
}

func (t tTestDo) Last() (*model.TTest, error) {
	if result, err := t.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.TTest), nil
	}
}

func (t tTestDo) Find() ([]*model.TTest, error) {
	result, err := t.DO.Find()
	return result.([]*model.TTest), err
}

func (t tTestDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.TTest, err error) {
	buf := make([]*model.TTest, 0, batchSize)
	err = t.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (t tTestDo) FindInBatches(result *[]*model.TTest, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return t.DO.FindInBatches(result, batchSize, fc)
}

func (t tTestDo) Attrs(attrs ...field.AssignExpr) *tTestDo {
	return t.withDO(t.DO.Attrs(attrs...))
}

func (t tTestDo) Assign(attrs ...field.AssignExpr) *tTestDo {
	return t.withDO(t.DO.Assign(attrs...))
}

func (t tTestDo) Joins(fields ...field.RelationField) *tTestDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Joins(_f))
	}
	return &t
}

func (t tTestDo) Preload(fields ...field.RelationField) *tTestDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Preload(_f))
	}
	return &t
}

func (t tTestDo) FirstOrInit() (*model.TTest, error) {
	if result, err := t.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.TTest), nil
	}
}

func (t tTestDo) FirstOrCreate() (*model.TTest, error) {
	if result, err := t.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.TTest), nil
	}
}

func (t tTestDo) FindByPage(offset int, limit int) (result []*model.TTest, count int64, err error) {
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

func (t tTestDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = t.Count()
	if err != nil {
		return
	}

	err = t.Offset(offset).Limit(limit).Scan(result)
	return
}

func (t tTestDo) Scan(result interface{}) (err error) {
	return t.DO.Scan(result)
}

func (t *tTestDo) withDO(do gen.Dao) *tTestDo {
	t.DO = *do.(*gen.DO)
	return t
}
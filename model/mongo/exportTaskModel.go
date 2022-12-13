package mongo

import (
	"github.com/globalsign/mgo/bson"
	cachec "github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/mongoc"
	"strings"
)

var prefixExportTaskCacheKey = "cache:my_zero:ExportTask:"

type ExportTaskModel interface {
	Insert(data *ExportTask) error
	BatchInsert(data []*ExportTask) error
	FindOne(id string) (*ExportTask, error)
	Update(data *ExportTask) error
	Delete(id string) error
	// query by any params
	FindByAny(query interface{}) (*ExportTask, error)
	List(query interface{}, opts ...mongoc.QueryOption) ([]ExportTask, error)
	Count(query interface{}) (int, error)
}

type defaultExportTaskModel struct {
	*mongoc.Model
}

func NewExportTaskModel(url, collection string, c cachec.CacheConf) ExportTaskModel {
	return &defaultExportTaskModel{
		Model: mongoc.MustNewModel(url, collection, c),
	}
}

func (m *defaultExportTaskModel) Insert(data *ExportTask) error {
	if !data.ID.Valid() {
		data.ID = bson.NewObjectId()
	}

	session, err := m.TakeSession()
	if err != nil {
		return err
	}

	defer m.PutSession(session)
	return m.GetCollection(session).Insert(data)
}

func (m *defaultExportTaskModel) BatchInsert(datas []*ExportTask) error {
	var docs []interface{}

	for _, doc := range datas {
		if !doc.ID.Valid() {
			doc.ID = bson.NewObjectId()
		}
		docs = append(docs, doc)
	}
	session, err := m.TakeSession()
	if err != nil {
		return err
	}

	defer m.PutSession(session)
	return m.GetCollection(session).Insert(docs...)
}

func (m *defaultExportTaskModel) FindOne(id string) (*ExportTask, error) {
	if !bson.IsObjectIdHex(id) {
		return nil, ErrInvalidObjectId
	}

	session, err := m.TakeSession()
	if err != nil {
		return nil, err
	}

	defer m.PutSession(session)
	var data ExportTask
	key := prefixExportTaskCacheKey + id
	err = m.GetCollection(session).FindOneId(&data, key, bson.ObjectIdHex(id))
	switch err {
	case nil:
		return &data, nil
	case mongoc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultExportTaskModel) cacheKeyFromQuery(query interface{}) (string, error) {
	mj, err := bson.MarshalJSON(query)
	if err != nil {
		return "", err
	}

	b := strings.Builder{}
	_, err = b.WriteString(prefixExportTaskCacheKey)
	if err != nil {
		return "", err
	}

	_, err = b.Write(mj)
	if err != nil {
		return "", err
	}

	return b.String(), nil
}

func (m *defaultExportTaskModel) FindByAny(query interface{}) (*ExportTask, error) {
	session, err := m.TakeSession()
	if err != nil {
		return nil, err
	}
	defer m.PutSession(session)

	var data ExportTask

	key, err := m.cacheKeyFromQuery(query)
	if err != nil {
		return nil, err
	}
	err = m.GetCollection(session).FindOne(&data, key, query)

	switch err {
	case nil:
		return &data, nil
	case mongoc.ErrNotFound:
		return nil, err
	default:
		return nil, err
	}
}

func (m *defaultExportTaskModel) Update(data *ExportTask) error {
	session, err := m.TakeSession()
	if err != nil {
		return err
	}

	defer m.PutSession(session)
	key := prefixExportTaskCacheKey + data.ID.Hex()
	return m.GetCollection(session).UpdateId(data.ID, data, key)
}

func (m *defaultExportTaskModel) Delete(id string) error {
	session, err := m.TakeSession()
	if err != nil {
		return err
	}

	defer m.PutSession(session)
	key := prefixExportTaskCacheKey + id
	return m.GetCollection(session).RemoveId(bson.ObjectIdHex(id), key)
}

func (m *defaultExportTaskModel) List(query interface{}, opts ...mongoc.QueryOption) (res []ExportTask, err error) {
	session, err := m.TakeSession()
	if err != nil {
		return nil, err
	}
	defer m.PutSession(session)
	err = m.GetCollection(session).FindAllNoCache(&res, query, opts...)
	return
}

func (m *defaultExportTaskModel) Count(query interface{}) (int, error) {
	session, err := m.TakeSession()
	if err != nil {
		return 0, err
	}
	defer m.PutSession(session)

	count, err := m.GetCollection(session).Count(query)
	if err != nil {
		return 0, err
	}
	return count, nil
}

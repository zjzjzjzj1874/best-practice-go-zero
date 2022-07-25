package model

import (
	"github.com/globalsign/mgo/bson"
	cachec "github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/mongoc"
	"strings"
)

var prefixMongoTestCacheKey = "cache:MongoTest:"

type MongoTestModel interface {
	Insert(data *MongoTest) error
	BatchInsert(data []*MongoTest) error
	FindOne(id string) (*MongoTest, error)
	Update(data *MongoTest) error
	Delete(id string) error
	// query by any params
	FindByAny(query interface{}) (*MongoTest, error)
	List(query interface{}, opts ...mongoc.QueryOption) ([]MongoTest, error)
	Count(query interface{}) (int, error)
}

type defaultMongoTestModel struct {
	*mongoc.Model
}

func NewMongoTestModel(url, collection string, c cachec.CacheConf) MongoTestModel {
	return &defaultMongoTestModel{
		Model: mongoc.MustNewModel(url, collection, c),
	}
}

func (m *defaultMongoTestModel) Insert(data *MongoTest) error {
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

func (m *defaultMongoTestModel) BatchInsert(datas []*MongoTest) error {
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

func (m *defaultMongoTestModel) FindOne(id string) (*MongoTest, error) {
	if !bson.IsObjectIdHex(id) {
		return nil, ErrInvalidObjectId
	}

	session, err := m.TakeSession()
	if err != nil {
		return nil, err
	}

	defer m.PutSession(session)
	var data MongoTest
	key := prefixMongoTestCacheKey + id
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

func (m *defaultMongoTestModel) cacheKeyFromQuery(query interface{}) (string, error) {
	mj, err := bson.MarshalJSON(query)
	if err != nil {
		return "", err
	}

	b := strings.Builder{}
	_, err = b.WriteString(prefixMongoTestCacheKey)
	if err != nil {
		return "", err
	}

	_, err = b.Write(mj)
	if err != nil {
		return "", err
	}

	return b.String(), nil
}

func (m *defaultMongoTestModel) FindByAny(query interface{}) (*MongoTest, error) {
	session, err := m.TakeSession()
	if err != nil {
		return nil, err
	}
	defer m.PutSession(session)

	var data MongoTest

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

func (m *defaultMongoTestModel) Update(data *MongoTest) error {
	session, err := m.TakeSession()
	if err != nil {
		return err
	}

	defer m.PutSession(session)
	key := prefixMongoTestCacheKey + data.ID.Hex()
	return m.GetCollection(session).UpdateId(data.ID, data, key)
}

func (m *defaultMongoTestModel) Delete(id string) error {
	session, err := m.TakeSession()
	if err != nil {
		return err
	}

	defer m.PutSession(session)
	key := prefixMongoTestCacheKey + id
	return m.GetCollection(session).RemoveId(bson.ObjectIdHex(id), key)
}

func (m *defaultMongoTestModel) List(query interface{}, opts ...mongoc.QueryOption) (res []MongoTest, err error) {
	session, err := m.TakeSession()
	if err != nil {
		return nil, err
	}
	defer m.PutSession(session)
	err = m.GetCollection(session).FindAllNoCache(&res, query, opts...)
	return
}

func (m *defaultMongoTestModel) Count(query interface{}) (int, error) {
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

package model

import (
    "github.com/globalsign/mgo/bson"
     {{if .Cache}}cachec "github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/mongoc"{{else}}"github.com/zeromicro/go-zero/core/stores/mongo"{{end}}
	"strings"
)

// #{{if .Cache}}var prefix{{.Type}}CacheKey = "cache:{{.Type}}:"{{end}}
{{if .Cache}}var prefix{{.Type}}CacheKey = "cache:{{Service}}:{{.Type}}:"{{end}}

type {{.Type}}Model interface{
	Insert(data *{{.Type}}) error
	BatchInsert(data []*{{.Type}}) error
	FindOne(id string) (*{{.Type}}, error)
	Update(data *{{.Type}}) error
	Delete(id string) error
	// query by any params
	FindByAny( query interface{}) (*{{.Type}}, error)
	List(query interface{}, opts ...mongoc.QueryOption) ([]{{.Type}}, error)
	Count(query interface{}) (int, error)
}

type default{{.Type}}Model struct {
    {{if .Cache}}*mongoc.Model{{else}}*mongo.Model{{end}}
}

func New{{.Type}}Model(url, collection string{{if .Cache}}, c cachec.CacheConf{{end}}) {{.Type}}Model {
	return &default{{.Type}}Model{
		Model: {{if .Cache}}mongoc.MustNewModel(url, collection, c){{else}}mongo.MustNewModel(url, collection){{end}},
	}
}


func (m *default{{.Type}}Model) Insert(data *{{.Type}}) error {
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

func (m *default{{.Type}}Model) BatchInsert(datas []*{{.Type}}) error {
    var docs []interface{}

   for _, doc := range datas {
    if !doc.ID.Valid() {
           doc.ID = bson.NewObjectId()
       }
       docs=append(docs,doc)
   }
    session, err := m.TakeSession()
    if err != nil {
        return err
    }

    defer m.PutSession(session)
    return m.GetCollection(session).Insert(docs...)
}


func (m *default{{.Type}}Model) FindOne(id string) (*{{.Type}}, error) {
    if !bson.IsObjectIdHex(id) {
        return nil, ErrInvalidObjectId
    }

    session, err := m.TakeSession()
    if err != nil {
        return nil, err
    }

    defer m.PutSession(session)
    var data {{.Type}}
    {{if .Cache}}key := prefix{{.Type}}CacheKey + id
    err = m.GetCollection(session).FindOneId(&data, key, bson.ObjectIdHex(id))
	{{- else}}
	err = m.GetCollection(session).FindId(bson.ObjectIdHex(id)).One(&data)
	{{- end}}
    switch err {
    case nil:
        return &data,nil
    case {{if .Cache}}mongoc.ErrNotFound{{else}}mongo.ErrNotFound{{end}}:
        return nil,ErrNotFound
    default:
        return nil,err
    }
}

{{if .Cache}}
func (m *default{{.Type}}Model) cacheKeyFromQuery(query interface{}) (string, error) {
	mj, err := bson.MarshalJSON(query)
	if err != nil {
		return "", err
	}

	b := strings.Builder{}
	_, err = b.WriteString(prefix{{.Type}}CacheKey)
	if err != nil {
		return "", err
	}

	_, err = b.Write(mj)
	if err != nil {
		return "", err
	}

	return b.String(), nil
}
{{- end}}

func (m *default{{.Type}}Model) FindByAny(query interface{}) (*{{.Type}}, error) {
	session, err := m.TakeSession()
	if err != nil {
		return nil, err
	}
	defer m.PutSession(session)

	var data {{.Type}}
	{{if .Cache}}
	key, err := m.cacheKeyFromQuery(query)
	if err != nil {
		return nil, err
	}
	err = m.GetCollection(session).FindOne(&data, key, query)
	{{- else}}
	err = m.GetCollection(session).Find(query).One(&data)
	{{- end}}

	switch err {
	case nil:
		return &data, nil
	case {{if .Cache}}mongoc.ErrNotFound{{else}}mongo.ErrNotFound{{end}}:
		return nil, err
	default:
		return nil, err
	}
}

func (m *default{{.Type}}Model) Update(data *{{.Type}}) error {
    session, err := m.TakeSession()
    if err != nil {
        return err
    }

    defer m.PutSession(session)
	{{if .Cache}}key := prefix{{.Type}}CacheKey + data.ID.Hex()
    return m.GetCollection(session).UpdateId(data.ID, data, key)
	{{- else}}
	return m.GetCollection(session).UpdateId(data.ID, data)
	{{- end}}
}

func (m *default{{.Type}}Model) Delete(id string) error {
    session, err := m.TakeSession()
    if err != nil {
        return err
    }

    defer m.PutSession(session)
    {{if .Cache}}key := prefix{{.Type}}CacheKey + id
    return m.GetCollection(session).RemoveId(bson.ObjectIdHex(id), key)
	{{- else}}
	return m.GetCollection(session).RemoveId(bson.ObjectIdHex(id))
	{{- end}}
}

func (m *default{{.Type}}Model) List(query interface{}, opts ...mongoc.QueryOption) (res []{{.Type}}, err error) {
	session, err := m.TakeSession()
	if err != nil {
		return nil, err
	}
	defer m.PutSession(session)
	err = m.GetCollection(session).FindAllNoCache(&res, query, opts...)
	return
}

func (m *default{{.Type}}Model) Count(query interface{}) (int, error) {
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
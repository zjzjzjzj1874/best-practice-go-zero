package helper

import (
	"errors"
	"time"

	"github.com/globalsign/mgo/bson"
	"github.com/jinzhu/copier"
)

const DEFAULT_TIMELAYOUT = "2006-01-02 15:04:05"

// String2BsonObject string转成bson的objectID
func String2BsonObject() copier.TypeConverter {
	return copier.TypeConverter{
		SrcType: copier.String,
		DstType: bson.ObjectId(""),
		Fn: func(src interface{}) (interface{}, error) {
			s, ok := src.(string)
			if !ok {
				return src, nil
			}

			return bson.ObjectId(s), nil
		},
	}
}

// BsonObject2String 将bson的objectID转换成string
func BsonObject2String() copier.TypeConverter {
	return copier.TypeConverter{
		SrcType: bson.ObjectId(""),
		DstType: copier.String,
		Fn: func(src interface{}) (interface{}, error) {
			obj, ok := src.(bson.ObjectId)
			if !ok {
				return src, nil
			}
			return obj.Hex(), nil
		},
	}
}

// Time2String 将time转换成string
func Time2String() copier.TypeConverter {
	return copier.TypeConverter{
		SrcType: time.Time{},
		DstType: copier.String,
		Fn: func(src interface{}) (interface{}, error) {
			s, ok := src.(time.Time)
			if !ok {
				return nil, errors.New("src type not matching")
			}
			return s.Format(DEFAULT_TIMELAYOUT), nil
		},
	}
}

// String2Time 将string转换成time
func String2Time() copier.TypeConverter {
	return copier.TypeConverter{
		SrcType: copier.String,
		DstType: time.Time{},
		Fn: func(src interface{}) (interface{}, error) {
			s, ok := src.(string)
			if !ok {
				return nil, errors.New("src type not matching")
			}
			return time.Parse(DEFAULT_TIMELAYOUT, s)
		},
	}
}

// String2Time 将string转换成time
func String2LocalTime() copier.TypeConverter {
	return copier.TypeConverter{
		SrcType: copier.String,
		DstType: time.Time{},
		Fn: func(src interface{}) (interface{}, error) {
			s, ok := src.(string)
			if !ok {
				return nil, errors.New("src type not matching")
			}
			return time.ParseInLocation(DEFAULT_TIMELAYOUT, s, time.Local)
		},
	}
}

// OutOption 数据库字段类型转出成外部参数类型
func OutOption() copier.Option {
	return copier.Option{
		IgnoreEmpty: true,
		DeepCopy:    true,
		Converters: []copier.TypeConverter{
			Time2String(),
		},
	}
}

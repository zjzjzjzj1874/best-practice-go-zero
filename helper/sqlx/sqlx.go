package sqlx

import (
	"database/sql/driver"
	"fmt"

	jsoniter "github.com/json-iterator/go"
)

// StringSlice 数据库数组存储
type StringSlice []string

func (s *StringSlice) Scan(src interface{}) error {
	return JSONScan(src, s)
}

func (s StringSlice) Value() (driver.Value, error) {
	return JSONValue(s)
}

func JSONScan(dbValue interface{}, value interface{}) error {
	switch v := dbValue.(type) {
	case []byte:
		bytes := v
		if len(bytes) > 0 {
			return jsoniter.Unmarshal(bytes, value)
		}
		return nil
	case string:
		str := v
		if str == "" {
			return nil
		}
		return jsoniter.UnmarshalFromString(str, value)
	case nil:
		return nil
	default:
		return fmt.Errorf("cannot sql.Scan() from: %#v", value)
	}
}

func JSONValue(value interface{}) (driver.Value, error) {
	if zeroCheck, ok := value.(interface {
		IsZero() bool
	}); ok {
		if zeroCheck.IsZero() {
			return "", nil
		}
	}
	bytes, err := jsoniter.Marshal(value)
	if err != nil {
		return "", err
	}
	str := string(bytes)
	if str == "null" {
		return "", nil
	}
	return str, nil
}

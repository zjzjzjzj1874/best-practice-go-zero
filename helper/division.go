package helper

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"path"
	"sync"

	_ "github.com/mattn/go-sqlite3"
)

var ConnPool sync.Pool

// Divisions sqlite divisions table structure
type Divisions struct {
	ProvinceName        string  `json:"provinceName,omitempty"`
	ProvinceCode        int64   `json:"provinceCode,omitempty"`
	ProvinceCoordinateX float64 `json:"provinceCoordinateX,omitempty"`
	ProvinceCoordinateY float64 `json:"provinceCoordinateY,omitempty"`
	CityName            string  `json:"cityName,omitempty"`
	CityCode            int64   `json:"cityCode,omitempty"`
	CityCoordinateX     float64 `json:"cityCoordinateX,omitempty"`
	CityCoordinateY     float64 `json:"cityCoordinateY,omitempty"`
	AreaName            string  `json:"areaName,omitempty"`
	AreaCode            int64   `json:"areaCode,omitempty"`
	AreaCoordinateX     float64 `json:"areaCoordinateX,omitempty"`
	AreaCoordinateY     float64 `json:"areaCoordinateY,omitempty"`
}

func init() {
	dir, _ := os.Getwd()
	dsn := path.Join(dir, "/divisions.sqlite")

	ConnPool = sync.Pool{
		New: func() interface{} {
			db, err := sql.Open("sqlite3", dsn)
			if err != nil {
				logrus.Errorf("Open sqlite failure:[dsn:%s,err:%v]", dsn, err)
				return nil
			}
			ctx := context.TODO()
			conn, err := db.Conn(ctx)
			if err != nil {
				logrus.Errorf("Conn sqlite failure:[dsn:%s,err:%v]", dsn, err)
				return nil
			}
			return conn
		},
	}
}

// level - 0 省
// level - 1 市
// level - 2 区
func QueryNameByCode(code, level int32) string {
	conn, ok := ConnPool.Get().(*sql.Conn)
	if !ok {
		return ""
	}
	defer ConnPool.Put(conn)

	query := "SELECT %s FROM divisions WHERE %s=%d limit 1;"
	switch level {
	case 1:
		query = fmt.Sprintf(query, "provinceName", "provinceCode", code)
	case 2:
		query = fmt.Sprintf(query, "cityName", "cityCode", code)
	default:
		query = fmt.Sprintf(query, "areaName", "areaCode", code)
	}
	row := conn.QueryRowContext(context.TODO(), query)
	name := ""
	err := row.Scan(&name)
	if err != nil {
		logrus.Errorf("scan failure:[query:%s,err:%s]", query, err.Error())
	}

	return name
}

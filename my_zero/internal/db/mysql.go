// Package db 本包用来跑对应服务数据库操作相关的
package db

import (
	"fmt"
	"gorm.io/gorm"

	"github.com/zjzjzjzj1874/best-pracrice-go-zero/databases"
)

type Mysql struct {
	*gorm.DB
}

// NewMysqlClient 返回一个mysql的client
func NewMysqlClient(mysql *gorm.DB) *Mysql {
	return &Mysql{
		DB: mysql,
	}
}

// MigrateWithApi 使用API来跑MySQL Migrate
func (db *Mysql) MigrateWithApi() error {
	tables := []interface{}{
		databases.Mock{},
		databases.User{},
		databases.Test{},
	}
	for _, table := range tables {
		err := db.AutoMigrate(&table)
		if err != nil {
			fmt.Printf("autoMigrate failure:[err:%s]", err.Error())
			return err
		}
	}

	return nil
}

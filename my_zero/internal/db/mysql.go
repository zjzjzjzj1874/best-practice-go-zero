// Package db 本包用来跑对应服务数据库操作相关的
package db

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"

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
func (db *Mysql) MigrateWithApi(tableNames []string) error {
	if len(tableNames) == 0 {
		return db.migrateAll()
	}
	//
	//tableMap := map[string]interface{}{
	//	databases.Mock{}.TableName(): databases.Mock{},
	//	databases.User{}.TableName(): databases.User{},
	//	databases.Test{}.TableName(): databases.Test{},
	//}

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

// MigrateWithApi 使用API来跑MySQL Migrate
func (db *Mysql) migrateAll() error {
	tables := []interface{}{
		databases.Mock{},
		databases.User{},
		databases.Test{},
	}
	ctx, cancel := context.WithTimeout(context.TODO(), 20*time.Second)
	defer cancel()
	tx := db.WithContext(ctx)
	for _, table := range tables {
		err := tx.AutoMigrate(&table)
		if err != nil {
			fmt.Printf("autoMigrate failure:[err:%s]", err.Error())
			return err
		}
	}

	return nil
}
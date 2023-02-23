package mysql

import (
	"gorm.io/gorm"
)

// User 用户表 1.联合索引构建 => 不同字段上index名字一样,即可创建联合索引;2.唯一索引的两种方式(邮箱和姓名字段)
//go:generate goctl model mysql datasource "mysql://root:admin123@tcp(127.0.0.1:3306)/test" -table user -dir .
type User struct {
	gorm.Model
	Name     string `gorm:"uniqueIndex:idx_name;type:varchar(64);not null;default:'';comment:姓名"`
	Email    string `gorm:"index:idx_email,unique;type:varchar(255);not null;default:'';comment:邮箱"`
	NickName string `gorm:"index:idx_nick_name_age;type:varchar(64);not null;default:'';comment:昵称"`
	Age      int    `gorm:"index:idx_nick_name_age;type:int(11);not null;default:0;comment:年龄"`
	Phone    string `gorm:"index:idx_nick_name_age;type:varchar(18);not null;default:'';comment:电话号码"`
}

func (User) TableName() string {
	return "t_user"
}

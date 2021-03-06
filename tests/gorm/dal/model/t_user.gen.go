// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameTUser = "t_user"

// TUser mapped from table <t_user>
type TUser struct {
	ID        int64          `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	CreatedAt time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
	Name      string         `gorm:"column:name;not null" json:"name"`           // 姓名
	Email     string         `gorm:"column:email;not null" json:"email"`         // 邮箱
	NickName  string         `gorm:"column:nick_name;not null" json:"nick_name"` // 昵称
	Age       int32          `gorm:"column:age;not null" json:"age"`             // 年龄
	Phone     string         `gorm:"column:phone;not null" json:"phone"`         // 电话号码
}

// TableName TUser's table name
func (*TUser) TableName() string {
	return TableNameTUser
}

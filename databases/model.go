package databases

import "gorm.io/gorm"

type Mock struct {
	gorm.Model
	TestName string `gorm:"test_name" json:"test_name"`
}

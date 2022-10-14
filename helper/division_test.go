package helper

import (
	"fmt"
	"testing"
)

func TestSqlite(t *testing.T) {
	t.Run("#查询省市区", func(t *testing.T) {
		cus := QueryNameByCode(320000, 1)

		fmt.Println(len(cus))
	})
}

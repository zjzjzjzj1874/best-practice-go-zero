package main

import (
	"fmt"

	"github.com/blastrain/vitess-sqlparser/sqlparser"
)

func main() {
	stmt, err := sqlparser.Parse("select * from user_items where user_id=1 group by user_id order by created_at limit 3 offset 10")
	if err != nil {
		panic(err)
	}

	fmt.Printf("stmt = %+v\n", stmt)
}

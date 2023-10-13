package main

import (
	"fmt"
)

func main() {
	client := newClient()
	resp, code, err := client.Json("LangDetect", nil, `{"TextList":["你好世界"]}`)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%d %s\n", code, string(resp))
}

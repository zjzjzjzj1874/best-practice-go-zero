package batch

import "fmt"

func yourLogic(tags []Meta) {
	if len(tags) == 0 {
		fmt.Println("Abort:本次没有任务数据处理,直接返回")
		return
	}

	fmt.Println("Next: 在这里做你的逻辑处理 => tags:", tags)
}

type Meta struct {
	Type int `json:"type,omitempty" in:"body"` // 类型
}

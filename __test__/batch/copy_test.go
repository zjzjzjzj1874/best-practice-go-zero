package batch

import (
	"fmt"
	"testing"
)

func TestTextTempReview_Test(t *testing.T) {
	t.Run("copy", func(t *testing.T) {
		src := []int{1, 2, 3, 4, 5, 6}
		dst := make([]int, len(src))

		copy(dst, src)

		fmt.Println(dst)

		src = append(src, 123)
		src[0] = 100

		fmt.Println(dst)
	})

	// 这个测试用例的作用:当您有很多数据(100w)时,一次性不能发送100w条,
	// 但是又不想一条一条处理,这个时候就可以批量处理,每次10000条,分100次发送,
	// 就需要用到下面分组的逻辑     =>
	// TODO 升级的可以先分好组,用协程去处理,waitGroup和另一个组,只有当都成功时才返回成功的,巴拉巴拉.
	t.Run("split", func(t *testing.T) {
		src := make([]Meta, 0, 5)
		texts := []int{
			1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
		}
		for _, text := range texts {
			src = append(src, Meta{
				Type: text,
			})
			if len(src) >= 5 {
				dst := make([]Meta, 5)
				copy(dst, src)
				src = make([]Meta, 0, 5)

				// do your logic
				yourLogic(dst)
			}
		}
		if len(src) != 0 {
			yourLogic(src) // 收尾工作
		}

	})
}

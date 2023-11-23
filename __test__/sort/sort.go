package main

import (
	"fmt"
	"golang.org/x/text/collate"
	"golang.org/x/text/language"
	"sort"
)

func main() {
	names := []string{"张三", "Amy", "王五", "李四", "爸爸", "John", "妈妈"} // 中英文混合排序	输出 [Amy 爸爸 John 李四 妈妈 王五 张三]

	// 创建排序器，这里使用的是简体中文的排序规则
	cl := collate.New(language.SimplifiedChinese, collate.Loose)
	sort.SliceStable(names, func(i, j int) bool {
		// 使用 collator 的 Compare 方法来比较字符串
		return cl.CompareString(names[i], names[j]) < 0
	})

	fmt.Println(names)
}

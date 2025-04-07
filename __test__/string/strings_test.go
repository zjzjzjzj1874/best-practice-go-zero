package string

import (
	"strings"
	"testing"
)

// 普通拼接测试：少量固定次数拼接
func BenchmarkConcatSimplePlus(b *testing.B) {
	str1 := "Hello"
	str2 := ", "
	str3 := "World!"
	for i := 0; i < b.N; i++ {
		_ = str1 + str2 + str3 // 编译器会优化为单次内存分配
	}
}

func BenchmarkConcatSimpleBuilder(b *testing.B) {
	str1 := "Hello"
	str2 := ", "
	str3 := "World!"
	for i := 0; i < b.N; i++ {
		var builder strings.Builder
		builder.WriteString(str1)
		builder.WriteString(str2)
		builder.WriteString(str3)
		_ = builder.String()
	}
}

// 循环拼接测试：大量动态拼接
func BenchmarkConcatLoopPlus(b *testing.B) {
	data := strings.Repeat("a", 100) // 单次拼接的字符串长度
	for i := 0; i < b.N; i++ {
		s := ""
		for j := 0; j < 1000; j++ { // 循环 1000 次
			s += data
		}
		_ = s
	}
}

func BenchmarkConcatLoopBuilder(b *testing.B) {
	data := strings.Repeat("a", 100) // 单次拼接的字符串长度
	for i := 0; i < b.N; i++ {
		var builder strings.Builder
		builder.Grow(100 * 1000)    // 预分配内存（已知总长度）
		for j := 0; j < 1000; j++ { // 循环 1000 次
			builder.WriteString(data)
		}
		_ = builder.String()
	}
}

// 测试结果
//goos: darwin
//goarch: arm64
//pkg: ~/best-pracrice-go-zero/__test__/string
//cpu: Apple M3 Pro
//BenchmarkConcatSimplePlus
//BenchmarkConcatSimplePlus-11       	85823354	        13.26 ns/op
//BenchmarkConcatSimpleBuilder
//BenchmarkConcatSimpleBuilder-11    	41003386	        28.48 ns/op
//BenchmarkConcatLoopPlus
//BenchmarkConcatLoopPlus-11         	     388	   3198098 ns/op
//BenchmarkConcatLoopBuilder
//BenchmarkConcatLoopBuilder-11      	  152229	      7788 ns/op
//PASS

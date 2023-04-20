package main

import (
	"testing"
)

func TestMarshalWithJson(t *testing.T) {
	MarshalWithJson()
}

func TestMarshalWithJsonEncodeAndBuf(t *testing.T) {
	MarshalWithJsonEncodeAndBuf()
}

func BenchmarkMarshalWithJson(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MarshalWithJson()
	}
}

func BenchmarkMarshalWithJsonEncodeAndBuf(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MarshalWithJsonEncodeAndBuf()
	}
}

// 基于内存的性能测试:go test -bench='Marshal' . -benchmem
// 普通性能测试:go test -bench='Marshal' .

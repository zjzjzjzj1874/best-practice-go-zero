package main

import (
	"testing"
)

//func TestMarshalWithJson(t *testing.T) {
//	MarshalWithJson()
//}
//
//func TestMarshalWithJsonEncodeAndBuf(t *testing.T) {
//	MarshalWithJsonEncodeAndBuf()
//}

func BenchmarkMarshalWithJson(b *testing.B) {
	//var memStats runtime.MemStats
	//runtime.ReadMemStats(&memStats)
	for i := 0; i < b.N; i++ {
		MarshalWithJson()
	}
	//runtime.ReadMemStats(&memStats)
	//b.Logf("Alloc: %d bytes, TotalAlloc: %d bytes, Sys: %d bytes, NumGC: %d",
	//	memStats.Alloc, memStats.TotalAlloc, memStats.Sys, memStats.NumGC)
}

func BenchmarkMarshalWithJsonEncodeAndBuf(b *testing.B) {
	//var memStats runtime.MemStats
	//runtime.ReadMemStats(&memStats)
	for i := 0; i < b.N; i++ {
		MarshalWithJsonEncodeAndBuf()
	}
	//runtime.ReadMemStats(&memStats)
	//b.Logf("Alloc: %d bytes, TotalAlloc: %d bytes, Sys: %d bytes, NumGC: %d",
	//	memStats.Alloc, memStats.TotalAlloc, memStats.Sys, memStats.NumGC)
}

// 基于内存的性能测试:go test -bench='Marshal' . -benchmem
// 普通性能测试:go test -bench='Marshal' .

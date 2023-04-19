package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"sync"
)

// Note:优化,之前Person使用json.marshal方法,如果json.marshal内存占用过高,可以使用以下优化:
// 1.json.NewEncoder代替json.marshal;
// 2.使用json.Encoder的底层缓冲区,减少内存分配和垃圾回收开销

type Person struct {
	Name string
	Age  int
}
type JsonPool struct {
	buf    *bytes.Buffer
	encode *json.Encoder
}

func (p *JsonPool) Reset() {
	p.buf.Reset()
}

var (
	encodePool = sync.Pool{New: func() any {
		var buf bytes.Buffer
		encode := json.NewEncoder(&buf)
		return &JsonPool{
			buf:    &buf,
			encode: encode,
		}
	}}

	person = Person{Name: "Alice", Age: 30}
)

func main() {
	jp, ok := encodePool.Get().(*JsonPool)
	if ok {
		jp.Reset()
	}
	defer encodePool.Put(jp)

	_ = jp.encode.Encode(&person)
	fmt.Println(jp.buf.String()) // string也是强转,有内存开销
	fmt.Println(jp.buf.Bytes())
}

func MarshalWithJson() {
	_, _ = json.Marshal(person)
}

func MarshalWithJsonEncodeAndBuf() {
	jp, ok := encodePool.Get().(*JsonPool)
	if ok {
		jp.Reset()
	}
	defer encodePool.Put(jp)

	_ = jp.encode.Encode(&person)
}

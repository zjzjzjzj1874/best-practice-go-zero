package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"net"
	"os"
	"testing"
	"time"
)

func Test_validateEmail(t *testing.T) {
	type args struct {
		email string
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "email", args: args{email: "user@gmail.com"}},
		{name: "email", args: args{email: "user@example.com"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			validateEmail(tt.args.email)
		})
	}
}
func Test_Other(t *testing.T) {
	t.Run("NSLOOKUP", func(t *testing.T) {
		ns, err := net.LookupNS("baidu.com")
		if err != nil {
			println("err:", err)
			return
		}

		for idx, n := range ns {
			println(idx, ":", n.Host)
		}
	})
}

func Test_NetPipe(t *testing.T) {
	t.Run("net.Pipe", func(t *testing.T) {
		reader, writer := net.Pipe()

		// 写入数据
		go func() {
			defer writer.Close()

			file, err := os.Open("/Users/zjzjzjzj1874/personal/best-practice-go-zero/go.csv")
			if err != nil {
				fmt.Println("err:", err)
				return
			}
			defer file.Close()

			r := csv.NewReader(file)
			//r.FieldsPerRecord
			for {
				// 逐行读取，
				// 1.因为如果数据量很大的情况下，全部读取占用的内存很大
				// 2.可以按照每行的值来读取，如果不满足第一行的格式，可以选择跳过，也可以选择继续加入
				record, err := r.Read()
				if err == io.EOF {
					break
				}
				if err != nil {
					fmt.Println("err:", err)
					fmt.Println("record:", record)
				}

				_, err = writer.Write([]byte(fmt.Sprintf("%v", record)))
				if err != nil {
					fmt.Println("write msg failure:", err)
					return
				}
				//fmt.Println(record)
			}

			//i := 1
			//for {
			//	msg := []byte(fmt.Sprintf("%d:Hello World!", i))
			//	_, err := writer.Write(msg)
			//	if err != nil {
			//		fmt.Println("write msg failure:", err)
			//		return
			//	}
			//
			//	time.Sleep(time.Second)
			//	i++
			//}
		}()

		// 读取数据
		go func() {
			defer reader.Close()
			for {
				buffer := make([]byte, 1024)
				msg := ""
				for {
					n, err := reader.Read(buffer)
					if err != nil {
						if err == io.EOF {
							msg += string(buffer[:n])
							break
						}
						fmt.Println("read msg failure:", err)
						return
					}
					if n < 1024 {
						msg += string(buffer[:n])
						break
					}
				}

				fmt.Println("msg === ", msg)
				time.Sleep(time.Second)
			}
		}()

		// 阻塞
		select {}
	})
}

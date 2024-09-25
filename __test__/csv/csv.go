package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

// csv:CSV（Comma-Separated Values，逗号分隔值，也不仅限于只能用逗号分隔）是一种用于存储表格数据的简单文本格式。
// 它常用于数据交换，将数据组织成行和列的形式，类似于电子表格或数据库表中的数据。
func main() {
	file, err := os.Create("go.csv")
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	w := csv.NewWriter(file)

	records := [][]string{
		{"first_name", "last_name", "username"},
		{"Rob", "Pike", "rob"},
		{"Ken", "Thompson", "ken"},
		{"Robert", "Griesemer", "gri"},
	}

	// 写文件
	for _, record := range records {
		err := w.Write(record)
		if err != nil {
			fmt.Println("err:", err)
		}
	}
	w.Flush()
	// 上面可以使用一行来解决
	//w.WriteAll(records)

	if err := w.Error(); err != nil {
		fmt.Println("err:", err)
	}

	_ = file.Close()
}

func readCsv(filename string) {
	file, err := os.Open(filename)
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

		fmt.Println(record)
	}

	return
}

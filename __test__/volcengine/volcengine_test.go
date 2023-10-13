package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"testing"

	"github.com/xuri/excelize/v2"
)

func TestVolc3(t *testing.T) {
	client := newClient()
	t.Run("#2CH", func(t *testing.T) {
		text := "인스타그램: https://instagram"
		req := transReq{
			TargetLanguage: "zh",
			TextList:       []string{text},
		}
		body, err := json.Marshal(req)
		if err != nil {
			log.Fatal(err)
		}
		resp, _, err := client.Json("TranslateText", nil, string(body))
		if err != nil {
			log.Fatal(err)
		}

		r := new(transResp)
		err = json.Unmarshal(resp, &r)
		if err != nil {
			log.Fatal(err)
		}

		if len(r.TranslationList) > 0 {
			log.Printf("Result:%+v", r.TranslationList)
			return
		}
		log.Printf("【volc 翻译】%#v", r.ResponseMetaData.Error)
	})
}

func TestVolc4(t *testing.T) {
	// 翻译香港
	t.Run("#test_xc_hunhe_hk_out_out.xlsx", func(t *testing.T) {
		// 打开Excel文件
		excel, err := excelize.OpenFile("./test_xc_hunhe_hk_out_out.xlsx")
		if err != nil {
			fmt.Println(err)
			return
		}
		// 获取Sheet1上所有单元格 => 下面-1是减去表头
		rows, err := excel.GetRows("Sheet1")

		// 创建一个新的Excel文件
		newFile := excelize.NewFile()
		// 创建新的工作表
		_ = newFile.NewSheet("Sheet1")
		// 循环读取每个工作表
		for idx, row := range rows {
			if idx == 0 { // 跳过表头
				_ = newFile.SetSheetRow("Sheet1", "A1", &row) // 设置Excel表header
				continue
			}

			if len(row) < 5 { // 格式不对
				continue
			}

			resultTitle := ""
			resultContent := ""
			// 翻译
			resultTitle = transIntoZH(row[4])
			if row[4] == row[5] {
				resultContent = resultTitle
			} else {
				resultContent = transIntoZH(row[5])
			}

			log.Printf("Idx:%+d,resultTitle:%+v,resultContent:%+v", idx, resultTitle, resultContent)
			row = append(row, []string{"", "", "", ""}...)
			row[14] = resultTitle
			row[15] = resultContent
			_ = newFile.SetSheetRow("Sheet1", fmt.Sprintf("A%d", idx+1), &row) // 设置Excel表header
		}

		// 保存新文件
		if err := newFile.SaveAs("test_xc_hunhe_hk_out_out_out.xlsx"); err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("Excel文件处理完成并保存为 test_xc_hunhe_hk_out_out_out.xlsx")
	})

	t.Run("#test_xc_hunhe_ht_out.xlsx", func(t *testing.T) {
		// 打开Excel文件
		excel, err := excelize.OpenFile("./test_xc_hunhe_ht_out.xlsx")
		if err != nil {
			fmt.Println(err)
			return
		}
		// 获取Sheet1上所有单元格 => 下面-1是减去表头
		rows, err := excel.GetRows("Sheet1")

		// 创建一个新的Excel文件
		newFile := excelize.NewFile()
		// 创建新的工作表
		_ = newFile.NewSheet("Sheet1")
		// 循环读取每个工作表
		for idx, row := range rows {
			if idx == 0 { // 跳过表头
				_ = newFile.SetSheetRow("Sheet1", "A1", &row) // 设置Excel表header
				continue
			}

			if len(row) < 5 { // 格式不对
				continue
			}
			resultContent := ""
			resultTitle := transIntoZH(row[4])
			if row[4] == row[5] {
				resultContent = resultTitle
			} else {
				resultContent = transIntoZH(row[5])
			}
			log.Printf("Idx:%+d,resultTitle:%+v,resultContent:%+v", idx, resultTitle, resultContent)
			row = append(row, []string{"", "", "", ""}...)
			row[14] = resultTitle
			row[15] = resultContent
			_ = newFile.SetSheetRow("Sheet1", fmt.Sprintf("A%d", idx+1), &row) // 设置Excel表header
		}

		// 保存新文件
		if err := newFile.SaveAs("test_xc_hunhe_ht_out_out.xlsx"); err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("Excel文件处理完成并保存为 test_xc_hunhe_ht_out_out.xlsx")
	})

	t.Run("#test_xc_hunhe_out.xlsx", func(t *testing.T) {
		// 打开Excel文件
		excel, err := excelize.OpenFile("./test_xc_hunhe_out.xlsx")
		if err != nil {
			fmt.Println(err)
			return
		}
		// 获取Sheet1上所有单元格 => 下面-1是减去表头
		rows, err := excel.GetRows("Sheet1")

		// 创建一个新的Excel文件
		newFile := excelize.NewFile()
		// 创建新的工作表
		_ = newFile.NewSheet("Sheet1")
		// 循环读取每个工作表
		for idx, row := range rows {
			if idx == 0 { // 跳过表头
				_ = newFile.SetSheetRow("Sheet1", "A1", &row) // 设置Excel表header
				continue
			}

			if len(row) < 5 { // 格式不对
				continue
			}
			resultContent := ""
			resultTitle := transIntoZH(row[4])
			if row[4] == row[5] {
				resultContent = resultTitle
			} else {
				resultContent = transIntoZH(row[5])
			}
			log.Printf("Idx:%+d,ID:%+v", idx, row[1])
			row = append(row, []string{"", "", "", ""}...)
			row[14] = resultTitle
			row[15] = resultContent
			_ = newFile.SetSheetRow("Sheet1", fmt.Sprintf("A%d", idx+1), &row) // 设置Excel表header
		}

		// 保存新文件
		if err := newFile.SaveAs("test_xc_hunhe_out_out.xlsx"); err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("Excel文件处理完成并保存为 test_xc_hunhe_out_out.xlsx")
	})
}

// 修复QPS超限问题
func TestQPSRecover(t *testing.T) {
	t.Run("#test_xc_hunhe_ht_out_out.xlsx", func(t *testing.T) {
		// 打开Excel文件
		excel, err := excelize.OpenFile("./test_xc_hunhe_ht_out_out.xlsx")
		if err != nil {
			fmt.Println(err)
			return
		}
		// 获取Sheet1上所有单元格 => 下面-1是减去表头
		rows, err := excel.GetRows("Sheet1")

		// 创建一个新的Excel文件
		newFile := excelize.NewFile()
		// 创建新的工作表
		_ = newFile.NewSheet("Sheet1")
		// 循环读取每个工作表
		for idx, row := range rows {
			if idx == 0 { // 跳过表头
				_ = newFile.SetSheetRow("Sheet1", "A1", &row) // 设置Excel表header
				continue
			}

			if len(row) < 16 { // 格式不对
				fmt.Println("格式不对")
				continue
			}

			if strings.Contains(row[14], "qps=") {
				row[14] = transIntoZH(row[4])
				log.Printf("Idx:%+d,ID:%+v", idx, row[1])
			}
			if strings.Contains(row[15], "qps=") {
				row[15] = transIntoZH(row[5])
				log.Printf("Idx:%+d,ID:%+v", idx, row[1])
			}
			_ = newFile.SetSheetRow("Sheet1", fmt.Sprintf("A%d", idx+1), &row) // 设置Excel表header
		}

		// 保存新文件
		if err := newFile.SaveAs("test_xc_hunhe_ht_out_out_out.xlsx"); err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("Excel文件处理完成并保存为 test_xc_hunhe_ht_out_out_out.xlsx")
	})
}

var client = newClient()

func transIntoZH(text string) string {
	req := transReq{
		TargetLanguage: "zh",
		TextList:       []string{text},
	}
	body, err := json.Marshal(req)
	if err != nil {
		//log.Fatal(err)
		return ""
	}
	resp, _, err := client.Json("TranslateText", nil, string(body))
	if err != nil {
		//log.Fatal(err)
		return ""
	}

	r := new(transResp)
	err = json.Unmarshal(resp, &r)
	if err != nil {
		return ""
	}

	if len(r.TranslationList) > 0 {
		return r.TranslationList[0].Translation
	}

	return r.ResponseMetaData.Error.Message
	//log.Printf("【volc 翻译】%#v", r.ResponseMetaData.Error)
}

package exporter

import (
	"fmt"
	"io/ioutil"
	"os"
	"runtime"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/xuri/excelize/v2"

	"github.com/zjzjzjzj1874/best-pracrice-go-zero/model/mongo"
	"github.com/zjzjzjzj1874/best-pracrice-go-zero/task/internal/svc"
)

func TestExporterCode_ExcelAssembly(t *testing.T) {
	t.Run("TestName", func(t *testing.T) {
		var (
			lists = []mongo.ExportTask{{
				ID:           "1",
				ExportType:   1,
				ExportState:  1,
				RequestArgs:  `{"name":"hello"}`,
				TargetEmails: []string{"test@email.com"},
			}}
			ctx *svc.ServiceContext
		)
		for _, list := range lists {
			logrus.Infof("%+v", list)
			fact := newExportFact(ctx, &list) // 根据不同报表类型创建不同类型

			fact.LoadStatsRecords() // 获取统计数据
			if err := fact.Err(); err != nil {
				continue
			}
			fact.ExcelAssembly() // 组装Excel报表
			if err := fact.Err(); err != nil {
				continue
			}
			fact.UploadObs() // 上传到obs
			if err := fact.Err(); err != nil {
				continue
			}

			fact.PostEmail() // 发送邮件(发邮件错误可以忽略,)
			fact.AfterPost() // 更新数据库
		}
	})

	t.Run("#GenExcel", func(t *testing.T) {
		var (
			idx       = 4
			sheetName = "统计"
			header    = []string{"日期", "调用类型", "条数"}
		)
		f := excelize.NewFile()
		defer func() {
			_ = f.Close()
			go runtime.GC()
		}()
		f.NewSheet(sheetName)
		fillSty, _ := f.NewStyle(&excelize.Style{
			Fill: excelize.Fill{
				// gradient： 渐变色    pattern   填充图案
				// Pattern: 1,                   // 填充样式  当类型是 pattern 0-18 填充图案  1 实体填充
				// Color:   []string{"#FF0000"}, // 当Type = pattern 时，只有一个
				Type:  "gradient",
				Color: []string{"#DCDCDC", "#DCDCDC"},
				// 类型是 gradient 使用 0-5 横向(每种颜色横向分布) 纵向 对角向上 对角向下 有外向内 由内向外
				Shading: 1,
			}})
		headStyle, err := f.NewStyle(&excelize.Style{
			Fill: excelize.Fill{
				// gradient： 渐变色    pattern   填充图案
				// Pattern: 1,                   // 填充样式  当类型是 pattern 0-18 填充图案  1 实体填充
				// Color:   []string{"#FF0000"}, // 当Type = pattern 时，只有一个
				Type:  "gradient",
				Color: []string{"#DCDCDC", "#DCDCDC"},
				// 类型是 gradient 使用 0-5 横向(每种颜色横向分布) 纵向 对角向上 对角向下 有外向内 由内向外
				Shading: 1,
			}, Alignment: &excelize.Alignment{
				Horizontal: "center",
				Vertical:   "center",
			},
		})
		// 定义行样式（通过JSON格式指定）
		//headerSty, err := f.NewStyle(`{
		//	"alignment":{
		//		"vertical":"center",
		//		"horizontal":"center"
		//	}
		//}`)
		if err != nil {
			return
		}
		_ = f.MergeCell(sheetName, "A1", "C1")                   // 合并单元格
		_ = f.SetCellValue(sheetName, "A1", "客户名称(张三)-内容审核用量统计") // 设置Excel表header
		_ = f.SetCellStyle(sheetName, "A1", "C1", headStyle)     // 设置Excel表header
		//_ = f.SetSheetRow(sheetName, "A1", &[]string{"客户名称-内容审核用量统计"})  // 设置Excel表header
		_ = f.MergeCell(sheetName, "A2", "C2")
		_ = f.SetCellValue(sheetName, "B2", "统计日期:2022.11.23")
		_ = f.MergeCell(sheetName, "A3", "C3")
		_ = f.SetCellValue(sheetName, "C3", "调用记录数:6991条,计费量:100223次")
		_ = f.SetSheetRow(sheetName, fmt.Sprintf("A%d", idx), &header)                           // 设置Excel表header
		_ = f.SetCellStyle(sheetName, fmt.Sprintf("A%d", idx), fmt.Sprintf("D%d", idx), fillSty) // 设置Excel表header

		for i := 0; i < 4; i++ {
			idx++
			err := f.SetSheetRow(sheetName, fmt.Sprintf("A%d", idx), &[]string{"2023-01-02", "图片", "11"})
			if err != nil {
				logrus.Errorf("[报表导出]SetSheetRow failrue:[err:%s]", err.Error())
				continue
			}
		}

		f.SetActiveSheet(1)
		f.DeleteSheet("Sheet1") // 默认的sheet表删除不要
		buf, err := f.WriteToBuffer()
		if err != nil {
			logrus.Errorf("[报表导出]write buf failure:[err:%s]", err.Error())
			return
		}

		_ = ioutil.WriteFile("2023.xlsx", buf.Bytes(), os.ModePerm)

		return
	})
}

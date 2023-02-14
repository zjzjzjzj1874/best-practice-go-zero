package exporter

import (
	"fmt"
	"runtime"

	"github.com/sirupsen/logrus"
	"github.com/xuri/excelize/v2"

	"github.com/zjzjzjzj1874/best-pracrice-go-zero/model/mongo"
	"github.com/zjzjzjzj1874/best-pracrice-go-zero/task/internal/svc"
)

type ExporterReport struct {
	*Exporter
}

func NewExporterReport(ctx *svc.ServiceContext, task *mongo.ExportTask) *ExporterReport {
	return &ExporterReport{NewExport(ctx, task)}
}

// LoadStatsRecords 组装Excel报表
func (e *ExporterReport) LoadStatsRecords() {
	// TODO find your stat data
}

// ExcelAssembly 组装Excel报表
func (e *ExporterReport) ExcelAssembly() {
	// 按条件组装
	for idx := range *(e.records) {
		logrus.Info(idx)
		// TODO add your logic to deal with data
	}

	data := e.readRecords()
	if e.err = e.genExcel(data); e.err != nil {
		logrus.Errorf("[ExporterReport]Failure to genExcel,[err:%v]", e.err.Error())
	}
}

// 读取数据库记录,以通道形式返回给下游消费
func (e *ExporterReport) readRecords() <-chan interface{} {
	var (
		offset int64 = 0
		size   int64 = 50
	)
	chanData := make(chan interface{}, 100) // 带缓存的通道

	go func() {
		defer close(chanData)
		for {
			if offset > 1<<10 {
				break
			}

			chanData <- "hello"
			offset += size
		}
	}()

	return chanData
}

// 生成Excel表格
func (e *ExporterReport) genExcel(chanData <-chan interface{}) error {
	var (
		index  = 1
		err    error
		isErr  bool
		header = []string{"测试", "数量"}
	)
	f := excelize.NewFile()
	_ = f.SetSheetRow("sheetName", fmt.Sprintf("A%d", index), &header)
	defer func() {
		_ = f.Close()
		go runtime.GC()
	}()
	// 全局style
	styleID, _ := f.NewStyle(&excelize.Style{Font: &excelize.Font{Color: "#000000"}, Alignment: &excelize.Alignment{Vertical: "center", Horizontal: "center"}})

	for {
		obj, ok := <-chanData
		if !ok {
			break
		}
		if err, isErr = obj.(error); isErr {
			return err
		}
		index++
		_ = f.SetSheetRow("sheetName", fmt.Sprintf("A%d", index), &[]string{"world", fmt.Sprintf("%v", obj)})
	}

	_ = f.SetRowStyle("sheetName", 1, 999, styleID)
	f.SetActiveSheet(1)     // 默认第一个sheet
	f.DeleteSheet("Sheet1") // 默认的sheet表删除不要
	e.buf, e.err = f.WriteToBuffer()
	if e.err != nil {
		logrus.Errorf("write buf failure:[err:%s]", e.err.Error())
	}
	return err
}

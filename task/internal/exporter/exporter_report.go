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

// ExcelAssembly 组装Excel报表
func (e *ExporterReport) ExcelAssembly() {
	// 按条件组装
	for idx := range *(e.records) {
		logrus.Info(idx)
		// TODO add your logic to deal with data
	}

	e.genExcel()
}

// 生成Excel表格
func (e *ExporterReport) genExcel( /*业务数据*/ ) {
	logrus.Info("===============开始生成xlsx===============")
	var (
		index  = 1
		header = []string{"Day", "Code", "数量"}
	)
	f := excelize.NewFile()
	defer func() {
		_ = f.Close()
		logrus.Infof("=============== 完成Excel生成,err:%s===============", e.err)
		go runtime.GC()
	}()
	// 全局style
	styleID, _ := f.NewStyle(&excelize.Style{Font: &excelize.Font{Color: "#000000"}, Alignment: &excelize.Alignment{Vertical: "center", Horizontal: "center"}})

	f.SetSheetRow("sheetName", fmt.Sprintf("A%d", index), &header)
	f.SetRowStyle("sheetName", 1, 999, styleID)
	f.SetActiveSheet(1)     // 默认第一个sheet
	f.DeleteSheet("Sheet1") // 默认的sheet表删除不要
	e.buf, e.err = f.WriteToBuffer()
	if e.err != nil {
		logrus.Errorf("write buf failure:[err:%s]", e.err.Error())
	}
}

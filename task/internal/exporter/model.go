package exporter

const (
	yyyyMMDD                        = "2006-01-02"          // 日期格式
	YYYYMMDD                        = "2006/01/02"          // 日期格式
	timeFormat                      = "2006-01-02 15:04:05" // 日期格式
	REPORT_FORMAT_CSV  ReportFormat = "csv"
	REPORT_FORMAT_XLSX ReportFormat = "xlsx"
)

// ReportFormat 导出参数设置
type ReportFormat string

// ContentTypes 导出不同格式对应的content-type
var ContentTypes = map[ReportFormat]string{
	REPORT_FORMAT_CSV:  "text/csv;charset=utf-8",
	REPORT_FORMAT_XLSX: "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet;charset=utf-8",
}

type ExportReq struct {
	Id        string `json:"id" description:"id" validate:"min=1"`
	StartDate string `json:"start_date" description:"开始日期，格式'2006-01-02'" validate:"required"`
	EndDate   string `json:"end_date" description:"结束日期，格式'2006-01-02'" validate:"required"`
}

// Meta 单元数据
type Meta struct {
	Date   string `json:"date" description:"日期"`
	Name   string `json:"name" description:"客户名称"`
	Finish int    `json:"finish" description:"审核完成总量"`
}

// UnitData 单元数据
type UnitData struct {
	Date string `json:"date" description:"日期"`
	Meta []Meta
}

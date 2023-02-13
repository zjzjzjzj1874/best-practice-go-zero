package types

//go:generate tools gen enum ExportType
// ExportType 导出类型
type ExportType int8

const (
	EXPORT_TYPE_UNKNOWN ExportType = iota
	EXPORT_TYPE__CODE              // 代码导出
	EXPORT_TYPE__REPORT            // 报表导出
)

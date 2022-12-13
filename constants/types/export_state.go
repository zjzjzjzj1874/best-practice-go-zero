package types

//go:generate tools gen enum ExportState
// ExportState 内容类型
type ExportState int8

const (
	EXPORT_STATE_UNKNOWN  ExportState = iota
	EXPORT_STATE__TODO                // 待导出
	EXPORT_STATE__ING                 // 导出中
	EXPORT_STATE__FAILURE             // 导出失败
	EXPORT_STATE__SUCCESS             // 导出成功
)

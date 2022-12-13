package types

import (
	bytes "bytes"
	database_sql_driver "database/sql/driver"
	errors "errors"

	github_com_go_courier_enumeration "github.com/go-courier/enumeration"
)

var InvalidExportState = errors.New("invalid ExportState type")

func ParseExportStateFromLabelString(s string) (ExportState, error) {
	switch s {
	case "":
		return EXPORT_STATE_UNKNOWN, nil
	case "待导出":
		return EXPORT_STATE__TODO, nil
	case "导出中":
		return EXPORT_STATE__ING, nil
	case "导出失败":
		return EXPORT_STATE__FAILURE, nil
	case "导出成功":
		return EXPORT_STATE__SUCCESS, nil
	}
	return EXPORT_STATE_UNKNOWN, InvalidExportState
}

func (v ExportState) String() string {
	switch v {
	case EXPORT_STATE_UNKNOWN:
		return ""
	case EXPORT_STATE__TODO:
		return "TODO"
	case EXPORT_STATE__ING:
		return "ING"
	case EXPORT_STATE__FAILURE:
		return "FAILURE"
	case EXPORT_STATE__SUCCESS:
		return "SUCCESS"
	}
	return "UNKNOWN"
}

func ParseExportStateFromString(s string) (ExportState, error) {
	switch s {
	case "":
		return EXPORT_STATE_UNKNOWN, nil
	case "TODO":
		return EXPORT_STATE__TODO, nil
	case "ING":
		return EXPORT_STATE__ING, nil
	case "FAILURE":
		return EXPORT_STATE__FAILURE, nil
	case "SUCCESS":
		return EXPORT_STATE__SUCCESS, nil
	}
	return EXPORT_STATE_UNKNOWN, InvalidExportState
}

func (v ExportState) Label() string {
	switch v {
	case EXPORT_STATE_UNKNOWN:
		return ""
	case EXPORT_STATE__TODO:
		return "待导出"
	case EXPORT_STATE__ING:
		return "导出中"
	case EXPORT_STATE__FAILURE:
		return "导出失败"
	case EXPORT_STATE__SUCCESS:
		return "导出成功"
	}
	return "UNKNOWN"
}

func (v ExportState) Int() int {
	return int(v)
}

func (ExportState) TypeName() string {
	return "github.com/zjzjzjzj1874/best-pracrice-go-zero/constants/types.ExportState"
}

func (ExportState) ConstValues() []github_com_go_courier_enumeration.IntStringerEnum {
	return []github_com_go_courier_enumeration.IntStringerEnum{EXPORT_STATE__TODO, EXPORT_STATE__ING, EXPORT_STATE__FAILURE, EXPORT_STATE__SUCCESS}
}

func (v ExportState) MarshalText() ([]byte, error) {
	str := v.String()
	if str == "UNKNOWN" {
		return nil, InvalidExportState
	}
	return []byte(str), nil
}

func (v *ExportState) UnmarshalText(data []byte) (err error) {
	*v, err = ParseExportStateFromString(string(bytes.ToUpper(data)))
	return
}

func (v ExportState) Value() (database_sql_driver.Value, error) {
	offset := 0
	if o, ok := (interface{})(v).(github_com_go_courier_enumeration.DriverValueOffset); ok {
		offset = o.Offset()
	}
	return int64(v) + int64(offset), nil
}

func (v *ExportState) Scan(src interface{}) error {
	offset := 0
	if o, ok := (interface{})(v).(github_com_go_courier_enumeration.DriverValueOffset); ok {
		offset = o.Offset()
	}

	i, err := github_com_go_courier_enumeration.ScanIntEnumStringer(src, offset)
	if err != nil {
		return err
	}
	*v = ExportState(i)
	return nil
}

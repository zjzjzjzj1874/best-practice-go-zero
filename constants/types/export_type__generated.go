package types

import (
	bytes "bytes"
	database_sql_driver "database/sql/driver"
	errors "errors"

	github_com_go_courier_enumeration "github.com/go-courier/enumeration"
)

var InvalidExportType = errors.New("invalid ExportType type")

func ParseExportTypeFromLabelString(s string) (ExportType, error) {
	switch s {
	case "":
		return EXPORT_TYPE_UNKNOWN, nil
	case "Code导出":
		return EXPORT_TYPE__CODE, nil
	}
	return EXPORT_TYPE_UNKNOWN, InvalidExportType
}

func (v ExportType) String() string {
	switch v {
	case EXPORT_TYPE_UNKNOWN:
		return ""
	case EXPORT_TYPE__CODE:
		return "CODE"
	}
	return "UNKNOWN"
}

func ParseExportTypeFromString(s string) (ExportType, error) {
	switch s {
	case "":
		return EXPORT_TYPE_UNKNOWN, nil
	case "CODE":
		return EXPORT_TYPE__CODE, nil
	}
	return EXPORT_TYPE_UNKNOWN, InvalidExportType
}

func (v ExportType) Label() string {
	switch v {
	case EXPORT_TYPE_UNKNOWN:
		return ""
	case EXPORT_TYPE__CODE:
		return "Code导出"
	}
	return "UNKNOWN"
}

func (v ExportType) Int() int {
	return int(v)
}

func (ExportType) TypeName() string {
	return "github.com/zjzjzjzj1874/best-pracrice-go-zero/constants/types.ExportType"
}

func (ExportType) ConstValues() []github_com_go_courier_enumeration.IntStringerEnum {
	return []github_com_go_courier_enumeration.IntStringerEnum{EXPORT_TYPE__CODE}
}

func (v ExportType) MarshalText() ([]byte, error) {
	str := v.String()
	if str == "UNKNOWN" {
		return nil, InvalidExportType
	}
	return []byte(str), nil
}

func (v *ExportType) UnmarshalText(data []byte) (err error) {
	*v, err = ParseExportTypeFromString(string(bytes.ToUpper(data)))
	return
}

func (v ExportType) Value() (database_sql_driver.Value, error) {
	offset := 0
	if o, ok := (interface{})(v).(github_com_go_courier_enumeration.DriverValueOffset); ok {
		offset = o.Offset()
	}
	return int64(v) + int64(offset), nil
}

func (v *ExportType) Scan(src interface{}) error {
	offset := 0
	if o, ok := (interface{})(v).(github_com_go_courier_enumeration.DriverValueOffset); ok {
		offset = o.Offset()
	}

	i, err := github_com_go_courier_enumeration.ScanIntEnumStringer(src, offset)
	if err != nil {
		return err
	}
	*v = ExportType(i)
	return nil
}

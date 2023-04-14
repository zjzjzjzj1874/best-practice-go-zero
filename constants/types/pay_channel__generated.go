package types

import (
	bytes "bytes"
	database_sql_driver "database/sql/driver"
	errors "errors"

	github_com_go_courier_enumeration "github.com/go-courier/enumeration"
)

var InvalidPayChannel = errors.New("invalid PayChannel type")

func ParsePayChannelFromLabelString(s string) (PayChannel, error) {
	switch s {
	case "":
		return PAY_CHANNEL_UNKNOWN, nil
	case "微信支付":
		return PAY_CHANNEL__WECHAT, nil
	case "支付宝支付":
		return PAY_CHANNEL__ALI, nil
	}
	return PAY_CHANNEL_UNKNOWN, InvalidPayChannel
}

func (v PayChannel) String() string {
	switch v {
	case PAY_CHANNEL_UNKNOWN:
		return ""
	case PAY_CHANNEL__WECHAT:
		return "WECHAT"
	case PAY_CHANNEL__ALI:
		return "ALI"
	}
	return "UNKNOWN"
}

func ParsePayChannelFromString(s string) (PayChannel, error) {
	switch s {
	case "":
		return PAY_CHANNEL_UNKNOWN, nil
	case "WECHAT":
		return PAY_CHANNEL__WECHAT, nil
	case "ALI":
		return PAY_CHANNEL__ALI, nil
	}
	return PAY_CHANNEL_UNKNOWN, InvalidPayChannel
}

func (v PayChannel) Label() string {
	switch v {
	case PAY_CHANNEL_UNKNOWN:
		return ""
	case PAY_CHANNEL__WECHAT:
		return "微信支付"
	case PAY_CHANNEL__ALI:
		return "支付宝支付"
	}
	return "UNKNOWN"
}

func (v PayChannel) Int() int {
	return int(v)
}

func (PayChannel) TypeName() string {
	return "anti_fake_api/constants/types.PayChannel"
}

func (PayChannel) ConstValues() []github_com_go_courier_enumeration.IntStringerEnum {
	return []github_com_go_courier_enumeration.IntStringerEnum{PAY_CHANNEL__WECHAT, PAY_CHANNEL__ALI}
}

func (v PayChannel) MarshalText() ([]byte, error) {
	str := v.String()
	if str == "UNKNOWN" {
		return nil, InvalidPayChannel
	}
	return []byte(str), nil
}

func (v *PayChannel) UnmarshalText(data []byte) (err error) {
	*v, err = ParsePayChannelFromString(string(bytes.ToUpper(data)))
	return
}

func (v PayChannel) Value() (database_sql_driver.Value, error) {
	offset := 0
	if o, ok := (interface{})(v).(github_com_go_courier_enumeration.DriverValueOffset); ok {
		offset = o.Offset()
	}
	return int64(v) + int64(offset), nil
}

func (v *PayChannel) Scan(src interface{}) error {
	offset := 0
	if o, ok := (interface{})(v).(github_com_go_courier_enumeration.DriverValueOffset); ok {
		offset = o.Offset()
	}

	i, err := github_com_go_courier_enumeration.ScanIntEnumStringer(src, offset)
	if err != nil {
		return err
	}
	*v = PayChannel(i)
	return nil
}

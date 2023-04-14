package types

import (
	bytes "bytes"
	database_sql_driver "database/sql/driver"
	errors "errors"

	github_com_go_courier_enumeration "github.com/go-courier/enumeration"
)

var InvalidTradeState = errors.New("invalid TradeState type")

func ParseTradeStateFromLabelString(s string) (TradeState, error) {
	switch s {
	case "":
		return TRADE_STATE_UNKNOWN, nil
	case "未支付":
		return TRADE_STATE__NOTPAY, nil
	case "用户支付中":
		return TRADE_STATE__USERPAYING, nil
	case "支付成功":
		return TRADE_STATE__SUCCESS, nil
	case "转入退款":
		return TRADE_STATE__REFUND, nil
	case "已关闭":
		return TRADE_STATE__CLOSED, nil
	case "已撤销":
		return TRADE_STATE__REVOKED, nil
	case "支付失败":
		return TRADE_STATE__PAYERROR, nil
	}
	return TRADE_STATE_UNKNOWN, InvalidTradeState
}

func (v TradeState) String() string {
	switch v {
	case TRADE_STATE_UNKNOWN:
		return ""
	case TRADE_STATE__NOTPAY:
		return "NOTPAY"
	case TRADE_STATE__USERPAYING:
		return "USERPAYING"
	case TRADE_STATE__SUCCESS:
		return "SUCCESS"
	case TRADE_STATE__REFUND:
		return "REFUND"
	case TRADE_STATE__CLOSED:
		return "CLOSED"
	case TRADE_STATE__REVOKED:
		return "REVOKED"
	case TRADE_STATE__PAYERROR:
		return "PAYERROR"
	}
	return "UNKNOWN"
}

func ParseTradeStateFromString(s string) (TradeState, error) {
	switch s {
	case "":
		return TRADE_STATE_UNKNOWN, nil
	case "NOTPAY":
		return TRADE_STATE__NOTPAY, nil
	case "USERPAYING":
		return TRADE_STATE__USERPAYING, nil
	case "SUCCESS":
		return TRADE_STATE__SUCCESS, nil
	case "REFUND":
		return TRADE_STATE__REFUND, nil
	case "CLOSED":
		return TRADE_STATE__CLOSED, nil
	case "REVOKED":
		return TRADE_STATE__REVOKED, nil
	case "PAYERROR":
		return TRADE_STATE__PAYERROR, nil
	}
	return TRADE_STATE_UNKNOWN, InvalidTradeState
}

func (v TradeState) Label() string {
	switch v {
	case TRADE_STATE_UNKNOWN:
		return ""
	case TRADE_STATE__NOTPAY:
		return "未支付"
	case TRADE_STATE__USERPAYING:
		return "用户支付中"
	case TRADE_STATE__SUCCESS:
		return "支付成功"
	case TRADE_STATE__REFUND:
		return "转入退款"
	case TRADE_STATE__CLOSED:
		return "已关闭"
	case TRADE_STATE__REVOKED:
		return "已撤销"
	case TRADE_STATE__PAYERROR:
		return "支付失败"
	}
	return "UNKNOWN"
}

func (v TradeState) Int() int {
	return int(v)
}

func (TradeState) TypeName() string {
	return "anti_fake_api/constants/types.TradeState"
}

func (TradeState) ConstValues() []github_com_go_courier_enumeration.IntStringerEnum {
	return []github_com_go_courier_enumeration.IntStringerEnum{TRADE_STATE__NOTPAY, TRADE_STATE__USERPAYING, TRADE_STATE__SUCCESS, TRADE_STATE__REFUND, TRADE_STATE__CLOSED, TRADE_STATE__REVOKED, TRADE_STATE__PAYERROR}
}

func (v TradeState) MarshalText() ([]byte, error) {
	str := v.String()
	if str == "UNKNOWN" {
		return nil, InvalidTradeState
	}
	return []byte(str), nil
}

func (v *TradeState) UnmarshalText(data []byte) (err error) {
	*v, err = ParseTradeStateFromString(string(bytes.ToUpper(data)))
	return
}

func (v TradeState) Value() (database_sql_driver.Value, error) {
	offset := 0
	if o, ok := (interface{})(v).(github_com_go_courier_enumeration.DriverValueOffset); ok {
		offset = o.Offset()
	}
	return int64(v) + int64(offset), nil
}

func (v *TradeState) Scan(src interface{}) error {
	offset := 0
	if o, ok := (interface{})(v).(github_com_go_courier_enumeration.DriverValueOffset); ok {
		offset = o.Offset()
	}

	i, err := github_com_go_courier_enumeration.ScanIntEnumStringer(src, offset)
	if err != nil {
		return err
	}
	*v = TradeState(i)
	return nil
}

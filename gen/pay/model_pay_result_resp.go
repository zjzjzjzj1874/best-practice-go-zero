/*
 * pay-api
 *
 * 支付模块
 *
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type PayResultResp struct {
	TradeNo string `json:"trade_no"`
	OrderNo string `json:"order_no"`
	OrderState int32 `json:"order_state"`
}

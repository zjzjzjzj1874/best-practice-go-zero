/*
 * pay-api
 *
 * 支付模块
 *
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type PayReq struct {
	Id int32 `json:"id"`
	PayChannel int32 `json:"pay_channel"`
	ReturnUrl string `json:"return_url,omitempty"`
}
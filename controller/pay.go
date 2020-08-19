package controller

import (
	"github.com/davveo/singleTsquare/utils/response"
	"github.com/gin-gonic/gin"
)

/*
/获取支付配置接口
*/
func PayConfig(ctx *gin.Context) {
	response.OkWithData(
		map[string]interface{}{
			"channels": map[string]interface{}{
				"alipay": map[string]interface{}{
					"img": "", // 支付图片
				},
				"wechat": map[string]interface{}{
					"img": "", // 图片
				},
				"yibao": map[string]interface{}{
					"img":  "", // 图片
					"code": "", // 一些支付编码
				},
			},
		}, ctx)
}

/*
/创建支付链接
POST http://127.0.0.1:8080/api/v1/pay/create
{
	trade_id: xxxx
}
*/
func CreatePay(ctx *gin.Context) {

}

/*
/ 支付退款接口
*/
func Refund(ctx *gin.Context) {

}

/*
/ 支付异步通知接口
*/
func Notify(ctx *gin.Context) {

}

/*
/ 支付同步通知接口
*/
func Return(ctx *gin.Context) {

}

/*
/ 支付查询接口
*/
func QueryTrade(ctx *gin.Context) {

}

/*
/ 退款查询接口
*/
func QueryRefund(ctx *gin.Context) {

}

/*
/ 查询账单接口
*/
func QueryBill(ctx *gin.Context) {

}

/*
/ 结账接口
*/
func QuerySettle(ctx *gin.Context) {

}

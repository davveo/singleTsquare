package controller

import (
	"github.com/davveo/singleTsquare/utils/response"
	"github.com/gin-gonic/gin"
)

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

func CreatePay(ctx *gin.Context) {

}

func Refund(ctx *gin.Context) {

}

func Notify(ctx *gin.Context) {

}

func Return(ctx *gin.Context) {

}

func QueryTrade(ctx *gin.Context) {

}

func QueryRefund(ctx *gin.Context) {

}

func QueryBill(ctx *gin.Context) {

}

func QuerySettle(ctx *gin.Context) {

}

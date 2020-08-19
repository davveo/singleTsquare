package router

import (
	"github.com/davveo/singleTsquare/controller"
	"github.com/gin-gonic/gin"
)

func PayRouter(router *gin.RouterGroup) {
	payRouter := router.Group("pay")
	{
		payRouter.GET("config", controller.PayConfig)                 // 支付配置
		payRouter.POST("create", controller.CreatePay)                // 发起支付
		payRouter.POST("refund", controller.Refund)                   // 退款接口
		payRouter.POST("notify/:channel/:tradeno", controller.Notify) // 支付异步通知
		payRouter.POST("return/:channel/:tradeno", controller.Return) // 支付同步通知
		payRouter.GET("query/trade", controller.QueryTrade)           // 交易查询
		payRouter.GET("query/bill", controller.QueryBill)             // 获取账单
		payRouter.GET("query/refund", controller.QueryRefund)         // 查询退款
		payRouter.GET("query/settle", controller.QuerySettle)         // 结算明显
	}

}

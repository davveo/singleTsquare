package router

import (
	"github.com/davveo/singleTsquare/controller"
	"github.com/gin-gonic/gin"
)

func CommonRouter(router *gin.RouterGroup) {
	// 前端路由
	commonRouter := router.Group("")
	{
		commonRouter.POST("code", controller.Code)
		commonRouter.POST("captcha", controller.Captcha)
		commonRouter.POST("qrcode", controller.QrCode)
		commonRouter.GET("health_check", controller.HealthCheck)

		// 第三登录回调url
		// qq授权登录回调
		commonRouter.GET("qqLogin", controller.QQLoginCallBack)
		// 微博授权登录回调
		commonRouter.GET("wbLogin", controller.WBLoginCallBack)
		// github授权登录回调
		commonRouter.GET("gbLogin", controller.GBLoginCallBack)
		// wechat授权登录回调
		commonRouter.GET("wcLogin", controller.WCLoginCallBack)
	}

}

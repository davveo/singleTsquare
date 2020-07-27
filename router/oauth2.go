package router

import (
	"github.com/davveo/singleTsquare/controller"
	"github.com/gin-gonic/gin"
)

func Oauth2Router(router *gin.RouterGroup) {
	// 前端路由
	commonRouter := router.Group("oauth")
	{
		// 第三登录回调url
		// qq授权登录回调
		commonRouter.GET("qqLogin", controller.QQLoginCallBack)
		// 微博授权登录回调
		commonRouter.GET("wbLogin", controller.WBLoginCallBack)
		// github授权登录回调
		commonRouter.GET("gbLogin", controller.GBLoginCallBack)
		// wechat授权登录回调
		commonRouter.GET("wcLogin", controller.WCLoginCallBack)

		commonRouter.POST("login", controller.OauthLogin) //扫码登录
	}

}

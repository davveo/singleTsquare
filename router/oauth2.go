package router

import (
	"github.com/davveo/singleTsquare/controller"
	"github.com/gin-gonic/gin"
)

func Oauth2Router(router *gin.RouterGroup) {
	// 前端路由
	commonRouter := router.Group("oauth")
	{
		// 第三方授权登录
		commonRouter.POST("login", controller.OauthLogin)
		// qq授权登录回调
		commonRouter.GET(":service/callback", controller.OauthCallBack)
	}

}

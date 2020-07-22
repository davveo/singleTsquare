package router

import (
	"github.com/davveo/singleTsquare/controller"
	"github.com/gin-gonic/gin"
)

func CommonRouter(router *gin.RouterGroup) {
	// 前端路由
	commonRouter := router.Group("common")
	{
		commonRouter.POST("code", controller.Code)
		commonRouter.POST("captcha", controller.Captcha)
		commonRouter.GET("health_check", controller.HealthCheck)
	}

}

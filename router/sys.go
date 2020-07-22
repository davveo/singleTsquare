package router

import (
	"github.com/davveo/singleTsquare/controller"
	"github.com/gin-gonic/gin"
)

func SysRouter(router *gin.RouterGroup) {
	// 系统路由
	sysRouter := router.Group("sys")
	{
		sysRouter.GET("menu", controller.Menu)
	}

}

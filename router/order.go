package router

import (
	"github.com/gin-gonic/gin"
)

func OrderRouter(router *gin.RouterGroup) {
	orderRouter := router.Group("order")
	{
		orderRouter.GET("create", func(context *gin.Context) {

		})
	}

}

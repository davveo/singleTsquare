package router

import (
	"github.com/davveo/singleTsquare/controller"
	"github.com/davveo/singleTsquare/middleware"
	"github.com/gin-gonic/gin"
)

func UserRouter(router *gin.RouterGroup) {
	userRouter := router.Group("user").Use(middleware.JwtMiddleWare())
	{
		userRouter.POST("create", controller.UserCreate)
		userRouter.GET("d/:userId", controller.UserGet)
		userRouter.PUT("d/:userId", controller.UserUpdate)
		userRouter.DELETE("d/:userId", controller.UserDelete)
		userRouter.GET("list", controller.UserList)
	}
}

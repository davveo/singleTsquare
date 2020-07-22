package router

import (
	"github.com/davveo/singleTsquare/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	// 中间件管理
	router.Use(middleware.CorsMiddleWare())
	router.Use(middleware.JwtMiddleWare())

	// 路由管理
	ApiGroup := router.Group("api")

	UserRouter(ApiGroup)
	CommonRouter(ApiGroup)

	return router
}

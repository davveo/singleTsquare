package router

import (
	"github.com/davveo/singleTsquare/controller"
	"github.com/davveo/singleTsquare/middleware"
	"github.com/gin-gonic/gin"
)

func UserRouter(router *gin.RouterGroup) {
	// 前端路由
	userOuterRouter := router.Group("user").Use(middleware.JwtMiddleWare())
	{
		userOuterRouter.POST("register", controller.Register)              //注册接口
		userOuterRouter.POST("login", controller.Login)                    //登录接口
		userOuterRouter.POST("fast_login", controller.FastLogin)           //快捷登录
		userOuterRouter.POST("scan", controller.ScanLogin)                 //第三方登录
		userOuterRouter.POST("logout", controller.Logout)                  //登出接口
		userOuterRouter.GET("user_id", controller.Get)                     //获取用户信息
		userOuterRouter.PUT("user_id", controller.Update)                  //修改用户信息
		userOuterRouter.POST("change_password", controller.ChangePassword) //修改用户信息

	}

	// 后台路由
	//userInnerRouter := router.Group("i_user").Use(middleware.JwtMiddleWare())
	//{
	//	userInnerRouter.POST("register", controller.Register)
	//	userInnerRouter.GET("list", controller.List)
	//	userInnerRouter.GET("captcha", controller.Captcha)
	//}

}

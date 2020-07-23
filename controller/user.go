package controller

import (
	"fmt"

	"github.com/davveo/singleTsquare/models"

	"github.com/davveo/singleTsquare/services"

	"github.com/davveo/singleTsquare/models/request"
	"github.com/davveo/singleTsquare/utils/ip"
	"github.com/davveo/singleTsquare/utils/response"
	"github.com/davveo/singleTsquare/utils/str"
	"github.com/gin-gonic/gin"
)

var (
	ErrorPassword    = fmt.Sprintf("两次密码不一致")
	ErrorVerifyCode  = fmt.Sprintf("验证码不正确")
	UserHasRegister  = fmt.Sprintf("用户名已经注册")
	PhoneHasRegister = fmt.Sprintf("手机号已经注册")
)

func Register(context *gin.Context) {
	var userRequest request.UserRequest
	clientIp := ip.ClientIP(context.Request)

	if err := context.ShouldBindJSON(&userRequest); err != nil {
		response.FailWithMessage(err.Error(), context)
		return
	}
	// 比较验证码
	verifycodestr := fmt.Sprintf("verifycode:%s", userRequest.Phone)
	bverifycode, _ := Cache.Get(verifycodestr)
	_ = Cache.Delete(verifycodestr)

	if str.ByteTostr(bverifycode) != userRequest.Code {
		response.FailWithMoreMessage("", ErrorVerifyCode, context)
		return
	}
	// 比较两次密码
	if userRequest.Password != userRequest.RepeatPassword {
		response.FailWithMoreMessage("", ErrorPassword, context)
		return
	}

	if services.UserService.UserExistByUsername(userRequest.UserName) {
		response.FailWithMoreMessage("", UserHasRegister, context)
		return
	}

	if services.UserService.UserExistByPhone(userRequest.Phone) {
		response.FailWithMoreMessage("", PhoneHasRegister, context)
		return
	}

	if _, err := services.UserService.Create(
		userRequest.UserName,
		userRequest.Password,
		userRequest.Phone, clientIp); err != nil {
		response.FailWithMessage(err.Error(), context)
		return
	}
	response.Ok(context)
}

// username/phone三者择一
// password	必传
func Login(context *gin.Context) {
	var (
		loginRequest request.LoginRequest
		accountUser  *models.AccountUser
	)

	clientIp := ip.ClientIP(context.Request)

	if err := context.ShouldBindJSON(&loginRequest); err != nil {
		response.FailWithMessage(err.Error(), context)
		return
	}

	if loginRequest.UserName != "" {
		accountUser, _ = services.UserService.FindUserByUsername(loginRequest.UserName)
	}

	if loginRequest.Phone != "" {
		accountUser, _ = services.UserService.FindUserByPhone(loginRequest.Phone)
	}

	if accountUser != nil {

		// 更新用户信息
		_ = services.UserService.UpdateUser(accountUser, clientIp)
	} else {
		// 用户不存在
	}
}

func FastLogin(context *gin.Context) {

}

func OauthLogin(context *gin.Context) {

}

func Logout(context *gin.Context) {

}

func Get(context *gin.Context) {

}

func Update(context *gin.Context) {
	// 获取uri参数
	userId := context.Param("userId")

	// 获取query参数
	name := context.DefaultQuery("name", "")

	fmt.Println(userId, name)
}

func List(context *gin.Context) {

}

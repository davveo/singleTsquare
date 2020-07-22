package controller

import (
	"fmt"
	"github.com/davveo/singleTsquare/models/request"
	"github.com/davveo/singleTsquare/services"
	"github.com/davveo/singleTsquare/utils/ip"
	"github.com/davveo/singleTsquare/utils/response"
	"github.com/davveo/singleTsquare/utils/str"
	"github.com/gin-gonic/gin"
)

var (
	userService services.UserService

	ErrorPassword    = fmt.Sprintf("两次密码不一致")
	ErrorVerifyCode  = fmt.Sprintf("验证码不正确")
	UserHasRegister  = fmt.Sprintf("用户名已经注册")
	PhoneHasRegister = fmt.Sprintf("手机号已经注册")
)

//  username	string	非必传	用户账号
// email	string	email/phone两者择一	用户邮箱
// phone	string	email/phone两者择一	用户手机号
// code	int	必传	验证码
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
	if str.ByteTostr(bverifycode) != userRequest.Code {
		response.FailWithMoreMessage("", ErrorVerifyCode, context)
		return
	}
	// 比较两次密码
	if userRequest.Password != userRequest.RepeatPassword {
		response.FailWithMoreMessage("", ErrorPassword, context)
		return
	}

	if userService.UserExistByUsername(userRequest.UserName) {
		response.FailWithMoreMessage("", UserHasRegister, context)
		return
	}

	if userService.UserExistByPhone(userRequest.Phone) {
		response.FailWithMoreMessage("", PhoneHasRegister, context)
		return
	}

	if _, err := userService.Create(
		userRequest.UserName,
		userRequest.Password,
		userRequest.Phone, clientIp); err != nil {
		response.FailWithMessage(err.Error(), context)
		return
	}
	response.Ok(context)
}

// username/email/phone三者择一
// password	必传
func Login(context *gin.Context) {

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

package controller

import (
	"fmt"
	"net/http"

	"github.com/davveo/singleTsquare/services"

	"github.com/davveo/singleTsquare/models/request"
	"github.com/gin-gonic/gin"
)

var (
	userService services.UserService
)

//username	string	非必传	用户账号
// email	string	email/phone两者择一	用户邮箱
// phone	string	email/phone两者择一	用户手机号
// code	int	必传	验证码
func Register(context *gin.Context) {
	// 包括前台和后台注册
	var userRequst request.UserRequest
	if err := context.ShouldBindJSON(&userRequst); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := userService.FindUserByUsername(userRequst.UserName)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err})
	}
	fmt.Println(user)
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

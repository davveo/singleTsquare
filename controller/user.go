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
	UserHasExist     = fmt.Sprintf("用户名已经存在")
	PhoneHasRegister = fmt.Sprintf("手机号已经注册")
	EmailHasExist    = fmt.Sprintf("邮箱已存在")
)

/*
入参
{
	"username": "davveo",
	"email": "test@qq.com",
	"phone": "1213213123",
	"code": "123123"
}

返回

*/
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

	if services.AccountService.ExistByUserName(userRequest.UserName) {
		response.FailWithMoreMessage("", UserHasExist, context)
		return
	}

	if services.AccountService.ExistByPhone(userRequest.Phone) {
		response.FailWithMoreMessage("", PhoneHasRegister, context)
		return
	}
	if services.AccountService.ExistByMail(userRequest.Email) {
		response.FailWithMoreMessage("", EmailHasExist, context)
		return
	}

	if _, err := services.AccountService.Create(
		userRequest.UserName,
		userRequest.Password,
		userRequest.Phone,
		userRequest.Email,
		clientIp); err != nil {
		response.FailWithMessage(err.Error(), context)
		return
	}
	response.Ok(context)
}

/*
入参
{
	"loginid": "",
	"password": ""
}
返回

*/

func Login(context *gin.Context) {
	var loginRequest request.LoginRequest
	clientIp := ip.ClientIP(context.Request)
	if err := context.ShouldBindJSON(&loginRequest); err != nil {
		response.FailWithMessage(err.Error(), context)
		return
	}
	// loginRequest.LoginId == email/username
	user, err := LoginTool(clientIp, loginRequest.LoginId)
	if err != nil {
		response.FailWithMessage(err.Error(), context)
		return
	}
	response.OkWithData(
		map[string]interface{}{
			"nickname": user.NickName,
			"avatar":   user.Avatar,
			"user_id":  user.ID,
		}, context)
}

/*
快速登录
入参
{
	"login_id": "123321312",  //手机号或者邮箱
	"code": "31313"
}
*/

func FastLogin(context *gin.Context) {
	var fastLoginRequest request.FastLoginRequest
	clientIp := ip.ClientIP(context.Request)
	if err := context.ShouldBindJSON(&fastLoginRequest); err != nil {
		response.FailWithMessage(err.Error(), context)
		return
	}
	verifycodestr := fmt.Sprintf("verifycode:%s", fastLoginRequest.Phone)
	bverifycode, _ := Cache.Get(verifycodestr)
	_ = Cache.Delete(verifycodestr)

	if str.ByteTostr(bverifycode) != fastLoginRequest.Code {
		response.FailWithMoreMessage("", ErrorVerifyCode, context)
		return
	}

	user, err := LoginTool(clientIp, fastLoginRequest.Phone)
	if err != nil {
		response.FailWithMessage(err.Error(), context)
		return
	}

	response.OkWithData(
		map[string]interface{}{
			"nickname": user.NickName,
			"avatar":   user.Avatar,
			"user_id":  user.ID,
		}, context)
}

// 扫码登录
func ScanLogin(context *gin.Context) {

}

func Logout(context *gin.Context) {

}

func Get(context *gin.Context) {
	userId := context.DefaultQuery("user_id", "")
	fmt.Println(userId)
}

func Update(context *gin.Context) {
	// 获取query参数
	userId := context.DefaultQuery("user_id", "")

	fmt.Println(userId)
}

func ChangePassword(context *gin.Context) {

}

func List(context *gin.Context) {

}

func LoginTool(clientIp, identify string) (user *models.User, err error) {
	account, err := services.AccountService.FindByLoginId(identify)
	if err != nil {
		return nil, err
	}

	user, err = services.UserService.FindByUid(account.ID)
	if err != nil {
		return nil, err
	}
	// 更新账户信息
	_ = services.AccountService.UpdateAccount(clientIp, account)

	return
}

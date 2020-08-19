package controller

import (
	"fmt"

	"github.com/davveo/singleTsquare/services"

	"github.com/davveo/singleTsquare/models/request"
	"github.com/davveo/singleTsquare/utils/ip"
	"github.com/davveo/singleTsquare/utils/response"
	"github.com/gin-gonic/gin"
)

var (
	shortAccountService = services.AccountService
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
	if !BindCheck(&userRequest, context) {
		response.FailWithMessage(response.ParamValidateFailed, context)
		return
	}
	if !VerifyCodeUtil(userRequest.Phone, userRequest.Code) {
		response.FailWithMoreMessage("", response.ErrorVerifyCode, context)
		return
	}
	// 比较两次密码
	if userRequest.Password != userRequest.RepeatPassword {
		response.FailWithMoreMessage("", response.ErrorPassword, context)
		return
	}

	if shortAccountService.ExistByUserName(userRequest.UserName) {
		response.FailWithMoreMessage("", response.UserHasExist, context)
		return
	}
	if shortAccountService.ExistByPhone(userRequest.Phone) {
		response.FailWithMoreMessage("", response.PhoneHasRegister, context)
		return
	}
	if shortAccountService.ExistByMail(userRequest.Email) {
		response.FailWithMoreMessage("", response.EmailHasExist, context)
		return
	}

	if _, err := shortAccountService.Create(
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
	if !BindCheck(&loginRequest, context) {
		response.FailWithMessage(response.ParamValidateFailed, context)
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

	if !BindCheck(&fastLoginRequest, context) {
		response.FailWithMessage(response.ParamValidateFailed, context)
		return
	}

	if !VerifyCodeUtil(fastLoginRequest.Phone, fastLoginRequest.Code) {
		response.FailWithMoreMessage("", response.ErrorVerifyCode, context)
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

/*
修改密码: 用户处于登录状态 login_id + password

忘记密码: 用户处于未登录状态, 流程为:
	先进行邮箱或者手机号验证-->发送code
	然后在进行密码修改
{
	"login_id": "" // email/phone/username
	"password": ""
}
*/
func ChangePassword(context *gin.Context) {
	var loginRequest request.LoginRequest
	if !BindCheck(&loginRequest, context) {
		response.FailWithMoreMessage(
			"", response.ParamValidateFailed, context)
		return
	}
	// loginRequest.LoginId == email/username/phone
	account, _ := shortAccountService.FindByLoginId(loginRequest.LoginId)
	err := shortAccountService.UpdateAccountPassword(loginRequest.Password, account)
	if err != nil {
		response.FailWithMoreMessage(
			err.Error(), response.FaildUpdateAccountPassword, context)
		return
	}
	response.Ok(context)
}

/*
POST
application/json
{
	"identify_id": "xxxx",
	"phone": "123213123",
	"code": "1232131",
	"login_id": "",
    "password": ""
}

identify_id 必传
phone + code
login_id + password
*/
func BindAccount(context *gin.Context) {
	// identify_id, phone, email, username
	// 将第三方的identify_id与系统phone email username进行绑定
	// 目前支持绑定手机号
	var (
		bindRequest *request.BindRequest
		err         error
	)
	if err := context.ShouldBindJSON(&bindRequest); err != nil {
		response.FailWithMessage(err.Error(), context)
		return
	}
	clientIp := ip.ClientIP(context.Request)
	accountPlatform, err := shortPlatformService.FindByIdentifyId(bindRequest.IdentifyId)
	if err != nil {
		response.FailWithMessage(err.Error(), context)
		return
	}
	if bindRequest.Phone != "" && bindRequest.Code != "" { // phone+code
		err = BindAccountByPhone(clientIp, bindRequest, accountPlatform)

	} else if bindRequest.LoginId != "" && bindRequest.Password != "" { // loginRequest.LoginId == email/username
		err = BindAccountByEmailOrUserName(clientIp, bindRequest, accountPlatform)
	}
	if err != nil {
		response.FailWithMoreMessage(err.Error(), "绑定失败!", context)
		return
	}
	response.OkDetailed(map[string]interface{}{
		"nickname":            accountPlatform.NickName,
		"avatar":              accountPlatform.Avatar,
		"account_platform_id": accountPlatform.ID,
		"account_id":          accountPlatform.AccountID,
	}, "绑定成功!", context)
}

/*
{
	"phone": "321312312",
	"code": "21312",
}
验证手机号+code
*/

func VerifyCode(context *gin.Context) {
	var verifyCodeRequest request.VerifyCodeRequest
	if !BindCheck(&verifyCodeRequest, context) {
		response.FailWithMessage(response.ParamValidateFailed, context)
		return
	}
	if !VerifyCodeUtil(verifyCodeRequest.Phone, verifyCodeRequest.Code) {
		response.FailWithMoreMessage("", response.ErrorVerifyCode, context)
		return
	}
	response.Ok(context)
}

func Logout(context *gin.Context) {

}

/*
根据user_id获取用户信息
*/
func UserGet(context *gin.Context) {
	userId := context.DefaultQuery("user_id", "")
	fmt.Println(userId)
}

func UserUpdate(context *gin.Context) {
	// 获取query参数
	userId := context.DefaultQuery("user_id", "")

	fmt.Println(userId)
}

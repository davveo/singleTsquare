package controller

import (
	"fmt"

	"github.com/davveo/singleTsquare/services"

	"github.com/davveo/singleTsquare/models/request"
	"github.com/davveo/singleTsquare/utils/ip"
	"github.com/davveo/singleTsquare/utils/response"
	"github.com/davveo/singleTsquare/utils/str"
	"github.com/gin-gonic/gin"
)

var (
	ErrorPassword      = fmt.Sprintf("两次密码不一致")
	ErrorVerifyCode    = fmt.Sprintf("验证码不正确")
	UserHasExist       = fmt.Sprintf("用户名已经存在")
	PhoneHasRegister   = fmt.Sprintf("手机号已经注册")
	EmailHasExist      = fmt.Sprintf("邮箱已存在")
	FaildCreateAccount = fmt.Sprintf("创建账户失败")
	FaildUpdateAccount = fmt.Sprintf("更新账户失败")
	BindFailed         = fmt.Sprintf("绑定失败")

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

	if shortAccountService.ExistByUserName(userRequest.UserName) {
		response.FailWithMoreMessage("", UserHasExist, context)
		return
	}
	if shortAccountService.ExistByPhone(userRequest.Phone) {
		response.FailWithMoreMessage("", PhoneHasRegister, context)
		return
	}
	if shortAccountService.ExistByMail(userRequest.Email) {
		response.FailWithMoreMessage("", EmailHasExist, context)
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

/*
{
	"account_id": "",
	"password": "",
}
*/
func ChangePassword(context *gin.Context) {

}

/*

 */
func ResetPassword(context *gin.Context) {

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
func BindAccountController(context *gin.Context) {
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

func List(context *gin.Context) {

}

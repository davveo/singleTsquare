package response

import (
	"fmt"
)

var (
	ErrorPassword              = fmt.Sprintf("两次密码不一致")
	ErrorVerifyCode            = fmt.Sprintf("验证码不正确")
	UserHasExist               = fmt.Sprintf("用户名已经存在")
	PhoneHasRegister           = fmt.Sprintf("手机号已经注册")
	EmailHasExist              = fmt.Sprintf("邮箱已存在")
	FaildCreateAccount         = fmt.Sprintf("创建账户失败")
	FaildUpdateAccount         = fmt.Sprintf("更新账户失败")
	FaildUpdateAccountPassword = fmt.Sprintf("更新账户密码失败")
	BindFailed                 = fmt.Sprintf("绑定失败")
	ParamValidateFailed        = fmt.Sprintf("参数校验失败")
	AccountNotExist            = fmt.Sprintf("不存在用户信息")
	PhoneNotExist              = fmt.Sprintf("手机号不存在")
	EmailNotExist              = fmt.Sprintf("邮箱不存在")
)

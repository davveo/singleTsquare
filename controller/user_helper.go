package controller

import (
	"errors"
	"fmt"

	"github.com/davveo/singleTsquare/models"
	"github.com/davveo/singleTsquare/models/request"
	"github.com/davveo/singleTsquare/services"
	"github.com/davveo/singleTsquare/utils/common"
	"github.com/davveo/singleTsquare/utils/randomstr"
	"github.com/davveo/singleTsquare/utils/str"
)

func VerifyCode(phone, code string) bool {
	verifycodestr := fmt.Sprintf("verifycode:%s", phone)
	bverifycode, _ := Cache.Get(verifycodestr)
	_ = Cache.Delete(verifycodestr)
	return str.ByteTostr(bverifycode) == code
}

func BindAccountByPhone(clientIp string,
	bindRequest *request.BindRequest,
	accountPlatform *models.AccountPlatform) error {
	if !VerifyCode(bindRequest.Phone, bindRequest.Code) {
		return errors.New(ErrorVerifyCode)
	}
	// 查找手机号
	// 如果没有找到, 则进行创建, 然后绑定; 否则直接更新就行
	userNameOrPassword := randomstr.GenRandomString(6)
	accountService, err := shortAccountService.FindByPhone(bindRequest.Phone)
	if err != nil { // 不存在
		accountService, err = shortAccountService.Create(
			userNameOrPassword, userNameOrPassword,
			bindRequest.Phone, "", clientIp)
		if err != nil {
			return errors.New(FaildCreateAccount)
		}
	}
	if err = shortPlatformService.UpdateAccountId(accountService.ID, accountPlatform); err != nil {
		return errors.New(FaildUpdateAccount)
	}
	return nil
}

func BindAccountByEmailOrUserName(clientIp string,
	bindRequest *request.BindRequest,
	accountPlatform *models.AccountPlatform) error {
	accountService, err := shortAccountService.FindByLoginId(bindRequest.LoginId)
	if err != nil { // 如果不存在呢?
		userNameOrPassword := randomstr.GenRandomString(6)
		// 需要判断login_id是邮箱还是username
		if common.IsEmail(bindRequest.LoginId) {
			accountService, err = shortAccountService.Create(
				userNameOrPassword, bindRequest.Password, "", bindRequest.LoginId, clientIp)
		} else {
			accountService, err = shortAccountService.Create(
				bindRequest.LoginId, bindRequest.Password, "", "", clientIp)
		}
		if err != nil {
			return errors.New(BindFailed)
		}
	}
	if err = shortPlatformService.UpdateAccountId(accountService.ID, accountPlatform); err != nil {
		return errors.New(FaildUpdateAccount)
	}
	return nil
}

func LoginTool(clientIp, identify string) (user *models.User, err error) {
	account, err := shortAccountService.FindByLoginId(identify)
	if err != nil {
		return nil, err
	}

	user, err = services.UserService.FindByAccountID(account.ID)
	if err != nil {
		return nil, err
	}
	// 更新账户信息
	_ = shortAccountService.UpdateAccountIp(clientIp, account)

	return
}

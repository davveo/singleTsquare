package controller

import (
	"fmt"
	"time"

	"github.com/davveo/singleTsquare/models"

	"github.com/davveo/singleTsquare/services"

	"github.com/davveo/singleTsquare/utils/common"
	"github.com/davveo/singleTsquare/utils/email"
	"github.com/davveo/singleTsquare/utils/oauth2/github"
	"github.com/davveo/singleTsquare/utils/oauth2/qq"
	"github.com/davveo/singleTsquare/utils/oauth2/weibo"

	"github.com/davveo/singleTsquare/utils/sms"

	"github.com/davveo/singleTsquare/utils/str"

	"github.com/davveo/singleTsquare/models/request"
	"github.com/davveo/singleTsquare/utils/log"

	"github.com/allegro/bigcache"

	"github.com/davveo/singleTsquare/utils/captcha"
	"github.com/davveo/singleTsquare/utils/code"
	"github.com/davveo/singleTsquare/utils/response"
	"github.com/gin-gonic/gin"
)

var (
	config = bigcache.Config{
		Shards:             16,              // 存储的条目数量，值必须是2的幂
		LifeWindow:         2 * time.Second, // 超时后条目被删除
		MaxEntriesInWindow: 0,               // 在 Life Window 中的最大数量
		MaxEntrySize:       0,               // 条目最大尺寸，以字节为单位
		HardMaxCacheSize:   0,               // 设置缓存最大值，以MB为单位，超过了不在分配内存。0表示无限制分配
	}
	Cache, _ = bigcache.NewBigCache(config)
)

func Code(context *gin.Context) {
	var loginRequestJson request.LoginRequestJson
	if err := context.ShouldBindJSON(&loginRequestJson); err != nil {
		response.FailWithMessage(err.Error(), context)
		return
	}

	// 生成验证码
	verifyCode := code.GenerateVerifyCode()
	_ = Cache.Set(
		fmt.Sprintf("verifycode:%s", loginRequestJson.LoginId), str.StrToByte(verifyCode))

	// login_id 可能是邮箱或者手机号
	// TODO 考虑使用策略模式
	// TODO 考虑异步发送
	if common.VerifyEmailFormat(loginRequestJson.LoginId) {
		subject := "登录注册码邮件"
		bodyMsg := fmt.Sprintf("登录注册码为: %s", verifyCode)
		if err := email.Send(loginRequestJson.LoginId, subject, bodyMsg); err != nil {
			response.FailWithMessage(err.Error(), context)
			return
		}
	}
	if common.VerifyMobileFormat(loginRequestJson.LoginId) {
		if err := sms.Send(loginRequestJson.LoginId); err != nil {
			response.FailWithMessage(err.Error(), context)
			return
		}
	}

	response.Ok(context)
}

func Captcha(context *gin.Context) {
	// captchaType=[audio, string, math, chinese]
	captchaType := context.DefaultQuery("captchaType", "")

	id, b64s, err := captcha.GenerateCaptcha(captchaType)
	if err != nil {
		log.ERROR.Println(err)
		response.FailWithMessage(err.Error(), context)
		return
	}
	response.OkWithData(
		map[string]interface{}{
			"data":      b64s,
			"captchaId": id,
		}, context)

}

func HealthCheck(context *gin.Context) {
	response.Ok(context)
}

func QrCode(context *gin.Context) {

}

// qq授权回调
func QQLoginCallBack(context *gin.Context) {
	codeStr := context.DefaultQuery("code", "")
	userInfo, err := qq.GetUserInfo(codeStr)
	if err != nil {
		response.FailWithMessage(err.Error(), context)
		return
	}
	// TODO 需要将用户的信息与第三方信息做一个绑定

	// qq头像存在多个figureurl_qq_2, figureurl_qq_1, figureurl, figureurl_1, figureurl_2
	accountPlatform, _ := services.AccountPlatformService.FindByIdentifyId(userInfo.OpenId)
	if accountPlatform != nil { // 如果存在, 则更新
		_ = services.AccountPlatformService.UpdateByIdentifyId(
			userInfo.AccessToken,
			userInfo.NickName,
			userInfo.Avatar,
			accountPlatform)
	}
	accountPlatformUser, err := services.AccountPlatformService.Create(
		0, // 默认, 等待后续绑定
		models.GetPlatformType("qq"),
		userInfo.OpenId,
		userInfo.AccessToken,
		userInfo.NickName,
		userInfo.Avatar)

	if err != nil {
		response.FailWithMessage(err.Error(), context)
		return
	}
	// 返回用户信息
	response.OkWithData(map[string]interface{}{
		"uid":         accountPlatformUser.Uid,
		"avatar":      accountPlatformUser.Avatar,
		"nickname":    accountPlatformUser.NickName,
		"identify_id": accountPlatformUser.IdentifyId,
	}, context)
}

// 微博授权回调
func WBLoginCallBack(context *gin.Context) {
	codeStr := context.DefaultQuery("code", "")
	// 获取token
	userInfo, err := weibo.GetUserInfo(codeStr)
	if err != nil {
		response.FailWithMessage(err.Error(), context)
		return
	}

	accountPlatform, _ := services.AccountPlatformService.FindByIdentifyId(userInfo.OpenId)
	if accountPlatform != nil { // 如果存在, 则更新
		_ = services.AccountPlatformService.UpdateByIdentifyId(
			userInfo.AccessToken,
			userInfo.NickName,
			userInfo.Avatar,
			accountPlatform)
	}
	accountPlatformUser, err := services.AccountPlatformService.Create(
		0, // 默认, 等待后续绑定
		models.GetPlatformType("weibo"),
		userInfo.OpenId,
		userInfo.AccessToken,
		userInfo.NickName,
		userInfo.Avatar)

	if err != nil {
		response.FailWithMessage(err.Error(), context)
		return
	}
	// 返回用户信息
	response.OkWithData(map[string]interface{}{
		"uid":         accountPlatformUser.Uid,
		"avatar":      accountPlatformUser.Avatar,
		"nickname":    accountPlatformUser.NickName,
		"identify_id": accountPlatformUser.IdentifyId,
	}, context)
}

// github授权回调
func GBLoginCallBack(context *gin.Context) {
	codeStr := context.DefaultQuery("code", "")
	userinfo, err := github.Oauth(codeStr)
	if err != nil {
		response.FailWithMessage(err.Error(), context)
		return
	}
	// TODO 用户信息入库
	fmt.Println(userinfo)
	response.Ok(context)
}

// 微信授权回调
func WCLoginCallBack(context *gin.Context) {

}

package controller

import (
	"net/http"

	"github.com/davveo/singleTsquare/services"
	"github.com/davveo/singleTsquare/utils/oauth2/base"
	"github.com/davveo/singleTsquare/utils/response"
	"github.com/gin-gonic/gin"
)

var (
	shortPlatformService = services.AccountPlatformService
)

// api/v1/user/oauth_login?service=qq
// service = qq/weibo/github/facebook/wechat
func OauthLogin(context *gin.Context) {
	service := context.DefaultQuery("service", "qq")
	oauthService, err := base.OauthService(service)
	if err != nil {
		response.FailWithMessage(err.Error(), context)
		return
	}

	redirectUrl := oauthService.GenRedirectURL()
	context.Redirect(http.StatusMovedPermanently, redirectUrl)
}

// qq授权回调
// 127.0.0.1:8080/api/v1/qq[qq/weibo/github/wechat]/callback?code=xxxx
func OauthCallBack(context *gin.Context) {
	service := context.Param("service")
	codeStr := context.DefaultQuery("code", "")

	if codeStr == "" {
		response.FailWithMoreMessage(
			"未获取到授权码, 请重试!", "登录失败!", context)
		return
	}

	oauthService, _ := base.OauthService(service)
	platformType := oauthService.GetPlatformType()
	userInfo, err := oauthService.GetUserInfo(codeStr)
	if err != nil {
		response.FailWithMoreMessage(
			err.Error(), "登录失败!", context)
		return
	}

	accountPlatform, _ := shortPlatformService.FindByIdentifyId(userInfo.OpenId)
	if accountPlatform != nil {
		_ = shortPlatformService.UpdateByIdentifyId(
			userInfo.AccessToken,
			userInfo.NickName,
			userInfo.Avatar,
			accountPlatform)
	}
	_, err = shortPlatformService.Create(
		0, // 默认, 等待后续绑定
		platformType,
		userInfo.OpenId,
		userInfo.AccessToken,
		userInfo.NickName,
		userInfo.Avatar)

	if err != nil {
		response.FailWithMessage(err.Error(), context)
		return
	}
	response.Ok(context)
}

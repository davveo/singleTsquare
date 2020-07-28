package controller

import (
	"net/http"

	"github.com/davveo/singleTsquare/models"
	"github.com/davveo/singleTsquare/services"
	"github.com/davveo/singleTsquare/utils/oauth2/base"
	"github.com/davveo/singleTsquare/utils/response"
	"github.com/gin-gonic/gin"
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
	oauthService, err := base.OauthService(service)
	if err != nil {
		response.FailWithMessage(err.Error(), context)
		return
	}
	userInfo, err := oauthService.GetUserInfo(codeStr)
	if err != nil {
		response.FailWithMessage(err.Error(), context)
		return
	}
	accountPlatform, _ := services.AccountPlatformService.FindByIdentifyId(userInfo.OpenId)
	if accountPlatform != nil {
		_ = services.AccountPlatformService.UpdateByIdentifyId(
			userInfo.AccessToken,
			userInfo.NickName,
			userInfo.Avatar,
			accountPlatform)
	}
	accountPlatformUser, err := services.AccountPlatformService.Create(
		0, // 默认, 等待后续绑定
		models.GetPlatformType(service),
		userInfo.OpenId,
		userInfo.AccessToken,
		userInfo.NickName,
		userInfo.Avatar)

	if err != nil {
		response.FailWithMessage(err.Error(), context)
		return
	}

	response.OkWithData(map[string]interface{}{
		"uid":         accountPlatformUser.Uid,
		"avatar":      accountPlatformUser.Avatar,
		"nickname":    accountPlatformUser.NickName,
		"identify_id": accountPlatformUser.IdentifyId,
	}, context)
}

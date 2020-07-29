package controller

import (
	"fmt"
	"time"

	"github.com/davveo/singleTsquare/utils/common"
	"github.com/davveo/singleTsquare/utils/email"
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
	if !BindCheck(&loginRequestJson, context) {
		response.FailWithMessage(response.ParamValidateFailed, context)
		return
	}

	// 如果是重置密码调用, 则需要进行login_id校验
	if loginRequestJson.Mode == "reset" {
		_, err := shortAccountService.FindByLoginId(loginRequestJson.LoginId)
		if err != nil {
			response.FailWithMoreMessage("", response.AccountNotExist, context)
			return
		}
	}

	// 生成验证码
	verifyCode := code.GenerateVerifyCode()
	_ = Cache.Set(
		fmt.Sprintf("verifycode:%s", loginRequestJson.LoginId), str.StrToByte(verifyCode))

	// login_id 可能是邮箱或者手机号
	// TODO 考虑使用策略模式
	// TODO 考虑异步发送
	if common.VerifyEmailFormat(loginRequestJson.LoginId) {
		subject := "验证码邮件"
		bodyMsg := fmt.Sprintf("验证码为: %s", verifyCode)
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

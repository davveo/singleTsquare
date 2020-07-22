package controller

import (
	"fmt"
	"time"

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
	cache, _ = bigcache.NewBigCache(config)
)

func Code(context *gin.Context) {
	var phoneRequrresJson request.PhoneRequestJson

	// 获取手机号
	if err := context.ShouldBindJSON(&phoneRequrresJson); err != nil {
		log.ERROR.Println(err)
	}

	// 生成验证码
	verifyCode := code.GenerateVerifyCode()
	_ = cache.Set(
		fmt.Sprintf("verifycode:%s", phoneRequrresJson.Phone),
		str.StrToByte(verifyCode))

	// 发送至手机 TODO
	sms.Send(phoneRequrresJson.Phone)
	response.Ok(context)
}

func Captcha(context *gin.Context) {
	// captchaType=[audio, string, math, chinese]
	captchaType := context.DefaultQuery("captchaType", "")

	id, b64s, err := captcha.GenerateCaptcha(captchaType)
	if err != nil {
		log.ERROR.Println(err)
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

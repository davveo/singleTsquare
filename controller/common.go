package controller

import (
	"log"

	"github.com/davveo/singleTsquare/utils/captcha"
	"github.com/davveo/singleTsquare/utils/code"
	"github.com/davveo/singleTsquare/utils/response"
	"github.com/gin-gonic/gin"
)

func Code(context *gin.Context) {
	verifyCode := code.GenerateVerifyCode()
	response.OkWithData(map[string]string{
		"verifyCode": verifyCode,
	}, context)
}

func Captcha(context *gin.Context) {
	// captchaType=[audio, string, math, chinese]
	captchaType := context.DefaultQuery("captchaType", "")

	id, b64s, err := captcha.GenerateCaptcha(captchaType)
	if err != nil {
		log.Println(err)
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

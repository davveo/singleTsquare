package controller

import (
	"log"
	"net/http"

	"github.com/davveo/singleTsquare/utils/captcha"
	"github.com/davveo/singleTsquare/utils/code"
	"github.com/gin-gonic/gin"
)

func Code(context *gin.Context) {
	verifyCode := code.GenerateVerifyCode()
	context.JSON(http.StatusOK, gin.H{
		"code": 0,
		"data": map[string]string{
			"verifyCode": verifyCode,
		},
		"msg": "操作成功",
	})
}

func Captcha(context *gin.Context) {
	// captchaType=[audio, string, math, chinese]
	captchaType := context.DefaultQuery("captchaType", "")

	id, b64s, err := captcha.GenerateCaptcha(captchaType)
	if err != nil {
		log.Println(err)
	}

	context.JSON(http.StatusOK, gin.H{
		"code": 0,
		"data": map[string]interface{}{
			"data":      b64s,
			"captchaId": id,
		},
		"errorMsg": "",
		"showMsg":  "success",
	})

}

func HealthCheck(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"code":     0,
		"data":     map[string]interface{}{},
		"errorMsg": "",
		"showMsg":  "success",
	})
}

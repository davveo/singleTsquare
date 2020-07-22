package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	ERROR   = -1
	SUCCESS = 0
)

func Result(code int, data interface{}, errorMsg, showMsg string, context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"code":     code,
		"data":     data,
		"errorMsg": errorMsg,
		"showMsg":  showMsg,
	})
}

func Ok(context *gin.Context) {
	Result(SUCCESS, map[string]interface{}{}, "", "操作成功", context)
}

func OkWithMessage(message string, context *gin.Context) {
	Result(SUCCESS, map[string]interface{}{}, "", message, context)
}

func OkWithData(data interface{}, context *gin.Context) {
	Result(SUCCESS, data, "", "操作成功", context)
}

func OkDetailed(data interface{}, message string, context *gin.Context) {
	Result(SUCCESS, data, "", message, context)
}

func Fail(context *gin.Context) {
	Result(ERROR, map[string]interface{}{}, "", "系统繁忙,请稍后再试!", context)
}

func FailWithMessage(message string, context *gin.Context) {
	Result(ERROR, map[string]interface{}{}, message, "系统繁忙,请稍后再试!", context)
}

func FailWithDetailed(code int, data interface{}, message string, context *gin.Context) {
	Result(code, data, message, "系统繁忙,请稍后再试!", context)
}

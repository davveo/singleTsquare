package controller

import (
	"github.com/gin-gonic/gin"
)

func BindCheck(obj interface{}, context *gin.Context) bool {
	if err := context.ShouldBindJSON(&obj); err != nil {
		return false
	}
	return true
}

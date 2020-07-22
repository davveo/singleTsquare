package controller

import (
	"fmt"
	"net/http"

	"github.com/davveo/singleTsquare/services"

	"github.com/davveo/singleTsquare/models/request"
	"github.com/gin-gonic/gin"
)

var (
	userService services.UserService
)

func UserCreate(context *gin.Context) {
	var userRequst request.UserRequest
	if err := context.ShouldBindJSON(&userRequst); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := userService.FindUserByUsername(userRequst.UserName)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err})
	}
	fmt.Println(user)
}

func UserGet(context *gin.Context) {

}

func UserUpdate(context *gin.Context) {
	// 获取uri参数
	userId := context.Param("userId")

	// 获取query参数
	name := context.DefaultQuery("name", "")

	fmt.Println(userId, name)
}

func UserDelete(context *gin.Context) {

}

func UserList(context *gin.Context) {

}

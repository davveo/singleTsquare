package controller

import (
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

}

func UserGet(context *gin.Context) {

}

func UserUpdate(context *gin.Context) {

}

func UserDelete(context *gin.Context) {

}

func UserList(context *gin.Context) {

}

package router

import (
	"github.com/gin-gonic/gin"
)

func TopicRouter(router *gin.RouterGroup) {
	topicRouter := router.Group("topic")
	{
		topicRouter.GET("create", func(context *gin.Context) {

		})
	}

}

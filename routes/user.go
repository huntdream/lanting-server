package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/huntdream/lanting-server/service"
)

// register routes of article
func registerUser(router *gin.RouterGroup) {
	user := router.Group("/user")
	{
		user.POST("", service.CreateUser)
	}
}

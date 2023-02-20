package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/huntdream/lanting-server/service"
)

// registerAuth routes for authentication
func registerAuth(router *gin.RouterGroup) {
	auth := router.Group("/auth")
	{
		auth.POST("/login", service.Login)
		auth.POST("/signup", service.SignUp)
	}
}

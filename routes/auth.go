package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/huntdream/lanting-server/service"
)

//registerAuth routes for authentication
func registerAuth(router *gin.RouterGroup) {
	auth := router.Group("/auth")
	{
		auth.POST("/signin", service.SignIn)
		auth.POST("/signup", service.SignUp)
	}
}

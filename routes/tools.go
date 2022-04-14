package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/huntdream/lanting-server/service"
)

//registerAuth routes for authentication
func registerTools(router *gin.RouterGroup) {
	tools := router.Group("/tools")
	{
		tools.POST("/uploadToken", service.GetUploadToken)
	}
}

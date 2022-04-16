package routes

import (
	"github.com/gin-gonic/gin"
)

//Register routes
func Register(router *gin.Engine) {

	v1 := router.Group("/api/v1")
	{
		registerArticle(v1)
		registerUser(v1)
		registerAuth(v1)
		registerTools(v1)
	}
}

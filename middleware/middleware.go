package middleware

import "github.com/gin-gonic/gin"

//EnableMiddleware initialize middlewares
func EnableMiddleware(router *gin.Engine) {
	router.Use(CORS())
}

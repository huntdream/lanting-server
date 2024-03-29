package middleware

import "github.com/gin-gonic/gin"

//EnableMiddleware initialize middlewares
func EnableMiddleware(router *gin.Engine) {
	router.Use(CORS())
	router.Use(User())
	router.Use(Log())
}

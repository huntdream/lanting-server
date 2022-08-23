package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/huntdream/lanting-server/service"
)

func Log() gin.HandlerFunc {
	return service.AddLog
}

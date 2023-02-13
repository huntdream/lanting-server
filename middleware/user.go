package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/huntdream/lanting-server/util"
	"strings"
)

func User() gin.HandlerFunc {
	return func(c *gin.Context) {
		authorization := c.GetHeader("Authorization")

		if authorization != "" {
			token := strings.TrimPrefix(authorization, "Bearer ")

			userId, _, _ := util.ParseToken(token)
			fmt.Println("userId", userId)
			c.Set("userId", userId)

			c.Next()
		}
	}
}

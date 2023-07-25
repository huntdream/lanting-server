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

			userId, _, err := util.ParseToken(token)

			if err != nil {
				c.Next()
			}

			fmt.Println("userId", userId)
			c.Set("userId", userId)

			c.Next()
		}
	}
}

package middleware

import (
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/huntdream/lanting-server/util"
)

//JWT middleware
func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {

		if strings.HasPrefix(c.FullPath(), "/api/v1/auth") {
			c.Next()

			return
		}

		authorization := c.GetHeader("Authorization")

		//check if Authorization header provided
		if authorization == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Unauthorized",
			})

			c.Abort()

			return
		}

		token := strings.TrimPrefix(authorization, "Bearer ")

		//check if token provided
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Unauthorized",
			})

			c.Abort()

			return
		}

		//parse token
		username, err := util.ParseToken(token)

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": err.Error(),
			})

			c.Abort()

			return
		}

		log.Println(token, username)

		c.Next()
	}

}

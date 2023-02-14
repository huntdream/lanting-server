package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/huntdream/lanting-server/app"
	"strings"
	"time"
)

type Log struct {
	Path   string `json:"path"`
	UA     string `json:"ua"`
	Method string `json:"method"`
	Time   string `json:"time"`
	UserId int    `json:"userId"`
}

//AddLog add log
func AddLog(c *gin.Context) {
	remoteAddr := c.ClientIP()
	userAgent := c.GetHeader("User-Agent")
	method := c.Request.Method
	path := c.Request.URL.Path
	userId := c.GetInt64("userId")

	if strings.HasPrefix(path, "/api/v1") {
		_, err := app.DB.Exec("insert into logs (user_id,ip, ua, method, time, path) values (?, ?, ? ,? ,? ,?)", userId, remoteAddr, userAgent, method, time.Now(), path)
		if err != nil {
			fmt.Println(err.Error())
		}
	}

	c.Next()
}

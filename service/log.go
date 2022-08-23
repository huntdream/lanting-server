package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/huntdream/lanting-server/app"
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

	_, err := app.DB.Exec("insert into logs (ip, ua, method, time, path) values (?, ? ,? ,? ,?)", remoteAddr, userAgent, method, time.Now(), path)
	if err != nil {
		fmt.Println(err.Error())
	}

	c.Next()
}

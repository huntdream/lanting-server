package main

import (
	"github.com/gin-gonic/gin"
	"github.com/huntdream/lanting-server/app"
	"github.com/huntdream/lanting-server/config"
	"github.com/huntdream/lanting-server/db"
	"github.com/huntdream/lanting-server/middleware"
	"github.com/huntdream/lanting-server/routes"
)

func main() {
	router := gin.Default()
	router.RemoveExtraSlash = true
	app.Config = config.ReadConfiguration()

	db.Initialize()

	middleware.EnableMiddleware(router)

	routes.Register(router)

	err := router.Run(":" + app.Config.Server.Port)

	if err != nil {
		return
	}
}

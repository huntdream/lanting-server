package main

import (
	"github.com/gin-gonic/gin"
	"github.com/huntdream/lanting-server/db"
	"github.com/huntdream/lanting-server/middleware"
	"github.com/huntdream/lanting-server/routes"
)

func main() {
	router := gin.Default()

	middleware.EnableMiddleware(router)

	routes.Register(router)

	db.Initialize()

	router.Run(":4000")
}

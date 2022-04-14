package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/huntdream/lanting-server/api"
	"github.com/huntdream/lanting-server/middleware"
)

// register routes of article
func registerArticle(router *gin.RouterGroup) {
	article := router.Group("/article")
	{
		article.GET("", api.GetArticles)
		article.GET("/:id", api.GetArticle)
		article.POST("", middleware.JWT(), api.AddArticle)
		article.POST("/", middleware.JWT(), api.AddArticle)
	}
}

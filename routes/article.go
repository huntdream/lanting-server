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
		article.GET("/:id", api.GetArticle)
		article.POST("/", middleware.JWT(), api.AddArticle)
		article.POST("/:id", middleware.JWT(), api.UpdateArticle)
	}

	articles := router.Group("/articles")
	{
		articles.GET("/", api.GetArticles)
		articles.GET("/:id", api.GetArticlesByUserId)
	}
}

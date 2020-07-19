package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/huntdream/lanting-server/service"
)

// register routes of article
func registerArticle(router *gin.RouterGroup) {
	article := router.Group("/article")
	{
		article.GET("", service.GetArticles)
		article.GET("/:id", service.GetArticle)
	}
}

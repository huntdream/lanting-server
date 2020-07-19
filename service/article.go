package service

import (
	"github.com/gin-gonic/gin"
	"github.com/huntdream/lanting-server/app"
	"github.com/huntdream/lanting-server/model"
)

//ArticlesRequest params
type ArticlesRequest struct {
	after int
	size  int
}

//GetArticles get articles
func GetArticles(c *gin.Context) {

	var articles []model.Article
	var total int
	var count int

	size := c.DefaultQuery("size", "10")
	after := c.DefaultQuery("after", "0")

	app.DB.Table("articles").Where("id > ?", after).Limit(size).Find(&articles).Count(&total)

	count = len(articles)

	c.JSON(200, gin.H{
		"data":  articles,
		"total": total,
		"count": count,
	})

	return
}

//GetArticle get article
func GetArticle(c *gin.Context) {
	id := c.Param("id")

	c.JSON(200, gin.H{
		"message": id,
	})

	return
}

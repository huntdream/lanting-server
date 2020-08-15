package api

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/huntdream/lanting-server/model"
	"github.com/huntdream/lanting-server/service"
)

//GetArticle get article by id
func GetArticle(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

		return
	}

	article, err := service.GetArticleByID(id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "article not found",
		})
	}

	c.JSON(http.StatusOK, article)
}

//GetArticles get articles
func GetArticles(c *gin.Context) {

	size := c.DefaultQuery("size", "10")
	after := c.DefaultQuery("after", "0")

	articles, total, count := service.GetArticles(size, after)

	c.JSON(http.StatusOK, gin.H{
		"data":  articles,
		"total": total,
		"count": count,
	})

	return
}

//AddArticle add article
func AddArticle(c *gin.Context) {
	var article model.Article

	if err := c.ShouldBind(&article); err != nil {
		log.Println(err)
	}

	savedArticle := service.AddArticle(article)

	c.JSON(http.StatusOK, savedArticle)
}

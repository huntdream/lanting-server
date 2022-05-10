package api

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/huntdream/lanting-server/model"
	"github.com/huntdream/lanting-server/service"
)

//GetArticle get article by id
func GetArticle(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

		return
	}

	article, err := service.GetArticleByID(id)
	user := service.GetCurrentUser(c)

	article.CanEdit = article.AuthorId == user.ID

	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusNotFound, gin.H{
			"message": "article not found",
		})

		return
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

	author := service.GetCurrentUser(c)
	fmt.Println(author, "???")
	article.AuthorId = author.ID

	savedArticle, err := service.AddArticle(article)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, savedArticle)
}

//UpdateArticle update article
func UpdateArticle(c *gin.Context) {
	var article model.Article

	if err := c.ShouldBind(&article); err != nil {
		log.Println(err)

		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

		return
	}

	savedArticle, err := service.UpdateArticle(c, article)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, savedArticle)
}

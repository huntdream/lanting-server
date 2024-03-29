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

// GetArticle get article by id
func GetArticle(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	userId := c.GetInt64("userId")

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

		return
	}

	article, err := service.GetArticleByID(id)

	if article.AuthorId != userId && article.Visibility != 1 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "article not found",
		})

		return
	}

	user, _ := service.FindUserById(userId)

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

// GetArticlesByUserId get my articles
func GetArticlesByUserId(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

		return
	}
	userId := c.GetInt64("userId")

	articles, total := service.GetArticlesByUserID(id, userId)

	c.JSON(http.StatusOK, gin.H{
		"data":  articles,
		"total": total,
	})

	return
}

// GetArticles get articles
func GetArticles(c *gin.Context) {
	userId := c.GetInt64("userId")
	isMe := c.Request.URL.Path

	fmt.Println(isMe)

	size := c.DefaultQuery("size", "10")
	after := c.DefaultQuery("after", "0")

	articles, total := service.GetArticles(userId, size, after)

	c.JSON(http.StatusOK, gin.H{
		"data":  articles,
		"total": total,
	})

	return
}

// AddArticle add article
func AddArticle(c *gin.Context) {
	var article model.Article
	authorId := c.GetInt64("userId")

	if err := c.ShouldBind(&article); err != nil {
		log.Println(err)
	}

	article.AuthorId = authorId

	savedArticle, err := service.AddArticle(article)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, savedArticle)
}

// UpdateArticle update article
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

func DeleteArticles(c *gin.Context) {
	var ids model.Ids

	if err := c.ShouldBind(&ids); err != nil {
		log.Println(err)

		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

		return
	}

	if len(ids.Ids) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ids is empty",
		})

		return
	}

	err := service.DeleteArticles(c, ids.Ids)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
	})

	return
}

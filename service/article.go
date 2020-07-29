package service

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/huntdream/lanting-server/app"
	"github.com/huntdream/lanting-server/model"
	"github.com/huntdream/lanting-server/util"
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

//GetArticleByID get article by id
func GetArticleByID(id int) (article model.Article, err error) {

	if err := app.DB.Table("articles").Where("id = ?", id).Find(&article).Error; err != nil {
		return article, err
	}

	return article, nil
}

//AddArticle add article
func AddArticle(c *gin.Context) {
	var article model.Article

	if err := c.ShouldBind(&article); err != nil {
		log.Println(err)
	}

	article.Content = util.Sanitize(article.Content)
	excerpt := []rune(util.ExtractText(article.Content))

	if len(excerpt) > 40 {
		excerpt = excerpt[:40]
	}

	article.Excerpt = string(excerpt)

	record := app.DB.Table("articles").Create(&article)

	c.JSON(http.StatusOK, record.Value)

	return
}

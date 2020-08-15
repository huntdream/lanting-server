package service

import (
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
func GetArticles(size string, after string) (articles []model.Article, total int, count int) {

	app.DB.Table("articles").Where("id < ?", after).Order("created_at desc").Limit(size).Find(&articles).Count(&total)

	count = len(articles)

	return articles, total, count
}

//GetArticleByID get article by id
func GetArticleByID(id int) (article model.Article, err error) {

	if err := app.DB.Table("articles").Where("id = ?", id).Find(&article).Error; err != nil {
		return article, err
	}

	return article, nil
}

//AddArticle add article
func AddArticle(article model.Article) (value interface{}) {

	article.Content = util.Sanitize(article.Content)
	excerpt := []rune(util.ExtractText(article.Content))

	if len(excerpt) > 40 {
		excerpt = excerpt[:40]
	}

	article.Excerpt = string(excerpt)

	record := app.DB.Table("articles").Create(&article)

	return record.Value
}

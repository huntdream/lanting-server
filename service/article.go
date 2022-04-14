package service

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/huntdream/lanting-server/app"
	"github.com/huntdream/lanting-server/model"
	"github.com/huntdream/lanting-server/util"
)

//GetArticles get articles
func GetArticles(size string, after string) (feed []model.FeedItem, total int, count int) {
	rows, err := app.DB.Query("select id, title, excerpt from articles order by id desc")

	if err != nil {
		return feed, 0, 0
	}

	defer rows.Close()

	for rows.Next() {
		var article model.FeedItem

		if err = rows.Scan(&article.ID, &article.Title, &article.Excerpt); err != nil {
			fmt.Println(article.Title, err.Error())

			return feed, 0, 0
		}

		feed = append(feed, article)
	}

	if err := rows.Err(); err != nil {
		return feed, 0, 0
	}

	count = len(feed)

	return feed, total, count
}

//GetArticleByID get article by id
func GetArticleByID(id int) (article model.Article, err error) {
	row := app.DB.QueryRow("select id, title, content from articles where id = ?", id)

	if err := row.Scan(&article.ID, &article.Title, &article.Content); err != nil {
		if err == sql.ErrNoRows {
			return article, errors.New("article not found")
		}

		return article, fmt.Errorf("articleById %d: %v", id, err)
	}

	return article, nil
}

//AddArticle add article
func AddArticle(article model.Article) (value interface{}, err error) {
	if article.Title == "" {
		return nil, errors.New("title is required")
	}

	// article.Content = util.Sanitize(article.Content)
	excerpt := []rune(util.ExtractText(article.Content))

	if len(excerpt) > 40 {
		excerpt = excerpt[:40]
	}

	article.Excerpt = string(excerpt)

	result, err := app.DB.Exec("insert into articles (title, content,excerpt) values (?, ?, ?)", article.Title, article.Content, article.Excerpt)

	if err != nil {
		return 0, fmt.Errorf("addArticle: %v", err)
	}
	id, err := result.LastInsertId()

	if err != nil {
		return 0, fmt.Errorf("addArticle: %v", err)
	}
	return id, nil

}

package service

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/huntdream/lanting-server/app"
	"github.com/huntdream/lanting-server/model"
)

func scanArticles(rows *sql.Rows) (articles []model.Article) {
	articles = []model.Article{}

	for rows.Next() {
		var article model.Article

		if err := rows.Scan(&article.ID, &article.Title, &article.Excerpt, &article.Visibility, &article.AuthorId, &article.CreatedAt); err != nil {
			fmt.Println(article.Title, err.Error())

			return articles
		}

		author, err := FindUserById(article.AuthorId)

		if err != nil {
			fmt.Println(err)
		}

		article.Author = author

		articles = append(articles, article)
	}

	return articles
}

// GetArticlesByUserID Get Articles By User ID
func GetArticlesByUserID(id int64, userId int64) (feed []model.Article, total int) {
	visibility := 0
	if id == userId {
		visibility = 1
	}

	rows, err := app.DB.Query("select id, title, excerpt, visibility,author_id, created_at from articles where visibility=? and author_id=? order by id desc", visibility, userId)

	if err != nil {
		return feed, 0
	}

	defer rows.Close()

	feed = scanArticles(rows)

	if err := rows.Err(); err != nil {
		fmt.Println(err.Error())

		return feed, 0
	}

	total = len(feed)

	return feed, total
}

// GetArticles get articles
func GetArticles(userId int64, size string, after string) (feed []model.Article, total int) {
	rows, err := app.DB.Query("select id, title, excerpt, visibility,author_id, created_at from articles where visibility=1 or author_id=? order by id desc limit ?", userId, size)

	if err != nil {
		return feed, 0
	}

	defer rows.Close()

	feed = scanArticles(rows)

	if err := rows.Err(); err != nil {
		fmt.Println(err.Error())

		return feed, 0
	}

	total = len(feed)

	return feed, total
}

// GetArticleByID get article by id
func GetArticleByID(id int64) (article model.Article, err error) {
	row := app.DB.QueryRow("select id, title, author_id, excerpt, content, visibility, created_at, updated_at from articles where id = ?", id)

	if err := row.Scan(&article.ID, &article.Title, &article.AuthorId, &article.Excerpt, &article.Content, &article.Visibility, &article.CreatedAt, &article.UpdatedAt); err != nil {
		if err == sql.ErrNoRows {
			return article, errors.New("article not found")
		}

		return article, fmt.Errorf("articleById %d: %v", id, err)
	}

	user, err := FindUserById(article.AuthorId)

	if err != nil {
		return article, nil
	}

	article.Author = user

	return article, nil
}

// AddArticle add article
func AddArticle(article model.Article) (value interface{}, err error) {
	if article.Title == "" {
		return nil, errors.New("title is required")
	}

	result, err := app.DB.Exec("insert into articles (title, content, text, author_id, excerpt, visibility) values (?,?,?,?,?,?)", article.Title, article.Content, article.Text, article.AuthorId, article.Excerpt, article.Visibility)

	if err != nil {
		return 0, fmt.Errorf("addArticle: %v", err)
	}

	id, err := result.LastInsertId()

	if err != nil {
		return 0, fmt.Errorf("addArticle: %v", err)
	}

	article, err = GetArticleByID(id)

	return article, nil
}

// UpdateArticle update article
func UpdateArticle(c *gin.Context, newArticle model.Article) (value interface{}, err error) {
	useId := c.GetInt64("userId")

	if newArticle.ID == 0 {
		return nil, errors.New("id is required")
	}

	if newArticle.Title == "" {
		return nil, errors.New("title is required")
	}

	article, err := GetArticleByID(newArticle.ID)

	if err != nil {
		return nil, err
	}

	if useId != article.AuthorId {
		return nil, errors.New("permission denied")
	}

	article.Title = newArticle.Title
	article.Content = newArticle.Content
	article.Excerpt = newArticle.Excerpt
	article.Visibility = newArticle.Visibility
	article.Text = newArticle.Text

	_, err = app.DB.Exec("update articles set title = ?, content = ?,text = ?, excerpt = ?, visibility = ? where id = ?", article.Title, article.Content, article.Text, article.Excerpt, article.Visibility, article.ID)

	if err != nil {
		return 0, fmt.Errorf("updateArticle: %v", err)
	}

	article, err = GetArticleByID(article.ID)

	return article, nil
}

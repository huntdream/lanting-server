package service

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/huntdream/lanting-server/app"
	"github.com/huntdream/lanting-server/model"
)

//GetArticles get articles
func GetArticles(userId int64, size string, after string) (feed []model.Article, total int, count int) {
	rows, err := app.DB.Query("select id, title, excerpt, visibility, created_at from articles where visibility=1 or author_id=? order by id desc", userId)

	if err != nil {
		return feed, 0, 0
	}

	defer rows.Close()

	for rows.Next() {
		var article model.Article

		if err = rows.Scan(&article.ID, &article.Title, &article.Excerpt, &article.Visibility, &article.CreatedAt); err != nil {
			fmt.Println(article.Title, err.Error())

			return feed, 0, 0
		}

		feed = append(feed, article)
	}

	if err := rows.Err(); err != nil {
		fmt.Println(err.Error())

		return feed, 0, 0
	}

	count = len(feed)

	return feed, total, count
}

//GetArticleByID get article by id
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

	article.Author.ID = user.ID
	article.Author.Username = user.Username

	return article, nil
}

//AddArticle add article
func AddArticle(article model.Article) (value interface{}, err error) {
	if article.Title == "" {
		return nil, errors.New("title is required")
	}

	result, err := app.DB.Exec("insert into articles (title, content, author_id, excerpt, visibility) values (?,?,?,?,?)", article.Title, article.Content, article.AuthorId, article.Excerpt, article.Visibility)

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

//UpdateArticle update article
func UpdateArticle(c *gin.Context, newArticle model.Article) (value interface{}, err error) {
	user := GetCurrentUser(c)

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

	if user.ID != article.AuthorId {
		return nil, errors.New("permission denied")
	}

	article.Title = newArticle.Title
	article.Content = newArticle.Content
	article.Excerpt = newArticle.Excerpt
	article.Visibility = newArticle.Visibility

	_, err = app.DB.Exec("update articles set title = ?, content = ?, excerpt = ?, visibility = ? where id = ?", article.Title, article.Content, article.Excerpt, article.Visibility, article.ID)

	if err != nil {
		return 0, fmt.Errorf("updateArticle: %v", err)
	}

	article, err = GetArticleByID(article.ID)

	return article, nil
}

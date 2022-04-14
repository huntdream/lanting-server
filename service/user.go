package service

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/huntdream/lanting-server/app"
	"github.com/huntdream/lanting-server/model"
)

// CreateUser create user
func CreateUser(c *gin.Context) {
	var user model.User

	if err := (c.ShouldBindJSON(&user)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid params",
		})

		return
	}

	app.DB.Exec("insert into users (username, password) values (?,? )", user.Username, user.Password)

	fmt.Println(user)

	c.JSON(200, user)

	return
}

//FindUser find user by username
func FindUser(username string) (user model.User, err error) {
	row := app.DB.QueryRow("select username, password from users where username = ?", username)

	if err := row.Scan(&user.Username, &user.Password); err != nil {
		if err == sql.ErrNoRows {
			return user, fmt.Errorf("user not found")
		}

		return user, fmt.Errorf("FindUser %v", err)
	}

	return user, nil
}

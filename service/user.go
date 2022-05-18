package service

import (
	"database/sql"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/huntdream/lanting-server/app"
	"github.com/huntdream/lanting-server/model"
	"github.com/huntdream/lanting-server/util"
)

// CreateUser create user
func CreateUser(c *gin.Context) {
	var user model.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

		return
	}

	_, err := app.DB.Exec("insert into users (username, password) values (?,? )", user.Username, user.Password)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})

		return
	}

	fmt.Println(user)

	c.JSON(http.StatusOK, user)

	return
}

//FindUserByUsername find user by username
func FindUserByUsername(username string) (user model.User, err error) {
	row := app.DB.QueryRow("select id, username, password from users where username = ?", username)

	if err := row.Scan(&user.ID, &user.Username, &user.Password); err != nil {
		if err == sql.ErrNoRows {
			return user, fmt.Errorf("user not found")
		}

		return user, fmt.Errorf("FindUser %v", err)
	}

	return user, nil
}

//FindUserById find user by id
func FindUserById(id int) (user model.User, err error) {
	row := app.DB.QueryRow("select id, username from users where id = ?", id)

	if err := row.Scan(&user.ID, &user.Username); err != nil {
		if err == sql.ErrNoRows {
			return user, fmt.Errorf("user not found")
		}

		return user, fmt.Errorf("FindUser %v", err)
	}

	return user, nil
}

//GetCurrentUser get current user
func GetCurrentUser(c *gin.Context) (user model.User) {
	user = model.User{}
	authorization := c.GetHeader("Authorization")

	//check if Authorization header provided
	if authorization == "" {
		return user
	}

	token := strings.TrimPrefix(authorization, "Bearer ")

	//check if token provided
	if token == "" {
		return user
	}

	//parse token
	username, err := util.ParseToken(token)

	if err != nil {
		return user
	}

	user, err = FindUserByUsername(username)

	if err != nil {
		return user
	}

	return user
}

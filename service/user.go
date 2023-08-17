package service

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/huntdream/lanting-server/app"
	"github.com/huntdream/lanting-server/model"
	"net/http"
	"strconv"
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

// FindUserByUsername find user by username
func FindUserByUsername(username string) (user model.User, err error) {
	row := app.DB.QueryRow("select id, username, password, avatar from users where username = ?", username)

	if err := row.Scan(&user.ID, &user.Username, &user.Password, &user.Avatar); err != nil {
		if err == sql.ErrNoRows {
			return user, fmt.Errorf("user not found")
		}

		return user, fmt.Errorf("FindUser %v", err)
	}

	return user, nil
}

// FindUserById find user by id
func FindUserById(id int64) (user model.User, err error) {
	row := app.DB.QueryRow("select id, username, avatar, name, bio, email from users where id = ?", id)

	if err := row.Scan(&user.ID, &user.Username, &user.Avatar, &user.Name, &user.Bio, &user.Email); err != nil {
		if err == sql.ErrNoRows {
			return user, fmt.Errorf("user not found")
		}

		return user, fmt.Errorf("FindUser %v", err)
	}

	return user, nil
}

// GetCurrentUser get current user
func GetCurrentUser(c *gin.Context) {
	userId := c.GetInt64("userId")

	user, err := FindUserById(userId)

	if err != nil {
		c.JSON(http.StatusOK, user)
		return
	}

	c.JSON(http.StatusOK, user)

	return
}

// GetUserById Get user by ID
func GetUserById(c *gin.Context) {
	var user = model.User{}

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

		return
	}

	user, err = FindUserById(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

		return
	}

	if user.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "user not found",
		})
	}

	c.JSON(http.StatusOK, user)
	return
}

// UpdateUser Update user info
func UpdateUser(c *gin.Context) {
	var user model.UserInfo

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

		return
	}
	userId := c.GetInt64("userId")

	if userId != user.ID {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Are you trying to update other user's profile ?",
		})

		return
	}

	_, err := app.DB.Exec("update users set avatar=?, name=? where id=?", user.Avatar, user.Name, user.ID)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

		return
	}

	updatedUser, err := FindUserById(user.ID)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, updatedUser)

	return
}

package service

import (
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

	app.DB.Table("users").Create(&user)

	fmt.Println(user)

	c.JSON(200, user)

	return
}

//FindUser find user by username
func FindUser(username string) (user model.User, err error) {
	if err = app.DB.Table("users").Where("username = ?", username).Find(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

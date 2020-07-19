package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/huntdream/lanting-server/app"
	"github.com/huntdream/lanting-server/model"
	"github.com/huntdream/lanting-server/util"
	"golang.org/x/crypto/bcrypt"
)

func hashPassword(password string) string {
	hashed, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)

	return string(hashed)
}

//SignUp sign up
func SignUp(c *gin.Context) {
	var userInfo model.AuthRequest

	if err := c.ShouldBind(&userInfo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
	}

	_, err := FindUser(userInfo.Username)

	if err == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "username already exists",
		})

		return
	}

	userInfo.Password = hashPassword(userInfo.Password)

	if err := app.DB.Table("users").Create(&userInfo).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

		return
	}

	user, err := FindUser(userInfo.Username)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
	}

	token, err := util.GenerateToken(userInfo.Username)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
	}

	c.JSON(http.StatusOK, model.AuthResponse{
		ID:       user.ID,
		Username: user.Username,
		Name:     user.Name,
		Bio:      user.Bio,
		Email:    user.Email,
		Token:    token,
	})

	return
}

//SignIn sign in
func SignIn(c *gin.Context) {
	var userInfo model.AuthRequest

	//get user provided info
	if err := c.ShouldBind(&userInfo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
	}

	//get user by username from database
	user, err := FindUser(userInfo.Username)

	//check if user exists in database
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "User not found",
		})

		return
	}

	//check if password provided match the database record
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userInfo.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "username or password not match",
		})

		return
	}

	token, err := util.GenerateToken(userInfo.Username)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
	}

	c.JSON(http.StatusOK, model.AuthResponse{
		ID:       user.ID,
		Username: user.Username,
		Name:     user.Name,
		Bio:      user.Bio,
		Email:    user.Email,
		Token:    token,
	})

	return
}

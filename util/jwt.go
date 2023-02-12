package util

import (
	"github.com/huntdream/lanting-server/model"
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("my_secret_key")

//Claims JWT claims
type Claims struct {
	Username string `json:"username"`
	ID       int64  `json:"id"`
	jwt.StandardClaims
}

//GenerateToken generate jwt token
func GenerateToken(user model.User) (tokenString string, err error) {
	//the token expiration time
	expirationTime := time.Now().Add(500 * time.Hour)

	claims := Claims{
		Username: user.Username,
		ID:       user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err = token.SignedString(jwtKey)

	if err != nil {
		return tokenString, err
	}

	return tokenString, nil
}

//ParseToken parse jwt token
func ParseToken(tokenString string) (userId int64, username string, err error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	claims := token.Claims.(*Claims)

	username = claims.Username
	userId = claims.ID

	log.Println(username, claims.StandardClaims.IssuedAt, claims.StandardClaims.ExpiresAt)

	return userId, username, err
}

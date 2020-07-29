package util

import (
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("my_secret_key")

//Claims JWT claims
type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

//GenerateToken generate jwt token
func GenerateToken(username string) (tokenString string, err error) {
	//the token expiration time
	expirationTime := time.Now().Add(5 * time.Hour)

	claims := Claims{
		Username: username,
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
func ParseToken(tokenString string) (username string, err error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	claims := token.Claims.(*Claims)

	username = claims.Username

	log.Println(username, claims.StandardClaims.IssuedAt, claims.StandardClaims.ExpiresAt)

	if err != nil {
		return username, err
	}

	return username, nil
}

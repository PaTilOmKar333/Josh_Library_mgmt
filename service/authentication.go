package service

import (
	"errors"
	"log"
	"project/config"
	"project/models"
	"time"

	"github.com/golang-jwt/jwt"
)

type AuthTokenInterface interface {
	GenerateToken(uID int, email, role string) (string, error)
	ValidateToken(tokenString string, authorizedRole []string) (claimss *models.Claims, err error)
}

type JwtToken struct {
}

func InitAuthService() AuthTokenInterface {

	return &JwtToken{}
}

func (j *JwtToken) GenerateToken(uID int, email, role string) (string, error) {
	var claims models.Claims

	var mySigningKey = []byte(config.SecretKey())

	claims.UserID = uID
	claims.Email = email
	claims.Role = role
	claims.StandardClaims = jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Minute * 90).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(mySigningKey)
	if err != nil {
		log.Println(err)
		err = errors.New("something went wrong")
		return "", err
	}
	return tokenString, nil
}

func (j *JwtToken) ValidateToken(tokenString string, authorizedRole []string) (claimss *models.Claims, err error) {

	var mySigningKey = []byte(config.SecretKey())

	token, err := jwt.ParseWithClaims(tokenString, &models.Claims{}, func(token *jwt.Token) (interface{}, error) {
		return mySigningKey, nil
	})
	if err != nil {
		log.Println(err)
		return
	}

	claimss, ok := token.Claims.(*models.Claims)
	if !ok {
		err = errors.New("error in genrating claimss")
		return
	}

	return
}

package jwt

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
	"time"
)

var signingKey = []byte(viper.GetString("jwt.signingKey"))

type CustomClaims struct {
	ID   int
	Name string
	jwt.RegisteredClaims
}

func GenerateToken(id int, name string) (string, error) {

	i := viper.GetInt64("jwt.tokenExpire")
	d := time.Duration(i) * time.Hour
	claims := CustomClaims{ID: id, Name: name, RegisteredClaims: jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(d)),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		Subject:   "Token",
	}}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(signingKey)
}

func ParseToken(token string) (CustomClaims, error) {
	var customClaims CustomClaims
	claims, err := jwt.ParseWithClaims(token, &customClaims, func(token *jwt.Token) (interface{}, error) {
		return signingKey, nil
	})
	if err != nil {
		err := errors.New("parse Token error: " + err.Error())
		return CustomClaims{}, err
	}
	if !claims.Valid {
		err := errors.New("invalid Token")
		return CustomClaims{}, err
	}
	return customClaims, nil
}

func IsTokenValid(token string) bool {
	_, err := ParseToken(token)
	if err != nil {
		return false
	}
	return true
}

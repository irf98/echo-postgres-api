package config

import (
	"awesomeProject/api/models"
	"github.com/golang-jwt/jwt"
	"time"
)

type Claim struct {
	models.User `json:"user"`
	jwt.StandardClaims
}

type ResponseToken struct {
	Token string `json:"token"`
}

func GenerateJWT(user models.User) string {
	claims := Claim{
		User: user,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 1).Unix(),
			Issuer:    "Golang Awesome API",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)

	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return err.Error()
	}

	return t
}

package utils

import (
	"example/app/models"
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)
var mySigningKey = []byte(os.Getenv("SECRET_KEY"))

func GenerateJWT(userData models.Users) (string, error) {
    token := jwt.New(jwt.SigningMethodHS256)
    claims := token.Claims.(jwt.MapClaims)

    claims["authorized"] = true
    claims["client"] = userData.Username
    claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

    tokenString, err := token.SignedString(mySigningKey)

    if err != nil {
        return "", err
    }

    return tokenString, nil
}

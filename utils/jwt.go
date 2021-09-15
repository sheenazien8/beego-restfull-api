package utils

import (
  "beego/models"
  jwt "github.com/dgrijalva/jwt-go"
  "time"
  "github.com/beego/beego/v2/server/web"
)

type Jwt struct {
  Token string
}

var jwtkey, _ = web.AppConfig.String("jwtkey")
var mySigningKey = []byte(jwtkey)

func (this *Jwt) GenerateJWT(userData models.Users) (string, error) {
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



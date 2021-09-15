package controllers

import (
  "beego/models"
  "beego/utils"
  "encoding/json"
  "net/http"
  "github.com/beego/beego/v2/client/orm"
  "github.com/beego/beego/v2/server/web"
  _ "github.com/go-sql-driver/mysql"
  "golang.org/x/crypto/bcrypt"
)

// Operations about Users
type AuthController struct {
  web.Controller
}

type LoginRequest struct {
  Username string `json:"username"`
  Password string `json:"password"`
}

type LoginSuccessResponse struct {
  Status bool `json:"status"`
  Code int `json:"code"`
  Token string `json:"token"`
}

var jwtkey, _ = web.AppConfig.String("jwtkey")
var mySigningKey = []byte(jwtkey)

// @Title Login
// @Description auth login
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} ""
// @Failure 403 body is empty
// @router /login [post]
func (this *AuthController) Login() {
  this.Ctx.ResponseWriter.Header().Set("Content-Type", "application/json; charset=UTF-8")
  var request LoginRequest
  json.Unmarshal(this.Ctx.Input.RequestBody, &request)
  o := orm.NewOrm()
  user := models.Users{Username: request.Username}
  err := o.Read(&user, "username")
  if err == orm.ErrNoRows {
    this.Ctx.ResponseWriter.WriteHeader(http.StatusNotFound)
    this.Data["json"] = utils.GlobalResponse{
      Messages: "Credentials anda tidak ditemukan",
      Status: false,
      Code: http.StatusNotFound,
    }
    this.ServeJSON()
    return
  }
  password := []byte(request.Password)
  hashedPassword := []byte(user.Password)
  isNotMatched := bcrypt.CompareHashAndPassword(hashedPassword, password)
  if isNotMatched != nil {
    this.Ctx.ResponseWriter.WriteHeader(http.StatusUnauthorized)
    this.Data["json"] = utils.GlobalResponse{
      Messages: "Password anda Salah",
      Status: false,
      Code: http.StatusUnauthorized,
    }
    this.ServeJSON()
    return
  }
  // generate jwt token
  var jwt utils.Jwt
  token, err := jwt.GenerateJWT(user)
  if err != nil {
    panic(err.Error())
  }
  this.Data["json"] = LoginSuccessResponse{
    Status: true,
    Code: http.StatusUnauthorized,
    Token: token,
  }
  this.ServeJSON()
  return
}

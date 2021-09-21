package controllers

import (
	"encoding/json"
	"example/app/models"
	"example/bootstrap"
	"example/utils"

	"golang.org/x/crypto/bcrypt"
)

type AuthController struct {
	BaseController
}

// @Title Login
// @Description Example login with return jwt token
// @Success 200 {object}
// @router /login [post]
func (this *AuthController) Login() {
	var requestBody models.Users
	json.Unmarshal(this.Ctx.Input.RequestBody, &requestBody)
	db := bootstrap.DB
	findedUser := db.Find(&user, "username = ?", requestBody.Username)
	if findedUser.RowsAffected == 0 {
		println("Not Found")
		this.Ctx.ResponseWriter.WriteHeader(404)
		response.Code = "404"
		response.Status = 1
		response.Message = "Sorry, username not found in were records."
		this.Data["json"] = response
		this.ServeJSON()
		return
	}
	password := []byte(requestBody.Password)
	hashedPassword := []byte(user.Password)
	isNotMatched := bcrypt.CompareHashAndPassword(hashedPassword, password)
	if isNotMatched != nil {
		println(isNotMatched.Error())
		this.Ctx.ResponseWriter.WriteHeader(401)
		response.Code = "401"
		response.Status = 1
		response.Message = "Sorry, password is didnt match."
		this.Data["json"] = response
		this.ServeJSON()
		return
	}
	token, err := utils.GenerateJWT(user)
	if err != nil {
		panic(err.Error())
	}
	this.Ctx.ResponseWriter.WriteHeader(200)
	response.Code = "200"
	response.Status = 1
	response.Message = "Success!"
	response.Data = map[string]string{
		"token": token,
	}
	this.Data["json"] = response
	this.ServeJSON()
	return
}


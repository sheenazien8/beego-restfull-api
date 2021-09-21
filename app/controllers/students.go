package controllers

import (
	"encoding/json"
	"example/bootstrap"
)

type StudentController struct {
	BaseController
}

// @Title Get
// @Description Example get data
// @Success 200 {object}
// @router / [get]
func (this *StudentController) Get() {
	db := bootstrap.DB
	db.Find(&users)
	response.Code = "200"
	response.Status = 1
	response.Message = "SUCCESS"
	response.Data = users
	response.Paginate = this.PaginateValidate()
	this.Data["json"] = response
	this.ServeJSON()
}

// @Title Create
// @Description Example create data
// @Success 200 {object}
// @router / [post]
func (this *StudentController) Create() {
	json.Unmarshal(this.Ctx.Input.RequestBody, &user)
	user.DescriptionRoles = "Ini Student"
	user.Roles = "student"
	db := bootstrap.DB
	db.Create(&user)
	response.Code = "200"
	response.Status = 1
	response.Message = "SUCCESS"
	response.Data = user
	this.Data["json"] = response
	this.ServeJSON()
}

// @Title GetOne
// @Description Example get data
// @Success 200 {object}
// @router /:id [get]
func (this *StudentController) GetOne(id int) {
	db := bootstrap.DB
	db.Find(&user, id)
	response.Code = "200"
	response.Status = 1
	response.Message = "SUCCESS"
	response.Data = user
	this.Data["json"] = response
	this.ServeJSON()
}

// @Title Update
// @Description Example update data
// @Success 200 {object}
// @router /:id [put]
func (this *StudentController) Update(id int) {
	db := bootstrap.DB
	db.Find(&user, id)
	json.Unmarshal(this.Ctx.Input.RequestBody, &user)
	db.Save(&user)
	response.Code = "200"
	response.Status = 1
	response.Message = "SUCCESS"
	response.Data = user
	this.Data["json"] = response
	this.ServeJSON()
}

// @Title Delete
// @Description Example delete data
// @Success 200 {object}
// @router /:id [delete]
func (this *StudentController) Deletes(id int) {
	db := bootstrap.DB
	db.Delete(&user, id)
	response.Code = "200"
	response.Status = 1
	response.Message = "SUCCESS"
	response.Data = map[string]int{
		"id": id,
	}
	this.Data["json"] = response
	this.ServeJSON()
}

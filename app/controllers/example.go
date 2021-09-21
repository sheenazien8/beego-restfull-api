package controllers

type ExampleController struct {
	BaseController
}

// @Title Example GET
// @Description Example get data
// @Success 200 {object}
// @router / [get]
func (c *ExampleController) Get() {
	c.Data["json"] = map[string]string{
		"message": "success",
	}
	c.ServeJSON()
}

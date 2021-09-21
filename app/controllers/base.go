package controllers

import (
  "example/app/mock/req"
  "example/app/mock/res"
  "github.com/beego/beego/v2/server/web"
)

type BaseController struct {
  web.Controller
}

func (bc *BaseController) PaginateValidate() req.Paginate {
  paginate := req.Paginate{
    Page:    1,
    Last:    0,
    Limit:   20,
    Total:   0,
    Keyword: "",
  }

  if page, err := bc.GetInt("page"); err == nil {
    paginate.Page = page
  }

  if limit, err := bc.GetInt("limit"); err == nil {
    paginate.Limit = limit

  }

  if keyword := bc.GetString("keyword"); keyword != "" {
    paginate.Keyword = keyword
  }

  return paginate
}

func (bc *BaseController) Response(result res.Result) {
  bc.Ctx.Output.SetStatus(result.Status)
  bc.Data["json"] = result
  bc.ServeJSON()
}

package models

import "github.com/beego/beego/v2/client/orm"

type Users struct {
  Id int
  Roles string
  DescriptionRoles string
  Username string
  Password string
  NickName string
  FullName string
  Nis string
}

func init() {
  orm.RegisterModel(new(Users))
}

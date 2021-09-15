package main

import (
  _ "beego/routers"
  "fmt"
  "github.com/beego/beego/v2/client/orm"
  beego "github.com/beego/beego/v2/server/web"
  _ "github.com/go-sql-driver/mysql"
)

func init() {
  fmt.Println("Init")
  orm.RegisterDriver("mysql", orm.DRMySQL)
  orm.RegisterDataBase("default", "mysql", "homestead:secret@tcp(192.168.10.10:3306)/schools?charset=utf8")
  //orm.RegisterModel(new(models.User))
}

func main() {
  if beego.BConfig.RunMode == "dev" {
    beego.BConfig.WebConfig.DirectoryIndex = true
    beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
  }
  // autosync
  // db alias
  name := "default"

  // drop table and re-create
  force := false

  // print log
  verbose := true

  // error
  err := orm.RunSyncdb(name, force, verbose)
  if err != nil {
    fmt.Println(err)
  }
  beego.Run()
}

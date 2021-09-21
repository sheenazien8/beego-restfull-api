package bootstrap

import (
	"github.com/beego/beego/v2/adapter"
	"github.com/beego/beego/v2/adapter/plugins/cors"
)

func (b *Bootstrap) Cors() {
	adapter.InsertFilter("*", adapter.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE"},
		AllowHeaders:     []string{"Origin", "public-key", "Authorization", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))
}

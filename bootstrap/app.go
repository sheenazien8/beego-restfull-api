package bootstrap

import (
	"github.com/beego/beego/v2/server/web"
)

type Bootstrap struct{}

func Setup() {
	bootstrap := Bootstrap{}

	bootstrap.Env()
	bootstrap.Cors()
	bootstrap.Config()
	bootstrap.Logger()
	bootstrap.Database()
	bootstrap.Documentation()

	web.Run()
}

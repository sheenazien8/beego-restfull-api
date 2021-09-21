package bootstrap

import (
	"os"

	"github.com/beego/beego/v2/server/web"
)

func (b *Bootstrap) Documentation() {
	debug := os.Getenv("NODE_ENV")
	if debug == "development" {
		web.BConfig.WebConfig.DirectoryIndex = true
		web.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
}

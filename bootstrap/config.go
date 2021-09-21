package bootstrap

import (
	"os"
	"strconv"

	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
)

func (b *Bootstrap) Config() {
	httpPort, err := strconv.Atoi(os.Getenv("APP_PORT"))

	if err != nil {
		logs.Error("Error PORT not setup or invalid!", err)
	}

	debug := os.Getenv("NODE_ENV")

	if debug == "development" {
		web.BConfig.RunMode = "dev"
	} else {
		web.BConfig.RunMode = "prod"
	}

	web.BConfig.CopyRequestBody = true
	web.BConfig.Listen.HTTPPort = httpPort
	web.BConfig.WebConfig.CommentRouterPath = "app/controllers"
}

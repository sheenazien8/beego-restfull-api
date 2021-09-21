package bootstrap

import (
	"github.com/beego/beego/v2/core/logs"
	"github.com/joho/godotenv"
)

func (b *Bootstrap) Env() {
	err := godotenv.Load()

	if err != nil {
		logs.Error("Error loading .env file", err)
	}
}

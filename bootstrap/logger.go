package bootstrap

import "github.com/beego/beego/v2/core/logs"

func (b *Bootstrap) Logger() {
	logs.SetLogger(logs.AdapterFile, `{"filename":"logs/output.log","level":7,"maxlines":0,"maxsize":0,"daily":true,"maxdays":10,"color":true}`)
}

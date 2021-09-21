package bootstrap

import (
	"os"

	"github.com/beego/beego/v2/core/logs"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func (b *Bootstrap) Database() {
	host := os.Getenv("DATABASE_HOST")
	port := os.Getenv("DATABASE_PORT")
	user := os.Getenv("DATABASE_USER")
	pass := os.Getenv("DATABASE_PASS")
	name := os.Getenv("DATABASE_NAME")

	sqlconn := user + ":" + pass + "@tcp(" + host + ":" + port + ")/" + name + "?"
	sqlconn += "charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(sqlconn), &gorm.Config{})

	if err != nil {
		logs.Error("Error connection database!")
	}

	DB = db
}

package database

import (
	"github.com/MarcosRoch4/prjdini/helpers"
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

func InitDatabase() {
	database, err := gorm.Open("postgres", "host=127.0.0.1 port=5432 user=postgres dbname=prjdini password=postgres sslmode=disable")
	helpers.HandlerErr(err)
	database.DB().SetMaxIdleConns(20)
	database.DB().SetMaxOpenConns(200)
	DB = database
}

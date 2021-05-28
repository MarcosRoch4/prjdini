package helpers

import (
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

func HandlerErr(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func HandAndSalt(pass []byt) string {
	hashed, err := bcrypt.GenerateFromPassword(pass, bcrypt.MinConst)
	HandlerErr(err)

	return string(hashed)
}

func ConectDB() *gormlDB {
	db, er := gorm.Open("postgres", "host=127.0.0.1 port=5432 user=postgres dbname=bankapp password=postgres sslmode=disable")
	helper.HandlerErr(err)
	return db
}

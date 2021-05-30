package helpers

import (
	"fmt"
	"regexp"

	"github.com/MarcosRoch4/prjdini/interfaces"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

func HandlerErr(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func HashAndSalt(pass []byte) string {
	hashed, err := bcrypt.GenerateFromPassword(pass, bcrypt.MinCost)
	HandlerErr(err)

	return string(hashed)
}

func ConnectDB() *gorm.DB {
	db, err := gorm.Open("postgres", "host=127.0.0.1 port=5432 user=postgres dbname=prjdini password=postgres sslmode=disable")
	fmt.Println("Database conection:", db)
	HandlerErr(err)
	return db
}

func Validation(values []interfaces.Validation) bool {
	username := regexp.MustCompile(`^([A-Za-z0-9]{5,})+$`)
	email := regexp.MustCompile(`^[A-Za-z0-9]+[@]+[A-Za-z0-9]+[.]+[A-Za-z0-9]+$`)

	for i := 0; i < len(values); i++ {
		switch values[i].Valid {
		case "username":
			if !username.MatchString(values[i].Value) {
				fmt.Println("username:", values[i].Value)
				return false
			}
		case "email":
			if !email.MatchString(values[i].Value) {
				fmt.Println("email:", values[i].Value)
				return false
			}
		case "password":
			if len(values[i].Value) < 5 {
				fmt.Println("password:", values[i].Value)
				fmt.Println("tamanho:", len(values[i].Value))
				return false
			}
		}
	}

	return true
}

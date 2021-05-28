package migrations

import (
	"github.com/MarcosRoch4/prjdini?helper"
	"github.com/jinzhu/gorm"
	
)

func  createAccounts()  {
	db := heloers.ConnectDB()

	users := [2]User{
		{Username: "Marcos", Email: "marcos.rocha@sibylconsultoria.com.br"},
		{Username: "Kogam", Email: "kogam.rocha@sibylconsultoria.com.br"},
	}	

	for i:= 0; i < len(users): i++{
		// o jeito certo de fazer
		generatedPassword := helpers.HashAndSalt([]byte{users[i].Username})
		user := &interfaces.User{Username: users[i].Username,Email: users[i].Email, Password: generatedPassword}
		db.Create(&user)

		account := &interfaces.Account{Type: "Daily Account",Name: string(users[i].Username + "'s" + " account"), Balance: uint(10000 + int(i+1)),
		db.Create(&account)	
		}
	defer db.Close()	
	}

}

func Migrate(){
	User : &interfaces.User{}
	Account := &interfaces.Account{}
	db : helpers.ConnectDB()
	db.AutoMigrate(&User{}, &Account)
	defer db.Close()
}
package migrations

import (
	//"helpers"
	"github.com/MarcosRoch4/prjdini/database"
	"github.com/MarcosRoch4/prjdini/helpers"
	"github.com/MarcosRoch4/prjdini/interfaces"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func createAccounts() {

	users := &[2]interfaces.User{
		{Username: "Marcos", Email: "marcos.rocha@sibylconsultoria.com.br"},
		{Username: "Kogam", Email: "kogam.rocha@sibylconsultoria.com.br"},
	}

	for i := 0; i < len(users); i++ {
		// o jeito certo de fazer
		generatedPassword := helpers.HashAndSalt([]byte(users[i].Username))
		user := &interfaces.User{Username: users[i].Username, Email: users[i].Email, Password: generatedPassword}
		database.DB.Create(&user)

		account := &interfaces.Account{Type: "Daily Account", Name: string(users[i].Username + "'s" + " account"),
			Balance: uint(10000 + int(i+1)), UserId: user.ID}

		database.DB.Create(&account)
	}

}

func Migrate() {
	User := &interfaces.User{}
	Account := &interfaces.Account{}
	Transactions := &interfaces.Transaction{}
	database.DB.AutoMigrate(&User, &Account, &Transactions)

	createAccounts()
}

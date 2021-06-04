package main

import (
	"github.com/MarcosRoch4/prjdini/api"
	"github.com/MarcosRoch4/prjdini/database"
)

//	"github.com/MarcosRoch4/prjdini/migrations"

func main() {
	//migrations.MigrateTransactions()

	database.InitDatabase()

	api.StartApi()
}

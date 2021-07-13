package main

import (
	app "happy_bank_simulator/app"
	"happy_bank_simulator/database"
	databaseHelpers "happy_bank_simulator/database/helpers"
)

func main() {
	database.InitDB()
	databaseHelpers.MigrateDB()

	app.InitApp()
}

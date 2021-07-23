package main

import (
	app "happy_bank_simulator/app"
	"happy_bank_simulator/database"
	databaseHelpers "happy_bank_simulator/database/helpers"
	"happy_bank_simulator/simulation"
)

// TODO
//
// - Fix duplicated lender assignation when more than two loans to create
// - Refactor output messages: remove messages from models and place them only on controller side

func main() {
	database.InitDB()
	databaseHelpers.MigrateDB()

	simulation.Prepare()
	app.InitApp()
}

package main

import (
	// app "happy_bank_simulator/app"
	"happy_bank_simulator/database"
	databaseHelpers "happy_bank_simulator/database/helpers"
	"happy_bank_simulator/simulation"
)

func main() {
	database.InitDB()
	databaseHelpers.MigrateDB()

	// simulation.Prepare()
	// app.InitApp()
	simulation.Run()
}

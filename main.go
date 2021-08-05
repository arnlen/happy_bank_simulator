package main

import (
	// app "happy_bank_simulator/app"
	"happy_bank_simulator/internal/database"
	"happy_bank_simulator/simulation"
)

func main() {
	database.SetupDB()

	simulation.Prepare()
	// app.InitApp()
	simulation.Run()
}

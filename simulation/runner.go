package simulation

import (
	databaseHelpers "happy_bank_simulator/database/helpers"
)

var quantityOfLoansToCreate int

func Prepare() {
	databaseHelpers.DropBD()
	databaseHelpers.MigrateDB()
	createLoans()
}

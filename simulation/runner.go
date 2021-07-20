package simulation

import (
	databaseHelpers "happy_bank_simulator/database/helpers"
	"happy_bank_simulator/models"
)

var (
	quantityOfLoansToCreate int
	depositAccount          models.DepositAccount
)

func Prepare() {
	databaseHelpers.DropBD()
	databaseHelpers.MigrateDB()

	depositAccount.Balance = 0

	createLoans()
}

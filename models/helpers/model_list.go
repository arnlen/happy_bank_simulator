package helpers

import "happy_bank_simulator/models"

var modelList = []interface{}{
	&models.Borrower{},
	&models.Insurer{},
	&models.Lender{},
	&models.Loan{},
}

func GetModelList() []interface{} {
	return modelList
}

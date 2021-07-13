package models

import ()

var modelList = []interface{}{
	&Borrower{},
	&Insurer{},
	&Lender{},
	&Loan{},
}

func GetModelList() []interface{} {
	return modelList
}

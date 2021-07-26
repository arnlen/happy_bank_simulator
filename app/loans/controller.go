package loans

import (
	"fmt"
	"happy_bank_simulator/models"
	"strconv"
)

type Controller struct{}

func (c *Controller) GetModelName(pluralize bool) string {
	loanModel := models.Loan{}
	if pluralize {
		return fmt.Sprintf("%ss", loanModel.ModelName())
	}
	return loanModel.ModelName()
}

func (c *Controller) GetLoanTableData() [][]string {
	loans := models.ListLoans()

	loanTableData := [][]string{
		{"ID", "Débiteur", "Créancier", "Assureur", "Montant", "Durée"}}

	for _, loan := range loans {
		lenders := loan.Lenders
		lendersString := fmt.Sprintf("%s lenders", strconv.Itoa(len(lenders)))
		insurers := loan.Insurers
		insurersString := fmt.Sprintf("%s insurers", strconv.Itoa(len(insurers)))

		loanRow := []string{
			strconv.Itoa(int(loan.ID)),
			loan.Borrower.Name,
			lendersString,
			insurersString,
			fmt.Sprintf("%1.2f €", loan.Amount),
			fmt.Sprintf("%s mois", strconv.Itoa(int(loan.Duration))),
		}

		loanTableData = append(loanTableData, loanRow)
	}

	return loanTableData
}

func (c *Controller) GetLoanStringList() []string {
	loans := models.ListLoans()
	var loanStringList []string

	for _, loan := range loans {
		string := fmt.Sprintf("%s - %1.2f € ", strconv.Itoa(int(loan.ID)), loan.Amount)
		loanStringList = append(loanStringList, string)
	}

	return loanStringList
}

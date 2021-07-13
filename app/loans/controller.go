package loans

import (
	"fmt"
	"happy_bank_simulator/database"
	"happy_bank_simulator/models"
	"strconv"

	"gorm.io/gorm/clause"
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
	var loans []models.Loan
	database.GetDB().Preload(clause.Associations).Find(&loans)

	loanTableData := [][]string{
		{"ID", "Débiteur", "Créancier", "Assureur", "Montant", "Durée"}}

	for _, loan := range loans {
		loanRow := []string{
			strconv.Itoa(int(loan.ID)),
			loan.Borrower.Name,
			loan.Lender.Name,
			loan.Insurer.Name,
			fmt.Sprintf("%8.0f €", loan.Amount),
			fmt.Sprintf("%s mois", strconv.Itoa(int(loan.Duration))),
		}

		loanTableData = append(loanTableData, loanRow)
	}

	return loanTableData
}

func (c *Controller) GetLoanStringList() []string {
	var loanStringList []string

	var loans []models.Loan
	database.GetDB().Preload(clause.Associations).Find(&loans)

	for _, loan := range loans {
		string := fmt.Sprintf("%s - %8.0f € ", strconv.Itoa(int(loan.ID)), loan.Amount)
		loanStringList = append(loanStringList, string)
	}

	return loanStringList
}

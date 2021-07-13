package loans

import (
	"fmt"
	"happy_bank_simulator/database"
	"happy_bank_simulator/models"
	"strconv"

	"gorm.io/gorm/clause"
)

type Controller struct{}

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

// func (c *Controller) Create(name string, balance float64) *models.Loan {
// 	loan := &models.Loan{
// 		Name:    name,
// 		Loans:   []models.Loan{},
// 		Balance: balance,
// 	}

// 	result := loan.Create()

// 	if loan.ID == 0 || result.RowsAffected == 0 {
// 		log.Fatal(result.Error)
// 	}

// 	return loan
// }

package borrowers

import (
	"fmt"
	"happy_bank_simulator/database"
	"happy_bank_simulator/models"
	"log"
	"strconv"

	"gorm.io/gorm/clause"
)

// Declare conformity with BaseController interface
// var _ BaseController = (*Borrowers)(nil)

type Controller struct{}

func (c *Controller) GetBorrowerTableData() [][]string {
	var borrowers []models.Borrower
	database.GetDB().Preload(clause.Associations).Find(&borrowers)

	borrowerTableData := [][]string{
		{"ID", "Name", "Balance"}}

	for _, borrower := range borrowers {
		borrowerRow := []string{
			strconv.Itoa(int(borrower.ID)),
			borrower.Name,
			fmt.Sprintf("%8.0f €", borrower.Balance),
		}

		borrowerTableData = append(borrowerTableData, borrowerRow)
	}

	return borrowerTableData
}

func (c *Controller) Create(name string, balance float64) *models.Borrower {
	borrower := &models.Borrower{
		Name:    name,
		Loans:   []models.Loan{},
		Balance: balance,
	}

	result := borrower.Create()

	if borrower.ID == 0 || result.RowsAffected == 0 {
		log.Fatal(result.Error)
	}

	return borrower
}

package borrowers

import (
	"fmt"
	"happy_bank_simulator/app/configs"
	"happy_bank_simulator/models"
	"strconv"
)

type Controller struct{}

func (c *Controller) GetBorrowerTableData() [][]string {
	borrowers := models.ListActors(configs.Actor.BorrowerString)

	borrowerTableData := [][]string{
		{"ID", "Name", "Initial Balance", "Balance"}}

	for _, borrower := range borrowers {
		borrowerRow := []string{
			strconv.Itoa(int(borrower.ID)),
			borrower.Name,
			fmt.Sprintf("%1.2f €", borrower.InitialBalance),
			fmt.Sprintf("%1.2f €", borrower.Balance),
		}

		borrowerTableData = append(borrowerTableData, borrowerRow)
	}

	return borrowerTableData
}

func (c *Controller) Create(name string, balance float64) *models.Actor {
	return models.CreateActor(configs.Actor.BorrowerString, name, balance)
}

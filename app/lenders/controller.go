package lenders

import (
	"fmt"
	"happy_bank_simulator/models"
	"strconv"
)

type Controller struct{}

func (c *Controller) GetModelName(pluralize bool) string {
	lenderModel := models.Lender{}
	if pluralize {
		return fmt.Sprintf("%ss", lenderModel.ModelName())
	}
	return lenderModel.ModelName()
}

func (c *Controller) GetLenderTableData() [][]string {
	lenders := models.ListLenders()

	lenderTableData := [][]string{
		{"ID", "Name", "Initial Balance", "Balance"}}

	for _, lender := range lenders {
		lenderRow := []string{
			strconv.Itoa(int(lender.ID)),
			lender.Name,
			fmt.Sprintf("%1.2f €", lender.InitialBalance),
			fmt.Sprintf("%1.2f €", lender.Balance),
		}

		lenderTableData = append(lenderTableData, lenderRow)
	}

	return lenderTableData
}

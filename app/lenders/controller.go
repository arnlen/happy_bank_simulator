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
		{"ID", "Name", "Balance"}}

	for _, lender := range lenders {
		lenderRow := []string{
			strconv.Itoa(int(lender.ID)),
			lender.Name,
			fmt.Sprintf("%s €", strconv.Itoa(lender.Balance)),
		}

		lenderTableData = append(lenderTableData, lenderRow)
	}

	return lenderTableData
}

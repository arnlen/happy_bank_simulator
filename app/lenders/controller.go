package lenders

import (
	"fmt"
	"happy_bank_simulator/app/configs"
	"happy_bank_simulator/models"
	"strconv"
)

type Controller struct{}

func (c *Controller) GetModelName(pluralize bool) string {
	lenderModel := models.Actor{}
	if pluralize {
		return fmt.Sprintf("%ss", lenderModel.Type)
	}
	return lenderModel.Type
}

func (c *Controller) GetLenderTableData() [][]string {
	lenders := models.ListActors(configs.Actor.Lender)

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

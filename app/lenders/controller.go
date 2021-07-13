package lenders

import (
	"fmt"
	"happy_bank_simulator/database"
	"happy_bank_simulator/models"
	"strconv"

	"gorm.io/gorm/clause"
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
	var lenders []models.Lender
	database.GetDB().Preload(clause.Associations).Find(&lenders)

	lenderTableData := [][]string{
		{"ID", "Name", "Balance"}}

	for _, lender := range lenders {
		lenderRow := []string{
			strconv.Itoa(int(lender.ID)),
			lender.Name,
			fmt.Sprintf("%8.0f €", lender.Balance),
		}

		lenderTableData = append(lenderTableData, lenderRow)
	}

	return lenderTableData
}

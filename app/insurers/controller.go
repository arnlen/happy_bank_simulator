package insurers

import (
	"fmt"
	"happy_bank_simulator/database"
	"happy_bank_simulator/models"
	"strconv"

	"gorm.io/gorm/clause"
)

type Controller struct{}

func (c *Controller) GetModelName(pluralize bool) string {
	insurerModel := models.Insurer{}
	if pluralize {
		return fmt.Sprintf("%ss", insurerModel.ModelName())
	}
	return insurerModel.ModelName()
}

func (c *Controller) GetInsurerTableData() [][]string {
	var insurers []models.Insurer
	database.GetDB().Preload(clause.Associations).Find(&insurers)

	insurerTableData := [][]string{
		{"ID", "Name", "Balance"}}

	for _, insurer := range insurers {
		insurerRow := []string{
			strconv.Itoa(int(insurer.ID)),
			insurer.Name,
			fmt.Sprintf("%8.0f €", insurer.Balance),
		}

		insurerTableData = append(insurerTableData, insurerRow)
	}

	return insurerTableData
}

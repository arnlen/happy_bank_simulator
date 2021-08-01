package insurers

import (
	"fmt"
	"happy_bank_simulator/app/configs"
	"happy_bank_simulator/models"
	"strconv"
)

type Controller struct{}

func (c *Controller) GetModelName(pluralize bool) string {
	insurerModel := models.Actor{}
	if pluralize {
		return fmt.Sprintf("%ss", insurerModel.Type)
	}
	return insurerModel.Type
}

func (c *Controller) GetInsurerTableData() [][]string {
	insurers := models.ListActors(configs.Actor.Insurer)

	insurerTableData := [][]string{
		{"ID", "Name", "Initial Balance", "Balance"}}

	for _, insurer := range insurers {
		insurerRow := []string{
			strconv.Itoa(int(insurer.ID)),
			insurer.Name,
			fmt.Sprintf("%1.2f €", insurer.InitialBalance),
			fmt.Sprintf("%1.2f €", insurer.Balance),
		}

		insurerTableData = append(insurerTableData, insurerRow)
	}

	return insurerTableData
}

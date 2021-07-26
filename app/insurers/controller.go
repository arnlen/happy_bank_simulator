package insurers

import (
	"fmt"
	"happy_bank_simulator/models"
	"strconv"
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
	insurers := models.ListInsurers()

	insurerTableData := [][]string{
		{"ID", "Name", "Balance"}}

	for _, insurer := range insurers {
		insurerRow := []string{
			strconv.Itoa(int(insurer.ID)),
			insurer.Name,
			fmt.Sprintf("%1.2f €", insurer.Balance),
		}

		insurerTableData = append(insurerTableData, insurerRow)
	}

	return insurerTableData
}

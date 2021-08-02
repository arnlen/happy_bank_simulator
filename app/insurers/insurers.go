package insurers

import (
	"fmt"
	"happy_bank_simulator/app/configs"
	"happy_bank_simulator/models"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func RenderIndex() *fyne.Container {
	insurerTableData := getInsurerTableData()
	fmt.Printf("%s", insurerTableData[2])

	table := widget.NewTable(
		func() (int, int) {
			return len(insurerTableData), len(insurerTableData[0])
		},
		func() fyne.CanvasObject {
			item := widget.NewLabel("Template")
			return item
		},
		func(cell widget.TableCellID, item fyne.CanvasObject) {
			item.(*widget.Label).SetText(insurerTableData[cell.Row][cell.Col])
		})

	table.SetColumnWidth(0, 50)
	table.SetColumnWidth(1, 250)
	table.SetColumnWidth(2, 150)
	table.SetColumnWidth(3, 150)

	return container.NewBorder(nil, nil, nil, nil, table)
}

func getInsurerTableData() [][]string {
	insurers := models.ListActors(configs.Actor.InsurerString)

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

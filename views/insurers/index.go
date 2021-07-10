package views

import (
	"fmt"
	"happy_bank_simulator/models"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func Index(insurers []models.Insurer) *fyne.Container {
	var insurersTableData = [][]string{
		{"ID", "Name", "Balance"}}

	for _, insurer := range insurers {
		insurerRow := []string{
			strconv.Itoa(int(insurer.ID)),
			insurer.Name,
			fmt.Sprintf("%8.0f €", insurer.Balance),
		}

		insurersTableData = append(insurersTableData, insurerRow)
	}

	table := widget.NewTable(
		func() (int, int) {
			return len(insurersTableData), len(insurersTableData[0])
		},
		func() fyne.CanvasObject {
			item := widget.NewLabel("Template")
			return item
		},
		func(cell widget.TableCellID, item fyne.CanvasObject) {
			item.(*widget.Label).SetText(insurersTableData[cell.Row][cell.Col])
		})

	table.SetColumnWidth(0, 50)
	table.SetColumnWidth(1, 250)
	table.SetColumnWidth(2, 100)

	newButton := widget.NewButton("Nouvel assureur", func() {
		fmt.Println("New button")
	})

	return container.NewBorder(newButton, nil, nil, nil, table)
}

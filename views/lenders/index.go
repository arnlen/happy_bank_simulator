package views

import (
	"fmt"
	"happy_bank_simulator/models"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func Index(lenders []models.Lender) *fyne.Container {
	var lendersTableData = [][]string{
		{"ID", "Name", "Balance"}}

	for _, lender := range lenders {
		lenderRow := []string{
			strconv.Itoa(int(lender.ID)),
			lender.Name,
			fmt.Sprintf("%8.0f €", lender.Balance),
		}

		lendersTableData = append(lendersTableData, lenderRow)
	}

	table := widget.NewTable(
		func() (int, int) {
			return len(lendersTableData), len(lendersTableData[0])
		},
		func() fyne.CanvasObject {
			item := widget.NewLabel("Template")
			return item
		},
		func(cell widget.TableCellID, item fyne.CanvasObject) {
			item.(*widget.Label).SetText(lendersTableData[cell.Row][cell.Col])
		})

	table.SetColumnWidth(0, 50)
	table.SetColumnWidth(1, 250)
	table.SetColumnWidth(2, 100)

	return container.NewBorder(nil, nil, nil, nil, table)
}

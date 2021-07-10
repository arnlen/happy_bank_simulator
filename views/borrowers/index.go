package views

import (
	"fmt"
	"happy_bank_simulator/models"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func Index(borrowers []models.Borrower) *fyne.Container {
	var borrowersTableData = [][]string{
		{"ID", "Name", "Balance"}}

	for _, borrower := range borrowers {
		borrowerRow := []string{
			strconv.Itoa(int(borrower.ID)),
			borrower.Name,
			fmt.Sprintf("%8.0f €", borrower.Balance),
		}

		borrowersTableData = append(borrowersTableData, borrowerRow)
	}

	table := widget.NewTable(
		func() (int, int) {
			return len(borrowersTableData), len(borrowersTableData[0])
		},
		func() fyne.CanvasObject {
			item := widget.NewLabel("Template")
			return item
		},
		func(cell widget.TableCellID, item fyne.CanvasObject) {
			item.(*widget.Label).SetText(borrowersTableData[cell.Row][cell.Col])
		})

	table.SetColumnWidth(0, 50)
	table.SetColumnWidth(1, 250)
	table.SetColumnWidth(2, 100)

	newButton := widget.NewButton("Nouveau débiteur", func() {
		fmt.Println("New button")
	})

	return container.NewBorder(newButton, nil, nil, nil, table)
}

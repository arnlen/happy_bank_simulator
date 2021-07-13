package views

import (
	"fmt"
	"happy_bank_simulator/models"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func Index(loans []models.Loan) *fyne.Container {
	var loansTableData = [][]string{
		{"ID", "Débiteur", "Créancier", "Assureur", "Montant", "Durée"}}

	for _, loan := range loans {
		loanRow := []string{
			strconv.Itoa(int(loan.ID)),
			loan.Borrower.Name,
			loan.Lender.Name,
			loan.Insurer.Name,
			fmt.Sprintf("%8.0f €", loan.Amount),
			fmt.Sprintf("%s mois", strconv.Itoa(int(loan.Duration))),
		}

		loansTableData = append(loansTableData, loanRow)
	}

	table := widget.NewTable(
		func() (int, int) {
			return len(loansTableData), len(loansTableData[0])
		},
		func() fyne.CanvasObject {
			item := widget.NewLabel("Template")
			return item
		},
		func(cell widget.TableCellID, item fyne.CanvasObject) {
			item.(*widget.Label).SetText(loansTableData[cell.Row][cell.Col])
		})

	table.SetColumnWidth(0, 50)
	table.SetColumnWidth(1, 250)
	table.SetColumnWidth(2, 250)
	table.SetColumnWidth(3, 250)
	table.SetColumnWidth(4, 100)

	newButton := widget.NewButton("Nouveau prêt", func() {
		fmt.Println("New button")
	})

	return container.NewBorder(newButton, nil, nil, nil, table)
}

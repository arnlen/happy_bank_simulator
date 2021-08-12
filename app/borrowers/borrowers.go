package borrowers

import (
	"fmt"
	"happy_bank_simulator/app/configs"
	"happy_bank_simulator/app/helpers"
	"happy_bank_simulator/app/loans"
	"happy_bank_simulator/models"
	"strconv"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func RenderIndex() *fyne.Container {
	borrowerTableData := updateTableData()

	table := widget.NewTable(
		func() (int, int) {
			return len(borrowerTableData), len(borrowerTableData[0])
		},
		func() fyne.CanvasObject {
			item := widget.NewLabel("Template")
			return item
		},
		func(cell widget.TableCellID, item fyne.CanvasObject) {
			item.(*widget.Label).SetText(borrowerTableData[cell.Row][cell.Col])
		})

	table.SetColumnWidth(0, 50)
	table.SetColumnWidth(1, 250)
	table.SetColumnWidth(2, 150)
	table.SetColumnWidth(3, 150)

	newButtonString := strings.Title(configs.Actor.BorrowerString)
	newButton := widget.NewButtonWithIcon(newButtonString, theme.ContentAddIcon(), func() {
		fmt.Println("newButton", newButtonString)
		RenderNew()
	})

	return container.NewBorder(nil, newButton, nil, nil, table)
}

func updateTableData() [][]string {
	return getBorrowerTableData()
}

func RenderNew() {
	loanList := loans.GetLoanStringList()

	nameEntry := widget.NewEntry()
	balanceEntry := widget.NewEntry()
	selectMenu := widget.Select{
		Options: loanList,
	}

	formItems := []*widget.FormItem{
		{Text: "Nom", Widget: nameEntry},
		{Text: "Balance", Widget: balanceEntry},
		{Text: "Prêt associé", Widget: &selectMenu},
	}

	dialog.ShowForm(strings.Title(configs.Actor.BorrowerString), "Créer", "Annuler", formItems, func(bool) {
		fmt.Println("Nom :", nameEntry.Text)
		fmt.Println("Balance :", balanceEntry.Text)
	}, helpers.GetMasterWindow())
}

func getBorrowerTableData() [][]string {
	borrowers := models.ListActors(configs.Actor.BorrowerString)

	borrowerTableData := [][]string{
		{"ID", "Name", "Initial Balance", "Balance"}}

	for _, borrower := range borrowers {
		borrowerRow := []string{
			strconv.Itoa(int(borrower.ID)),
			borrower.Name,
			fmt.Sprintf("%1.2f €", borrower.InitialBalance),
			fmt.Sprintf("%1.2f €", borrower.Balance),
		}

		borrowerTableData = append(borrowerTableData, borrowerRow)
	}

	return borrowerTableData
}

func Create(name string, balance float64) *models.Actor {
	return models.CreateBorrower()
}

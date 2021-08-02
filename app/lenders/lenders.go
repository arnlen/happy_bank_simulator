package lenders

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
	lenderTableData := getLenderTableData()

	table := widget.NewTable(
		func() (int, int) {
			return len(lenderTableData), len(lenderTableData[0])
		},
		func() fyne.CanvasObject {
			item := widget.NewLabel("Template")
			return item
		},
		func(cell widget.TableCellID, item fyne.CanvasObject) {
			item.(*widget.Label).SetText(lenderTableData[cell.Row][cell.Col])
		})

	table.SetColumnWidth(0, 50)
	table.SetColumnWidth(1, 250)
	table.SetColumnWidth(2, 150)
	table.SetColumnWidth(3, 150)

	newButtonString := strings.Title(configs.Actor.LenderString)
	newButton := widget.NewButtonWithIcon(newButtonString, theme.ContentAddIcon(), func() {
		fmt.Println("newButton", newButtonString)
		RenderNew()
	})

	return container.NewBorder(nil, newButton, nil, nil, table)
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

	dialog.ShowForm(strings.Title(configs.Actor.LenderString), "Créer", "Annuler", formItems, func(bool) {
		fmt.Println("Nom :", nameEntry.Text)
		fmt.Println("Balance :", balanceEntry.Text)
	}, helpers.GetMasterWindow())
}

func getLenderTableData() [][]string {
	lenders := models.ListActors(configs.Actor.LenderString)

	lenderTableData := [][]string{
		{"ID", "Name", "Initial Balance", "Balance"}}

	for _, lender := range lenders {
		lenderRow := []string{
			strconv.Itoa(int(lender.ID)),
			lender.Name,
			fmt.Sprintf("%1.2f €", lender.InitialBalance),
			fmt.Sprintf("%1.2f €", lender.Balance),
		}

		lenderTableData = append(lenderTableData, lenderRow)
	}

	return lenderTableData
}

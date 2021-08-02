package transactions

import (
	"fmt"
	"strconv"
	"strings"

	"happy_bank_simulator/app/configs"
	"happy_bank_simulator/models"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func RenderIndex() *fyne.Container {
	transactionTableData := getTransactionTableData()

	table := widget.NewTable(
		func() (int, int) {
			return len(transactionTableData), len(transactionTableData[0])
		},
		func() fyne.CanvasObject {
			item := widget.NewLabel("Template")
			return item
		},
		func(cell widget.TableCellID, item fyne.CanvasObject) {
			item.(*widget.Label).SetText(transactionTableData[cell.Row][cell.Col])
		})

	table.SetColumnWidth(0, 50)
	table.SetColumnWidth(1, 250)
	table.SetColumnWidth(2, 150)
	table.SetColumnWidth(3, 150)

	newButtonString := strings.Title(configs.Transaction.String)
	newButton := widget.NewButtonWithIcon(newButtonString, theme.ContentAddIcon(), func() {
		fmt.Println("newButton", newButtonString)
		RenderNew()
	})

	return container.NewBorder(nil, newButton, nil, nil, table)
}

func RenderNew() {
}

func getTransactionTableData() [][]string {
	transactions := models.ListTransactions()

	transactionTableData := [][]string{
		{"ID", "Sender", "Receiver", "Amount"}}

	for _, transaction := range transactions {
		sender := fmt.Sprintf("%s #%s", transaction.SenderType, strconv.Itoa(transaction.SenderID))
		receiver := fmt.Sprintf("%s #%s", transaction.ReceiverType, strconv.Itoa(transaction.ReceiverID))

		transactionRow := []string{
			strconv.Itoa(int(transaction.ID)),
			sender,
			receiver,
			fmt.Sprintf("%1.2f €", transaction.Amount),
		}

		transactionTableData = append(transactionTableData, transactionRow)
	}

	return transactionTableData
}

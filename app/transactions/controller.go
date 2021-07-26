package transactions

import (
	"fmt"
	"happy_bank_simulator/models"
	"strconv"
)

type Controller struct{}

func (c *Controller) GetModelName(pluralize bool) string {
	transactionModel := models.Transaction{}
	if pluralize {
		return fmt.Sprintf("%ss", transactionModel.ModelName())
	}
	return transactionModel.ModelName()
}

func (c *Controller) GetTransactionTableData() [][]string {
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

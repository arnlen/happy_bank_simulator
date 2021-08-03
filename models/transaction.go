package models

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"happy_bank_simulator/database"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Transaction struct {
	gorm.Model
	SenderType   string
	SenderID     int
	ReceiverType string
	ReceiverID   int
	Amount       float64
}

func (instance *Transaction) ModelName() string {
	return "transaction"
}

func (instance *Transaction) Save() {
	result := database.GetDB().Save(instance)

	if instance.ID == 0 || result.RowsAffected == 0 {
		log.Fatal(result.Error)
	}

	instance.refresh()
}

func (instance *Transaction) Print() {
	instance.refresh()

	sender := "INCOMES"
	if instance.SenderID != 0 {
		sender = fmt.Sprintf("%s #%s",
			strings.Title(instance.SenderType),
			strconv.Itoa(int(instance.SenderID)),
		)
	}

	receiver := "DEPOSIT"
	if instance.ReceiverID != 0 {
		receiver = fmt.Sprintf("%s #%s",
			strings.Title(instance.ReceiverType),
			strconv.Itoa(int(instance.ReceiverID)),
		)
	}

	fmt.Printf("🔁 Transaction #%s: [%s] == %1.2f € ==> [%s]\n",
		strconv.Itoa(int(instance.ID)),
		sender,
		instance.Amount,
		receiver)
}

func (instance *Transaction) refresh() {
	database.GetDB().Preload(clause.Associations).Find(&instance)
}

// ------- Package methods -------

func ListTransactions() []*Transaction {
	var transactions []*Transaction
	database.GetDB().Preload(clause.Associations).Find(&transactions)
	return transactions
}

func CreateTransaction(sender Actor, receiver Actor, amount float64) *Transaction {
	sender.UpdateBalance(-amount)
	receiver.UpdateBalance(amount)

	transaction := &Transaction{
		SenderID:     int(sender.GetID()),
		SenderType:   sender.Type,
		ReceiverID:   int(receiver.GetID()),
		ReceiverType: receiver.Type,
		Amount:       amount,
	}

	transaction.Save()
	return transaction
}

func CreateDepositTransaction(borrower Actor, amount float64) *Transaction {
	borrower.UpdateBalance(-amount)

	depositTransaction := &Transaction{
		SenderID:     int(borrower.GetID()),
		SenderType:   borrower.Type,
		ReceiverID:   0,
		ReceiverType: "deposit",
		Amount:       amount,
	}

	depositTransaction.Save()
	return depositTransaction
}

func CreateIncomeTransaction(borrower Actor, amount float64) *Transaction {
	borrower.UpdateBalance(amount)

	incomeTransaction := &Transaction{
		SenderID:     0,
		SenderType:   "income",
		ReceiverID:   int(borrower.GetID()),
		ReceiverType: borrower.Type,
		Amount:       amount,
	}

	incomeTransaction.Save()
	return incomeTransaction
}

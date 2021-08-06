package models

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"happy_bank_simulator/internal/global"

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

func (instance *Transaction) Save() {
	result := global.Db.Save(instance)

	if instance.ID == 0 || result.RowsAffected == 0 {
		log.Fatal(result.Error)
	}

	instance.Refresh()
}

func (instance *Transaction) Refresh() {
	global.Db.Preload(clause.Associations).Find(&instance)
}

func (instance *Transaction) Print() {
	instance.Refresh()

	senderString := "INCOMES"
	if instance.SenderID != 0 {
		senderString = fmt.Sprintf("%s #%s",
			strings.Title(instance.SenderType),
			strconv.Itoa(int(instance.SenderID)),
		)
	}

	receiverString := "DEPOSIT"
	if instance.ReceiverID != 0 {
		receiverString = fmt.Sprintf("%s #%s",
			strings.Title(instance.ReceiverType),
			strconv.Itoa(int(instance.ReceiverID)),
		)
	}

	fmt.Printf("🔁 Transaction #%s: [%s] == %1.2f € ==> [%s]\n",
		strconv.Itoa(int(instance.ID)),
		senderString,
		instance.Amount,
		receiverString)
}

// ------- Package methods -------

func ListTransactions() []*Transaction {
	var transactions []*Transaction
	global.Db.Preload(clause.Associations).Find(&transactions)
	return transactions
}

func CreateTransaction(sender Actor, receiver Actor, amount float64) *Transaction {
	sender.UpdateBalance(-amount)
	receiver.UpdateBalance(amount)

	transaction := &Transaction{
		SenderID:     int(sender.ID),
		SenderType:   sender.Type,
		ReceiverID:   int(receiver.ID),
		ReceiverType: receiver.Type,
		Amount:       amount,
	}

	transaction.Save()
	return transaction
}

func CreateDepositTransaction(borrower Actor, amount float64) *Transaction {
	borrower.UpdateBalance(-amount)

	depositTransaction := &Transaction{
		SenderID:     int(borrower.ID),
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
		ReceiverID:   int(borrower.ID),
		ReceiverType: borrower.Type,
		Amount:       amount,
	}

	incomeTransaction.Save()
	return incomeTransaction
}

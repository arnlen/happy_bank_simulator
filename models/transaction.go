package models

import (
	"happy_bank_simulator/database"
	"log"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// Declare conformity with Actor interface
var _ ModelBase = (*Transaction)(nil)

type Transaction struct {
	gorm.Model
	SenderType   string
	SenderID     int
	ReceiverType string
	ReceiverID   int
	Amount       int
	isDeposit    bool
}

func (instance *Transaction) ModelName() string {
	return "transaction"
}

func (instance *Transaction) Refresh() {
	database.GetDB().Preload(clause.Associations).Find(&instance)
}

func (instance *Transaction) Save() {
	result := database.GetDB().Save(instance)

	if instance.ID == 0 || result.RowsAffected == 0 {
		log.Fatal(result.Error)
	}

	instance.Refresh()
}

// ------- Package methods -------

func ListTransactions() []*Transaction {
	var transactions []*Transaction
	database.GetDB().Preload(clause.Associations).Find(&transactions)
	return transactions
}

func NewTransaction(sender Actor, receiver Actor, amount int) *Transaction {
	sender.UpdateBalance(-amount)
	receiver.UpdateBalance(amount)

	return &Transaction{
		SenderID:     int(sender.GetID()),
		SenderType:   sender.ModelName(),
		ReceiverID:   int(receiver.GetID()),
		ReceiverType: receiver.ModelName(),
		Amount:       amount,
		isDeposit:    false,
	}
}

func NewDepositTransaction(borrower Borrower, amount int) *Transaction {
	borrower.UpdateBalance(-amount)

	return &Transaction{
		SenderID:     int(borrower.GetID()),
		SenderType:   borrower.ModelName(),
		ReceiverType: "",
		ReceiverID:   0,
		Amount:       amount,
		isDeposit:    true,
	}
}

func CreateDepositTransaction(borrower Borrower, amount int) *Transaction {
	depositTransaction := NewDepositTransaction(borrower, amount)
	depositTransaction.Save()
	return depositTransaction
}

func (instance *Transaction) Create() *gorm.DB {
	return database.GetDB().Create(instance)
}

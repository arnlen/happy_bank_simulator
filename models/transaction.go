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
	From   interface{}
	To     interface{}
	Amount int
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

func NewTransaction(from Actor, to Actor, amount int) *Transaction {
	from.UpdateBalance(-amount)
	to.UpdateBalance(amount)

	return &Transaction{
		From:   from,
		To:     to,
		Amount: amount,
	}
}

func (instance *Transaction) Create() *gorm.DB {
	return database.GetDB().Create(instance)
}

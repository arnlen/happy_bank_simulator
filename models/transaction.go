package models

import (
	"happy_bank_simulator/database"
	"log"

	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	From   interface{}
	To     interface{}
	Amount float64
}

func (instance *Transaction) ModelName() string {
	return "transaction"
}

func (instance *Transaction) Save() *Transaction {
	result := database.GetDB().Save(instance)

	if instance.ID == 0 || result.RowsAffected == 0 {
		log.Fatal(result.Error)
	}

	return instance
}

func NewTransaction(from interface{}, to interface{}, amount float64) *Transaction {
	return &Transaction{
		From:   from,
		To:     to,
		Amount: amount,
	}
}

func (instance *Transaction) Create() *gorm.DB {
	return database.GetDB().Create(instance)
}

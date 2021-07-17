package models

import (
	"happy_bank_simulator/app/configs"
	"happy_bank_simulator/database"
	"log"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"syreclabs.com/go/faker"
)

type Lender struct {
	gorm.Model
	Name    string
	Loans   []*Loan `gorm:"many2many:loan_lenders;"`
	Balance int
}

func (instance *Lender) ModelName() string {
	return "prêteur"
}

func ListLenders() []*Lender {
	var lenders []*Lender
	database.GetDB().Preload(clause.Associations).Find(&lenders)
	return lenders
}

func (instance *Lender) Save() *Lender {
	result := database.GetDB().Save(instance)

	if instance.ID == 0 || result.RowsAffected == 0 {
		log.Fatal(result.Error)
	}

	return instance
}

func NewLender(name string, balance int) *Lender {
	return &Lender{
		Name:    name,
		Loans:   []*Loan{},
		Balance: balance,
	}
}

func NewDefaultLender() *Lender {
	return &Lender{
		Name:    faker.Name().Name(),
		Loans:   []*Loan{},
		Balance: configs.Lender.InitialBalance,
	}
}

func CreateLender(name string, balance int) *Lender {
	lender := NewLender(name, balance)
	result := database.GetDB().Create(&lender)

	if lender.ID == 0 || result.RowsAffected == 0 {
		log.Fatal(result.Error)
	}

	return lender
}

package models

import (
	"happy_bank_simulator/app/configs"
	"happy_bank_simulator/database"
	"log"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"syreclabs.com/go/faker"
)

type Insurer struct {
	gorm.Model
	Name    string
	Loans   []*Loan `gorm:"many2many:loan_insurers;"`
	Balance int
}

func (instance *Insurer) ModelName() string {
	return "assureur"
}

func ListInsurers() []*Insurer {
	var insurers []*Insurer
	database.GetDB().Preload(clause.Associations).Find(&insurers)
	return insurers
}

func (instance *Insurer) Save() *Insurer {
	result := database.GetDB().Save(instance)

	if instance.ID == 0 || result.RowsAffected == 0 {
		log.Fatal(result.Error)
	}

	return instance
}

func NewInsurer(name string, balance int) *Insurer {
	return &Insurer{
		Name:    name,
		Loans:   []*Loan{},
		Balance: balance,
	}
}

func NewDefaultInsurer() *Insurer {
	return &Insurer{
		Name:    faker.Name().Name(),
		Loans:   []*Loan{},
		Balance: configs.Insurer.InitialBalance,
	}
}

func CreateInsurer(name string, balance int) *Insurer {
	insurer := NewInsurer(name, balance)
	result := database.GetDB().Create(&insurer)

	if insurer.ID == 0 || result.RowsAffected == 0 {
		log.Fatal(result.Error)
	}

	return insurer
}

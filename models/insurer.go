package models

import (
	"fmt"
	"happy_bank_simulator/app/configs"
	"happy_bank_simulator/database"
	"log"
	"strconv"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"syreclabs.com/go/faker"
)

// Declare conformity with Actor interface
var _ ModelBase = (*Insurer)(nil)
var _ Actor = (*Insurer)(nil)

type Insurer struct {
	gorm.Model
	Name    string
	Loans   []*Loan `gorm:"many2many:loan_insurers;"`
	Balance int
}

// ------- Instance methods -------

func (instance *Insurer) ModelName() string {
	return "assureur"
}

func (instance *Insurer) Refresh() {
	database.GetDB().Preload(clause.Associations).Find(&instance)
}

func (instance *Insurer) Save() {
	result := database.GetDB().Save(instance)

	if instance.ID == 0 || result.RowsAffected == 0 {
		log.Fatal(result.Error)
	}

	instance.Refresh()
}

func (instance *Insurer) UpdateBalance(amount int) {
	instance.Balance += amount
	instance.Save()
}

func (instance *Insurer) GetID() uint {
	return instance.ID
}

// ------- Package methods -------

func ListInsurers() []*Insurer {
	var insurers []*Insurer
	database.GetDB().Preload(clause.Associations).Find(&insurers)
	return insurers
}

func ListInsurersWithPositiveBalance() []*Insurer {
	insurers := ListInsurers()
	var insurersWithPositiveBalance []*Insurer
	for _, insurer := range insurers {
		if insurer.Balance > 0 {
			insurersWithPositiveBalance = append(insurersWithPositiveBalance, insurer)
		}
	}
	fmt.Printf("%s insurers with a positive balance\n", strconv.Itoa(len(insurersWithPositiveBalance)))
	return insurersWithPositiveBalance
}

func ListInsurersWithoutLoan(insurers []*Insurer) []*Insurer {
	var availableInsurersWithoutLoan []*Insurer
	for _, insurer := range insurers {
		if len(insurer.Loans) == 0 {
			availableInsurersWithoutLoan = append(availableInsurersWithoutLoan, insurer)
		}
	}
	fmt.Printf("%s insurers without any loans are available\n", strconv.Itoa(len(availableInsurersWithoutLoan)))
	return availableInsurersWithoutLoan
}

func ListInsurersWithLoanOtherThan(insurers []*Insurer, loan *Loan) []*Insurer {
	var availableInsurersWithLoan []*Insurer
	for _, insurer := range insurers {
		if len(insurer.Loans) != 0 {
			for _, insurerLoan := range insurer.Loans {
				if insurerLoan.ID != loan.ID {
					availableInsurersWithLoan = append(availableInsurersWithLoan, insurer)
				}
			}
		}
	}
	fmt.Printf("%s insurers wit loans different than the current one are available\n", strconv.Itoa(len(availableInsurersWithLoan)))
	return availableInsurersWithLoan
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

func CreateDefaultInsurer() *Insurer {
	insurer := NewDefaultInsurer()
	insurer.Save()
	return insurer
}

func CreateInsurer(name string, balance int) *Insurer {
	insurer := NewInsurer(name, balance)
	result := database.GetDB().Create(&insurer)

	if insurer.ID == 0 || result.RowsAffected == 0 {
		log.Fatal(result.Error)
	}

	return insurer
}

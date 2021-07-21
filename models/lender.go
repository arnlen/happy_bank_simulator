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
var _ ModelBase = (*Lender)(nil)
var _ Actor = (*Lender)(nil)

type Lender struct {
	gorm.Model
	Name    string
	Loans   []*Loan `gorm:"many2many:loan_lenders;"`
	Balance int
}

// ------- Instance methods -------

func (instance *Lender) ModelName() string {
	return "prêteur"
}

func (instance *Lender) Refresh() {
	database.GetDB().Preload(clause.Associations).Find(&instance)
}

func (instance *Lender) Save() {
	result := database.GetDB().Save(instance)

	if instance.ID == 0 || result.RowsAffected == 0 {
		log.Fatal(result.Error)
	}

	instance.Refresh()
}

func (instance *Lender) UpdateBalance(amount int) {
	instance.Balance += amount
	instance.Save()
}

// ------- Package methods -------

func ListLenders() []*Lender {
	var lenders []*Lender
	database.GetDB().Preload(clause.Associations).Find(&lenders)
	return lenders
}

func ListLendersWithPositiveBalance() []*Lender {
	lenders := ListLenders()
	var lendersWithPositiveBalance []*Lender
	for _, lender := range lenders {
		if lender.Balance > 0 {
			lendersWithPositiveBalance = append(lendersWithPositiveBalance, lender)
		}
	}
	fmt.Printf("%s lenders with a positive balance\n", strconv.Itoa(len(lendersWithPositiveBalance)))
	return lendersWithPositiveBalance
}

func ListLendersWithoutLoan(lenders []*Lender) []*Lender {
	var availableLendersWithoutLoan []*Lender
	for _, lender := range lenders {
		if len(lender.Loans) == 0 {
			availableLendersWithoutLoan = append(availableLendersWithoutLoan, lender)
		}
	}
	fmt.Printf("%s lenders without any loans are available\n", strconv.Itoa(len(availableLendersWithoutLoan)))
	return availableLendersWithoutLoan
}

func ListLendersWithLoanOtherThan(lenders []*Lender, loan *Loan) []*Lender {
	var availableLendersWithLoan []*Lender
	for _, lender := range lenders {
		if len(lender.Loans) != 0 {
			for _, lenderLoan := range lender.Loans {
				if lenderLoan.ID != loan.ID {
					availableLendersWithLoan = append(availableLendersWithLoan, lender)
				}
			}
		}
	}
	fmt.Printf("%s lenders wit loans different than the current one are available\n", strconv.Itoa(len(availableLendersWithLoan)))
	return availableLendersWithLoan
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

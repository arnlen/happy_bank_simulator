package models

import (
	"log"

	"happy_bank_simulator/app/configs"
	"happy_bank_simulator/database"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"syreclabs.com/go/faker"
)

// Declare conformity with interfaces
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

func (instance *Lender) GetID() uint {
	return instance.ID
}

// ------- Package methods -------

func ListLenders() []*Lender {
	var lenders []*Lender
	database.GetDB().Preload(clause.Associations).Find(&lenders)
	return lenders
}

// Duplicate with Insurer: same method
func ListLendersWithPositiveBalance() []*Lender {
	lenders := ListLenders()
	var lendersWithPositiveBalance []*Lender
	for _, lender := range lenders {
		if lender.Balance > 0 {
			lendersWithPositiveBalance = append(lendersWithPositiveBalance, lender)
		}
	}
	return lendersWithPositiveBalance
}

// Duplicate with Insurer: same method
func ListLendersWithoutLoan(lenders []*Lender) []*Lender {
	var availableLendersWithoutLoan []*Lender
	for _, lender := range lenders {
		if len(lender.Loans) == 0 {
			availableLendersWithoutLoan = append(availableLendersWithoutLoan, lender)
		}
	}
	return availableLendersWithoutLoan
}

// Duplicate with Insurer: same method
func ListLendersWithLoanOtherThan(lenders []*Lender, loan *Loan) []*Lender {
	var availableLendersWithLoan []*Lender
	for _, lender := range lenders {
		if len(lender.Loans) != 0 {
			for _, lenderLoan := range lender.Loans {
				if lenderLoan.ID != loan.ID && !isLenderAlreadyInSlice(*lender, lenders) {
					availableLendersWithLoan = append(availableLendersWithLoan, lender)
				}
			}
		}
	}
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

func CreateDefaultLender() *Lender {
	lender := NewDefaultLender()
	lender.Save()
	return lender
}

func CreateLender(name string, balance int) *Lender {
	lender := NewLender(name, balance)
	result := database.GetDB().Create(&lender)

	if lender.ID == 0 || result.RowsAffected == 0 {
		log.Fatal(result.Error)
	}

	return lender
}

// Duplicate with Insurer: same method
func isLenderAlreadyInSlice(newLender Lender, lenders []*Lender) bool {
	for _, lender := range lenders {
		if lender.ID == newLender.ID {
			return true
		}
	}

	return false
}

//
// NOT USED ANYMORE - TO DELETE TO COMPLETE REFACTORING
// Refatored inside models.Actor
//

package models

import (
// "log"

// "happy_bank_simulator/app/configs"
// "happy_bank_simulator/database"

// "gorm.io/gorm"
// "gorm.io/gorm/clause"
// "syreclabs.com/go/faker"
)

// Declare conformity with interfaces
// var _ ModelBase = (*Lender)(nil)
// var _ Actor = (*Lender)(nil)

// type Lender struct {
// 	gorm.Model
// 	Name           string
// 	Loans          []*Loan `gorm:"many2many:loan_lenders;"`
// 	InitialBalance float64
// 	Balance        float64
// }

// ------- Instance methods -------

// func (instance *Lender) ModelName() string {
// 	return "lender"
// }

// func (instance *Lender) Refresh() {
// 	database.GetDB().Preload(clause.Associations).Find(&instance)
// }

// func (instance *Lender) Save() {
// 	result := database.GetDB().Save(instance)

// 	if instance.ID == 0 || result.RowsAffected == 0 {
// 		log.Fatal(result.Error)
// 	}

// 	instance.Refresh()
// }

// func (instance *Lender) UpdateBalance(amount float64) {
// 	instance.Balance += amount
// 	instance.Save()
// }

// func (instance *Lender) GetID() uint {
// 	return instance.ID
// }

// ------- Package methods -------

// func ListLenders() []*Lender {
// 	var lenders []*Lender
// 	database.GetDB().Preload(clause.Associations).Find(&lenders)
// 	return lenders
// }

// // Duplicate with Insurer: same method
// func ListLendersWithPositiveBalance() []*Lender {
// 	lenders := ListLenders()
// 	var lendersWithPositiveBalance []*Lender
// 	for _, lender := range lenders {
// 		if lender.Balance > 0 {
// 			lendersWithPositiveBalance = append(lendersWithPositiveBalance, lender)
// 		}
// 	}
// 	return lendersWithPositiveBalance
// }

// // Duplicate with Insurer: same method
// func ListLendersWithoutLoan(lenders []*Lender) []*Lender {
// 	var availableLendersWithoutLoan []*Lender
// 	for _, lender := range lenders {
// 		if len(lender.Loans) == 0 {
// 			availableLendersWithoutLoan = append(availableLendersWithoutLoan, lender)
// 		}
// 	}
// 	return availableLendersWithoutLoan
// }

// // Duplicate with Insurer: same method
// func ListLendersWithLoanOtherThan(lenders []*Lender, loan *Loan) []*Lender {
// 	var availableLendersWithLoan []*Lender
// 	for _, lender := range lenders {
// 		if len(lender.Loans) != 0 {
// 			for _, lenderLoan := range lender.Loans {
// 				if lenderLoan.ID != loan.ID && !isLenderAlreadyInSlice(*lender, lenders) {
// 					availableLendersWithLoan = append(availableLendersWithLoan, lender)
// 				}
// 			}
// 		}
// 	}
// 	return availableLendersWithLoan
// }

// func NewLender(name string, balance float64) *Lender {
// 	return &Lender{
// 		Name:           name,
// 		Loans:          []*Loan{},
// 		InitialBalance: balance,
// 		Balance:        balance,
// 	}
// }

// func NewDefaultLender() *Lender {
// 	return &Lender{
// 		Name:           faker.Name().Name(),
// 		Loans:          []*Loan{},
// 		InitialBalance: configs.Lender.InitialBalance,
// 		Balance:        configs.Lender.InitialBalance,
// 	}
// }

// func CreateDefaultLender() *Lender {
// 	lender := NewDefaultLender()
// 	lender.Save()
// 	return lender
// }

// func CreateLender(name string, balance float64) *Lender {
// 	lender := NewLender(name, balance)
// 	result := database.GetDB().Create(&lender)

// 	if lender.ID == 0 || result.RowsAffected == 0 {
// 		log.Fatal(result.Error)
// 	}

// 	return lender
// }

// Duplicate with Insurer: same method
// func isLenderAlreadyInSlice(newLender Lender, lenders []*Lender) bool {
// 	for _, lender := range lenders {
// 		if lender.ID == newLender.ID {
// 			return true
// 		}
// 	}

// 	return false
// }

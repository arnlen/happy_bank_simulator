package models

import (
// "log"

// "happy_bank_simulator/app/configs"
// "happy_bank_simulator/database"

// "gorm.io/gorm"
// "gorm.io/gorm/clause"
// "syreclabs.com/go/faker"
)

// Declare conformity with Actor interface
// var _ ModelBase = (*Insurer)(nil)
// var _ Actor = (*Insurer)(nil)

// type Insurer struct {
// 	gorm.Model
// 	Name           string
// 	Loans          []*Loan `gorm:"many2many:loan_insurers;"`
// 	InitialBalance float64
// 	Balance        float64
// }

// ------- Instance methods -------

// func (instance *Insurer) ModelName() string {
// 	return "insurer"
// }

// func (instance *Insurer) Refresh() {
// 	database.GetDB().Preload(clause.Associations).Find(&instance)
// }

// func (instance *Insurer) Save() {
// 	result := database.GetDB().Save(instance)

// 	if instance.ID == 0 || result.RowsAffected == 0 {
// 		log.Fatal(result.Error)
// 	}

// 	instance.Refresh()
// }

// func (instance *Insurer) UpdateBalance(amount float64) {
// 	instance.Balance += amount
// 	instance.Save()
// }

// func (instance *Insurer) GetID() uint {
// 	return instance.ID
// }

// ------- Package methods -------

// func ListInsurers() []*Insurer {
// 	var insurers []*Insurer
// 	database.GetDB().Preload(clause.Associations).Find(&insurers)
// 	return insurers
// }

// Duplicate with Lender: same method
// func ListInsurersWithPositiveBalance() []*Insurer {
// 	insurers := ListInsurers()
// 	var insurersWithPositiveBalance []*Insurer
// 	for _, insurer := range insurers {
// 		if insurer.Balance > 0 {
// 			insurersWithPositiveBalance = append(insurersWithPositiveBalance, insurer)
// 		}
// 	}
// 	return insurersWithPositiveBalance
// }

// Duplicate with Lender: same method
// func ListInsurersWithoutLoan(insurers []*Insurer) []*Insurer {
// 	var availableInsurersWithoutLoan []*Insurer
// 	for _, insurer := range insurers {
// 		if len(insurer.Loans) == 0 {
// 			availableInsurersWithoutLoan = append(availableInsurersWithoutLoan, insurer)
// 		}
// 	}
// 	return availableInsurersWithoutLoan
// }

// Duplicate with Lender: same method
// func ListInsurersWithLoanOtherThan(insurers []*Insurer, loan *Loan) []*Insurer {
// 	var availableInsurersWithLoan []*Insurer
// 	for _, insurer := range insurers {
// 		if len(insurer.Loans) != 0 {
// 			for _, insurerLoan := range insurer.Loans {
// 				if insurerLoan.ID != loan.ID && !isInsurerAlreadyInSlice(*insurer, insurers) {
// 					availableInsurersWithLoan = append(availableInsurersWithLoan, insurer)
// 				}
// 			}
// 		}
// 	}
// 	return availableInsurersWithLoan
// }

// func NewInsurer(name string, balance float64) *Insurer {
// 	return &Insurer{
// 		Name:           name,
// 		Loans:          []*Loan{},
// 		InitialBalance: balance,
// 		Balance:        balance,
// 	}
// }

// func NewDefaultInsurer() *Insurer {
// 	return &Insurer{
// 		Name:           faker.Name().Name(),
// 		Loans:          []*Loan{},
// 		InitialBalance: configs.Insurer.InitialBalance,
// 		Balance:        configs.Insurer.InitialBalance,
// 	}
// }

// func CreateDefaultInsurer() *Insurer {
// 	insurer := NewDefaultInsurer()
// 	insurer.Save()
// 	return insurer
// }

// func CreateInsurer(name string, balance float64) *Insurer {
// 	insurer := NewInsurer(name, balance)
// 	result := database.GetDB().Create(&insurer)

// 	if insurer.ID == 0 || result.RowsAffected == 0 {
// 		log.Fatal(result.Error)
// 	}

// 	return insurer
// }

// // Duplicate with Lender: same method
// func isInsurerAlreadyInSlice(newInsurer Insurer, insurers []*Insurer) bool {
// 	for _, insurer := range insurers {
// 		if insurer.ID == newInsurer.ID {
// 			return true
// 		}
// 	}

// 	return false
// }

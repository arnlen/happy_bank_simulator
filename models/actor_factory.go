package models

import (
	"log"

	"happy_bank_simulator/app/configs"
	"happy_bank_simulator/internal/global"

	"syreclabs.com/go/faker"
)

func CreateBorrowers(quantity int) []*Actor {
	var borrowers []*Actor
	for i := 0; i < quantity; i++ {
		borrower := createDefaultActor("borrower")
		borrowers = append(borrowers, borrower)
	}
	return borrowers
}

func CreateBorrower() *Actor {
	return CreateBorrowers(1)[0]
}

func CreateBorrowersWithLoan(quantity int) []*Actor {
	var borrowers []*Actor
	for i := 0; i < quantity; i++ {
		borrower := CreateBorrower()
		loan := CreateLoan()
		borrower.AssignLoan(loan)
		borrowers = append(borrowers, borrower)
	}
	return borrowers
}

func CreateBorrowerWithLoan() *Actor {
	return CreateBorrowersWithLoan(1)[0]
}

func CreateLenders(quantity int) []*Actor {
	var lenders []*Actor
	for i := 0; i < quantity; i++ {
		lender := createDefaultActor("lender")
		lenders = append(lenders, lender)
	}
	return lenders
}

func CreateLender() *Actor {
	return CreateLenders(1)[0]
}

func CreateLendersWithLoan(quantity int) []*Actor {
	var lenders []*Actor
	for i := 0; i < quantity; i++ {
		lender := CreateLender()
		loan := CreateLoan()
		lender.AssignLoan(loan)
		lenders = append(lenders, lender)
	}
	return lenders
}

func CreateLenderWithLoan() *Actor {
	return CreateLendersWithLoan(1)[0]
}

func CreateInsurers(quantity int) []*Actor {
	var insurers []*Actor
	for i := 0; i < quantity; i++ {
		insurer := createDefaultActor("insurer")
		insurers = append(insurers, insurer)
	}
	return insurers
}

func CreateInsurer() *Actor {
	return CreateInsurers(1)[0]
}

func CreateInsurersWithLoan(quantity int) []*Actor {
	var insurers []*Actor
	for i := 0; i < quantity; i++ {
		insurer := CreateInsurer()
		loan := CreateLoan()
		insurer.AssignLoan(loan)
		insurers = append(insurers, insurer)
	}
	return insurers
}

func CreateInsurerWithLoan() *Actor {
	return CreateInsurersWithLoan(1)[0]
}

// ----- PRIVATE METHODS -----

func newActor(actorType string, name string, balance float64) *Actor {
	return &Actor{
		Name:           name,
		Loans:          []*Loan{},
		Balance:        balance,
		InitialBalance: balance,
		Type:           actorType,
	}
}

func createActor(actorType string, name string, balance float64) *Actor {
	actor := newActor(actorType, name, balance)
	result := global.Db.Create(&actor)

	if actor.ID == 0 || result.RowsAffected == 0 {
		log.Fatal(result.Error)
	}

	return actor
}

func newDefaultActor(actorType string) *Actor {
	name := faker.Name().Name()
	balance := configs.Actor.InitialBalance
	return newActor(actorType, name, balance)
}

func createDefaultActor(actorType string) *Actor {
	actor := newDefaultActor(actorType)
	actor.Save()
	return actor
}

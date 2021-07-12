package controllers

import (
	"fmt"
	"happy_bank_simulator/database"
	"happy_bank_simulator/factories"
	"happy_bank_simulator/initializers"
	"happy_bank_simulator/models"

	"gorm.io/gorm/clause"
)

// Declare conformity with BaseController interface
var _ BaseController = (*Borrowers)(nil)

type Overview struct{}

func (c *Overview) GetCounters() []int {
	var borrowers []models.Borrower
	var lenders []models.Lender
	var insurers []models.Insurer
	var loans []models.Loan

	database.GetDB().Preload(clause.Associations).Find(&borrowers)
	database.GetDB().Preload(clause.Associations).Find(&lenders)
	database.GetDB().Preload(clause.Associations).Find(&insurers)
	database.GetDB().Preload(clause.Associations).Find(&loans)

	return []int{len(borrowers), len(lenders), len(insurers), len(loans)}
}

func (c *Overview) PopulateDatabase() {
	factories.CreateSeedState()
	fmt.Println("Database populated")
}

func (c *Overview) WipeDatabase() {
	database.DropBD()
	fmt.Println("Database dropped")
	initializers.InitDB()
	fmt.Println("Database initialized")
}

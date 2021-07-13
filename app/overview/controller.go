package overview

import (
	"fmt"
	"happy_bank_simulator/factories"
	"happy_bank_simulator/models"
	"happy_bank_simulator/services/database"

	"gorm.io/gorm/clause"
)

// Declare conformity with BaseController interface
// var _ BaseController = (*Controller)(nil)

type Controller struct{}

func (c *Controller) GetCounters() []int {
	db := database.GetDB()

	var borrowers []models.Borrower
	var lenders []models.Lender
	var insurers []models.Insurer
	var loans []models.Loan

	db.Preload(clause.Associations).Find(&borrowers)
	db.Preload(clause.Associations).Find(&lenders)
	db.Preload(clause.Associations).Find(&insurers)
	db.Preload(clause.Associations).Find(&loans)

	return []int{len(borrowers), len(lenders), len(insurers), len(loans)}
}

func (c *Controller) PopulateDatabase() {
	factories.CreateSeedState()
	fmt.Println("Database populated")
}

func (c *Controller) WipeDatabase() {
	database.DropBD()
	fmt.Println("Database dropped")
	database.InitDB()
	fmt.Println("Database initialized")
	database.MigrateDB()
	fmt.Println("Database migrated")
}

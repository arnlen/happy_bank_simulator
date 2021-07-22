package overview

import (
	"fmt"
	"happy_bank_simulator/database"
	databaseHelpers "happy_bank_simulator/database/helpers"
	"happy_bank_simulator/models"

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
	fmt.Println("Deactivated")
}

func (c *Controller) WipeDatabase() {
	databaseHelpers.DropBD()
	fmt.Println("Database dropped")
	database.InitDB()
	fmt.Println("Database initialized")
	databaseHelpers.MigrateDB()
	fmt.Println("Database migrated")
}

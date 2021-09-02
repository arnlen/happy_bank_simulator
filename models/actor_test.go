package models_test

import (
	"testing"

	"happy_bank_simulator/app/configs"
	"happy_bank_simulator/internal/database"
	"happy_bank_simulator/models"

	"github.com/stretchr/testify/assert"
)

func TestActor_NetBalance(t *testing.T) {
	database.ResetDB()
	assert := assert.New(t)

	borrowerWithLoan := models.CreateBorrowerWithLoan()
	lenderWithLoan := models.CreateLenderWithLoan()
	insurerWithLoan := models.CreateInsurerWithLoan()

	assert.Equal(borrowerWithLoan.NetBalance(), borrowerWithLoan.Balance-borrowerWithLoan.Loans[0].Amount)
	assert.Equal(lenderWithLoan.NetBalance(), lenderWithLoan.Balance)
	assert.Equal(insurerWithLoan.NetBalance(), insurerWithLoan.Balance-insurerWithLoan.Loans[0].Amount)
}

func TestActor_TotalAmountAssigned(t *testing.T) {
	database.ResetDB()
	assert := assert.New(t)

	borrowerWithLoan := models.CreateBorrowerWithLoan()

	assert.Equal(borrowerWithLoan.TotalAmountAssigned(), borrowerWithLoan.Loans[0].Amount)
}

func TestActorFactory_UpdateBalance(t *testing.T) {
	database.ResetDB()
	assert := assert.New(t)

	actor := models.CreateInsurer()
	actor.Balance = 0
	actor.UpdateBalance(1000.0)

	assert.Equal(actor.Balance, 1000.0, "Balance of the actor should be increased be 1000")

	actor.UpdateBalance(-500.0)

	assert.Equal(actor.Balance, 500.0, "Balance of the actor should be decreased be 500")
}

func TestActorFactory_UpdateMontlyIncomes(t *testing.T) {
	database.ResetDB()
	assert := assert.New(t)

	// Bad case when actor isn't a borrower
	// => Not implemented

	// Good case when actor is a borrower
	borrower := models.CreateBorrower()
	borrower.UpdateMontlyIncomes(1000)

	assert.Equal(1000.0, borrower.MonthlyIncomes)
}

func TestActorFactory_AssignLoan(t *testing.T) {
	database.ResetDB()
	assert := assert.New(t)

	borrower := models.CreateBorrower()
	loan := models.CreateLoan()
	borrower.AssignLoan(loan)

	assert.Equal(loan.Borrower.ID, borrower.ID, "Loan's borrower ID match borrower's ID")
	assert.Equal(borrower.Loans[0].ID, loan.ID, "Borrower's loan ID match loan's ID")

	lender := models.CreateLender()
	loan = models.CreateLoan()
	lender.AssignLoan(loan)

	assert.Len(loan.Lenders, 1, "Loan should have one lender")
	assert.Equal(loan.Lenders[0].ID, lender.ID, "Loan's lender ID should be our loan ID")
	assert.Len(lender.Loans, 1, "Lender should have one loan")
	assert.Equal(lender.Loans[0].ID, loan.ID, "Lender's loan ID match loan's ID")

	insurer := models.CreateInsurer()
	loan = models.CreateLoan()
	insurer.AssignLoan(loan)

	assert.Len(loan.Insurers, 1, "Loan should have one insurer")
	assert.Equal(loan.Insurers[0].ID, insurer.ID, "Loan's insurer ID should be our loan ID")
	assert.Len(insurer.Loans, 1, "Insurer should have one loan")
	assert.Equal(insurer.Loans[0].ID, loan.ID, "Insurer's loan ID match loan's ID")
}

func TestActorFactory_CanTakeThisLoan(t *testing.T) {
	database.ResetDB()
	assert := assert.New(t)

	targetLoan := models.CreateLoan()
	targetLoan.Amount = 1000
	targetLoan.Save()
	anotherLoan := models.CreateLoan()
	anotherLoan.Amount = 1000
	anotherLoan.Save()

	// ----------------------------
	// Actor: Borrower
	//
	borrower := models.CreateBorrower()

	// Invalid case 1: When balance == 0
	borrower.UpdateBalance(-borrower.Balance)
	assert.False(borrower.CanTakeThisLoan(*targetLoan))

	// Invalid case 2: When balance leverage ratio is not met
	configs.Actor.BorrowerBalanceLeverageRatio = 1.0
	borrower.UpdateBalance(1000)
	borrower.AssignLoan(anotherLoan)
	assert.False(borrower.CanTakeThisLoan(*targetLoan))

	// Valid case
	borrower.UpdateBalance(1000)
	assert.True(borrower.CanTakeThisLoan(*targetLoan))

	// ----------------------------
	// Actor: Lender
	//
	lender := models.CreateLender()

	// Invalid case: when balance < amount to lend
	amountPerLender := targetLoan.GetAmountLentPerLender()
	lender.Balance = amountPerLender - 1
	lender.Save()
	assert.False(lender.CanTakeThisLoan(*targetLoan))

	// Valid case
	lender.UpdateBalance(1)
	assert.True(lender.CanTakeThisLoan(*targetLoan))

	// ----------------------------
	// Actor: Insurer
	//
	insurer := models.CreateInsurer()

	// Invalid case 1: when balance < amount to insure
	amountPerInsurer := targetLoan.GetAmountInsuredPerInsurer()
	insurer.Balance = amountPerInsurer - 1
	insurer.Save()
	assert.False(insurer.CanTakeThisLoan(*targetLoan))

	// Invalid case 2: when net balance < amount to insure
	insurer.UpdateBalance(1)
	insurer.AssignLoan(anotherLoan)
	insurer.Save()
	assert.False(insurer.CanTakeThisLoan(*targetLoan))

	// Valid case
	insurer.UpdateBalance(amountPerInsurer)
	assert.True(insurer.CanTakeThisLoan(*targetLoan))
}

func TestActorFactory_IsBorrower(t *testing.T) {
	database.ResetDB()
	assert := assert.New(t)

	borrower := models.CreateBorrower()

	assert.True(borrower.IsBorrower())
	assert.False(borrower.IsLender())
	assert.False(borrower.IsInsurer())
}

func TestActorFactory_IsLender(t *testing.T) {
	database.ResetDB()
	assert := assert.New(t)

	lender := models.CreateLender()

	assert.False(lender.IsBorrower())
	assert.True(lender.IsLender())
	assert.False(lender.IsInsurer())
}

func TestActorFactory_IsInsurer(t *testing.T) {
	database.ResetDB()
	assert := assert.New(t)

	insurer := models.CreateInsurer()

	assert.False(insurer.IsBorrower())
	assert.False(insurer.IsLender())
	assert.True(insurer.IsInsurer())
}

func TestActorFactory_ListActors(t *testing.T) {
	database.ResetDB()
	assert := assert.New(t)

	models.CreateBorrowers(3)
	models.CreateLenders(4)
	models.CreateInsurers(5)

	borrowers := models.ListActors(configs.Actor.BorrowerString)
	lenders := models.ListActors(configs.Actor.LenderString)
	insurers := models.ListActors(configs.Actor.InsurerString)

	assert.Equal(3, len(borrowers))
	assert.True(borrowers[0].IsBorrower())

	assert.Equal(4, len(lenders))
	assert.True(lenders[0].IsLender())

	assert.Equal(5, len(insurers))
	assert.True(insurers[0].IsInsurer())
}

func TestActorFactory_ListActorsWithPositiveBalance(t *testing.T) {
	database.ResetDB()
	assert := assert.New(t)

	borrowerWithNullBalance := models.CreateBorrower()
	borrowerWithNullBalance.UpdateBalance(-borrowerWithNullBalance.Balance)
	borrowersWithPositiveBalance := models.CreateBorrowers(2)
	borrowers := models.ListActorsWithPositiveBalance(configs.Actor.BorrowerString)

	lenderWithNullBalance := models.CreateLender()
	lenderWithNullBalance.UpdateBalance(-lenderWithNullBalance.Balance)
	lendersWithPositiveBalance := models.CreateLenders(3)
	lenders := models.ListActorsWithPositiveBalance(configs.Actor.LenderString)

	insurerWithNullBalance := models.CreateInsurer()
	insurerWithNullBalance.UpdateBalance(-insurerWithNullBalance.Balance)
	insurersWithPositiveBalance := models.CreateInsurers(4)
	insurers := models.ListActorsWithPositiveBalance(configs.Actor.InsurerString)

	assert.Equal(2, len(borrowers))
	assert.True(borrowers[0].IsBorrower())
	assert.Equal(borrowersWithPositiveBalance[0].ID, borrowers[0].ID)

	assert.Equal(3, len(lenders))
	assert.True(lenders[0].IsLender())
	assert.Equal(lendersWithPositiveBalance[0].ID, lenders[0].ID)

	assert.Equal(4, len(insurers))
	assert.True(insurers[0].IsInsurer())
	assert.Equal(insurersWithPositiveBalance[0].ID, insurers[0].ID)
}

func TestActorFactory_ListActorsWithoutLoan(t *testing.T) {
	database.ResetDB()
	assert := assert.New(t)

	models.CreateBorrowers(2)
	models.CreateBorrowersWithLoan(4)
	borrowers := models.ListActorsWithoutLoan(configs.Actor.BorrowerString)

	models.CreateLenders(3)
	models.CreateLendersWithLoan(3)
	lenders := models.ListActorsWithoutLoan(configs.Actor.LenderString)

	models.CreateInsurers(4)
	models.CreateInsurersWithLoan(2)
	insurers := models.ListActorsWithoutLoan(configs.Actor.InsurerString)

	assert.Equal(2, len(borrowers))
	assert.Equal(3, len(lenders))
	assert.Equal(4, len(insurers))
}

func TestActorFactory_ListActorsWithLoan(t *testing.T) {
	database.ResetDB()
	assert := assert.New(t)

	models.CreateBorrowers(2)
	models.CreateBorrowersWithLoan(4)
	borrowers := models.ListActorsWithLoan(configs.Actor.BorrowerString)

	models.CreateLenders(3)
	models.CreateLendersWithLoan(3)
	lenders := models.ListActorsWithLoan(configs.Actor.LenderString)

	models.CreateInsurers(4)
	models.CreateInsurersWithLoan(2)
	insurers := models.ListActorsWithLoan(configs.Actor.InsurerString)

	assert.Equal(4, len(borrowers))
	assert.Equal(3, len(lenders))
	assert.Equal(2, len(insurers))
}

func TestActorFactory_ListActorsWithLoanOtherThanTarget(t *testing.T) {
	database.ResetDB()
	assert := assert.New(t)

	loan := models.CreateLoan()

	borrowersWithLoan := models.CreateBorrowersWithLoan(3)
	borrowersWithLoan[0].AssignLoan(loan)
	borrowers := models.ListActorsWithLoanOtherThanTarget(configs.Actor.BorrowerString, loan)

	lendersWithLoan := models.CreateLendersWithLoan(4)
	lendersWithLoan[0].AssignLoan(loan)
	lenders := models.ListActorsWithLoanOtherThanTarget(configs.Actor.LenderString, loan)

	insurersWithLoan := models.CreateInsurersWithLoan(5)
	insurersWithLoan[0].AssignLoan(loan)
	insurers := models.ListActorsWithLoanOtherThanTarget(configs.Actor.InsurerString, loan)

	assert.Equal(2, len(borrowers))
	assert.Equal(3, len(lenders))
	assert.Equal(4, len(insurers))
}

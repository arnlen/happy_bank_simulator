package models_test

import (
	"testing"

	"happy_bank_simulator/app/configs"
	"happy_bank_simulator/factories"
	"happy_bank_simulator/internal/database"
	"happy_bank_simulator/models"

	"github.com/stretchr/testify/assert"
)

func TestActor_NetBalance(t *testing.T) {
	database.ResetDB()
	assert := assert.New(t)

	borrowerWithLoan := factories.NewBorrowerWithLoan()
	lenderWithLoan := factories.NewLenderWithLoan()
	insurerWithLoan := factories.NewInsurerWithLoan()

	assert.Equal(borrowerWithLoan.NetBalance(), borrowerWithLoan.Balance-borrowerWithLoan.Loans[0].Amount)
	assert.Equal(lenderWithLoan.NetBalance(), lenderWithLoan.Balance)
	assert.Equal(insurerWithLoan.NetBalance(), insurerWithLoan.Balance-insurerWithLoan.Loans[0].Amount)
}

func TestActor_TotalAmountAssigned(t *testing.T) {
	database.ResetDB()
	assert := assert.New(t)

	borrowerWithLoan := factories.NewBorrowerWithLoan()

	assert.Equal(borrowerWithLoan.TotalAmountAssigned(), borrowerWithLoan.Loans[0].Amount)
}

func TestActorFactory_UpdateBalance(t *testing.T) {
	database.ResetDB()
	assert := assert.New(t)

	actor := factories.NewInsurer()
	actor.Balance = 0
	actor.UpdateBalance(1000.0)

	assert.Equal(actor.Balance, 1000.0, "Balance of the actor should be increased be 1000")

	actor.UpdateBalance(-500.0)

	assert.Equal(actor.Balance, 500.0, "Balance of the actor should be decreased be 500")
}

func TestActorFactory_UpdateMontlyIncomes(t *testing.T) {
	database.ResetDB()
	assert := assert.New(t)

	borrower := factories.NewBorrower()
	borrower.UpdateMontlyIncomes(1000)

	assert.Equal(borrower.MonthlyIncomes, 1000.0, "Borrower's monthly income should be updated")
}

func TestActorFactory_AssignLoan(t *testing.T) {
	database.ResetDB()
	assert := assert.New(t)

	borrower := factories.NewBorrower()
	loan := factories.NewLoan()
	borrower.AssignLoan(loan)

	assert.Equal(loan.Borrower.ID, borrower.ID, "Loan's borrower ID match borrower's ID")
	assert.Equal(borrower.Loans[0].ID, loan.ID, "Borrower's loan ID match loan's ID")

	lender := factories.NewLender()
	loan = factories.NewLoan()
	lender.AssignLoan(loan)

	assert.Equal(len(loan.Lenders), 1, "Loan should have one lender")
	assert.Equal(loan.Lenders[0].ID, lender.ID, "Loan's lender ID should be our loan ID")
	assert.Equal(len(lender.Loans), 1, "Lender should have one loan")
	assert.Equal(lender.Loans[0].ID, loan.ID, "Lender's loan ID match loan's ID")

	insurer := factories.NewInsurer()
	loan = factories.NewLoan()
	insurer.AssignLoan(loan)

	assert.Equal(len(loan.Insurers), 1, "Loan should have one insurer")
	assert.Equal(loan.Insurers[0].ID, insurer.ID, "Loan's insurer ID should be our loan ID")
	assert.Equal(len(insurer.Loans), 1, "Insurer should have one loan")
	assert.Equal(insurer.Loans[0].ID, loan.ID, "Insurer's loan ID match loan's ID")
}

func TestActorFactory_CanTakeThisLoan(t *testing.T) {
	database.ResetDB()
	assert := assert.New(t)

	targetLoan := factories.NewLoan()
	targetLoan.Amount = 1000
	targetLoan.Save()
	anotherLoan := factories.NewLoan()
	anotherLoan.Amount = 1000
	anotherLoan.Save()

	// ----------------------------
	// Actor: Borrower
	//
	borrower := factories.NewBorrower()

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
	lender := factories.NewLender()

	// Invalid case: when balance < amount to lend
	amountPerLender := targetLoan.AmountPerLender()
	lender.Balance = amountPerLender - 1
	lender.Save()
	assert.False(lender.CanTakeThisLoan(*targetLoan))

	// Valid case
	lender.UpdateBalance(1)
	assert.True(lender.CanTakeThisLoan(*targetLoan))

	// ----------------------------
	// Actor: Insurer
	//
	insurer := factories.NewInsurer()

	// Invalid case 1: when balance < amount to insure
	amountPerInsurer := targetLoan.AmountPerInsurer()
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

func TestActorFactory_ListActors(t *testing.T) {
	database.ResetDB()
	assert := assert.New(t)

	factories.NewBorrowers(3)
	factories.NewLenders(4)
	factories.NewInsurers(5)

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

	borrowerWithNullBalance := factories.NewBorrower()
	borrowerWithNullBalance.UpdateBalance(-borrowerWithNullBalance.Balance)
	borrowersWithPositiveBalance := factories.NewBorrowers(2)
	borrowers := models.ListActorsWithPositiveBalance(configs.Actor.BorrowerString)

	lenderWithNullBalance := factories.NewLender()
	lenderWithNullBalance.UpdateBalance(-lenderWithNullBalance.Balance)
	lendersWithPositiveBalance := factories.NewLenders(3)
	lenders := models.ListActorsWithPositiveBalance(configs.Actor.LenderString)

	insurerWithNullBalance := factories.NewInsurer()
	insurerWithNullBalance.UpdateBalance(-insurerWithNullBalance.Balance)
	insurersWithPositiveBalance := factories.NewInsurers(4)
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

	factories.NewBorrowers(2)
	factories.NewBorrowersWithLoan(4)
	borrowers := models.ListActorsWithoutLoan(configs.Actor.BorrowerString)

	factories.NewLenders(3)
	factories.NewLendersWithLoan(3)
	lenders := models.ListActorsWithoutLoan(configs.Actor.LenderString)

	factories.NewInsurers(4)
	factories.NewInsurersWithLoan(2)
	insurers := models.ListActorsWithoutLoan(configs.Actor.InsurerString)

	assert.Equal(2, len(borrowers))
	assert.Equal(3, len(lenders))
	assert.Equal(4, len(insurers))
}

func TestActorFactory_ListActorsWithLoan(t *testing.T) {
	database.ResetDB()
	assert := assert.New(t)

	factories.NewBorrowers(2)
	factories.NewBorrowersWithLoan(4)
	borrowers := models.ListActorsWithLoan(configs.Actor.BorrowerString)

	factories.NewLenders(3)
	factories.NewLendersWithLoan(3)
	lenders := models.ListActorsWithLoan(configs.Actor.LenderString)

	factories.NewInsurers(4)
	factories.NewInsurersWithLoan(2)
	insurers := models.ListActorsWithLoan(configs.Actor.InsurerString)

	assert.Equal(4, len(borrowers))
	assert.Equal(3, len(lenders))
	assert.Equal(2, len(insurers))
}

func TestActorFactory_ListActorsWithLoanOtherThanTarget(t *testing.T) {
	database.ResetDB()
	assert := assert.New(t)

	loan := factories.NewLoan()

	borrowersWithLoan := factories.NewBorrowersWithLoan(3)
	borrowersWithLoan[0].AssignLoan(loan)
	borrowers := models.ListActorsWithLoanOtherThanTarget(configs.Actor.BorrowerString, loan)

	lendersWithLoan := factories.NewLendersWithLoan(4)
	lendersWithLoan[0].AssignLoan(loan)
	lenders := models.ListActorsWithLoanOtherThanTarget(configs.Actor.LenderString, loan)

	insurersWithLoan := factories.NewInsurersWithLoan(5)
	insurersWithLoan[0].AssignLoan(loan)
	insurers := models.ListActorsWithLoanOtherThanTarget(configs.Actor.InsurerString, loan)

	assert.Equal(2, len(borrowers))
	assert.Equal(3, len(lenders))
	assert.Equal(4, len(insurers))
}

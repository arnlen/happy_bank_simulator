package simulation

import (
	"fmt"
	"math"
	"strconv"

	"happy_bank_simulator/app/configs"
	"happy_bank_simulator/models"
)

func calculateLendersQuantityRequired(amount float64) int {
	maxAmountPerBorrower := configs.Actor.MaxAmountPerLoan
	quantity := int(math.Ceil(float64(amount) / float64(maxAmountPerBorrower)))
	fmt.Printf("%s lenders are required\n", strconv.Itoa(quantity))
	return quantity
}

func calculateInsurersQuantityRequired(amount float64) int {
	maxAmountPerLoan := configs.Actor.MaxAmountPerLoan
	return int(math.Ceil(float64(amount) / float64(maxAmountPerLoan)))
}

func calculateAmountToLendForLender(amount float64) float64 {
	lenderQuantity := calculateLendersQuantityRequired(amount)
	return amount / float64(lenderQuantity)
}

func canThisBorrowerTakeThisLoan(borrower *models.Actor, loan *models.Loan) bool {
	loans := borrower.Loans
	totalAmountBorrowed := borrower.GetTotalAmountBorrowed()
	fmt.Printf("Borrower #%s has %s loans for a total of %1.2f €\n", strconv.Itoa(int(borrower.ID)), strconv.Itoa(len(loans)), totalAmountBorrowed)

	borrowerNetBalance := borrower.GetNetBalance()
	fmt.Printf("Borrower #%s net balance is %1.2f €\n", strconv.Itoa(int(borrower.ID)), borrowerNetBalance)

	loanAmount := loan.Amount
	balanceLeverageRatio := configs.Actor.BorrowerBalanceLeverageRatio

	ratio := float64(borrowerNetBalance / loanAmount)

	fmt.Println("Balance leverage ratio set to", balanceLeverageRatio)
	fmt.Println("Ratio net balance / loan amont =", ratio)

	if ratio >= balanceLeverageRatio {
		fmt.Printf("Borrower #%s can take the loan #%s\n", strconv.Itoa(int(borrower.ID)), strconv.Itoa(int(loan.ID)))
		return true
	} else {
		fmt.Printf("Borrower #%s cannot take the loan #%s: net balance to low\n", strconv.Itoa(int(borrower.ID)), strconv.Itoa(int(loan.ID)))
		return false
	}
}

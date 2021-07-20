package simulation

import (
	"fmt"
	"math"
	"strconv"

	"happy_bank_simulator/app/configs"
	"happy_bank_simulator/models"
)

func calculateLendersQuantityRequired(amount int) int {
	maxAmountPerBorrower := configs.Lender.MaxAmountPerLoan
	quantity := int(math.Ceil(float64(amount) / float64(maxAmountPerBorrower)))
	fmt.Printf("%s lenders are required\n", strconv.Itoa(quantity))
	return quantity
}

func calculateInsurersQuantityRequired(amount int) int {
	maxAmountPerLoan := configs.Insurer.MaxAmountPerLoan
	return int(math.Ceil(float64(amount) / float64(maxAmountPerLoan)))
}

func canThisBorrowerTakeThisLoan(borrower *models.Borrower, loan *models.Loan) bool {
	borrowerNetBalance := borrower.GetNetBalance()
	loanAmount := loan.Amount
	balanceLeverageRatio := configs.Borrower.BalanceLeverageRatio

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

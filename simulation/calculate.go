package simulation

import (
	"fmt"
	"math"
	"strconv"

	"happy_bank_simulator/app/configs"
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

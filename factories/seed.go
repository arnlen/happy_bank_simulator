package factories

import (
	"fmt"
	"happy_bank_simulator/models"
)

func CreateSeedState() {
	var lenders []*models.Lender
	var borrowers []*models.Borrower
	var insurers []*models.Insurer
	var loans []*models.Loan

	for i := 0; i < 10; i++ {
		name := fmt.Sprintf("Lender %d", i+1)
		lenders = append(lenders, models.CreateLender(name, 10000))
	}

	for i := 0; i < 10; i++ {
		name := fmt.Sprintf("Borrower %d", i+1)
		borrowers = append(borrowers, models.CreateBorrower(name, 10000))
	}

	for i := 0; i < 10; i++ {
		name := fmt.Sprintf("Insurer %d", i+1)
		insurers = append(insurers, models.CreateInsurer(name, 10000))
	}

	for i := 0; i < 10; i++ {
		startDate := "27/06/2021"
		endDate := "27/06/2022"
		duration := 12
		amount := 10000
		loans = append(loans, models.CreateLoan(startDate, endDate, int32(duration), float64(amount), borrowers[i], lenders[i], insurers[i]))
	}
}

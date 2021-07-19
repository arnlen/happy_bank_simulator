package simulation

import (
	"fmt"
	"strconv"

	"happy_bank_simulator/models"
)

func getLendersWithPositiveBalance() []*models.Lender {
	lenders := models.ListLenders()
	var lendersWithPositiveBalance []*models.Lender
	for _, lender := range lenders {
		if lender.Balance > 0 {
			lendersWithPositiveBalance = append(lendersWithPositiveBalance, lender)
		}
	}
	fmt.Printf("%s lenders with a positive balance\n", strconv.Itoa(len(lendersWithPositiveBalance)))
	return lendersWithPositiveBalance
}

func getInsurersWithPositiveBalance() []*models.Insurer {
	insurers := models.ListInsurers()
	var insurersWithPositiveBalance []*models.Insurer
	for _, insurer := range insurers {
		if insurer.Balance > 0 {
			insurersWithPositiveBalance = append(insurersWithPositiveBalance, insurer)
		}
	}
	fmt.Printf("%s insurers with a positive balance\n", strconv.Itoa(len(insurersWithPositiveBalance)))
	return insurersWithPositiveBalance
}

func getLendersWithoutLoan(lenders []*models.Lender) []*models.Lender {
	var availableLendersWithoutLoan []*models.Lender
	for _, lender := range lenders {
		if len(lender.Loans) == 0 {
			availableLendersWithoutLoan = append(availableLendersWithoutLoan, lender)
		}
	}
	fmt.Printf("%s lenders without any loans are available\n", strconv.Itoa(len(availableLendersWithoutLoan)))
	return availableLendersWithoutLoan
}

func getInsurersWithoutLoan(insurers []*models.Insurer) []*models.Insurer {
	var availableInsurersWithoutLoan []*models.Insurer
	for _, insurer := range insurers {
		if len(insurer.Loans) == 0 {
			availableInsurersWithoutLoan = append(availableInsurersWithoutLoan, insurer)
		}
	}
	fmt.Printf("%s insurers without any loans are available\n", strconv.Itoa(len(availableInsurersWithoutLoan)))
	return availableInsurersWithoutLoan
}

func getLendersWithLoanOtherThan(lenders []*models.Lender, loan *models.Loan) []*models.Lender {
	var availableLendersWithLoan []*models.Lender
	for _, lender := range lenders {
		if len(lender.Loans) != 0 {
			for _, lenderLoan := range lender.Loans {
				if lenderLoan.ID != loan.ID {
					availableLendersWithLoan = append(availableLendersWithLoan, lender)
				}
			}
		}
	}
	fmt.Printf("%s lenders wit loans different than the current one are available\n", strconv.Itoa(len(availableLendersWithLoan)))
	return availableLendersWithLoan
}

func getInsurersWithLoanOtherThan(insurers []*models.Insurer, loan *models.Loan) []*models.Insurer {
	var availableInsurersWithLoan []*models.Insurer
	for _, insurer := range insurers {
		if len(insurer.Loans) != 0 {
			for _, insurerLoan := range insurer.Loans {
				if insurerLoan.ID != loan.ID {
					availableInsurersWithLoan = append(availableInsurersWithLoan, insurer)
				}
			}
		}
	}
	fmt.Printf("%s insurers wit loans different than the current one are available\n", strconv.Itoa(len(availableInsurersWithLoan)))
	return availableInsurersWithLoan
}

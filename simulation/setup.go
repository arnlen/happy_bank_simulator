package simulation

import (
	"fmt"
	"strconv"

	"happy_bank_simulator/app/configs"
	"happy_bank_simulator/models"
)

func setupBorrowerForLoan(loan *models.Loan) {
	borrower := createDefaultBorrower()
	assignBorrowerToLoan(borrower, loan)
}

func setupLendersForLoan(loan *models.Loan) {
	var availableLenders []*models.Lender
	defaultLoanAmount := configs.Loan.DefaultAmount

	lendersQuantityRequired := calculateLendersQuantityRequired(defaultLoanAmount)
	lendersWithPositiveBalance := getLendersWithPositiveBalance()

	lendersWithoutLoan := getLendersWithoutLoan(lendersWithPositiveBalance)
	availableLenders = append(availableLenders, lendersWithoutLoan...)

	if len(availableLenders) < lendersQuantityRequired {
		missingLendersQuantity := lendersQuantityRequired - len(availableLenders)
		fmt.Printf("Not enough available lenders: missing %s lenders\n", strconv.Itoa(missingLendersQuantity))
		fmt.Println("Trying to find available lenders inside lenders with already at least 1 loan")

		lendersWithLoan := getLendersWithLoanOtherThan(lendersWithPositiveBalance, loan)
		availableLenders = append(availableLenders, lendersWithLoan...)
	}

	fmt.Printf("%s total lenders available, including lender with other loans than the current one\n", strconv.Itoa(len(availableLenders)))

	if len(availableLenders) < lendersQuantityRequired {
		missingLendersQuantity := lendersQuantityRequired - len(availableLenders)
		fmt.Printf("Not enough available lenders: missing %s lenders\n", strconv.Itoa(missingLendersQuantity))
		fmt.Printf("Creating %s new lenders\n", strconv.Itoa(missingLendersQuantity))

		availableLenders = createMissingLenders(missingLendersQuantity, availableLenders)
	}

	assignLendersToLoan(availableLenders, loan)
}

func setupInsurersForLoan(loan *models.Loan) {
	var availableInsurers []*models.Insurer
	defaultLoanAmount := configs.Loan.DefaultAmount

	insurersQuantityRequired := calculateInsurersQuantityRequired(defaultLoanAmount)
	InsurersWithPositiveBalance := getInsurersWithPositiveBalance()

	insurersWithoutLoan := getInsurersWithoutLoan(InsurersWithPositiveBalance)
	availableInsurers = append(availableInsurers, insurersWithoutLoan...)

	if len(availableInsurers) < insurersQuantityRequired {
		missingInsurersQuantity := insurersQuantityRequired - len(availableInsurers)
		fmt.Printf("Not enough available Insurers: missing %s Insurers\n", strconv.Itoa(missingInsurersQuantity))
		fmt.Println("Trying to find available Insurers inside Insurers with already at least 1 loan")

		InsurersWithLoan := getInsurersWithLoanOtherThan(InsurersWithPositiveBalance, loan)
		availableInsurers = append(availableInsurers, InsurersWithLoan...)
	}

	fmt.Printf("%s total Insurers available, including Insurer with other loans than the current one\n", strconv.Itoa(len(availableInsurers)))

	if len(availableInsurers) < insurersQuantityRequired {
		missingInsurersQuantity := insurersQuantityRequired - len(availableInsurers)
		fmt.Printf("Not enough available Insurers: missing %s Insurers\n", strconv.Itoa(missingInsurersQuantity))
		fmt.Printf("Creating %s new Insurers\n", strconv.Itoa(missingInsurersQuantity))

		availableInsurers = createMissingInsurers(missingInsurersQuantity, availableInsurers)
	}

	assignInsurersToLoan(availableInsurers, loan)
}

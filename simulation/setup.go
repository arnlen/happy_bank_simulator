package simulation

import (
	"fmt"
	"strconv"

	"happy_bank_simulator/app/configs"
	"happy_bank_simulator/models"
)

func setupLendersForLoan(loan *models.Loan) {
	var availableLenders []*models.Actor
	defaultLoanAmount := configs.Loan.DefaultAmount

	lendersQuantityRequired := calculateLendersQuantityRequired(defaultLoanAmount)
	lendersWithPositiveBalance := models.ListActorsWithPositiveBalance(configs.Actor.LenderString)
	fmt.Printf("%s lenders with a positive balance\n", strconv.Itoa(len(lendersWithPositiveBalance)))

	lendersWithoutLoan := models.ListActorsWithoutLoan(configs.Actor.LenderString)
	fmt.Printf("%s lenders without any loans are available\n", strconv.Itoa(len(lendersWithoutLoan)))
	availableLenders = append(availableLenders, lendersWithoutLoan...)

	if len(availableLenders) < lendersQuantityRequired {
		missingLendersQuantity := lendersQuantityRequired - len(availableLenders)
		fmt.Printf("Not enough available lenders: missing %s lenders\n", strconv.Itoa(missingLendersQuantity))
		fmt.Println("Trying to find available lenders inside lenders with already at least 1 loan")

		lendersWithLoan := models.ListActorsWithLoanOtherThanTarget(configs.Actor.LenderString, loan)
		fmt.Printf("%s lenders wit loans different than the current one are available\n", strconv.Itoa(len(lendersWithLoan)))
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
	var availableInsurers []*models.Actor
	defaultLoanAmount := configs.Loan.DefaultAmount

	insurersQuantityRequired := calculateInsurersQuantityRequired(defaultLoanAmount)
	insurersWithPositiveBalance := models.ListActorsWithPositiveBalance(configs.Actor.InsurerString)
	fmt.Printf("%s insurers with a positive balance\n", strconv.Itoa(len(insurersWithPositiveBalance)))

	insurersWithoutLoan := models.ListActorsWithoutLoan(configs.Actor.InsurerString)
	fmt.Printf("%s insurers without any loans are available\n", strconv.Itoa(len(insurersWithoutLoan)))
	availableInsurers = append(availableInsurers, insurersWithoutLoan...)

	if len(availableInsurers) < insurersQuantityRequired {
		missingInsurersQuantity := insurersQuantityRequired - len(availableInsurers)
		fmt.Printf("Not enough available Insurers: missing %s Insurers\n", strconv.Itoa(missingInsurersQuantity))
		fmt.Println("Trying to find available Insurers inside Insurers with already at least 1 loan")

		insurersWithLoan := models.ListActorsWithLoanOtherThanTarget(configs.Actor.InsurerString, loan)
		fmt.Printf("%s insurers with loans different than the current one are available\n", strconv.Itoa(len(insurersWithLoan)))
		availableInsurers = append(availableInsurers, insurersWithLoan...)
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

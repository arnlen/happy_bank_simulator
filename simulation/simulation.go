package simulation

import (
	"fmt"
	"math"
	"strconv"

	"happy_bank_simulator/app/configs"
	"happy_bank_simulator/database/helpers"
	"happy_bank_simulator/models"
)

var quantityOfLoansToCreate int

func Prepare() {
	helpers.DropBD()
	helpers.MigrateDB()
	createLoans()
}

// TODO
//
// For Loan
// - Determine if the loan should be insured
//
// For borrower
// - Determine if the borrower can take this loan (BalanceLeverageRatio)
// - Place the initial security deposit
// - Determine if the borrower should fail
//
// For lenders and insurers
// - Take their balance into account

func createLoans() {
	quantityOfLoansToCreate = configs.Loan.InitialQuantity

	defaultLoanAmount := configs.Loan.DefaultAmount
	fmt.Println("Quantity of Loans to create:", quantityOfLoansToCreate)

	for i := 0; i < quantityOfLoansToCreate; i++ {
		// Create empty loan
		var currentLoan = models.NewDefaultLoan()
		currentLoan.Save()
		fmt.Printf("%s - Loan #%s created\n", strconv.Itoa(i+1), strconv.Itoa(int(currentLoan.ID)))

		// Create one new default borrower
		borrower := models.NewDefaultBorrower()
		borrower.Save()
		fmt.Printf("Borrower #%s created\n", strconv.Itoa(int(borrower.ID)))

		// Assign borrower to loan
		currentLoan.Borrower = *borrower
		currentLoan.Save()
		borrower.Refresh()
		fmt.Printf("Borrower assigned: Loan #%s's borrower = #%s\n", strconv.Itoa(int(currentLoan.ID)), strconv.Itoa(int(currentLoan.BorrowerID)))
		fmt.Printf("Borrower #%s has now %s loans\n", strconv.Itoa(int(borrower.ID)), strconv.Itoa(int(len(borrower.Loans))))

		// ------ LENDERS --------

		// How many lenders required for this loan?
		lendersQuantityRequired := calculateLendersQuantityRequired(defaultLoanAmount)
		fmt.Printf("%s lenders are required\n", strconv.Itoa(lendersQuantityRequired))

		// Find or create lenders

		// 1. Find all lenders + prepare slice of available lenders
		lenders := models.ListLenders()
		fmt.Printf("%s lenders in the system\n", strconv.Itoa(len(lenders)))
		var availableLenders []*models.Lender

		// 2. Split lenders betweens those with loans and those without
		var lendersWithLoan []*models.Lender
		for _, lender := range lenders {
			if len(lender.Loans) == 0 {
				availableLenders = append(availableLenders, &lender)
			} else {
				lendersWithLoan = append(lendersWithLoan, &lender)
			}
		}
		lenderAvailableQuantity := len(availableLenders)
		fmt.Printf("%s lenders without any loans are available\n", strconv.Itoa(lenderAvailableQuantity))

		if lenderAvailableQuantity < lendersQuantityRequired {
			missingLendersQuantity := lendersQuantityRequired - lenderAvailableQuantity
			fmt.Printf("Not enough available lenders: missing %s lenders\n", strconv.Itoa(missingLendersQuantity))
			fmt.Println("Trying to find available lenders inside lenders with already at least 1 loan")

			// 3. Within lenders with loan, check which are still available for the current loan
			for _, lenderWithLoan := range lendersWithLoan {
				for _, lenderLoan := range lenderWithLoan.Loans {
					if lenderLoan.ID != currentLoan.ID {
						availableLenders = append(availableLenders, lenderWithLoan)
					}
				}
			}
			lenderAvailableQuantity = len(availableLenders)
			fmt.Printf("%s total lenders available, including lender with other loans than the current one\n", strconv.Itoa(lenderAvailableQuantity))
		}

		if lenderAvailableQuantity < lendersQuantityRequired {
			missingLendersQuantity := lendersQuantityRequired - lenderAvailableQuantity
			fmt.Printf("Not enough available lenders: missing %s lenders\n", strconv.Itoa(missingLendersQuantity))
			fmt.Printf("Creating %s new lenders\n", strconv.Itoa(missingLendersQuantity))

			for i := 0; i < missingLendersQuantity; i++ {
				lender := models.NewDefaultLender()
				lender.Save()
				availableLenders = append(availableLenders, lender)
				fmt.Printf("%s/%s - Lender #%s created\n", strconv.Itoa(i+1), strconv.Itoa(missingLendersQuantity), strconv.Itoa(int(lender.ID)))
			}

			lenderAvailableQuantity = len(availableLenders)
			fmt.Printf("%s total lenders now available\n", strconv.Itoa(lenderAvailableQuantity))
		}

		// Assign lenders to loan
		fmt.Printf("Assigning those %s lenders to loan #%s\n", strconv.Itoa(lenderAvailableQuantity), strconv.Itoa(int(currentLoan.ID)))
		for _, availableLender := range availableLenders {
			currentLoan.AddLender(availableLender)
			fmt.Printf("- Lender %s assigned\n", strconv.Itoa(int(availableLender.ID)))
		}

		// currentLoan.Refresh()
		currentLoanLenders := currentLoan.Lenders
		fmt.Printf("Loan #%s has now %s lenders\n", strconv.Itoa(int(currentLoan.ID)), strconv.Itoa(len(currentLoanLenders)))

		// ------ INSURERS --------

		// How many insurers?
		insurersQuantityRequired := calculateInsurersQuantityRequired(defaultLoanAmount)
		fmt.Printf("%s insurers are required\n", strconv.Itoa(insurersQuantityRequired))

		// 1. Find all insurers + prepare slice of available insurers
		insurers := models.ListInsurers()
		fmt.Printf("%s insurers in the system\n", strconv.Itoa(len(insurers)))
		var availableInsurers []*models.Insurer

		// 2. Split insurers betweens those with loans and those without
		var insurersWithLoan []*models.Insurer
		for _, insurer := range insurers {
			if len(insurer.Loans) == 0 {
				availableInsurers = append(availableInsurers, &insurer)
			} else {
				insurersWithLoan = append(insurersWithLoan, &insurer)
			}
		}
		insurerAvailableQuantity := len(availableInsurers)
		fmt.Printf("%s insurers without any loans are available\n", strconv.Itoa(insurerAvailableQuantity))

		if insurerAvailableQuantity < insurersQuantityRequired {
			missingInsurersQuantity := insurersQuantityRequired - insurerAvailableQuantity
			fmt.Printf("Not enough available insurers: missing %s insurers\n", strconv.Itoa(missingInsurersQuantity))
			fmt.Println("Trying to find available insurers inside insurers with already at least 1 loan")

			// 3. Within insurers with loan, check which are still available for the current loan
			for _, insurerWithLoan := range insurersWithLoan {
				for _, insurerLoan := range insurerWithLoan.Loans {
					if insurerLoan.ID != currentLoan.ID {
						availableInsurers = append(availableInsurers, insurerWithLoan)
					}
				}
			}
			insurerAvailableQuantity = len(availableInsurers)
			fmt.Printf("%s total insurers available, including insurer with other loans than the current one\n", strconv.Itoa(insurerAvailableQuantity))
		}

		if insurerAvailableQuantity < insurersQuantityRequired {
			missingInsurersQuantity := insurersQuantityRequired - insurerAvailableQuantity
			fmt.Printf("Not enough available insurers: missing %s insurers\n", strconv.Itoa(missingInsurersQuantity))
			fmt.Printf("Creating %s new insurers\n", strconv.Itoa(missingInsurersQuantity))

			for i := 0; i < missingInsurersQuantity; i++ {
				insurer := models.NewDefaultInsurer()
				insurer.Save()
				availableInsurers = append(availableInsurers, insurer)
				fmt.Printf("%s/%s - insurer #%s created\n", strconv.Itoa(i+1), strconv.Itoa(missingInsurersQuantity), strconv.Itoa(int(insurer.ID)))
			}

			insurerAvailableQuantity = len(availableInsurers)
			fmt.Printf("%s total insurers now available\n", strconv.Itoa(insurerAvailableQuantity))
		}

		// Assign insurers to loan
		fmt.Printf("Assigning those %s insurers to loan #%s\n", strconv.Itoa(insurerAvailableQuantity), strconv.Itoa(int(currentLoan.ID)))
		for _, availableInsurer := range availableInsurers {
			currentLoan.AddInsurer(availableInsurer)
			fmt.Printf("- Insurer %s assigned\n", strconv.Itoa(int(availableInsurer.ID)))
		}

		// currentLoan.Refresh()
		currentLoaninsurers := currentLoan.Insurers
		fmt.Printf("Loan #%s has now %s insurers\n", strconv.Itoa(int(currentLoan.ID)), strconv.Itoa(len(currentLoaninsurers)))

		// ------ SUMMARY --------

		fmt.Printf("Summary for Loan #%s:\n", strconv.Itoa(int(currentLoan.ID)))
		fmt.Printf("- 1 borrower: %s (#%s)\n", currentLoan.Borrower.Name, strconv.Itoa(int(currentLoan.Borrower.ID)))
		fmt.Printf("- %s lenders:\n", strconv.Itoa(len(currentLoan.Lenders)))
		for _, lender := range currentLoan.Lenders {
			fmt.Printf("--- %s (#%s)\n", lender.Name, strconv.Itoa(int(lender.ID)))
		}
		fmt.Printf("- %s insurers:\n", strconv.Itoa(len(currentLoan.Insurers)))
		for _, insurer := range currentLoan.Insurers {
			fmt.Printf("--- %s (#%s)\n", insurer.Name, strconv.Itoa(int(insurer.ID)))
		}
	}
}

func calculateLendersQuantityRequired(amount int) int {
	maxAmountPerBorrower := configs.Lender.MaxAmountPerLoan
	return int(math.Ceil(float64(amount) / float64(maxAmountPerBorrower)))
}

func calculateInsurersQuantityRequired(amount int) int {
	maxAmountPerLoan := configs.Insurer.MaxAmountPerLoan
	return int(math.Ceil(float64(amount) / float64(maxAmountPerLoan)))
}

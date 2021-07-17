package simulation

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"time"

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
// For borrower
// - Determine if the borrower can take this loan (BalanceLeverageRatio)
// - Place the initial security deposit
// - Determine when the borrower will fail
//

func createLoans() {
	quantityOfLoansToCreate = configs.Loan.InitialQuantity
	fmt.Println("Quantity of Loans to create:", quantityOfLoansToCreate)

	for i := 0; i < quantityOfLoansToCreate; i++ {
		loan := createEmptyLoan()

		setupBorrowerForLoan(loan)
		setupLendersForLoan(loan)

		isThisLoanInsured := positiveForProbability(configs.Loan.InsuredQuantityRatio)
		if isThisLoanInsured {
			fmt.Println("This loan is insured")
			setupInsurersForLoan(loan)
		} else {
			fmt.Println("This loan is NOT insured 🚨")
		}

		printSummaryForLoan(*loan)
	}
}

func createEmptyLoan() *models.Loan {
	var loan = models.NewDefaultLoan()
	loan.Save()
	fmt.Printf("Loan #%s created\n", strconv.Itoa(int(loan.ID)))
	return loan
}

func createDefaultBorrower() *models.Borrower {
	borrower := models.NewDefaultBorrower()
	borrower.Save()
	fmt.Printf("Borrower #%s created\n", strconv.Itoa(int(borrower.ID)))
	return borrower
}

func assignBorrowerToLoan(borrower *models.Borrower, loan *models.Loan) {
	loan.Borrower = *borrower
	loan.Save()
	borrower.Refresh()
	fmt.Printf("Borrower assigned: Loan #%s's borrower = #%s\n", strconv.Itoa(int(loan.ID)), strconv.Itoa(int(loan.BorrowerID)))
	fmt.Printf("Borrower #%s has now %s loans\n", strconv.Itoa(int(borrower.ID)), strconv.Itoa(int(len(borrower.Loans))))
}

func setupBorrowerForLoan(loan *models.Loan) {
	borrower := createDefaultBorrower()
	willItFail := positiveForProbability(configs.Borrower.FailureRate)
	if willItFail {
		fmt.Println("This borrower will fail ❌")
		// TODO: add date of failure to model
	} else {
		fmt.Println("This borrower is strong")
	}
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

func createMissingLenders(missingQuantity int, availableLenders []*models.Lender) []*models.Lender {
	for i := 0; i < missingQuantity; i++ {
		lender := models.NewDefaultLender()
		lender.Save()
		availableLenders = append(availableLenders, lender)
		fmt.Printf("%s/%s - Lender #%s created\n", strconv.Itoa(i+1), strconv.Itoa(missingQuantity), strconv.Itoa(int(lender.ID)))
	}
	fmt.Printf("%s total lenders now available\n", strconv.Itoa(len(availableLenders)))
	return availableLenders
}

func createMissingInsurers(missingQuantity int, availableInsurers []*models.Insurer) []*models.Insurer {
	for i := 0; i < missingQuantity; i++ {
		insurer := models.NewDefaultInsurer()
		insurer.Save()
		availableInsurers = append(availableInsurers, insurer)
		fmt.Printf("%s/%s - Insurer #%s created\n", strconv.Itoa(i+1), strconv.Itoa(missingQuantity), strconv.Itoa(int(insurer.ID)))
	}
	fmt.Printf("%s total insurers now available\n", strconv.Itoa(len(availableInsurers)))
	return availableInsurers
}

func assignLendersToLoan(lenders []*models.Lender, loan *models.Loan) {
	fmt.Printf("Assigning %s lenders to loan #%s\n", strconv.Itoa(len(lenders)), strconv.Itoa(int(loan.ID)))
	for _, availableLender := range lenders {
		loan.AddLender(availableLender)
		fmt.Printf("- Lender #%s assigned\n", strconv.Itoa(int(availableLender.ID)))
	}
	fmt.Printf("Loan #%s has now %s lenders\n", strconv.Itoa(int(loan.ID)), strconv.Itoa(len(loan.Lenders)))
}

func assignInsurersToLoan(insurers []*models.Insurer, loan *models.Loan) {
	fmt.Printf("Assigning %s insurers to loan #%s\n", strconv.Itoa(len(insurers)), strconv.Itoa(int(loan.ID)))
	for _, availableInsurer := range insurers {
		loan.AddInsurer(availableInsurer)
		fmt.Printf("- Insurer #%s assigned\n", strconv.Itoa(int(availableInsurer.ID)))
	}
	fmt.Printf("Loan #%s has now %s insurers\n", strconv.Itoa(int(loan.ID)), strconv.Itoa(len(loan.Insurers)))
}

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

func printSummaryForLoan(loan models.Loan) {
	fmt.Printf("Summary for Loan #%s:\n", strconv.Itoa(int(loan.ID)))
	fmt.Printf("- 1 borrower: %s (#%s)\n", loan.Borrower.Name, strconv.Itoa(int(loan.Borrower.ID)))
	fmt.Printf("- %s lenders:\n", strconv.Itoa(len(loan.Lenders)))
	for _, lender := range loan.Lenders {
		fmt.Printf("--- %s (#%s)\n", lender.Name, strconv.Itoa(int(lender.ID)))
	}
	fmt.Printf("- %s insurers:\n", strconv.Itoa(len(loan.Insurers)))
	for _, insurer := range loan.Insurers {
		fmt.Printf("--- %s (#%s)\n", insurer.Name, strconv.Itoa(int(insurer.ID)))
	}
}

func positiveForProbability(probability float64) bool {
	probability = probability * 100

	rand.Seed(time.Now().UnixNano())
	randomNumber := rand.Intn(100)

	if randomNumber < int(probability) {
		return true
	} else {
		return false
	}
}

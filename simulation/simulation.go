package simulation

import (
	"fmt"
	"math"
	"strconv"

	"happy_bank_simulator/app/configs"
	"happy_bank_simulator/database/helpers"
	"happy_bank_simulator/models"

	"syreclabs.com/go/faker"
)

var (
	quantityOfLoansToCreate     int
	quantityOfBorrowersToCreate int
	lendersQuantityRequired     int
	quantityOfInsurersToCreate  int
)

func Prepare() {
	helpers.DropBD()
	helpers.MigrateDB()
	createLoans()
}

func createLoans() {
	quantityOfLoansToCreate = configs.Loan.InitialQuantity
	quantityOfBorrowersToCreate = quantityOfLoansToCreate

	defaultLoanAmount := configs.Loan.DefaultAmount
	fmt.Println("Quantity of Loans to create:", quantityOfLoansToCreate)

	for i := 0; i < quantityOfLoansToCreate; i++ {
		// Create empty loan
		var loan = models.NewDefaultLoan()
		loan.Save()
		fmt.Printf("%s - Loan #%s created\n", strconv.Itoa(i+1), strconv.Itoa(int(loan.ID)))

		// Create one new default borrower
		borrower := models.NewDefaultBorrower()
		borrower.Save()
		fmt.Printf("Borrower #%s created\n", strconv.Itoa(int(borrower.ID)))

		// Assign borrower to loan
		loan.Borrower = *borrower
		loan.Save()
		borrower.Refresh()
		fmt.Printf("Borrower assigned: Loan #%s's borrower = #%s\n", strconv.Itoa(int(loan.ID)), strconv.Itoa(int(loan.BorrowerID)))
		fmt.Printf("Borrower #%s has now %s loans\n", strconv.Itoa(int(borrower.ID)), strconv.Itoa(int(len(borrower.Loans))))

		// How many lenders required for this loan?
		lendersQuantityRequired := calculateLendersQuantityRequired(defaultLoanAmount)
		fmt.Println("lendersQuantityRequired", lendersQuantityRequired)

		// Find or create lenders
		lendersWithoutLoans := models.ListLendersWithoutLoan()
		lenders := models.ListLenders()
		currentLenderQuantity := len(lenders)
		fmt.Println("lendersWithoutLoans", len(lendersWithoutLoans))
		fmt.Println("currentLenderQuantity", currentLenderQuantity)

		if currentLenderQuantity < lendersQuantityRequired {
			lenderQuantityToCreate := lendersQuantityRequired - currentLenderQuantity
			fmt.Printf("Missing %s lenders\n", strconv.Itoa(lenderQuantityToCreate))

			for i := 0; i < lenderQuantityToCreate; i++ {
				lender := models.NewDefaultLender()
				lender.Save()
				fmt.Printf("Lender #%s created\n", strconv.Itoa(int(lender.ID)))
			}

			lenders = models.ListLenders()
			currentLenderQuantity := len(lenders)
			fmt.Println("New currentLenderQuantity", currentLenderQuantity)
		}

		// Assign lenders to loan
		// for lenders

		// How many insurers?

		// Find or create insurers

		// Determine if the borrower should fail

		// Update loan with borrower, lenders and insurers
	}
}

func prepareLoanCreation() {
	loanConfig := configs.Loan
	defaultLoanAmount := loanConfig.DefaultAmount

	quantityOfInsurersToCreate := calculateRequiredInsurersFor(quantityOfLoansToCreate)

	fmt.Println("defaultLoanAmount:", defaultLoanAmount)
	fmt.Println("quantityOfLoansToCreate:", quantityOfLoansToCreate)
	fmt.Println("quantityOfBorrowersToCreate:", quantityOfBorrowersToCreate)
	fmt.Println("lendersQuantityRequired:", lendersQuantityRequired)
	fmt.Println("quantityOfInsurersToCreate:", quantityOfInsurersToCreate)

	createBorrowers(quantityOfBorrowersToCreate)
	createLenders(lendersQuantityRequired)
	createInsurers(quantityOfInsurersToCreate)
}

func calculateLendersQuantityRequired(amount int) int {
	maxAmountPerBorrower := configs.Lender.MaxAmountPerLoan
	return int(math.Ceil(float64(amount) / float64(maxAmountPerBorrower)))
}

func calculateRequiredInsurersFor(amount int) int {
	insuredLoansQuantityRate := configs.Loan.InsuredQuantityRate
	return int(math.Ceil(float64(amount) * float64(insuredLoansQuantityRate)))
}

func createBorrowers(quantityToCreate int) {
	for i := 0; i < quantityToCreate; i++ {
		models.CreateBorrower(faker.Name().Name(), configs.Borrower.InitialBalance)
	}
	fmt.Println(quantityToCreate, "borrowers created")
}

func createLenders(quantityToCreate int) {
	for i := 0; i < quantityToCreate; i++ {
		models.CreateLender(faker.Name().Name(), configs.Lender.InitialBalance)
	}
	fmt.Println(quantityToCreate, "lenders created")
}

func createInsurers(quantityToCreate int) {
	for i := 0; i < quantityToCreate; i++ {
		models.CreateInsurer(faker.Name().Name(), configs.Insurer.InitialBalance)
	}
	fmt.Println(quantityToCreate, "insurers created")
}

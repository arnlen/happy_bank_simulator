package simulation

import (
	"fmt"
	"strconv"

	"happy_bank_simulator/app/configs"
	"happy_bank_simulator/factories"
	"happy_bank_simulator/helpers"
	"happy_bank_simulator/models"
)

func createInitialLoans() []*models.Loan {
	loans := factories.NewLoans(configs.Loan.InitialQuantity)
	fmt.Printf("Initial loans created: %s loans\n", strconv.Itoa(len(loans)))

	for _, loan := range loans {
		fmt.Printf("Loan #%s setup:\n", strconv.Itoa(int(loan.ID)))

		isThisLoanInsured := helpers.GetResultForProbability(configs.Loan.InsuredQuantityRatio)
		if isThisLoanInsured {
			fmt.Println("- This loan is insured")
			loan.IsInsured = true
		} else {
			fmt.Println("- This loan is NOT insured 🚨")
			loan.IsInsured = false
		}

		willThisLoanFail := helpers.GetResultForProbability(configs.Loan.FailureRate)
		if willThisLoanFail {
			fmt.Println("- This loan will fail 🚨")
			loan.SetRandomNumberOfMonthsBeforeFailure()
			fmt.Printf("- The failure will occure after %s months, on %s\n",
				strconv.Itoa(loan.NumberOfMonthsBeforeFailure), loan.WillFailOnString())
		}

		loans = append(loans, loan)
	}
	return loans
}

package factories

import (
	"happy_bank_simulator/database"
	"happy_bank_simulator/models"

	"happy_bank_simulator/services"

	"syreclabs.com/go/faker"
)

func CreateSeedState() {
	db := database.GetDB()

	startDate := "27/06/2021"
	endDate := "27/06/2022"
	duration := 12
	amount := 10000
	initialDeposit := amount / 10
	creditRate := 0.3
	insuranceRate := 0.3
	monthlyCredit := services.CalculateMonthlyCreditPayment(creditRate, float64(duration), float64(amount))
	monthlyInsurance := services.CalculateMonthlyInsurancePayment(insuranceRate, float64(duration), float64(amount))

	for i := 0; i < 10; i++ {
		db.Create(&models.Loan{
			Borrower:         models.Borrower{Name: faker.Name().Name(), Balance: float64(faker.Number().NumberInt(5))},
			Lender:           models.Lender{Name: faker.Name().Name(), Balance: float64(faker.Number().NumberInt(5))},
			Insurer:          models.Insurer{Name: faker.Name().Name(), Balance: float64(faker.Number().NumberInt(5))},
			StartDate:        startDate,
			EndDate:          endDate,
			Duration:         int32(duration),
			Amount:           float64(amount),
			InitialDeposit:   float64(initialDeposit),
			CreditRate:       creditRate,
			InsuranceRate:    insuranceRate,
			MonthlyCredit:    monthlyCredit,
			MonthlyInsurance: monthlyInsurance,
		})
	}
}

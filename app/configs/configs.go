package configs

import (
	"fmt"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// ----- GENERAL CONFIGS -----

type general struct {
	StartDate             string  // Simulation start date
	Duration              int     // Simulation duration (in months)
	CreditInterestRate    float64 // Interest rate of the credit part of a loan
	InsuranceInterestRate float64 // Interest rate of the insurance part of a loan
}

var General = general{
	StartDate:             "07/2022",
	Duration:              60,
	CreditInterestRate:    0.03,
	InsuranceInterestRate: 0.03,
}

// ----- LOAN CONFIGS -----

type loan struct {
	InitialQuantity      int     // How many loans should exist at the beginning at the simulation
	DefaultAmount        float64 // Default amount for a new loan
	DefaultDuration      int     // Default duration for a new loan
	SecurityDepositRate  float64 // For a given loan amout, how much % a borrower must stake
	InsuredQuantityRatio float64 // How many loans are insured, in % of the total
	FailureRate          float64 // How many loans should fail, in % of the total
	String               string
}

var Loan = loan{
	InitialQuantity:      1,
	DefaultAmount:        5000,
	DefaultDuration:      24,
	SecurityDepositRate:  0.1,
	InsuredQuantityRatio: 1,
	FailureRate:          1,
	String:               "loan",
}

// ----- ACTOR CONFIGS -----

type actor struct {
	MaxAmountPerLoan             float64 // Maximum amout of money a lender of an insurer can lend/insure per loan
	InitialBalance               float64 // Initial balance
	BorrowerBalanceLeverageRatio float64 // Ratio between the balance of the borrower and the amount he can borrow
	BorrowerString               string
	LenderString                 string
	InsurerString                string
}

var Actor = actor{
	MaxAmountPerLoan:             1000,
	InitialBalance:               5000,
	BorrowerBalanceLeverageRatio: 1.0,
	BorrowerString:               "borrower",
	LenderString:                 "lender",
	InsurerString:                "insurer",
}

// ----- TRANSACTION CONFIGS -----

type transaction struct {
	String string
}

var Transaction = transaction{
	String: "transaction",
}

// ----- RENDER -----

func RenderEdit() *fyne.Container {
	var (
		labels           []*widget.Label
		entries          []*widget.Entry
		borderContainers []*fyne.Container
	)

	// General fields

	startDateLabel := widget.NewLabel("Date de démarrage de la simulation")
	startDateEntry := widget.NewEntry()
	startDateEntry.SetText(General.StartDate)
	startDateEntry.OnChanged = func(value string) {
		startDate := value
		General.StartDate = startDate
		fmt.Println("General.StartDate updated to", value)
	}
	labels = append(labels, startDateLabel)
	entries = append(entries, startDateEntry)

	durationLabel := widget.NewLabel("Durée de la simulation (en mois)")
	durationEntry := widget.NewEntry()
	durationEntry.SetText(strconv.Itoa(General.Duration))
	durationEntry.OnChanged = func(value string) {
		duration, _ := strconv.Atoi(value)
		General.Duration = duration
		fmt.Println("General.Duration updated to", value)
	}
	labels = append(labels, durationLabel)
	entries = append(entries, durationEntry)

	creditInterestRateLabel := widget.NewLabel("Taux d'intérêt du crédit")
	creditInterestRateEntry := widget.NewEntry()
	creditInterestRateEntry.SetText(fmt.Sprintf("%1.2f", General.CreditInterestRate))
	creditInterestRateEntry.OnChanged = func(value string) {
		creditInterestRate, _ := strconv.ParseFloat(value, 64)
		General.CreditInterestRate = creditInterestRate
		fmt.Println("General.CreditInterestRate updated to", value)
	}
	labels = append(labels, creditInterestRateLabel)
	entries = append(entries, creditInterestRateEntry)

	insuranceInterestRateLabel := widget.NewLabel("Taux d'intérêt de l'assurance")
	insuranceInterestRateEntry := widget.NewEntry()
	insuranceInterestRateEntry.SetText(fmt.Sprintf("%1.2f", General.InsuranceInterestRate))
	insuranceInterestRateEntry.OnChanged = func(value string) {
		insuranceInterestRate, _ := strconv.ParseFloat(value, 64)
		General.InsuranceInterestRate = insuranceInterestRate
		fmt.Println("General.InsuranceInterestRate updated to", value)
	}
	labels = append(labels, insuranceInterestRateLabel)
	entries = append(entries, insuranceInterestRateEntry)

	// Loan fields

	initialLoanQuantityLabel := widget.NewLabel("Nombre d'emprunts au démarrage de la simulation")
	initialLoanQuantityEntry := widget.NewEntry()
	initialLoanQuantityEntry.SetText(strconv.Itoa(Loan.InitialQuantity))
	initialLoanQuantityEntry.OnChanged = func(value string) {
		initialLoanQuantity, _ := strconv.Atoi(value)
		Loan.InitialQuantity = initialLoanQuantity
		fmt.Println("Loan.InitialQuantity updated to", value)
	}
	labels = append(labels, initialLoanQuantityLabel)
	entries = append(entries, initialLoanQuantityEntry)

	loanDefaultAmountLabel := widget.NewLabel("Montant des emprunts par défaut")
	loanDefaultAmountEntry := widget.NewEntry()
	loanDefaultAmountEntry.SetText(fmt.Sprintf("%1.2f", Loan.DefaultAmount))
	loanDefaultAmountEntry.OnChanged = func(value string) {
		loanDefaultAmount, _ := strconv.ParseFloat(value, 64)
		Loan.DefaultAmount = loanDefaultAmount
		fmt.Println("Loan.DefaultAmount updated to", value)
	}
	labels = append(labels, loanDefaultAmountLabel)
	entries = append(entries, loanDefaultAmountEntry)

	defaultDurationLabel := widget.NewLabel("Durée par défaut d'un emprunt")
	defaultDurationEntry := widget.NewEntry()
	defaultDurationEntry.SetText(strconv.Itoa(Loan.DefaultDuration))
	defaultDurationEntry.OnChanged = func(value string) {
		defaultDuration, _ := strconv.Atoi(value)
		Loan.DefaultDuration = defaultDuration
		fmt.Println("Loan.DefaultDuration updated to", value)
	}
	labels = append(labels, defaultDurationLabel)
	entries = append(entries, defaultDurationEntry)

	securityDepositRateLabel := widget.NewLabel("Dépôt de garantie")
	securityDepositRateEntry := widget.NewEntry()
	securityDepositRateEntry.SetText(fmt.Sprintf("%1.2f", Loan.SecurityDepositRate))
	securityDepositRateEntry.OnChanged = func(value string) {
		securityDepositRate, _ := strconv.ParseFloat(value, 64)
		Loan.SecurityDepositRate = securityDepositRate
		fmt.Println("Loan.SecurityDepositRate updated to", value)
	}
	labels = append(labels, securityDepositRateLabel)
	entries = append(entries, securityDepositRateEntry)

	insuredQuantityRatioLabel := widget.NewLabel("Quantité d'emprunts assurés")
	insuredQuantityRatioEntry := widget.NewEntry()
	insuredQuantityRatioEntry.SetText(fmt.Sprintf("%1.2f", Loan.InsuredQuantityRatio))
	insuredQuantityRatioEntry.OnChanged = func(value string) {
		insuredQuantityRatio, _ := strconv.ParseFloat(value, 64)
		Loan.InsuredQuantityRatio = insuredQuantityRatio
		fmt.Println("Loan.InsuredQuantityRatio updated to", value)
	}
	labels = append(labels, insuredQuantityRatioLabel)
	entries = append(entries, insuredQuantityRatioEntry)

	borrowerFailureRateLabel := widget.NewLabel("Taux de défaut des prêts")
	borrowerFailureRateEntry := widget.NewEntry()
	borrowerFailureRateEntry.SetText(fmt.Sprintf("%1.2f", Loan.FailureRate))
	borrowerFailureRateEntry.OnChanged = func(value string) {
		borrowerFailureRate, _ := strconv.ParseFloat(value, 64)
		Loan.FailureRate = borrowerFailureRate
		fmt.Println("Loan.FailureRate updated to", value)
	}
	labels = append(labels, borrowerFailureRateLabel)
	entries = append(entries, borrowerFailureRateEntry)

	// Actor fields

	maxAmountPerLoanLabel := widget.NewLabel("Montant max par prêt par acteur")
	maxAmountPerLoanEntry := widget.NewEntry()
	maxAmountPerLoanEntry.SetText(fmt.Sprintf("%1.2f", Actor.MaxAmountPerLoan))
	maxAmountPerLoanEntry.OnChanged = func(value string) {
		lenderMaxAmountPerLoan, _ := strconv.ParseFloat(value, 64)
		Actor.MaxAmountPerLoan = lenderMaxAmountPerLoan
		fmt.Println("Actor.MaxAmountPerLoan updated to", value)
	}
	labels = append(labels, maxAmountPerLoanLabel)
	entries = append(entries, maxAmountPerLoanEntry)

	initialBalanceLabel := widget.NewLabel("Solde initial des acteurs")
	initialBalanceEntry := widget.NewEntry()
	initialBalanceEntry.SetText(fmt.Sprintf("%1.2f", Actor.InitialBalance))
	initialBalanceEntry.OnChanged = func(value string) {
		borrowerInitialBalance, _ := strconv.ParseFloat(value, 64)
		Actor.InitialBalance = borrowerInitialBalance
		fmt.Println("Actor.InitialBalance updated to", value)
	}
	labels = append(labels, initialBalanceLabel)
	entries = append(entries, initialBalanceEntry)

	balanceLeverageRatioLabel := widget.NewLabel("Ratio montant prêté sur balance")
	balanceLeverageRatioEntry := widget.NewEntry()
	balanceLeverageRatioEntry.SetText(fmt.Sprintf("%1.2f", Actor.BorrowerBalanceLeverageRatio))
	balanceLeverageRatioEntry.OnChanged = func(value string) {
		balanceLeverageRatio, _ := strconv.ParseFloat(value, 64)
		Actor.BorrowerBalanceLeverageRatio = balanceLeverageRatio
		fmt.Println("Actor.BorrowerBalanceLeverageRatio updated to", value)
	}
	labels = append(labels, balanceLeverageRatioLabel)
	entries = append(entries, balanceLeverageRatioEntry)

	// Master Container: where everything is bounded together

	for index, entry := range entries {
		borderContainer := container.NewBorder(nil, nil, labels[index], nil, entry)
		borderContainers = append(borderContainers, borderContainer)
	}

	vBox := container.NewVBox()
	for _, hbox := range borderContainers {
		vBox.Add(hbox)
	}

	masterContainer := container.NewBorder(nil, nil, nil, nil, vBox)
	return masterContainer
}

package configs

import (
	"fmt"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func RenderConfigs() *fyne.Container {
	var (
		labels           []*widget.Label
		entries          []*widget.Entry
		borderContainers []*fyne.Container
	)

	// General

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

	// Loan

	loanDefaultAmountLabel := widget.NewLabel("Montant des emprunts par défaut")
	loanDefaultAmountEntry := widget.NewEntry()
	loanDefaultAmountEntry.SetText(strconv.Itoa(Loan.DefaultAmount))
	loanDefaultAmountEntry.OnChanged = func(value string) {
		loanDefaultAmount, _ := strconv.Atoi(value)
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

	insuredQuantityRateLabel := widget.NewLabel("Quantité d'emprunts assurés")
	insuredQuantityRateEntry := widget.NewEntry()
	insuredQuantityRateEntry.SetText(fmt.Sprintf("%1.2f", Loan.InsuredQuantityRate))
	insuredQuantityRateEntry.OnChanged = func(value string) {
		insuredQuantityRate, _ := strconv.ParseFloat(value, 64)
		Loan.InsuredQuantityRate = insuredQuantityRate
		fmt.Println("Loan.InsuredQuantityRate updated to", value)
	}
	labels = append(labels, insuredQuantityRateLabel)
	entries = append(entries, insuredQuantityRateEntry)

	// Borrower

	borrowerInitialBalanceLabel := widget.NewLabel("Solde initial des emprunteurs")
	borrowerInitialBalanceEntry := widget.NewEntry()
	borrowerInitialBalanceEntry.SetText(strconv.Itoa(Borrower.InitialBalance))
	borrowerInitialBalanceEntry.OnChanged = func(value string) {
		borrowerInitialBalance, _ := strconv.Atoi(value)
		Borrower.InitialBalance = borrowerInitialBalance
		fmt.Println("Borrower.InitialBalance updated to", value)
	}
	labels = append(labels, borrowerInitialBalanceLabel)
	entries = append(entries, borrowerInitialBalanceEntry)

	borrowerFailureRateLabel := widget.NewLabel("Taux de défaut des emprunteurs")
	borrowerFailureRateEntry := widget.NewEntry()
	borrowerFailureRateEntry.SetText(fmt.Sprintf("%1.2f", Borrower.FailureRate))
	borrowerFailureRateEntry.OnChanged = func(value string) {
		borrowerFailureRate, _ := strconv.ParseFloat(value, 64)
		Borrower.FailureRate = borrowerFailureRate
		fmt.Println("Borrower.FailureRate updated to", value)
	}
	labels = append(labels, borrowerFailureRateLabel)
	entries = append(entries, borrowerFailureRateEntry)

	// Lender

	lenderInitialBalanceLabel := widget.NewLabel("Solde initial des prêteurs")
	lenderInitialBalanceEntry := widget.NewEntry()
	lenderInitialBalanceEntry.SetText(strconv.Itoa(Lender.InitialBalance))
	lenderInitialBalanceEntry.OnChanged = func(value string) {
		lenderInitialBalance, _ := strconv.Atoi(value)
		Lender.InitialBalance = lenderInitialBalance
		fmt.Println("Lender.InitialBalance updated to", value)
	}
	labels = append(labels, lenderInitialBalanceLabel)
	entries = append(entries, lenderInitialBalanceEntry)

	lenderMaxAmountPerLoanLabel := widget.NewLabel("Montant max par prêt par prêteur")
	lenderMaxAmountPerLoanEntry := widget.NewEntry()
	lenderMaxAmountPerLoanEntry.SetText(strconv.Itoa(Lender.MaxAmountPerLoan))
	lenderMaxAmountPerLoanEntry.OnChanged = func(value string) {
		lenderMaxAmountPerLoan, _ := strconv.Atoi(value)
		Lender.MaxAmountPerLoan = lenderMaxAmountPerLoan
		fmt.Println("Lender.MaxAmountPerLoan updated to", value)
	}
	labels = append(labels, lenderMaxAmountPerLoanLabel)
	entries = append(entries, lenderMaxAmountPerLoanEntry)

	// Insurer

	insurerInitialBalanceLabel := widget.NewLabel("Solde initial des assureurs")
	insurerInitialBalanceEntry := widget.NewEntry()
	insurerInitialBalanceEntry.SetText(strconv.Itoa(Insurer.InitialBalance))
	insurerInitialBalanceEntry.OnChanged = func(value string) {
		insurerInitialBalance, _ := strconv.Atoi(value)
		Insurer.InitialBalance = insurerInitialBalance
		fmt.Println("Insurer.InitialBalance updated to", value)
	}
	labels = append(labels, insurerInitialBalanceLabel)
	entries = append(entries, insurerInitialBalanceEntry)

	insurerMaxAmountPerLoanLabel := widget.NewLabel("Montant max par prêt par assureur")
	insurerMaxAmountPerLoanEntry := widget.NewEntry()
	insurerMaxAmountPerLoanEntry.SetText(strconv.Itoa(Insurer.MaxAmountPerLoan))
	insurerMaxAmountPerLoanEntry.OnChanged = func(value string) {
		insurerMaxAmountPerLoan, _ := strconv.Atoi(value)
		Insurer.MaxAmountPerLoan = insurerMaxAmountPerLoan
		fmt.Println("Insurer.MaxAmountPerLoan updated to", value)
	}
	labels = append(labels, insurerMaxAmountPerLoanLabel)
	entries = append(entries, insurerMaxAmountPerLoanEntry)

	for index, entry := range entries {
		borderContainer := container.NewBorder(nil, nil, labels[index], nil, entry)
		borderContainers = append(borderContainers, borderContainer)
	}

	masterContainer := container.NewVBox()
	for _, hbox := range borderContainers {
		masterContainer.Add(hbox)
	}

	return masterContainer
}

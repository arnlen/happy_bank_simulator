package views

import (
	"fmt"
	"strconv"

	"happy_bank_simulator/app/configs"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func RenderEdit() *fyne.Container {
	var (
		labels           []*widget.Label
		entries          []*widget.Entry
		borderContainers []*fyne.Container
	)

	// configs.General fields

	startDateLabel := widget.NewLabel("Date de démarrage de la simulation")
	startDateEntry := widget.NewEntry()
	startDateEntry.SetText(configs.General.StartDate)
	startDateEntry.OnChanged = func(value string) {
		startDate := value
		configs.General.StartDate = startDate
		fmt.Println("configs.General.StartDate updated to", value)
	}
	labels = append(labels, startDateLabel)
	entries = append(entries, startDateEntry)

	durationLabel := widget.NewLabel("Durée de la simulation (en mois)")
	durationEntry := widget.NewEntry()
	durationEntry.SetText(strconv.Itoa(configs.General.Duration))
	durationEntry.OnChanged = func(value string) {
		duration, _ := strconv.Atoi(value)
		configs.General.Duration = duration
		fmt.Println("configs.General.Duration updated to", value)
	}
	labels = append(labels, durationLabel)
	entries = append(entries, durationEntry)

	creditInterestRateLabel := widget.NewLabel("Taux d'intérêt du crédit")
	creditInterestRateEntry := widget.NewEntry()
	creditInterestRateEntry.SetText(fmt.Sprintf("%1.2f", configs.General.CreditInterestRate))
	creditInterestRateEntry.OnChanged = func(value string) {
		creditInterestRate, _ := strconv.ParseFloat(value, 64)
		configs.General.CreditInterestRate = creditInterestRate
		fmt.Println("configs.General.CreditInterestRate updated to", value)
	}
	labels = append(labels, creditInterestRateLabel)
	entries = append(entries, creditInterestRateEntry)

	insuranceInterestRateLabel := widget.NewLabel("Taux d'intérêt de l'assurance")
	insuranceInterestRateEntry := widget.NewEntry()
	insuranceInterestRateEntry.SetText(fmt.Sprintf("%1.2f", configs.General.InsuranceInterestRate))
	insuranceInterestRateEntry.OnChanged = func(value string) {
		insuranceInterestRate, _ := strconv.ParseFloat(value, 64)
		configs.General.InsuranceInterestRate = insuranceInterestRate
		fmt.Println("configs.General.InsuranceInterestRate updated to", value)
	}
	labels = append(labels, insuranceInterestRateLabel)
	entries = append(entries, insuranceInterestRateEntry)

	// configs.Loan fields

	initialLoanQuantityLabel := widget.NewLabel("Nombre d'emprunts au démarrage de la simulation")
	initialLoanQuantityEntry := widget.NewEntry()
	initialLoanQuantityEntry.SetText(strconv.Itoa(configs.Loan.InitialQuantity))
	initialLoanQuantityEntry.OnChanged = func(value string) {
		initialLoanQuantity, _ := strconv.Atoi(value)
		configs.Loan.InitialQuantity = initialLoanQuantity
		fmt.Println("configs.Loan.InitialQuantity updated to", value)
	}
	labels = append(labels, initialLoanQuantityLabel)
	entries = append(entries, initialLoanQuantityEntry)

	loanDefaultAmountLabel := widget.NewLabel("Montant des emprunts par défaut")
	loanDefaultAmountEntry := widget.NewEntry()
	loanDefaultAmountEntry.SetText(fmt.Sprintf("%1.2f", configs.Loan.DefaultAmount))
	loanDefaultAmountEntry.OnChanged = func(value string) {
		loanDefaultAmount, _ := strconv.ParseFloat(value, 64)
		configs.Loan.DefaultAmount = loanDefaultAmount
		fmt.Println("configs.Loan.DefaultAmount updated to", value)
	}
	labels = append(labels, loanDefaultAmountLabel)
	entries = append(entries, loanDefaultAmountEntry)

	defaultDurationLabel := widget.NewLabel("Durée par défaut d'un emprunt")
	defaultDurationEntry := widget.NewEntry()
	defaultDurationEntry.SetText(strconv.Itoa(configs.Loan.DefaultDuration))
	defaultDurationEntry.OnChanged = func(value string) {
		defaultDuration, _ := strconv.Atoi(value)
		configs.Loan.DefaultDuration = defaultDuration
		fmt.Println("configs.Loan.DefaultDuration updated to", value)
	}
	labels = append(labels, defaultDurationLabel)
	entries = append(entries, defaultDurationEntry)

	securityDepositRateLabel := widget.NewLabel("Dépôt de garantie")
	securityDepositRateEntry := widget.NewEntry()
	securityDepositRateEntry.SetText(fmt.Sprintf("%1.2f", configs.Loan.SecurityDepositRate))
	securityDepositRateEntry.OnChanged = func(value string) {
		securityDepositRate, _ := strconv.ParseFloat(value, 64)
		configs.Loan.SecurityDepositRate = securityDepositRate
		fmt.Println("configs.Loan.SecurityDepositRate updated to", value)
	}
	labels = append(labels, securityDepositRateLabel)
	entries = append(entries, securityDepositRateEntry)

	insuredQuantityRatioLabel := widget.NewLabel("Quantité d'emprunts assurés")
	insuredQuantityRatioEntry := widget.NewEntry()
	insuredQuantityRatioEntry.SetText(fmt.Sprintf("%1.2f", configs.Loan.InsuredQuantityRatio))
	insuredQuantityRatioEntry.OnChanged = func(value string) {
		insuredQuantityRatio, _ := strconv.ParseFloat(value, 64)
		configs.Loan.InsuredQuantityRatio = insuredQuantityRatio
		fmt.Println("configs.Loan.InsuredQuantityRatio updated to", value)
	}
	labels = append(labels, insuredQuantityRatioLabel)
	entries = append(entries, insuredQuantityRatioEntry)

	borrowerFailureRateLabel := widget.NewLabel("Taux de défaut des prêts")
	borrowerFailureRateEntry := widget.NewEntry()
	borrowerFailureRateEntry.SetText(fmt.Sprintf("%1.2f", configs.Loan.FailureRate))
	borrowerFailureRateEntry.OnChanged = func(value string) {
		borrowerFailureRate, _ := strconv.ParseFloat(value, 64)
		configs.Loan.FailureRate = borrowerFailureRate
		fmt.Println("configs.Loan.FailureRate updated to", value)
	}
	labels = append(labels, borrowerFailureRateLabel)
	entries = append(entries, borrowerFailureRateEntry)

	// configs.Actor fields

	maxAmountPerLoanLabel := widget.NewLabel("Montant max par prêt par acteur")
	maxAmountPerLoanEntry := widget.NewEntry()
	maxAmountPerLoanEntry.SetText(fmt.Sprintf("%1.2f", configs.Actor.MaxAmountPerLoan))
	maxAmountPerLoanEntry.OnChanged = func(value string) {
		lenderMaxAmountPerLoan, _ := strconv.ParseFloat(value, 64)
		configs.Actor.MaxAmountPerLoan = lenderMaxAmountPerLoan
		fmt.Println("configs.Actor.MaxAmountPerLoan updated to", value)
	}
	labels = append(labels, maxAmountPerLoanLabel)
	entries = append(entries, maxAmountPerLoanEntry)

	initialBalanceLabel := widget.NewLabel("Solde initial des acteurs")
	initialBalanceEntry := widget.NewEntry()
	initialBalanceEntry.SetText(fmt.Sprintf("%1.2f", configs.Actor.InitialBalance))
	initialBalanceEntry.OnChanged = func(value string) {
		borrowerInitialBalance, _ := strconv.ParseFloat(value, 64)
		configs.Actor.InitialBalance = borrowerInitialBalance
		fmt.Println("configs.Actor.InitialBalance updated to", value)
	}
	labels = append(labels, initialBalanceLabel)
	entries = append(entries, initialBalanceEntry)

	balanceLeverageRatioLabel := widget.NewLabel("Ratio montant prêté sur balance")
	balanceLeverageRatioEntry := widget.NewEntry()
	balanceLeverageRatioEntry.SetText(fmt.Sprintf("%1.2f", configs.Actor.BorrowerBalanceLeverageRatio))
	balanceLeverageRatioEntry.OnChanged = func(value string) {
		balanceLeverageRatio, _ := strconv.ParseFloat(value, 64)
		configs.Actor.BorrowerBalanceLeverageRatio = balanceLeverageRatio
		fmt.Println("configs.Actor.BorrowerBalanceLeverageRatio updated to", value)
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

package models

import (
	"fmt"
	"log"
	"strconv"

	"happy_bank_simulator/app/configs"
	"happy_bank_simulator/database"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"syreclabs.com/go/faker"
)

type Borrower struct {
	gorm.Model
	Name    string
	Loans   []Loan
	Balance int
}

// ------- Instance methods -------

func (instance *Borrower) ModelName() string {
	return "emprunteur"
}

func (instance *Borrower) Refresh() {
	database.GetDB().Preload(clause.Associations).Find(&instance)
}

func (instance *Borrower) Save() {
	db := database.GetDB()
	result := db.Save(instance)

	if instance.ID == 0 || result.RowsAffected == 0 {
		log.Fatal(result.Error)
	}

	instance.Refresh()
}

func (instance *Borrower) GetNetBalance() int {
	netBalance := instance.Balance - instance.GetTotalAmountBorrowed()
	fmt.Printf("Borrower #%s net balance is %s €\n", strconv.Itoa(int(instance.ID)), strconv.Itoa(netBalance))
	return netBalance
}

func (instance *Borrower) GetTotalAmountBorrowed() int {
	loans := instance.Loans
	totalAmoutBorrowed := 0

	for _, loan := range loans {
		totalAmoutBorrowed += loan.Amount
	}

	fmt.Printf("Borrower #%s has %s loans for a total of %s €\n", strconv.Itoa(int(instance.ID)), strconv.Itoa(len(loans)), strconv.Itoa(totalAmoutBorrowed))
	return totalAmoutBorrowed
}

// ------- Package methods -------

func FindBorrower(id int) *Borrower {
	var borrower Borrower
	database.GetDB().Preload(clause.Associations).First(&borrower, id)
	return &borrower
}

func ListBorrowers() []Borrower {
	var borrowers []Borrower
	database.GetDB().Preload(clause.Associations).Find(&borrowers)
	return borrowers
}

func NewBorrower(name string, balance int) *Borrower {
	return &Borrower{
		Name:    name,
		Loans:   []Loan{},
		Balance: balance,
	}
}

func NewDefaultBorrower() *Borrower {
	return &Borrower{
		Name:    faker.Name().Name(),
		Loans:   []Loan{},
		Balance: configs.Borrower.InitialBalance,
	}
}

func CreateBorrower(name string, balance int) *Borrower {
	borrower := NewBorrower(name, balance)
	result := database.GetDB().Create(&borrower)

	if borrower.ID == 0 || result.RowsAffected == 0 {
		log.Fatal(result.Error)
	}

	return borrower
}

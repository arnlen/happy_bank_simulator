package models

import (
	"gorm.io/gorm"
)

type Borrower struct {
	gorm.Model
	Loans []Loan
}

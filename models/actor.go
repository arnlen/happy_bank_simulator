package models

import ()

type ModelBase interface {
	ModelName() string
	Refresh()
	Save()
}

type Actor interface {
	ModelBase
	UpdateBalance(amount float64)
	GetID() uint
}

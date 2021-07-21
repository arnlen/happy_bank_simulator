package models

import ()

type ModelBase interface {
	ModelName() string
	Refresh()
	Save()
}

type Actor interface {
	UpdateBalance(amount int)
}

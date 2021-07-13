package helpers

import (
	"happy_bank_simulator/database"
	modelHelpers "happy_bank_simulator/models/helpers"

	"gorm.io/gorm"
)

func DropBD() {
	for _, model := range modelHelpers.GetModelList() {
		database.GetDB().Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(model)
	}
}

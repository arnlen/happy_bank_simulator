package helpers

import (
	"happy_bank_simulator/database"
	modelHelpers "happy_bank_simulator/models/helpers"
)

func MigrateDB() {
	modelList := modelHelpers.GetModelList()
	database.GetDB().AutoMigrate(modelList...)
}

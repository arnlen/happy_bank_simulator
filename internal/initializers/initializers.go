package initializers

import (
	"fmt"
	"log"
	"os"
	"time"

	"happy_bank_simulator/internal/global"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	db  *gorm.DB
	err error
)

func InitDB() *gorm.DB {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Warn, // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			Colorful:                  false,       // Disable color
		},
	)

	databasePath := fmt.Sprintf("%s/database/happy_dev.db", global.ProjectPath)
	db, err = gorm.Open(sqlite.Open(databasePath), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Database initialized")
	return db
}

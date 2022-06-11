package database

import (
	"fmt"

	"github.com/CarlFlo/blacklisterBot/src/config"
	"github.com/CarlFlo/malm"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

const resetDatabaseOnStart = false

func connectToDB() error {

	var err error
	DB, err = gorm.Open(sqlite.Open(config.CONFIG.Database.FileName), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		return err
	}

	if resetDatabaseOnStart {

		malm.Info("Resetting database...")

		// Populates the database with the default values

		tableList := []string{
			(&Blacklist{}).TableName()}

		for _, name := range tableList {
			DB.Exec(fmt.Sprintf("DROP TABLE %s", name))
		}
	}

	return DB.AutoMigrate(&Blacklist{})

}

func InitializeDatabase() error {

	err := connectToDB()
	if err != nil {
		return err
	}

	return nil
}

// Loads the database.
// Any problems will be logged
func Load() {

	if err := InitializeDatabase(); err != nil {
		malm.Fatal("Database initialization error: %s", err)
	}
}

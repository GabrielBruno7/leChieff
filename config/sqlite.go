package config

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"leChief/schemas"
	"os"
)

func InitializeSQLite() (*gorm.DB, error) {
	logger := GetLogger("sqlite")

	databasePath := "./db/main.db"

	_, err := os.Stat(databasePath)
	if os.IsNotExist(err) {
		logger.Info("Database file does not exist, creating...")

		err = os.MkdirAll("./db", os.ModePerm)
		if err != nil {
			return nil, err
		}

		file, err := os.Create(databasePath)
		if err != nil {
			return nil, err
		}

		file.Close()
	}

	database, err := gorm.Open(sqlite.Open(databasePath), &gorm.Config{})
	if err != nil {
		logger.ErrorFormatted("Error opening SQLite database: %v", err)
		return nil, err
	}

	err = database.AutoMigrate(&schemas.Order{})
	if err != nil {
		logger.ErrorFormatted("Error migrating database: %v", err)
		return nil, err
	}

	return database, nil
}

package config

import (
	"fmt"
	"gorm.io/gorm"
)

var (
	database *gorm.DB
	logger   *Logger
)

func InitializeDatabase() error {
	var err error

	database, err = InitializeMySQL()
	if err != nil {
		return fmt.Errorf("Error initializing database: %v", err)
	}

	return nil
}

func GetDatabase() *gorm.DB {
	return database
}

func GetLogger(prefix string) *Logger {
	logger = NewLogger(prefix)
	return logger
}

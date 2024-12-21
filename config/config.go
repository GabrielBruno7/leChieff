package config

import (
	"errors"
	"gorm.io/gorm"
)

var (
	database *gorm.DB
	logger   *Logger
)

func InitializeDatabase() error {
	return errors.New("Deu erro paizao")
}

func GetLogger(prefix string) *Logger {
	logger = NewLogger(prefix)
	return logger
}

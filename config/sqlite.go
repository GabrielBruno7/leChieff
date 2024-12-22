package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"leChief/schemas"
)

func InitializeMySQL() (*gorm.DB, error) {
	logger := GetLogger("mysql")

	// Incluindo parseTime=True na DSN
	dsn := "root:wyaNaujpNbEZKZdwBFFCALelEApuJHmu@tcp(junction.proxy.rlwy.net:22781)/railway?charset=utf8mb4&parseTime=True&loc=Local"

	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.ErrorFormatted("Error connecting to MySQL: %v", err)
		return nil, err
	}

	logger.Info("Successfully connected to MySQL")

	// Migrar a estrutura de tabelas
	err = database.AutoMigrate(&schemas.Order{})
	if err != nil {
		logger.ErrorFormatted("Error migrating database: %v", err)
		return nil, err
	}

	return database, nil
}

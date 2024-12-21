package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"leChief/schemas"
)

func InitializeSQLite() (*gorm.DB, error) {
	logger := GetLogger("mysql")

	database, err := gorm.Open(mysql.New(mysql.Config{
		DriverName: "mysql",
		DSN:        "root:wyaNaujpNbEZKZdwBFFCALelEApuJHmu@tcp(junction.proxy.rlwy.net:22781)/railway",
	}), &gorm.Config{})

	if err != nil {
		logger.ErrorFormatted("Error connecting to MySQL: %v", err)
		return nil, err
	}

	if err == nil {
		logger.Info("Successfully connected to MySQL")
	}

	err = database.AutoMigrate(&schemas.Order{})
	if err != nil {
		logger.ErrorFormatted("Error migrating database: %v", err)
		return nil, err
	}

	return database, nil
}

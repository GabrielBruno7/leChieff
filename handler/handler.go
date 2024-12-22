package handler

import (
	"gorm.io/gorm"
	"leChief/config"
)

var (
	Logger   *config.Logger
	Database *gorm.DB
)

func InitializeHandler() {
	Logger = config.GetLogger("handler")
	Database = config.GetDatabase()
}

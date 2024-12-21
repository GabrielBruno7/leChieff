package handler

import (
	"gorm.io/gorm"
	"leChief/config"
)

var (
	logger   *config.Logger
	database *gorm.DB
)

func InitializeHandler() {
	logger = config.GetLogger("handler")
	database = config.GetDatabase()
}

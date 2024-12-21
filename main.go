package main

import (
	"leChief/config"
	"leChief/router"
)

var (
	logger *config.Logger
)

func main() {
	logger := config.GetLogger("main")
	err := config.InitializeDatabase()

	if err != nil {
		logger.ErrorFormatted("Config initialization failed: %v", err)
		return
	}

	router.Initialize()

}

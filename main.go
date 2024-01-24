package main

import (
	"github.com/madfelps/go-wallet/config"
	"github.com/madfelps/go-wallet/router"
)

var (
	logger config.Logger
)

func main() {

	logger = *config.GetLogger("main")

	// Initialize Configs
	err := config.Init()
	if err != nil {
		logger.Errorf("config initialization error: %v", err)
		return
	}

	// Initialize Router
	router.Initialize()
}

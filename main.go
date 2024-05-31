package main

import (
	"github.com/forbatnew/gopportunities/config"
	"github.com/forbatnew/gopportunities/router"
)

var (
	logger config.Logger
)

func main() {
	logger = *config.GetLogger("main")
	//Initialize configuration
	err := config.Init()

	if err != nil {
		//	panic(err)
		logger.Errorf("Config initialization error: %v", err)
		return
	}

	//Initialize router
	router.Initialize()
}

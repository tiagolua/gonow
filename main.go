package main

import (
	"github.com/tiagolua/gonow.git/config"
	"github.com/tiagolua/gonow.git/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")
	// initialize config
	err := config.Init()
	if err != nil {
		logger.Errorf("config initialization error: %v", err)
		return
	}

	//initialize router
	router.Initializer()
}

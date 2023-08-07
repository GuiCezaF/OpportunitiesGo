package main

import (
	"github.com/GuiCezaF/OpportunitiesGo/config"
	"github.com/GuiCezaF/OpportunitiesGo/router"
)

var (
	logger config.Logger
)

func main() {
	logger = *config.GetLogger("main")
	//Initialize Configs
	err := config.Init()
	if err != nil {
		logger.Errorf("config initialization error: %v", err)
		return
	}

	//Initialize Routes
	router.Initialize()
}

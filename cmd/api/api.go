package main

import (
	"fmt"

	"github.com/coomsy/mmo/cmd/api/app"
	"github.com/coomsy/mmo/internal/config"
	"github.com/coomsy/mmo/pkg/logger"
)

func init() {
	if err := config.InitAppConfig(); err != nil {
		logger.Fatal(err.Error())
	}

	logger.Info("Config loaded")
	logger.Debug(fmt.Sprintf("Config Values %+v\n", config.AppConfig))
}

func main() {
	server, err := app.NewApp()

	if err != nil {
		logger.Panic(err.Error())
	}
	if err := server.Run(); err != nil {
		logger.Fatal(err.Error())
	}
}

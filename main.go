package main

import (
	"fmt"
	"github.com/KCKT0112/GoWeb/app/config"
	"github.com/KCKT0112/GoWeb/app/utils"
	"go.uber.org/zap"
	"github.com/KCKT0112/GoWeb/app/routers"
)

func main() {
	config.InitConfig()

	port := config.AppConfig.Server.Port
	if port == 0 {
		port = 8083 // Default port
	}

	// Initialize the logger with the configuration
	utils.InitializeLogger()

	logger := utils.Logger
	logger.Info("Starting server", zap.String("port", fmt.Sprintf("%d", port)))

	router := routes.InitRouter()

	router.Run(fmt.Sprintf(":%d", port))
}

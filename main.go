package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/KCKT0112/GoWeb/app/config"
	"github.com/KCKT0112/GoWeb/app/utils"
	"github.com/gin-gonic/gin"
)

func main() {
	config.InitConfig()

	logLevel := config.AppConfig.Logger.Level
	if logLevel == "" {
		logLevel = "info"
	}

	// Initialize the logger with the configuration
    utils.InitializeLogger(logLevel)

    // Get the global logger
    logger := utils.Logger

    err := doSomething()
    if err != nil {
        logger.Error("Error occurred", zap.Error(err))
    }

	r := gin.Default()

	r.Use(MiddWare())

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, fmt.Sprintf("it's working! %v", time.Now()))
	})

	port := config.AppConfig.Server.Port
	if port == 0 {
		port = 8083 // Default port
	}

	r.Run(fmt.Sprintf(":%d", port))
}

func MiddWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		// t := time.Now()
		// fmt.Println("before middleware")
		c.Set("Request", "Middleware")

		// c.Next()

		// status := c.Writer.Status()
		// fmt.Println("after middleware")
		// fmt.Printf("status: %d\n", status)
		// t2 := time.Since(t)
		// fmt.Printf("time: %v\n", t2)
	}
}

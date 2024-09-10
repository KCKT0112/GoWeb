package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/KCKT0112/GoWeb/app/config"
	"github.com/KCKT0112/GoWeb/app/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
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

	r := gin.Default()

	r.Use(MiddWare())

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, fmt.Sprintf("it's working! %v", time.Now()))
	})

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

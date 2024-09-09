package main

import (
	"fmt"
	"net/http"
	"time"
	"github.com/gin-gonic/gin"
	"github.com/KCKT0112/GoWeb/app/config"
)

func main() {
	config.InitConfig()

	r := gin.Default()

	r.Use(MiddWare())

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, fmt.Sprintf("it's working! %v", time.Now()))
	})

	r.Run(fmt.Sprintf(":%d", config.AppConfig.Server.Port ?: 8083))
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

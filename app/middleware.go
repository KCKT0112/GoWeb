package app

import (
	"github.com/gin-gonic/gin"
)

func MiddWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("Request", "Middleware")
		c.Next()
	}
}

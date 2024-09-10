package routes

import "github.com/gin-gonic/gin"

func InitRouter() *gin.Engine {
    router := gin.Default()

	// Register middleware
	router.Use(MiddWare())

    // Register routes
    IndexRoutes(router)

    return router
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
package routes

import (
	"github.com/KCKT0112/GoWeb/app/controllers"
	"github.com/KCKT0112/GoWeb/app/services"
	"github.com/gin-gonic/gin"
)

func IndexRoutes(router *gin.Engine) {
	indexService := services.NewIndexService()
	indexController := controllers.NewIndexController(indexService)

	indexGroup := router.Group("/")
	{
		indexGroup.GET("/", indexController.GetIndex)
	}
}

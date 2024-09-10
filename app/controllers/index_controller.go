package controllers

import (
	"net/http"

	"github.com/KCKT0112/GoWeb/app/services"
	"github.com/gin-gonic/gin"
)

type IndexController struct {
	indexService services.IndexService
}

func NewIndexController(service services.IndexService) *IndexController {
	return &IndexController{
		indexService: service,
	}
}

func (uc *IndexController) GetIndex(c *gin.Context) {
	c.Writer.WriteHeader(http.StatusOK)
	c.Writer.Write([]byte("it works!"))
}

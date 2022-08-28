package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tgs266/neon/neon/api"
	"github.com/tgs266/neon/neon/services"
)

func CreateProduct(c *gin.Context) {
	var req api.CreateProductRequest
	c.BindJSON(&req)
	services.CreateProduct(req)
	c.JSON(http.StatusOK, req)
}

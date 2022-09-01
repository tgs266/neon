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

func ListProducts(c *gin.Context) {
	resp := services.ListProducts()
	c.JSON(http.StatusOK, resp)
}

func GetProduct(c *gin.Context) {
	productName := c.Param("name")
	resp := services.GetProductByName(c, productName)
	c.JSON(http.StatusOK, resp)
}

func GetProductReleases(c *gin.Context) {
	productName := c.Param("name")
	resp := services.GetProductByName(c, productName).Releases
	c.JSON(http.StatusOK, resp)
}

func GetProductInstalls(c *gin.Context) {
	productName := c.Param("name")
	resp := services.GetProductByName(c, productName).Installs
	c.JSON(http.StatusOK, resp)
}

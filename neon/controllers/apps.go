package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/tgs266/neon/neon/api"
	"github.com/tgs266/neon/neon/services"
)

func CreateApp(c *gin.Context) {
	var req api.CreateAppRequest
	c.BindJSON(&req)
	services.CreateApp(c, req)
	c.JSON(http.StatusOK, req)
}

func AddProductToApp(c *gin.Context) {
	var req api.AddProductRequest
	name := c.Param("name")
	c.BindJSON(&req)
	services.AddProductToApp(c, name, req)
	c.JSON(http.StatusOK, req)
}

func ListApps(c *gin.Context) {
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "0"))
	offset, _ := strconv.Atoi(c.DefaultQuery("offset", "0"))
	name := c.DefaultQuery("name", "")
	resp := services.ListApps(c, name, limit, offset)
	c.JSON(http.StatusOK, resp)
}

func GetApp(c *gin.Context) {
	name := c.Param("name")
	resp := services.GetAppByName(c, name)
	c.JSON(http.StatusOK, resp)
}

func GetAppInstall(c *gin.Context) {
	name := c.Param("name")
	productName := c.Param("productName")
	resp := services.GetAppInstall(c, name, productName)
	c.JSON(http.StatusOK, resp)
}

func GetAppInstallResources(c *gin.Context) {
	name := c.Param("name")
	productName := c.Param("productName")
	resp := services.GetAppInstallResources(c, name, productName)
	c.JSON(http.StatusOK, resp)
}

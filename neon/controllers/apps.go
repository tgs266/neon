package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/tgs266/neon/neon/api"
	"github.com/tgs266/neon/neon/services"
)

func ApplyApp(c *gin.Context) {
	var req api.ApplyAppRequest
	c.BindJSON(&req)
	services.ApplyApp(c, req)
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
	resp := services.GetAppByName(name)
	c.JSON(http.StatusOK, resp)
}

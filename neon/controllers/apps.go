package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tgs266/neon/neon/api"
	"github.com/tgs266/neon/neon/services"
)

func CreateApp(c *gin.Context) {
	var req api.CreateAppRequest
	c.BindJSON(&req)
	services.CreateApp(req)
	c.JSON(http.StatusOK, req)
}

func ListApps(c *gin.Context) {
	resp := services.ListApps()
	c.JSON(http.StatusOK, resp)
}

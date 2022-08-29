package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tgs266/neon/neon/api"
	"github.com/tgs266/neon/neon/services"
)

func CreateRelease(c *gin.Context) {
	var req api.CreateReleaseRequest
	c.BindJSON(&req)
	services.CreateRelease(req)
	c.JSON(http.StatusOK, req)
}

package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/tgs266/neon/neon/api"
)

func ListCredentials(c *gin.Context) {
	// resp := services.ListCredentials(c, app, limit, offset)
	// c.JSON(http.StatusOK, resp)
}

func AddCredentials(c *gin.Context) {
	var req api.AddCredentialsRequest
	c.BindJSON(&req)
	req.Validate()
	// resp := services.AddCredentials(c, req)
	// c.JSON(http.StatusOK, resp)
}

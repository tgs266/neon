package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tgs266/neon/neon/api"
	"github.com/tgs266/neon/neon/services"
)

func ListCredentials(c *gin.Context) {
	resp := services.GetCredentials(c)
	c.JSON(http.StatusOK, resp)
}

func AddCredentials(c *gin.Context) {
	var req api.AddCredentialsRequest
	c.BindJSON(&req)
	req.Validate(c)
	resp := services.AddCredentials(c, req)
	c.JSON(http.StatusOK, resp)
}

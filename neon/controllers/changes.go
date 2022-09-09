package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/tgs266/neon/neon/services"
)

func ListStoredChanges(c *gin.Context) {
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "0"))
	offset, _ := strconv.Atoi(c.DefaultQuery("offset", "0"))
	app := c.Param("name")
	resp := services.ListStoredChanges(c, app, limit, offset)
	c.JSON(http.StatusOK, resp)
}

func ListQueuedChanges(c *gin.Context) {
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "0"))
	offset, _ := strconv.Atoi(c.DefaultQuery("offset", "0"))
	app := c.Param("name")
	resp := services.ListQueuedChanges(c, app, limit, offset)
	c.JSON(http.StatusOK, resp)
}

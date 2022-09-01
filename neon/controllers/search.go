package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/tgs266/neon/neon/services"
)

func ProductSearch(c *gin.Context) {
	name := c.DefaultQuery("name", "")
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "0"))
	offset, _ := strconv.Atoi(c.DefaultQuery("offset", "0"))
	resp := services.FindProducts(c, name, limit, offset)
	c.JSON(http.StatusOK, resp)
}

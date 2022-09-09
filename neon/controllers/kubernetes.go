package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tgs266/neon/neon/kubernetes"
)

func GetPod(c *gin.Context) {
	name := c.Param("name")
	ns := c.Param("namespace")
	resp := kubernetes.Pods(c, ns).GetPod(name)
	resp.ManagedFields = nil
	c.JSON(http.StatusOK, resp)
}

func GetPodStatus(c *gin.Context) {
	name := c.Param("name")
	ns := c.Param("namespace")
	resp := kubernetes.Pods(c, ns).GetPodStatus(name)
	c.JSON(http.StatusOK, resp)
}

package neon

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tgs266/neon/neon/controllers"
	"github.com/tgs266/neon/neon/store"
)

func Start(host, username, password, port string) {
	store.CreateStore(host, username, password)
	r := gin.Default()
	r.GET("/api/v1/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"postgres": store.IsConnected(),
		})
	})
	controllers.Routes(r)
	r.Run(":" + port)
}

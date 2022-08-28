package neon

import (
	"github.com/gin-gonic/gin"
)

func Start(host, username, password, port string) {
	// store.CreateStore(host, username, password)
	r := gin.Default()
	// r.GET("/api/v1/health", func(c *gin.Context) {
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"postgres": store.IsConnected(),
	// 	})
	// })
	// controllers.Routes(r)
	r.Run(":" + port)
}

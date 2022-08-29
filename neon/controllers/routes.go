package controllers

import "github.com/gin-gonic/gin"

func Routes(r *gin.Engine) {
	r.GET("/api/v1/products", ListProducts)
	r.GET("/api/v1/products/:name", GetProduct)
	r.GET("/api/v1/products/:name/releases", GetProductReleases)
	r.POST("/api/v1/products", CreateProduct)

	r.POST("/api/v1/releases", CreateRelease)

	r.GET("/api/v1/apps", ListApps)
	r.POST("/api/v1/apps", CreateApp)
}

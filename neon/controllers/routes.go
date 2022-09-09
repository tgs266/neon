package controllers

import "github.com/gin-gonic/gin"

func Routes(r *gin.Engine) {
	r.GET("/api/v1/products", ListProducts)
	r.POST("/api/v1/products", CreateProduct)
	r.GET("/api/v1/products/:name", GetProduct)
	r.GET("/api/v1/products/:name/releases", GetProductReleases)
	r.GET("/api/v1/products/:name/installs", GetProductInstalls)

	r.POST("/api/v1/releases", CreateRelease)

	r.GET("/api/v1/apps", ListApps)
	r.GET("/api/v1/apps/:name", GetApp)
	r.POST("/api/v1/apps", ApplyApp)

	r.GET("/api/v1/apps/:name/installs/:productName/resources", GetAppInstallResources)
	r.GET("/api/v1/apps/:name/installs/:productName", GetAppInstall)
	r.GET("/api/v1/apps/:name/changes/stored", ListStoredChanges)
	r.GET("/api/v1/apps/:name/changes/queued", ListQueuedChanges)
}

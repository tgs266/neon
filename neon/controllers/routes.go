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
	r.POST("/api/v1/apps", CreateApp)

	r.POST("/api/v1/apps/:name/products", AddProductToApp)

	r.GET("/api/v1/apps/:name/installs/:productName/resources", GetAppInstallResources)
	r.GET("/api/v1/apps/:name/installs/:productName", GetAppInstall)
	r.GET("/api/v1/apps/:name/changes/stored", ListStoredChanges)
	r.GET("/api/v1/apps/:name/changes/queued", ListQueuedChanges)

	r.GET("/api/v1/kubernetes/pods/:namespace/:name", GetPod)
	r.GET("/api/v1/kubernetes/pods/:namespace/:name/status", GetPodStatus)

	r.GET("/api/v1/credentials", ListCredentials)
	r.POST("/api/v1/credentials", AddCredentials)
}

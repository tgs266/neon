package controllers

import "github.com/gin-gonic/gin"

func Routes(r *gin.Engine) {
	r.POST("/api/v1/products", CreateProduct)
}

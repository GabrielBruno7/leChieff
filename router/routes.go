package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func initializeRoute(router *gin.Engine) {
	v1 := router.Group("/api/v1/")
	{
		v1.GET("/orders", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{
				"message": "Get all orders",
			})
		})

		v1.GET("/order/{id}", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{
				"message": "Get orders by id",
			})
		})

		v1.POST("/order", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{
				"message": "Create order",
			})
		})

		v1.DELETE("/order/{id}", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{
				"message": "Delete order",
			})
		})

		v1.PUT("/order/{id}", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{
				"message": "Update order",
			})
		})
	}
}

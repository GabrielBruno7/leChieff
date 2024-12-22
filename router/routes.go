package router

import (
	"github.com/gin-gonic/gin"
	"leChief/handler"
)

func initializeRoute(router *gin.Engine) {
	handler.InitializeHandler()

	v1 := router.Group("/api/v1/")
	{
		v1.GET("/orders", handler.ListOrdersHandler)

		v1.POST("/order", handler.CreateOrderHandler)

		v1.GET("/order", handler.ShowOrderHandler)

		v1.PUT("/order", handler.UpdateOrderHandler)

		v1.DELETE("/order", handler.DeleteOrderHandler)
	}
}

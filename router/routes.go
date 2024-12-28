package router

import (
	"github.com/gin-gonic/gin"
	"leChief/handler"
	"leChief/handler/customers"
	"leChief/handler/orders"
)

func initializeRoute(router *gin.Engine) {
	handler.InitializeHandler()

	v1 := router.Group("/api/v1/")
	{
		v1.GET("/orders", orders.ListOrdersHandler)
		v1.POST("/order", orders.CreateOrderHandler)
		v1.GET("/order", orders.ShowOrderHandler)
		v1.PUT("/order", orders.UpdateOrderHandler)
		v1.DELETE("/order", orders.DeleteOrderHandler)

		v1.GET("/customers", customers.ListCustomersHandler)
		v1.POST("/customer", customers.CreateCustomerHandler)
		v1.GET("/customer", customers.ShowCustomerHandler)
		v1.PUT("/customer", customers.UpdateCustomerHandler)
		v1.DELETE("/customer", customers.DeleteCustomerHandler)
	}
}

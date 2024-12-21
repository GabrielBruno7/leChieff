package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateOrderHandler(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "Create order",
	})
}

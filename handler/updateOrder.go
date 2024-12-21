package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func UpdateOrderHandler(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "Update order",
	})
}

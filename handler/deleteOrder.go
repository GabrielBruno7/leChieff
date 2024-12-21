package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func DeleteOrderHandler(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "Delete order",
	})
}

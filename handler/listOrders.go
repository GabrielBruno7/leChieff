package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ListOrdersHandler(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "Get all orders",
	})
}

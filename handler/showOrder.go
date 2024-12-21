package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ShowOrderHandler(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "Get order by id",
	})
}

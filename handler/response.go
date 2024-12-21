package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func sendSuccessResponse(context *gin.Context, operation string, data interface{}) {
	context.Header("content-type", "application/json")
	context.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Successfully sent response: %s", operation),
		"data":    data,
	})
}

func sendErrorResponse(context *gin.Context, code int, message string) {
	context.Header("content-type", "application/json")
	context.JSON(code, gin.H{
		"code":    code,
		"message": message,
	})
}

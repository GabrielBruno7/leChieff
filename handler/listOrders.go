package handler

import (
	"github.com/gin-gonic/gin"
	"leChief/schemas"
	"net/http"
)

func ListOrdersHandler(context *gin.Context) {
	orders := []schemas.Order{}

	if err := database.Find(&orders).Order("created_at ASC").Error; err != nil {
		sendErrorResponse(context, http.StatusInternalServerError, "Error listing orders")
		return
	}

	sendSuccessResponse(context, "Listing orders", orders)
}

func ListOrdersWhenStatusIsLoadingHandler(context *gin.Context) {
	var results []schemas.Order

	query := "SELECT * FROM orders WHERE status LIKE ?"

	if err := database.Raw(query, "loading").Scan(&results).Error; err != nil {
		sendErrorResponse(context, http.StatusInternalServerError, "Error listing orders")
		return
	}

	sendSuccessResponse(context, "Listing orders", results)
}

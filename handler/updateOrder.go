package handler

import (
	"github.com/gin-gonic/gin"
	"leChief/schemas"
	"net/http"
)

func UpdateOrderHandler(context *gin.Context) {
	request := updateOrderRequest{}

	context.BindJSON(&request)
	if err := request.Validate(); err != nil {
		logger.ErrorFormatted("Validation Error: %v", err.Error())
		sendErrorResponse(context, http.StatusBadRequest, err.Error())
		return
	}

	id := context.Query("id")
	if id == "" {
		sendErrorResponse(context, http.StatusBadRequest, "Missing id")
		return
	}

	order := schemas.Order{}
	if err := database.First(&order, "id = ?", id).Error; err != nil {
		sendErrorResponse(context, http.StatusNotFound, "Order not found")
		return
	}

	if request.Status != "" {
		order.Status = request.Status
	}

	if request.Notes != "" {
		order.Notes = request.Notes
	}

	if err := database.Save(&order).Error; err != nil {
		sendErrorResponse(context, http.StatusInternalServerError, "Error saving order")
		return
	}

	sendSuccessResponse(context, "Updated order", order)
}

package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"leChief/schemas"
	"net/http"
)

func CreateOrderHandler(context *gin.Context) {
	request := createOrderRequest{}

	err := context.BindJSON(&request)
	if err != nil {
		logger.ErrorFormatted("Error on bind json process: %v", err.Error())
		sendErrorResponse(context, http.StatusInternalServerError, err.Error())
		return
	}

	if err := request.Validate(); err != nil {
		logger.ErrorFormatted("Validation error: %v", err.Error())
		sendErrorResponse(context, http.StatusBadRequest, err.Error())
		return
	}

	order := schemas.Order{
		Id:     uuid.New(),
		Status: request.Status,
		Notes:  request.Notes,
	}

	if err := database.Create(&order).Error; err != nil {
		logger.ErrorFormatted("Failed to create a order: %v", request)
		sendErrorResponse(context, http.StatusInternalServerError, err.Error())
		return
	}

	sendSuccessResponse(context, "create-order", order)
}

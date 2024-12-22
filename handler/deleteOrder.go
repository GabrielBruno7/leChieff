package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"leChief/schemas"
	"net/http"
)

func DeleteOrderHandler(context *gin.Context) {
	orderId := context.Query("id")
	if orderId == "" {
		sendErrorResponse(context, http.StatusBadRequest, checkIfParamIsRequired("id", "Query Parameter").Error())
		return
	}

	order := schemas.Order{}

	if err := database.Where("id = ?", orderId).First(&order).Error; err != nil {
		logger.ErrorFormatted("Error in delete order process", err)
		sendErrorResponse(context, http.StatusNotFound, fmt.Sprintf("Order with id %s not found", orderId))
		return
	}

	if err := database.Delete(&order).Error; err != nil {
		logger.ErrorFormatted("Internal server error", err)
		sendErrorResponse(context, http.StatusInternalServerError, err.Error())
		return
	}

	sendSuccessResponse(context, "deleted", order)
}

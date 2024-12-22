package handler

import (
	"github.com/gin-gonic/gin"
	"leChief/schemas"
	"net/http"
)

func ShowOrderHandler(context *gin.Context) {
	orderId := context.Query("id")
	if orderId == "" {
		sendErrorResponse(context, http.StatusBadRequest, checkIfParamIsRequired("id", "Query Parameter").Error())
		return
	}

	order := schemas.Order{}

	if err := database.First(&order, "id = ?", orderId).Error; err != nil {
		sendErrorResponse(context, http.StatusBadRequest, err.Error())
	}

	sendSuccessResponse(context, "Show order", order)
}

package orders

import (
	"github.com/gin-gonic/gin"
	"leChief/handler"
	"leChief/schemas"
	"net/http"
)

func UpdateOrderHandler(context *gin.Context) {
	request := handler.UpdateOrderRequest{}

	context.BindJSON(&request)
	if err := request.ValidateOrder(); err != nil {
		handler.Logger.ErrorFormatted("Validation Error: %v", err.Error())
		handler.SendErrorResponse(context, http.StatusBadRequest, err.Error())
		return
	}

	id := context.Query("id")
	if id == "" {
		handler.SendErrorResponse(context, http.StatusBadRequest, "Missing id")
		return
	}

	order := schemas.Order{}
	if err := handler.Database.First(&order, "id = ?", id).Error; err != nil {
		handler.SendErrorResponse(context, http.StatusNotFound, "Order not found")
		return
	}

	if request.Status != "" {
		order.Status = request.Status
	}

	if err := handler.Database.Save(&order).Error; err != nil {
		handler.SendErrorResponse(context, http.StatusInternalServerError, "Error saving order")
		return
	}

	handler.SendSuccessResponse(context, "Updated order", order)
}

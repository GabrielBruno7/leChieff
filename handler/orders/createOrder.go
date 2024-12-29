package orders

import (
	"github.com/gin-gonic/gin"
	"leChief/handler"
	"leChief/schemas"
	"net/http"
)

func CreateOrderHandler(context *gin.Context) {
	request := handler.CreateOrderRequest{}

	err := context.BindJSON(&request)
	if err != nil {
		handler.Logger.ErrorFormatted("Error on bind json process: %v", err.Error())
		handler.SendErrorResponse(context, http.StatusInternalServerError, err.Error())
		return
	}

	var customer schemas.Customer
	if err := handler.Database.First(&customer, "id = ?", request.Customer_id).Error; err != nil {
		handler.Logger.ErrorFormatted("Customer not found: %v", request.Customer_id)
		handler.SendErrorResponse(context, http.StatusBadRequest, "Invalid customer ID")
		return
	}

	if err := request.Validate(); err != nil {
		handler.Logger.ErrorFormatted("Validation error: %v", err.Error())
		handler.SendErrorResponse(context, http.StatusBadRequest, err.Error())
		return
	}

	order := schemas.Order{
		Status:     request.Status,
		Notes:      request.Notes,
		CustomerID: request.Customer_id,
	}

	if err := handler.Database.Create(&order).Error; err != nil {
		handler.Logger.ErrorFormatted("Failed to create a order: %v", request)
		handler.SendErrorResponse(context, http.StatusInternalServerError, err.Error())
		return
	}

	handler.SendSuccessResponse(context, "create-order", order)
}

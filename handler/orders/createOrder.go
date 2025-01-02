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
		CustomerID: request.Customer_id,
	}

	if err := handler.Database.Create(&order).Error; err != nil {
		handler.Logger.ErrorFormatted("Failed to create a order: %v", request)
		handler.SendErrorResponse(context, http.StatusInternalServerError, err.Error())
		return
	}

	for _, product := range request.Products {
		var productDatabase schemas.Product
		if err := handler.Database.First(&productDatabase, "id = ?", product.ProductID).Error; err != nil {
			handler.Logger.ErrorFormatted("Product not found: %v", product.ProductID)
			handler.SendErrorResponse(context, http.StatusBadRequest, "Invalid product ID")
			return
		}

		productOrder := schemas.ProductOrder{
			OrderID:   order.ID,
			ProductID: product.ProductID,
			Quantity:  product.Quantity,
		}

		if err := handler.Database.Create(&productOrder).Error; err != nil {
			handler.Logger.ErrorFormatted("Failed to create a product order: %v", productOrder)
			handler.SendErrorResponse(context, http.StatusInternalServerError, err.Error())
			return
		}
	}

	handler.SendSuccessResponse(context, "Create order", order)
}

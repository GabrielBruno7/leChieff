package customers

import (
	"github.com/gin-gonic/gin"
	"leChief/handler"
	"leChief/schemas"
	"net/http"
)

func UpdateCustomerHandler(context *gin.Context) {
	request := handler.UpdateCustomerRequest{}

	context.BindJSON(&request)
	if err := request.ValidateCustomer(); err != nil {
		handler.Logger.ErrorFormatted("Validation Error: %v", err.Error())
		handler.SendErrorResponse(context, http.StatusBadRequest, err.Error())
		return
	}

	id := context.Query("id")
	if id == "" {
		handler.SendErrorResponse(context, http.StatusBadRequest, "Missing id")
		return
	}

	customer := schemas.Customer{}
	if err := handler.Database.First(&customer, "id = ?", id).Error; err != nil {
		handler.SendErrorResponse(context, http.StatusNotFound, "Customer not found")
		return
	}

	if request.Name != "" {
		customer.Name = request.Name
	}

	if request.Email != "" {
		customer.Email = request.Email
	}

	if request.Number != "" {
		customer.Number = request.Number
	}

	if request.Cep != "" {
		customer.Cep = request.Cep
	}

	if err := handler.Database.Save(&customer).Error; err != nil {
		handler.SendErrorResponse(context, http.StatusInternalServerError, "Error saving customer")
		return
	}

	handler.SendSuccessResponse(context, "Updated customer", customer)
}

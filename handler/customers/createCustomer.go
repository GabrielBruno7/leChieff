package customers

import (
	"github.com/gin-gonic/gin"
	"leChief/handler"
	"leChief/schemas"
	"net/http"
)

func CreateCustomerHandler(context *gin.Context) {
	request := handler.CreateCustomerRequest{}

	err := context.BindJSON(&request)
	if err != nil {
		handler.Logger.ErrorFormatted("Error on bind json process: %v", err.Error())
		handler.SendErrorResponse(context, http.StatusInternalServerError, err.Error())
		return
	}

	if err := request.Validate(); err != nil {
		handler.Logger.ErrorFormatted("Validation error: %v", err.Error())
		handler.SendErrorResponse(context, http.StatusBadRequest, err.Error())
		return
	}

	customer := schemas.Customer{
		Name:   request.Name,
		Number: request.Number,
		Email:  request.Email,
		Cep:    request.Cep,
	}

	if err := handler.Database.Create(&customer).Error; err != nil {
		handler.Logger.ErrorFormatted("Failed to create a order: %v", request)
		handler.SendErrorResponse(context, http.StatusInternalServerError, err.Error())
		return
	}

	handler.SendSuccessResponse(context, "Create customer", customer)
}

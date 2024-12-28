package customers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"leChief/handler"
	"leChief/schemas"
	"net/http"
)

func DeleteCustomerHandler(context *gin.Context) {
	customerId := context.Query("id")
	if customerId == "" {
		handler.SendErrorResponse(context, http.StatusBadRequest, handler.CheckIfParamIsRequired("id", "Query Parameter").Error())
		return
	}

	customer := schemas.Customer{}

	if err := handler.Database.Where("id = ?", customerId).First(&customer).Error; err != nil {
		handler.Logger.ErrorFormatted("Error in delete customer process", err)
		handler.SendErrorResponse(context, http.StatusNotFound, fmt.Sprintf("Customer with id %s not found", customerId))
		return
	}

	if err := handler.Database.Delete(&customer).Error; err != nil {
		handler.Logger.ErrorFormatted("Internal server error", err)
		handler.SendErrorResponse(context, http.StatusInternalServerError, err.Error())
		return
	}

	handler.SendSuccessResponse(context, "Customer deleted", customer)
}

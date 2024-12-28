package customers

import (
	"github.com/gin-gonic/gin"
	"leChief/handler"
	"leChief/schemas"
	"net/http"
)

func ShowCustomerHandler(context *gin.Context) {
	customerId := context.Query("id")

	if customerId == "" {
		handler.SendErrorResponse(context, http.StatusBadRequest, handler.CheckIfParamIsRequired("id", "Query Parameter").Error())
		return
	}

	customer := schemas.Customer{}

	if err := handler.Database.First(&customer, "id = ?", customerId).Error; err != nil {
		handler.SendErrorResponse(context, http.StatusBadRequest, err.Error())
		return
	}

	handler.SendSuccessResponse(context, "Show customer", customer)
}

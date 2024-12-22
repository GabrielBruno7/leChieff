package orders

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"leChief/handler"
	"leChief/schemas"
	"net/http"
)

func DeleteOrderHandler(context *gin.Context) {
	orderId := context.Query("id")
	if orderId == "" {
		handler.SendErrorResponse(context, http.StatusBadRequest, handler.CheckIfParamIsRequired("id", "Query Parameter").Error())
		return
	}

	order := schemas.Order{}

	if err := handler.Database.Where("id = ?", orderId).First(&order).Error; err != nil {
		handler.Logger.ErrorFormatted("Error in delete order process", err)
		handler.SendErrorResponse(context, http.StatusNotFound, fmt.Sprintf("Order with id %s not found", orderId))
		return
	}

	if err := handler.Database.Delete(&order).Error; err != nil {
		handler.Logger.ErrorFormatted("Internal server error", err)
		handler.SendErrorResponse(context, http.StatusInternalServerError, err.Error())
		return
	}

	handler.SendSuccessResponse(context, "deleted", order)
}

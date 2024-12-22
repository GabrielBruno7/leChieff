package orders

import (
	"github.com/gin-gonic/gin"
	"leChief/handler"
	"leChief/schemas"
	"net/http"
)

func ShowOrderHandler(context *gin.Context) {
	orderId := context.Query("id")
	if orderId == "" {
		handler.SendErrorResponse(context, http.StatusBadRequest, handler.CheckIfParamIsRequired("id", "Query Parameter").Error())
		return
	}

	order := schemas.Order{}

	if err := handler.Database.First(&order, "id = ?", orderId).Error; err != nil {
		handler.SendErrorResponse(context, http.StatusBadRequest, err.Error())
	}

	handler.SendSuccessResponse(context, "Show order", order)
}

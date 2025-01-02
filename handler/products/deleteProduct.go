package products

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"leChief/handler"
	"leChief/schemas"
	"net/http"
)

func DeleteProductHandler(context *gin.Context) {
	productId := context.Query("id")
	if productId == "" {
		handler.SendErrorResponse(context, http.StatusBadRequest, handler.CheckIfParamIsRequired("id", "Query Parameter").Error())
		return
	}

	product := schemas.Product{}

	if err := handler.Database.Where("id = ?", productId).First(&product).Error; err != nil {
		handler.Logger.ErrorFormatted("Error in delete product process", err)
		handler.SendErrorResponse(context, http.StatusNotFound, fmt.Sprintf("Order with id %s not found", productId))
		return
	}

	if err := handler.Database.Delete(&product).Error; err != nil {
		handler.Logger.ErrorFormatted("Internal server error", err)
		handler.SendErrorResponse(context, http.StatusInternalServerError, err.Error())
		return
	}

	context.String(http.StatusOK, "null")
}

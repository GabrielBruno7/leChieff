package products

import (
	"github.com/gin-gonic/gin"
	"leChief/handler"
	"leChief/schemas"
	"net/http"
)

func ShowProductsHandler(context *gin.Context) {
	productId := context.Query("id")
	if productId == "" {
		handler.SendErrorResponse(context, http.StatusBadRequest, handler.CheckIfParamIsRequired("id", "Query Parameter").Error())
		return
	}

	product := schemas.Product{}

	if err := handler.Database.First(&product, "id = ?", productId).Error; err != nil {
		handler.SendErrorResponse(context, http.StatusBadRequest, err.Error())
		return
	}

	context.JSON(http.StatusOK, product)
}

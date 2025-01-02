package products

import (
	"github.com/gin-gonic/gin"
	"leChief/handler"
	"leChief/schemas"
	"net/http"
)

func CreateProductHandler(context *gin.Context) {
	request := handler.CreateProductRequest{}

	err := context.BindJSON(&request)
	if err != nil {
		handler.Logger.ErrorFormatted("Error on bind json process: %v", err.Error())
		handler.SendErrorResponse(context, http.StatusInternalServerError, err.Error())
		return
	}

	if err := request.ValidateProduct(); err != nil {
		handler.Logger.ErrorFormatted("Validation error: %v", err.Error())
		handler.SendErrorResponse(context, http.StatusBadRequest, err.Error())
		return
	}

	product := schemas.Product{
		Name:        request.Name,
		Description: request.Description,
		Value:       request.Value,
		Type:        request.Type,
	}

	if err := handler.Database.Create(&product).Error; err != nil {
		handler.Logger.ErrorFormatted("Failed to create a product: %v", request)
		handler.SendErrorResponse(context, http.StatusInternalServerError, err.Error())
		return
	}

	handler.SendSuccessResponse(context, "Create product", product)
}

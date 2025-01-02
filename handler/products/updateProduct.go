package products

import (
	"github.com/gin-gonic/gin"
	"leChief/handler"
	"leChief/schemas"
	"net/http"
)

func UpdateProductHandler(context *gin.Context) {
	request := handler.UpdateProductRequest{}

	context.BindJSON(&request)
	if err := request.ValidateProduct(); err != nil {
		handler.Logger.ErrorFormatted("Validation Error: %v", err.Error())
		handler.SendErrorResponse(context, http.StatusBadRequest, err.Error())
		return
	}

	id := context.Query("id")
	if id == "" {
		handler.SendErrorResponse(context, http.StatusBadRequest, "Missing id")
		return
	}

	product := schemas.Product{}
	if err := handler.Database.First(&product, "id = ?", id).Error; err != nil {
		handler.SendErrorResponse(context, http.StatusNotFound, "Product not found")
		return
	}

	if request.Name != "" {
		product.Name = request.Name
	}

	if request.Description != "" {
		product.Description = request.Description
	}

	if request.Value != 0 {
		product.Value = request.Value
	}

	if request.Type != "" {
		product.Type = request.Type
	}

	if err := handler.Database.Save(&product).Error; err != nil {
		handler.SendErrorResponse(context, http.StatusInternalServerError, "Error saving product")
		return
	}

	handler.SendSuccessResponse(context, "Updated product", product)
}

package products

import (
	"github.com/gin-gonic/gin"
	"leChief/handler"
	"leChief/schemas"
	"math"
	"net/http"
)

func PriceProductHandler(context *gin.Context) {
	request := handler.PriceProductRequest{}

	context.BindJSON(&request)
	if err := request.ValidatePriceProduct(); err != nil {
		handler.Logger.ErrorFormatted("Validation Error: %v", err.Error())
		handler.SendErrorResponse(context, http.StatusBadRequest, err.Error())
		return
	}

	id := context.Query("id")
	if id == "" {
		handler.SendErrorResponse(context, http.StatusBadRequest, "Missing product id")
		return
	}

	product := schemas.Product{}
	if err := handler.Database.First(&product, "id = ?", id).Error; err != nil {
		handler.SendErrorResponse(context, http.StatusNotFound, "Product not found")
		return
	}

	product = calculateProductPrice(request, product)

	if err := handler.Database.Save(&product).Error; err != nil {
		handler.SendErrorResponse(context, http.StatusInternalServerError, "Error saving product")
		return
	}

	handler.SendSuccessResponse(context, "Product Priced !", product)
}

func calculateProductPrice(request handler.PriceProductRequest, product schemas.Product) schemas.Product {
	var ingredientsTotalValue float32

	for _, value := range request.Ingredients {
		ingredientsTotalValue += value
	}

	totalCost := ingredientsTotalValue + request.Labor + request.Expenses

	product.Value = float32(math.Round(float64(totalCost*(1+request.Profit/100))*100) / 100)

	return product
}

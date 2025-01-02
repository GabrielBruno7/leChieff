package products

import (
	"github.com/gin-gonic/gin"
	"leChief/handler"
	"leChief/schemas"
	"net/http"
)

func ListProductsHandler(context *gin.Context) {
	product := []schemas.Product{}

	if err := handler.Database.Find(&product).Order("created_at ASC").Error; err != nil {
		handler.SendErrorResponse(context, http.StatusInternalServerError, "Error listing products")
		return
	}

	handler.SendSuccessResponse(context, "Listing products", product)
}

func ListOrdersWhenStatusIsLoadingHandler(context *gin.Context) {
	var results []schemas.Order

	query := "SELECT * FROM orders WHERE status LIKE ?"

	if err := handler.Database.Raw(query, "loading").Scan(&results).Error; err != nil {
		handler.SendErrorResponse(context, http.StatusInternalServerError, "Error listing products")
		return
	}

	handler.SendSuccessResponse(context, "Listing orders", results)
}

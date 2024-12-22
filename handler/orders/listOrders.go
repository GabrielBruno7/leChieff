package orders

import (
	"github.com/gin-gonic/gin"
	"leChief/handler"
	"leChief/schemas"
	"net/http"
)

func ListOrdersHandler(context *gin.Context) {
	orders := []schemas.Order{}

	if err := handler.Database.Find(&orders).Order("created_at ASC").Error; err != nil {
		handler.SendErrorResponse(context, http.StatusInternalServerError, "Error listing orders")
		return
	}

	handler.SendSuccessResponse(context, "Listing orders", orders)
}

func ListOrdersWhenStatusIsLoadingHandler(context *gin.Context) {
	var results []schemas.Order

	query := "SELECT * FROM orders WHERE status LIKE ?"

	if err := handler.Database.Raw(query, "loading").Scan(&results).Error; err != nil {
		handler.SendErrorResponse(context, http.StatusInternalServerError, "Error listing orders")
		return
	}

	handler.SendSuccessResponse(context, "Listing orders", results)
}

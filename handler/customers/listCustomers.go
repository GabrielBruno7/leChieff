package customers

import (
	"github.com/gin-gonic/gin"
	"leChief/handler"
	"leChief/schemas"
	"net/http"
)

func ListCustomersHandler(context *gin.Context) {
	customers := []schemas.Customer{}

	if err := handler.Database.Find(&customers).Order("created_at ASC").Error; err != nil {
		handler.SendErrorResponse(context, http.StatusInternalServerError, "Error listing customers")
		return
	}

	handler.SendSuccessResponse(context, "Listing customers", customers)
}

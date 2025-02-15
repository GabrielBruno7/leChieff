package customers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"leChief/handler"
	"leChief/schemas"
	"net/http"
	"regexp"
)

const MAX_EMAIL_LENGTH = 55

const NUMBER_REGEX = `^\d{11}$`
const NAME_REGEX = `^[a-zA-ZÀ-ÿ\s]{1,35}$`
const EMAIL_REGEX = `^[a-zA-Z0-9._%+-]+@([a-zA-Z0-9-]+\.)+[a-zA-Z]{2,}$`

func CreateCustomerHandler(context *gin.Context) {
	request := handler.CreateCustomerRequest{}

	err := context.BindJSON(&request)
	if err != nil {
		handler.Logger.ErrorFormatted("Error on bind json process: %v", err.Error())
		handler.SendErrorResponse(context, http.StatusInternalServerError, err.Error())
		return
	}

	if err := request.Validate(); err != nil {
		handler.Logger.ErrorFormatted("Validation error: %v", err.Error())
		handler.SendErrorResponse(context, http.StatusBadRequest, err.Error())
		return
	}

	customer := schemas.Customer{
		Name:   request.Name,
		Number: request.Number,
		Email:  request.Email,
		Cep:    request.Cep,
	}

	if err := validateCustomerFields(customer); err != nil {
		handler.SendErrorResponse(context, http.StatusBadRequest, err.Error())
		return
	}

	if err := handler.Database.Create(&customer).Error; err != nil {
		handler.Logger.ErrorFormatted("Failed to create a order: %v", request)
		handler.SendErrorResponse(context, http.StatusInternalServerError, err.Error())
		return
	}

	handler.SendSuccessResponse(context, "Create customer", customer)
}

func validateCustomerFields(customer schemas.Customer) error {
	if !isValidField(customer.Name, NAME_REGEX) {
		handler.Logger.ErrorFormatted("Customer name '%v' is invalid", customer.Name)
		return fmt.Errorf("invalid customer name. It should contain only letters")
	}

	if !isValidField(customer.Number, NUMBER_REGEX) {
		handler.Logger.ErrorFormatted("Customer number '%v' is invalid", customer.Number)
		return fmt.Errorf("invalid customer number. It should be exactly 11 digits and contain only numbers")
	}

	if !isValidField(customer.Email, EMAIL_REGEX) {
		handler.Logger.ErrorFormatted("Customer e-mail '%v' is invalid", customer.Email)
		return fmt.Errorf("invalid customer email format")
	}

	if len(customer.Email) > MAX_EMAIL_LENGTH {
		return fmt.Errorf("invalid customer email. It should have at most %d characters", MAX_EMAIL_LENGTH)
	}

	return nil
}

func isValidField(field string, regex string) bool {
	re := regexp.MustCompile(regex)
	return re.MatchString(field)
}

package handler

import (
	"fmt"
)

type CreateOrderRequest struct {
	Status      string           `json:"status"`
	Customer_id string           `json:"customer_id"`
	Products    []ProductRequest `json:"products" binding:"required,dive"`
}

type ProductRequest struct {
	ProductID string `json:"product_id" binding:"required"`
	Quantity  int    `json:"quantity" binding:"required"`
}

func (request *CreateOrderRequest) Validate() error {
	if request.Status == "" {
		return CheckIfParamIsRequired("status", "string")
	}
	if request.Customer_id == "" {
		return CheckIfParamIsRequired("customer_id", "string")
	}

	return nil
}

type UpdateOrderRequest struct {
	Status string `json:"status"`
	Notes  string `json:"notes"`
}

func (request *UpdateOrderRequest) ValidateOrder() error {
	if request.Status != "" || request.Notes != "" {
		return nil
	}

	return fmt.Errorf("at least one valid field must provided")
}

func (request *UpdateCustomerRequest) ValidateCustomer() error {
	if request.Name != "" || request.Email != "" || request.Cep != "" || request.Number != "" {
		return nil
	}

	return fmt.Errorf("at least one valid field must provided")
}

func CheckIfParamIsRequired(param, typ string) error {
	return fmt.Errorf("the param: %s (type: %s) is required", param, typ)
}

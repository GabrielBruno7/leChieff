package handler

import (
	"fmt"
)

type createOrderRequest struct {
	Status string `json:"status"`
	Notes  string `json:"notes"`
}

func (request *createOrderRequest) Validate() error {
	if request.Status == "" {
		return checkIfParamIsRequired("status", "string")
	}

	if request.Notes == "" {
		return checkIfParamIsRequired("notes", "string")
	}

	return nil
}

type updateOrderRequest struct {
	Status string `json:"status"`
	Notes  string `json:"notes"`
}

func (request *updateOrderRequest) Validate() error {
	if request.Status != "" || request.Notes != "" {
		return nil
	}

	return fmt.Errorf("at least one valid field must provided")
}

func checkIfParamIsRequired(param, typ string) error {
	return fmt.Errorf("the param: %s (type: %s) is required", param, typ)
}

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

func checkIfParamIsRequired(param, typ string) error {
	return fmt.Errorf("the param: %s (type: %s) is required", param, typ)
}

package handler

import "leChief/schemas"

type CreateProductRequest struct {
	Name        string           `json:"name"`
	Description string           `json:"description"`
	Value       float32          `json:"value"`
	Type        schemas.FoodType `json:"type"`
}

type UpdateProductRequest struct {
	Name        string           `json:"name"`
	Description string           `json:"description"`
	Value       float32          `json:"value"`
	Type        schemas.FoodType `json:"type"`
}

func (request *CreateProductRequest) ValidateProduct() error {
	if request.Name == "" {
		return CheckIfParamIsRequired("name", "string")
	}

	if request.Description == "" {
		return CheckIfParamIsRequired("description", "string")
	}

	if request.Value <= 0 {
		return CheckIfParamIsRequired("value", "float32 (greater than 0)")
	}

	productType := request.Type
	if productType == "" {
		return CheckIfParamIsRequired("type", "string (valid type)")
	}
	if productType != schemas.RUSK && productType != schemas.CAKE && productType != schemas.PANETTONE {
		return CheckIfParamIsRequired("type", "The type of product can be 'RUSK', 'CAKE', or 'PANETTONE'")
	}

	return nil
}
func (request *UpdateProductRequest) ValidateProduct() error {
	if request.Name == "" {
		return CheckIfParamIsRequired("name", "string")
	}

	if request.Description == "" {
		return CheckIfParamIsRequired("description", "string")
	}

	if request.Value <= 0 {
		return CheckIfParamIsRequired("value", "float32 (greater than 0)")
	}

	productType := request.Type
	if productType == "" {
		return CheckIfParamIsRequired("type", "string (valid type)")
	}
	if productType != schemas.RUSK && productType != schemas.CAKE && productType != schemas.PANETTONE {
		return CheckIfParamIsRequired("type", "The type of product can be 'RUSK', 'CAKE', or 'PANETTONE'")
	}

	return nil
}

package handler

import (
	"leChief/schemas"
)

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

type PriceProductRequest struct {
	Profit      float32            `json:"profit"`
	Labor       float32            `json:"labor"`
	Expenses    float32            `json:"expenses"`
	Ingredients map[string]float32 `json:"ingredients"`
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

func (request *PriceProductRequest) ValidatePriceProduct() error {
	if request.Profit <= 0 {
		return CheckIfParamIsRequired("profit", "float32 (greater than 0)")
	}

	if request.Labor <= 0 {
		return CheckIfParamIsRequired("labor", "float32 (greater than 0)")
	}

	if request.Expenses <= 0 {
		return CheckIfParamIsRequired("expenses", "float32 (greater than 0)")
	}

	if request.Ingredients == nil || len(request.Ingredients) == 0 {
		return CheckIfParamIsRequired("ingredients", "map[string]float32 (at least one ingredient required)")
	}

	for name, value := range request.Ingredients {
		if value <= 0 {
			return CheckIfParamIsRequired("ingredients."+name, "float32 (greater than 0)")
		}
	}

	return nil
}

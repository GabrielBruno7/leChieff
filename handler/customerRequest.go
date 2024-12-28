package handler

type CreateCustomerRequest struct {
	Name   string `json:"name"`
	Number string `json:"number"`
	Email  string `json:"email"`
	Cep    string `json:"cep"`
}

func (request *CreateCustomerRequest) Validate() error {
	if request.Name == "" {
		return CheckIfParamIsRequired("name", "string")
	}
	if request.Number == "" {
		return CheckIfParamIsRequired("number", "string")
	}
	if request.Email == "" {
		return CheckIfParamIsRequired("email", "string")
	}
	if request.Cep == "" {
		return CheckIfParamIsRequired("cep", "string")
	}

	return nil
}

type UpdateCustomerRequest struct {
	Name   string `json:"name"`
	Number string `json:"number"`
	Email  string `json:"email"`
	Cep    string `json:"cep"`
}

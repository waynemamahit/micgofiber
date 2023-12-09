package lib

import "github.com/go-playground/validator/v10"

type ValidateResponse struct {
	FailedField string
	Tag         string
	Value       string
}

func (resp *ValidateResponse) ValidateStruct(obj any) []*ValidateResponse {
	validate := validator.New()
	var errors []*ValidateResponse
	err := validate.Struct(obj)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			resp.FailedField = err.StructNamespace()
			resp.Tag = err.Tag()
			resp.Value = err.Param()
			errors = append(errors, resp)
		}
	}
	return errors
}

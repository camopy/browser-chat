package validator

import (
	"fmt"

	"github.com/go-playground/validator"
)

type ErrResponse struct {
	Errors []string `json:"errors"`
}

func New() *validator.Validate {
	validate := validator.New()
	validate.SetTagName("form")

	return validate
}

func ToErrResponse(err error) *ErrResponse {
	if fieldErrors, ok := err.(validator.ValidationErrors); ok {
		resp := ErrResponse{
			Errors: make([]string, len(fieldErrors)),
		}

		for i, err := range fieldErrors {
			switch err.Tag() {
			case "required":
				resp.Errors[i] = fmt.Sprintf("%s is required", err.Field())
			default:
				resp.Errors[i] = fmt.Sprintf("Something is wrong with the field %s; %s", err.Field(), err.Tag())
			}
		}

		return &resp
	}

	return nil
}

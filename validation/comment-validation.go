package validation

import (
	"github.com/go-playground/validator/v10"
	"github.com/husfuu/go-gram/dto"
)

func ValidateCreateComment(input dto.RequestComment) error {
	err := validate.Struct(input)
	if err != nil {
		validationErrors := err.(validator.ValidationErrors)
		for _, fieldError := range validationErrors {
			return fieldError
		}
	}
	return nil
}

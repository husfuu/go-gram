package validation

import (
	"github.com/go-playground/validator/v10"
	"github.com/husfuu/go-gram/dto"
	repository "github.com/husfuu/go-gram/repository/socialmediaRepository"
)

func ValidateIsSocialMediaExist(socialmediaID string, r repository.SocialMediaRepository) error {
	err := r.IsSocialMediaExist(socialmediaID)

	if err != nil {
		return nil
	}

	return err
}

func ValidateCreateSocialMedia(input dto.RequestSocialMedia) error {
	err := validate.Struct(input)
	if err != nil {
		validationErrors := err.(validator.ValidationErrors)
		for _, fieldError := range validationErrors {
			return fieldError
		}
	}
	return nil
}

func ValidateUpdateSocialMedia(input dto.RequestSocialMedia, r repository.SocialMediaRepository) error {
	err := validate.Struct(input)
	if err != nil {
		validationErrors := err.(validator.ValidationErrors)
		for _, fieldError := range validationErrors {
			return fieldError
		}
	}

	err = ValidateIsSocialMediaExist(input.ID, r)
	if err != nil {
		return err
	}

	return nil
}

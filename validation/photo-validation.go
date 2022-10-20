package validation

import (
	"github.com/go-playground/validator/v10"
	"github.com/husfuu/go-gram/dto"
	repository "github.com/husfuu/go-gram/repository/photoRepository"
)

func ValidateIsPhotoExist(photoID string, r repository.PhotoRepository) error {
	err := r.IsPhotoExist(photoID)
	return err
}

func ValidateCreatePhoto(input dto.RequestPhoto) error {
	err := validate.Struct(input)
	if err != nil {
		validationErrors := err.(validator.ValidationErrors)
		for _, fieldError := range validationErrors {
			return fieldError
		}
	}
	return nil
}

func ValidateUpdatePhoto(input dto.RequestPhoto, r repository.PhotoRepository) error {
	err := validate.Struct(input)
	if err != nil {
		validationErrors := err.(validator.ValidationErrors)
		for _, fieldError := range validationErrors {
			return fieldError
		}
	}
	err = ValidateIsPhotoExist(input.ID, r)
	if err != nil {
		return err
	}

	return nil
}

package photoService

import (
	"github.com/google/uuid"
	"github.com/husfuu/go-gram/dto"
	"github.com/husfuu/go-gram/entity"
	"github.com/husfuu/go-gram/repository/photoRepository"
	"github.com/husfuu/go-gram/validation"
	"github.com/jinzhu/copier"
)

type PhotoService interface {
	Create(input dto.RequestPhoto) (dto.ResponseCreatePhoto, error)
	GetPhotos() ([]dto.ResponseGetPhoto, error)
	Update(input dto.RequestPhoto, photoID string) (dto.ResponseUpdatePhoto, error)
	Delete(photoID string) error
}

type service struct {
	repository photoRepository.PhotoRepository
}

func NewPhotoService(repository photoRepository.PhotoRepository) PhotoService {
	return &service{repository: repository}
}

func (s service) Create(input dto.RequestPhoto) (dto.ResponseCreatePhoto, error) {
	err := validation.ValidateCreatePhoto(input)
	if err != nil {
		return dto.ResponseCreatePhoto{}, err
	}

	photo := new(entity.Photo)
	copier.Copy(&photo, &input)

	photo.ID = uuid.New().String()

	newPhoto, err := s.repository.Create(*photo)

	if err != nil {
		return dto.ResponseCreatePhoto{}, err
	}

	response := dto.ResponseCreatePhoto{}
	copier.Copy(&response, &newPhoto)

	return response, nil
}

func (s service) GetPhotos() ([]dto.ResponseGetPhoto, error) {
	photos, err := s.repository.GetPhotos()
	if err != nil {
		return []dto.ResponseGetPhoto{}, nil
	}

	var response []dto.ResponseGetPhoto
	for _, photo := range photos {
		tempPhoto := dto.ResponseGetPhoto{}
		tempPhoto.ID = photo.ID
		tempPhoto.Title = photo.Title
		tempPhoto.Caption = photo.Caption
		tempPhoto.PhotoURL = photo.PhotoURL
		//tempPhoto.CreatedAt = photo.CreatedAt
		tempPhoto.User.Username = photo.User.Username
		tempPhoto.User.Email = photo.User.Email
		response = append(response, tempPhoto)
	}
	return response, nil
}

func (s service) Update(input dto.RequestPhoto, photoID string) (dto.ResponseUpdatePhoto, error) {
	err := validation.ValidateUpdatePhoto(input, s.repository)

	if err != nil {
		return dto.ResponseUpdatePhoto{}, err
	}

	photo := new(entity.Photo)
	copier.Copy(&photo, &input)
	photo.ID = photoID

	updatedPhoto, err := s.repository.Update(*photo)
	if err != nil {
		return dto.ResponseUpdatePhoto{}, err
	}
	response := dto.ResponseUpdatePhoto{}
	copier.Copy(&response, &updatedPhoto)

	return response, nil
}

func (s service) Delete(photoID string) error {
	err := s.repository.DeleteByID(photoID)
	if err != nil {
		return nil
	}
	return nil
}

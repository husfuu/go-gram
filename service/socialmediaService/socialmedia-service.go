package socialmediaService

import (
	"errors"

	"github.com/google/uuid"
	"github.com/husfuu/go-gram/dto"
	"github.com/husfuu/go-gram/entity"
	"github.com/husfuu/go-gram/helper"
	"github.com/husfuu/go-gram/repository/photoRepository"
	"github.com/husfuu/go-gram/repository/socialmediaRepository"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

type SocialMediaService interface {
	Create(input dto.RequestSocialMedia) (dto.ResponseCreateSocialMedia, error)
	GetSocialMedias() ([]dto.ResponseGetSocialMedias, error)
	Update(input dto.RequestSocialMedia) (dto.ResponseCreateSocialMedia, error)
	Delete(id string) error
}

type service struct {
	socialMediaRepository socialmediaRepository.SocialMediaRepository
	photoRepository       photoRepository.PhotoRepository
}

func NewSocialMediaService(
	repository socialmediaRepository.SocialMediaRepository,
	photo_repository photoRepository.PhotoRepository,
) SocialMediaService {
	return &service{
		socialMediaRepository: repository,
		photoRepository:       photo_repository,
	}
}

func (s service) Create(input dto.RequestSocialMedia) (dto.ResponseCreateSocialMedia, error) {
	socialMedia := new(entity.SocialMedia)
	copier.Copy(&socialMedia, &input)

	socialMedia.ID = uuid.New().String()
	socialMedia.CreatedAt = helper.TimeNowMillis
	socialMedia.UpdatedAt = helper.TimeNowMillis
	newSocialMedia, err := s.socialMediaRepository.Create(*socialMedia)

	if err != nil {
		return dto.ResponseCreateSocialMedia{}, err
	}
	response := dto.ResponseCreateSocialMedia{}
	copier.Copy(&response, &newSocialMedia)

	return response, nil
}

func (s service) GetSocialMedias() ([]dto.ResponseGetSocialMedias, error) {
	socialMedias, err := s.socialMediaRepository.GetSocialMedias()

	if err != nil {
		return []dto.ResponseGetSocialMedias{}, err
	}

	response := []dto.ResponseGetSocialMedias{}
	for _, socialMedia := range socialMedias {
		// get photo by userID
		photo, err := s.photoRepository.GetPhotoByUserID(socialMedia.UserID)

		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return []dto.ResponseGetSocialMedias{}, err
		}
		temp := new(dto.ResponseGetSocialMedias)
		copier.Copy(&temp, &socialMedia)
		if photo.PhotoURL != "" {
			temp.User.ProfileImageUrl = photo.PhotoURL
		}
		response = append(response, *temp)
	}

	return response, nil
}

func (s service) Update(input dto.RequestSocialMedia) (dto.ResponseCreateSocialMedia, error) {
	socialMedia := new(entity.SocialMedia)

	copier.Copy(&socialMedia, &input)
	socialMedia.UpdatedAt = helper.TimeNowMillis
	updatedSocialMedia, err := s.socialMediaRepository.Update(*socialMedia)
	if err != nil {
		return dto.ResponseCreateSocialMedia{}, err
	}

	response := dto.ResponseCreateSocialMedia{}
	copier.Copy(&response, &updatedSocialMedia)

	return response, nil
}

func (s service) Delete(id string) error {
	err := s.socialMediaRepository.DeleteByID(id)
	if err != nil {
		return err
	}
	return nil
}

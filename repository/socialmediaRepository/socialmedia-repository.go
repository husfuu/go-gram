package socialmediaRepository

import (
	"fmt"

	"github.com/husfuu/go-gram/entity"
	"gorm.io/gorm"
)

type SocialMediaRepository interface {
	Create(socialMedia entity.SocialMedia) (entity.SocialMedia, error)
	GetSocialMedias() ([]entity.SocialMedia, error)
	Update(socialMedia entity.SocialMedia) (entity.SocialMedia, error)
	DeleteByID(id string) error
}

type repository struct {
	db *gorm.DB
}

func NewSocialMediaRepository(db *gorm.DB) SocialMediaRepository {
	return &repository{db: db}
}

func (r repository) Create(socialMedia entity.SocialMedia) (entity.SocialMedia, error) {
	err := r.db.Debug().Create(&socialMedia).Error
	if err != nil {
		return entity.SocialMedia{}, err
	}
	return socialMedia, nil
}

func (r repository) GetSocialMedias() ([]entity.SocialMedia, error) {
	var socialmedias []entity.SocialMedia
	err := r.db.Debug().Preload("User").Find(&socialmedias).Error
	if err != nil {
		return []entity.SocialMedia{}, err
	}
	fmt.Println("ini sosmed dari repo", socialmedias)
	return socialmedias, nil
}

func (r repository) Update(socialMedia entity.SocialMedia) (entity.SocialMedia, error) {
	err := r.db.Debug().Where("id = ?", socialMedia.ID).Updates(
		entity.SocialMedia{
			Name:           socialMedia.Name,
			SocialMediaUrl: socialMedia.SocialMediaUrl,
		},
	).Error

	if err != nil {
		return entity.SocialMedia{}, err
	}
	return socialMedia, nil
}

func (r repository) DeleteByID(id string) error {
	socialMedia := entity.SocialMedia{}
	socialMedia.ID = id
	return r.db.Debug().First(&socialMedia).Where("id = ?", socialMedia.ID).Delete(&socialMedia).Error
}

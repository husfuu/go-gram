package socialmediaRepository

import (
	"errors"
	"fmt"

	"github.com/husfuu/go-gram/entity"
	"gorm.io/gorm"
)

type SocialMediaRepository interface {
	Create(socialMedia entity.SocialMedia) (entity.SocialMedia, error)
	GetSocialMedias() ([]entity.SocialMedia, error)
	Update(socialMedia entity.SocialMedia) (entity.SocialMedia, error)
	DeleteByID(id string) error
	IsSocialMediaExist(id string) error
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
		return entity.SocialMedia{}, errors.New("user already have social media")
	}
	return socialMedia, nil
}

func (r repository) GetSocialMedias() ([]entity.SocialMedia, error) {
	var socialmedias []entity.SocialMedia
	err := r.db.Debug().Preload("User").Find(&socialmedias).Error
	if err != nil {
		return []entity.SocialMedia{}, err
	}

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

func (r repository) IsSocialMediaExist(id string) error {
	var socialmedia entity.SocialMedia

	err := r.db.Where("id = ?", id).First(&socialmedia).Error
	fmt.Println(socialmedia)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("social media doesn't exists")
		}
		return err
	}
	return nil
}

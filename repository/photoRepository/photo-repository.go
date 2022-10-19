package photoRepository

import (
	"github.com/husfuu/go-gram/entity"
	"gorm.io/gorm"
)

type PhotoRepository interface {
	Create(photo entity.Photo) (entity.Photo, error)
	GetPhotos() ([]entity.Photo, error)
	Update(photo entity.Photo) (entity.Photo, error)
	DeleteByID(id string) error
	GetPhotoByUserID(id string) (entity.Photo, error)
}

type repository struct {
	db *gorm.DB
}

func NewPhotoRepository(db *gorm.DB) PhotoRepository {
	return &repository{db}
}

func (r repository) Create(photo entity.Photo) (entity.Photo, error) {
	err := r.db.Debug().Create(&photo).Error

	if err != nil {
		return entity.Photo{}, err
	}
	return photo, err
}

func (r repository) GetPhotos() ([]entity.Photo, error) {
	var photos []entity.Photo
	err := r.db.Debug().Preload("User").Find(&photos).Error
	if err != nil {
		return []entity.Photo{}, err
	}
	return photos, nil
}

func (r repository) Update(photo entity.Photo) (entity.Photo, error) {
	err := r.db.Debug().Where("id = ?", photo.ID).Updates(
		entity.Photo{
			Title:    photo.Title,
			Caption:  photo.Caption,
			PhotoURL: photo.PhotoURL,
		}).Error
	if err != nil {
		return entity.Photo{}, err
	}
	return photo, nil
}

func (r repository) DeleteByID(id string) error {
	photo := entity.Photo{}
	photo.ID = id
	return r.db.Debug().First(&photo).Where("id = ?", photo.ID).Delete(&photo).Error
}

func (r repository) GetPhotoByUserID(userID string) (entity.Photo, error) {
	var photo entity.Photo
	err := r.db.Where("user_id = ?", userID).Find(&photo).Error
	if err != nil {
		return entity.Photo{}, err
	}
	return photo, nil
}

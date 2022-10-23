package userRepository

import (
	"errors"
	"github.com/husfuu/go-gram/entity"
	"github.com/husfuu/go-gram/helper"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user entity.User) (entity.User, error)
	FindByEmail(email string) (entity.User, error)
	FindByID(ID string) (entity.User, error)
	Update(user entity.User) (entity.User, error)
	DeleteByID(ID string) error
	IsEmailExist(email string) error
	IsUsernameExist(username string) error
}

type repository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &repository{db: db}
}

func (r *repository) IsUsernameExist(username string) error {
	var user entity.User
	err := r.db.Where("username = ?", username).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil
		}
		return err
	}
	return helper.ErrorUsernameAlreadyExists
}

func (r *repository) IsEmailExist(email string) error {
	var user entity.User
	err := r.db.Where("email = ?", email).First(&user).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil
		}
		return err
	}
	return helper.ErrorEmailAlreadyExists
}

func (r *repository) Create(user entity.User) (entity.User, error) {
	err := r.db.Create(&user).Error

	if err != nil {
		return entity.User{}, err
	}
	return user, nil
}

func (r *repository) FindByEmail(email string) (entity.User, error) {
	var user entity.User

	err := r.db.Where("email = ?", email).Find(&user).Error

	if err != nil {
		return entity.User{}, err
	}

	return user, nil
}

func (r *repository) FindByID(ID string) (entity.User, error) {
	var user entity.User
	err := r.db.Where("id = ?", ID).Find(&user).Error

	if err != nil {
		return entity.User{}, err
	}

	return user, nil
}

func (r *repository) Update(user entity.User) (entity.User, error) {

	err := r.db.Debug().Where("id = ?", user.ID).Updates(
		entity.User{Username: user.Username, Email: user.Email, Password: user.Password, Age: user.Age},
	).Error

	if err != nil {
		return entity.User{}, err
	}
	return user, nil
}

func (r *repository) DeleteByID(ID string) error {
	user := new(entity.User)
	user.ID = ID
	return r.db.Debug().First(&user).Where("id = ?", user.ID).Delete(&user).Error
}

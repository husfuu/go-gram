package userService

import (
	"github.com/google/uuid"
	"github.com/husfuu/go-gram/auth"
	"github.com/husfuu/go-gram/dto"
	"github.com/husfuu/go-gram/entity"
	"github.com/husfuu/go-gram/helper"
	"github.com/husfuu/go-gram/repository/userRepository"
	"github.com/husfuu/go-gram/validation"
	"github.com/jinzhu/copier"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	RegisterUser(input dto.RequestRegister) (dto.Response, error)
	Login(input dto.RequestLogin) (dto.ResponseLogin, error)
	Update(input dto.RequestRegister) (dto.Response, error)
	DeleteByID(input string) error
}

type service struct {
	repository userRepository.UserRepository
}

func NewUserService(repository userRepository.UserRepository) UserService {
	return &service{repository: repository}
}

func (s *service) RegisterUser(input dto.RequestRegister) (dto.Response, error) {
	err := validation.ValidateUserCreate(input, s.repository)

	if err != nil {
		return dto.Response{}, err
	}

	user := entity.User{}

	user.ID = uuid.New().String()
	user.Username = input.Username
	user.Email = input.Email
	user.Age = input.Age

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if err != nil {
		return dto.Response{}, err
	}

	user.Password = string(passwordHash)

	newUser, err := s.repository.Create(user)
	if err != nil {
		return dto.Response{}, err
	}
	response := dto.Response{}
	copier.Copy(&response, &newUser)
	return response, nil
}

func (s *service) Login(input dto.RequestLogin) (dto.ResponseLogin, error) {
	err := validation.ValidateUserLogin(input)
	if err != nil {
		return dto.ResponseLogin{}, err
	}
	email := input.Email
	password := input.Password
	user, err := s.repository.FindByEmail(email)

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return dto.ResponseLogin{}, helper.ErrorInvalidLogin
	}

	jwtService := auth.NewJWTService()

	token, err := jwtService.GenerateToken(user.ID)

	response := dto.ResponseLogin{}
	response.Token = token
	return response, nil
}

func (s *service) Update(input dto.RequestRegister) (dto.Response, error) {
	user := entity.User{}
	user.ID = input.ID
	user.Username = input.Username
	user.Email = input.Email
	user.Age = input.Age
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)

	if err != nil {
		return dto.Response{}, err
	}

	user.Password = string(passwordHash)

	updatedUser, err := s.repository.Update(user)
	if err != nil {
		return dto.Response{}, err
	}
	response := dto.Response{}
	copier.Copy(&response, &updatedUser)

	return response, nil
}

func (s *service) DeleteByID(id string) error {
	return s.repository.DeleteByID(id)
}

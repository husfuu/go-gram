package userRepository

import (
	"testing"

	"github.com/husfuu/go-gram/config"
	"github.com/husfuu/go-gram/entity"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
)

const EnvTestPath = "../../.env.test"

type UserRepositoryTestSuite struct {
	suite.Suite
	db             *gorm.DB
	repository     UserRepository
	defaultPayload entity.User
}

func TestUserRepository(t *testing.T) {
	err := godotenv.Load(EnvTestPath)
	assert.NoError(t, err)
	// DB initialization
	db, err := config.NewDbInit()
	assert.NoError(t, err)

	err = db.AutoMigrate(entity.User{})
	assert.NoError(t, err)

	deleteDBUser(db)

	repository := NewUserRepository(db)

	defaultPayload := entity.User{
		ID:       "2f48e1a6-b403-47c1-a48b-26ad61f8c0b2",
		Username: "husni_fu_fu_fu",
		Email:    "husfuudevTEST@gmail.com",
		Age:      20,
		Password: "husni_gans",
	}

	testSuite := &UserRepositoryTestSuite{
		db:             db,
		repository:     repository,
		defaultPayload: defaultPayload,
	}
	suite.Run(t, testSuite)
}

func (suite *UserRepositoryTestSuite) Test_A_repository_create() {
	suite.T().Run("Create user success", func(t *testing.T) {
		newUser, err := suite.repository.Create(suite.defaultPayload)
		assert.NoError(t, err)
		assert.NotEmpty(t, newUser.ID)
	})

	suite.T().Run("Create user with the same ID", func(t *testing.T) {
		_, err := suite.repository.Create(suite.defaultPayload)
		assert.Error(t, err)
	})
}

func (suite *UserRepositoryTestSuite) Test_B_repository_IsUsernameExists() {
	suite.T().Run("username exist", func(t *testing.T) {
		err := suite.repository.IsUsernameExist(suite.defaultPayload.Username)
		assert.Error(t, err)
	})

	suite.T().Run("username doesn't exists", func(t *testing.T) {
		err := suite.repository.IsUsernameExist("doesnt_exists_username")
		assert.NoError(t, err)
	})
}

func (suite *UserRepositoryTestSuite) Test_C_repository_IsEmailExists() {
	suite.T().Run("email exists", func(t *testing.T) {
		err := suite.repository.IsEmailExist(suite.defaultPayload.Email)
		assert.Error(t, err)
	})

	suite.T().Run("email doesn't exists", func(t *testing.T) {
		err := suite.repository.IsEmailExist("emailgakada@gmail.com")
		assert.NoError(t, err)
	})
}

func (suite *UserRepositoryTestSuite) Test_D_repository_Update() {
	payloadUpdate := entity.User{
		ID:       suite.defaultPayload.ID,
		Username: "tanakafuu",
		Email:    "tanakafuu@yahoo.com",
		Password: "tanakafuuPassword",
	}

	suite.T().Run("update success", func(t *testing.T) {
		updatedUser, err := suite.repository.Update(payloadUpdate)
		assert.NoError(t, err)
		assert.NotEmpty(t, updatedUser)
		assert.NotEqual(t, updatedUser.Username, suite.defaultPayload.Username)
		assert.NotEqual(t, updatedUser.Email, suite.defaultPayload.Email)
		assert.Equal(t, updatedUser.UpdatedAt, suite.defaultPayload.UpdatedAt)
	})
}

func (suite *UserRepositoryTestSuite) Test_E_repository_Delete() {
	suite.T().Run("delete success", func(t *testing.T) {
		err := suite.repository.DeleteByID(suite.defaultPayload.ID)
		assert.NoError(t, err)
	})
}

func deleteDBUser(db *gorm.DB) {
	db.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(entity.User{})
}

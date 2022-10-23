package photoRepository

import (
	"github.com/husfuu/go-gram/config"
	"github.com/husfuu/go-gram/entity"
	"github.com/husfuu/go-gram/helper"
	"github.com/husfuu/go-gram/repository/userRepository"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
	"testing"
)

const EnvTestPath = "../../.env.test"

type PhotoRepositoryTestSuite struct {
	suite.Suite
	db             *gorm.DB
	repository     PhotoRepository
	defaultPayload entity.Photo
}

func TestPhotoRepository(t *testing.T) {
	err := godotenv.Load(EnvTestPath)
	assert.NoError(t, err)
	// DB initialization
	db, err := config.NewDbInit()
	assert.NoError(t, err)

	err = db.AutoMigrate(entity.User{})
	assert.NoError(t, err)
	deleteData(db)

	repository := NewPhotoRepository(db)

	initUser := entity.User{
		ID:       "2f48e1a6-b403-47c1-a48b-26ad61f8c0b2",
		Username: "husni_fu_fu_fu",
		Email:    "husfuudevTEST@gmail.com",
		Age:      20,
		Password: "husfuuPass",
	}

	userRepository := userRepository.NewUserRepository(db)
	newUser, err := userRepository.Create(initUser)
	assert.NoError(t, err)

	defaultPayload := entity.Photo{
		ID:       helper.UUID,
		Title:    "waifuu",
		Caption:  "looks my pretty waifu",
		PhotoURL: "images.com/waifuu",
		UserID:   newUser.ID,
	}
	testSuite := &PhotoRepositoryTestSuite{
		db:             db,
		repository:     repository,
		defaultPayload: defaultPayload,
	}
	suite.Run(t, testSuite)
}

func (suite *PhotoRepositoryTestSuite) Test_A_repository_create() {
	suite.T().Run("create photo success", func(t *testing.T) {
		newPhoto, err := suite.repository.Create(suite.defaultPayload)
		assert.NoError(t, err)
		assert.NotEmpty(t, newPhoto)
		// update defaultPayload ID
		suite.defaultPayload.ID = newPhoto.ID
	})
}

func (suite *PhotoRepositoryTestSuite) Test_B_repository_isPhotoExists() {
	suite.T().Run("photo exists", func(t *testing.T) {
		err := suite.repository.IsPhotoExist(suite.defaultPayload.ID)
		assert.NoError(t, err)
	})

	suite.T().Run("photo doesn't exists", func(t *testing.T) {
		err := suite.repository.IsPhotoExist("does_not_exists_id")
		assert.Error(t, err)
	})
}

func (suite *PhotoRepositoryTestSuite) Test_C_repository_getPhotos() {
	suite.T().Run("get all photos success", func(t *testing.T) {
		photos, err := suite.repository.GetPhotos()
		assert.NoError(t, err)
		photosExpected := []entity.Photo{}
		photosExpected = append(photosExpected, suite.defaultPayload)

		assert.Equal(t, photosExpected[0].ID, photos[0].ID)
		assert.Equal(t, photosExpected[0].UserID, photos[0].UserID)
		assert.Equal(t, photosExpected[0].Title, photos[0].Title)
		assert.Equal(t, photosExpected[0].Caption, photos[0].Caption)
		assert.Equal(t, photosExpected[0].PhotoURL, photos[0].PhotoURL)
	})
}

func (suite *PhotoRepositoryTestSuite) Test_D_repository_getPhotobyUserID() {
	suite.T().Run("get photo by userID success", func(t *testing.T) {
		photo, err := suite.repository.GetPhotoByUserID(suite.defaultPayload.UserID)
		assert.NoError(t, err)
		assert.Equal(t, suite.defaultPayload.ID, photo.ID)
	})
}

func (suite *PhotoRepositoryTestSuite) Test_E_repository_update() {
	payloadUpdate := entity.Photo{
		ID:       suite.defaultPayload.ID,
		Title:    "kei karuizawa my waifuu",
		Caption:  "guys, look my new waifuu",
		PhotoURL: "image.com/karuizawa_waifuu",
	}

	suite.T().Run("update photo success", func(t *testing.T) {
		updatedPhoto, err := suite.repository.Update(payloadUpdate)
		assert.NoError(t, err)
		assert.NotEmpty(t, updatedPhoto)
		assert.NotEqual(t, updatedPhoto.Title, suite.defaultPayload.Title)
		assert.NotEqual(t, updatedPhoto.PhotoURL, suite.defaultPayload.PhotoURL)
		assert.Equal(t, updatedPhoto.UpdatedAt, suite.defaultPayload.UpdatedAt)
	})
}

func (suite *PhotoRepositoryTestSuite) Test_F_repository_delete() {
	suite.T().Run("delete photo success", func(t *testing.T) {
		err := suite.repository.DeleteByID(suite.defaultPayload.ID)
		assert.NoError(t, err)
	})
}

func deleteData(db *gorm.DB) {
	db.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(entity.Photo{})
	db.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(entity.User{})
}

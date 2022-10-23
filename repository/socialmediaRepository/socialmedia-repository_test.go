package socialmediaRepository

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

type SocialMediaRepositoryTestSuite struct {
	suite.Suite
	db             *gorm.DB
	repository     SocialMediaRepository
	defaultPayload entity.SocialMedia
}

func TestSocialMediaRepository(t *testing.T) {
	err := godotenv.Load(EnvTestPath)
	assert.NoError(t, err)
	// DB initialization
	db, err := config.NewDbInit()
	assert.NoError(t, err)

	err = db.AutoMigrate(entity.User{})
	assert.NoError(t, err)
	deleteData(db)

	repository := NewSocialMediaRepository(db)

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

	defaultPayload := entity.SocialMedia{
		ID:             helper.UUID,
		Name:           "tanakafuu",
		SocialMediaUrl: "twitter.com/tanakafuu",
		UserID:         newUser.ID,
	}
	testSuite := &SocialMediaRepositoryTestSuite{
		db:             db,
		repository:     repository,
		defaultPayload: defaultPayload,
	}

	suite.Run(t, testSuite)
}

func (suite *SocialMediaRepositoryTestSuite) Test_A_repository_create() {
	suite.T().Run("Create social media success", func(t *testing.T) {
		newSocialMedia, err := suite.repository.Create(suite.defaultPayload)
		assert.NoError(t, err)
		assert.NotEmpty(t, newSocialMedia)
		// update defaultPayload ID
		suite.defaultPayload.ID = newSocialMedia.ID
	})

	suite.T().Run("user already have social media", func(t *testing.T) {
		_, err := suite.repository.Create(suite.defaultPayload)
		assert.Error(t, err)
	})
}

func (suite *SocialMediaRepositoryTestSuite) Test_B_reposity_getsocialmedias() {
	suite.T().Run("get social media success", func(t *testing.T) {
		socialMedias, err := suite.repository.GetSocialMedias()
		assert.NoError(t, err)
		socialMediasExpected := []entity.SocialMedia{}
		socialMediasExpected = append(socialMediasExpected, suite.defaultPayload)

		assert.Equal(t, socialMediasExpected[0].ID, socialMedias[0].ID)
		assert.Equal(t, socialMediasExpected[0].UserID, socialMedias[0].UserID)
		assert.Equal(t, socialMediasExpected[0].Name, socialMedias[0].Name)
	})
}

//func (suite *SocialMediaRepositoryTestSuite) Test_C_IsSocialMediaExist() {
//	suite.T().Run("social media is exists", func(t *testing.T) {
//		err := suite.repository.IsSocialMediaExist(suite.defaultPayload.ID)
//		assert.Error(t, err)
//	})
//
//	suite.T().Run("social media is not exists", func(t *testing.T) {
//		err := suite.repository.IsSocialMediaExist("this_id_is_not_exists")
//		assert.NoError(t, err)
//	})
//}

func (suite *SocialMediaRepositoryTestSuite) Test_D_repository_updateSocialMedia() {
	payloadUpdate := entity.SocialMedia{
		ID:             suite.defaultPayload.ID,
		Name:           "keikaruizawa",
		SocialMediaUrl: "twitter.com/kei_karuizawa",
	}

	suite.T().Run("update social media success", func(t *testing.T) {
		updatedSocialMedia, err := suite.repository.Update(payloadUpdate)
		assert.NoError(t, err)
		assert.NotEmpty(t, updatedSocialMedia)
		assert.NotEqual(t, updatedSocialMedia.Name, suite.defaultPayload.Name)
		assert.NotEqual(t, updatedSocialMedia.SocialMediaUrl, suite.defaultPayload.SocialMediaUrl)
		assert.Equal(t, updatedSocialMedia.UpdatedAt, suite.defaultPayload.UpdatedAt)
	})
}

func (suite *SocialMediaRepositoryTestSuite) Test_E_repository_deleteSocialMedia() {
	suite.T().Run("delete social media success", func(t *testing.T) {
		err := suite.repository.DeleteByID(suite.defaultPayload.ID)
		assert.NoError(t, err)
	})
}

func deleteData(db *gorm.DB) {
	db.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(entity.SocialMedia{})
	db.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(entity.User{})
}

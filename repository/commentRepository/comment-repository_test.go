package commentRepository

import (
	"github.com/husfuu/go-gram/config"
	"github.com/husfuu/go-gram/entity"
	"github.com/husfuu/go-gram/helper"
	"github.com/husfuu/go-gram/repository/photoRepository"
	"github.com/husfuu/go-gram/repository/userRepository"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
	"testing"
)

const EnvTestPath = "../../.env.test"

type CommentRepositoryTestSuite struct {
	suite.Suite
	db             *gorm.DB
	repository     CommentRepository
	defaultPayload entity.Comment
}

func TestCommentRepository(t *testing.T) {
	err := godotenv.Load(EnvTestPath)
	assert.NoError(t, err)
	// DB initialization
	db, err := config.NewDbInit()
	assert.NoError(t, err)

	err = db.AutoMigrate(entity.User{})
	assert.NoError(t, err)
	deleteData(db)

	repository := NewCommentRepository(db)

	initUser := entity.User{
		ID:       "2f48e1a6-b403-47c1-a48b-26ad61f8c0b2",
		Username: "husni_fu_fu_fu",
		Email:    "husfuudevTEST@gmail.com",
		Age:      20,
		Password: "husfuuPass",
	}
	userRepo := userRepository.NewUserRepository(db)
	newUser, err := userRepo.Create(initUser)
	assert.NoError(t, err)

	initPhoto := entity.Photo{
		ID:       helper.UUID,
		Title:    "waifuu",
		Caption:  "looks my pretty waifu",
		PhotoURL: "images.com/waifuu",
		UserID:   newUser.ID,
	}
	photoRepo := photoRepository.NewPhotoRepository(db)
	newPhoto, err := photoRepo.Create(initPhoto)
	assert.NoError(t, err)

	defaultPayload := entity.Comment{
		ID:      helper.UUID,
		UserID:  initUser.ID,
		PhotoID: newPhoto.ID,
		Message: "your waifu has been claimed",
	}

	testSuite := &CommentRepositoryTestSuite{
		db:             db,
		repository:     repository,
		defaultPayload: defaultPayload,
	}
	suite.Run(t, testSuite)
}

func (suite *CommentRepositoryTestSuite) Test_A_repository_create() {
	suite.T().Run("create comment success", func(t *testing.T) {
		newComment, err := suite.repository.Create(suite.defaultPayload)
		assert.NoError(t, err)
		assert.NotEmpty(t, newComment)
		// update defaultPayload ID
		suite.defaultPayload.ID = newComment.ID
	})
}

func (suite *CommentRepositoryTestSuite) Test_B_repository_getComments() {
	suite.T().Run("get all comments success", func(t *testing.T) {
		comments, err := suite.repository.Get()
		assert.NoError(t, err)

		commentsExpected := []entity.Comment{}
		commentsExpected = append(commentsExpected, suite.defaultPayload)

		assert.Equal(t, commentsExpected[0].ID, comments[0].ID)
		assert.Equal(t, commentsExpected[0].UserID, comments[0].UserID)
		assert.Equal(t, commentsExpected[0].PhotoID, comments[0].PhotoID)
		assert.Equal(t, commentsExpected[0].Message, comments[0].Message)
	})
}

func (suite *CommentRepositoryTestSuite) Test_C_repository_update() {
	payloadUpdate := entity.Comment{
		ID:      suite.defaultPayload.ID,
		Message: "your waifuu has been claimed by me",
	}

	suite.T().Run("update comment success", func(t *testing.T) {
		updatedComment, err := suite.repository.Update(payloadUpdate)
		assert.NoError(t, err)
		assert.NotEmpty(t, updatedComment)
		assert.NotEqual(t, updatedComment.Message, suite.defaultPayload.Message)
	})
}

func (suite *CommentRepositoryTestSuite) Test_D_repository_delete() {
	suite.T().Run("delete comment success", func(t *testing.T) {
		err := suite.repository.Delete(suite.defaultPayload.ID)
		assert.NoError(t, err)
	})
}

func deleteData(db *gorm.DB) {
	db.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(entity.Comment{})
	db.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(entity.Photo{})
	db.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(entity.User{})
}

package commentRepository

import (
	"github.com/husfuu/go-gram/entity"
	"gorm.io/gorm"
)

type CommentRepository interface {
	Create(comment entity.Comment) (entity.Comment, error)
	Get() ([]entity.Comment, error)
	Update(comment entity.Comment) (entity.Comment, error)
	Delete(id string) error
}

type repository struct {
	db *gorm.DB
}

func NewCommentRepository(db *gorm.DB) CommentRepository {
	return &repository{db}
}

func (r repository) Create(comment entity.Comment) (entity.Comment, error) {
	err := r.db.Create(&comment).Error
	if err != nil {
		return entity.Comment{}, err
	}
	return comment, nil
}

func (r repository) Get() ([]entity.Comment, error) {
	var comment []entity.Comment
	err := r.db.Preload("User").Preload("Photo").Find(&comment).Error
	if err != nil {
		return []entity.Comment{}, err
	}
	return comment, nil
}

func (r repository) Update(comment entity.Comment) (entity.Comment, error) {
	err := r.db.Debug().Where("id = ?", comment.ID).Updates(&comment).First(&comment).Error
	if err != nil {
		return entity.Comment{}, err
	}
	return comment, nil
}

func (r repository) Delete(id string) error {
	comment := entity.Comment{}
	comment.ID = id

	return r.db.Debug().First(&comment).Where("id = ?", comment.ID).Delete(&comment).Error
}

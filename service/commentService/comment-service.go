package commentservice

import (
	"github.com/google/uuid"
	"github.com/husfuu/go-gram/dto"
	"github.com/husfuu/go-gram/entity"
	"github.com/husfuu/go-gram/helper"
	"github.com/husfuu/go-gram/repository/commentRepository"
	"github.com/husfuu/go-gram/validation"
	"github.com/jinzhu/copier"
)

type CommentService interface {
	Create(input dto.RequestComment) (dto.ResponseCreateComment, error)
	GetComments() ([]dto.ResponseGetComment, error)
	Update(input dto.RequestCommentUpdate, commentID string) (dto.ResponseCreateComment, error)
	Delete(commentID string) error
}

type service struct {
	repository commentRepository.CommentRepository
}

func NewCommentService(repository commentRepository.CommentRepository) CommentService {
	return &service{repository: repository}
}

func (s service) Create(input dto.RequestComment) (dto.ResponseCreateComment, error) {
	err := validation.ValidateCreateComment(input)

	if err != nil {
		return dto.ResponseCreateComment{}, err
	}

	var comment entity.Comment
	copier.Copy(&comment, &input)
	comment.ID = uuid.New().String()
	comment.CreatedAt = helper.TimeNowMillis
	comment.UpdatedAt = helper.TimeNowMillis
	newComment, err := s.repository.Create(comment)
	if err != nil {
		return dto.ResponseCreateComment{}, err
	}
	var response dto.ResponseCreateComment
	copier.Copy(&response, &newComment)
	return response, nil
}

func (s service) GetComments() ([]dto.ResponseGetComment, error) {
	comments, err := s.repository.Get()

	if err != nil {
		return []dto.ResponseGetComment{}, err
	}
	response := []dto.ResponseGetComment{}
	for _, comment := range comments {

		var temp dto.ResponseGetComment
		copier.Copy(&temp, &comment)
		response = append(response, temp)
	}
	return response, err
}

func (s service) Update(input dto.RequestCommentUpdate, commentID string) (dto.ResponseCreateComment, error) {
	var comment entity.Comment
	copier.Copy(&comment, &input)
	comment.ID = commentID
	comment.UpdatedAt = helper.TimeNowMillis

	updatedComment, err := s.repository.Update(comment)

	if err != nil {
		return dto.ResponseCreateComment{}, err
	}
	var response dto.ResponseCreateComment
	copier.Copy(&response, &updatedComment)
	return response, nil
}

func (s service) Delete(commentID string) error {
	err := s.repository.Delete(commentID)
	if err != nil {
		return err
	}
	return nil
}

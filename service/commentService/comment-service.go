package commentservice

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/husfuu/go-gram/dto"
	"github.com/husfuu/go-gram/entity"
	"github.com/husfuu/go-gram/repository/commentRepository"
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
	var comment entity.Comment
	copier.Copy(&comment, &input)
	comment.ID = uuid.New().String()

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
	fmt.Println("ini semua comment:", comments)
	if err != nil {
		return []dto.ResponseGetComment{}, err
	}
	var response []dto.ResponseGetComment

	for _, comment := range comments {
		fmt.Println("ini setiap comment:  ", comment)
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

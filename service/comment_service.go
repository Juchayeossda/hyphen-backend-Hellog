package service

import (
	"context"
	"hyphen-backend-hellog/cerrors"
	"hyphen-backend-hellog/model"
	"hyphen-backend-hellog/repository"
	"hyphen-backend-hellog/verifier"
)

type CommentService interface {
	Create(ctx context.Context, m *model.CommentCreate)
}

func NewCommentService(commentRepo repository.CommentRepository, userRepo repository.UserRepository, postRepo repository.PostRepository) CommentService {
	return &IPcommentService{commentRepo, userRepo, postRepo}
}

type IPcommentService struct {
	repository.CommentRepository
	repository.UserRepository
	repository.PostRepository
}

func (s *IPcommentService) Create(ctx context.Context, m *model.CommentCreate) {
	verifier.Validate(m)

	author, err := s.UserRepository.QueryByID(ctx, 4)
	if err != nil {
		panic(cerrors.NotFoundError{ErrorMessage: err.Error()})
	}

	post, err := s.PostRepository.QueryByID(ctx, m.PostID)
	if err != nil {
		panic(cerrors.NotFoundError{ErrorMessage: err.Error()})
	}

	if m.ParentID != 0 {

		parent, err := s.CommentRepository.QueryByID(ctx, m.ParentID)
		if err != nil {
			panic(cerrors.NotFoundError{ErrorMessage: err.Error()})
		}

		if _, err := s.CommentRepository.Create(ctx, m.Content, post, author, parent); err != nil {
			panic(cerrors.NotCreateError{ErrorMessage: err.Error()})
		}

	} else {
		if _, err := s.CommentRepository.Create(ctx, m.Content, post, author, nil); err != nil {
			panic(cerrors.NotCreateError{ErrorMessage: err.Error()})
		}
	}

}

// func (s *IPcommentService) Update(ctx context.Context, m *model.UpdateComment)

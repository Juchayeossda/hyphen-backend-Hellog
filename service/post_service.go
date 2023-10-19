package service

import (
	"context"
	"fmt"
	"hyphen-backend-hellog/cerrors"
	"hyphen-backend-hellog/model"
	"hyphen-backend-hellog/repository"
	"hyphen-backend-hellog/request/siss"
	"hyphen-backend-hellog/verifier"
)

type PostService interface {
	Create(ctx context.Context, m *model.PostCreate)
	SelectByID(ctx context.Context, m *model.PostSelecByID) *model.PostSelectByIDReturn
	Update(ctx context.Context, m *model.PostUpdate)
}

func NewPostRepository(postRepo repository.PostRepository, userRepo repository.UserRepository, commentRepo repository.CommentRepository) PostService {
	return &IPostService{postRepo, userRepo, commentRepo}
}

type IPostService struct {
	repository.PostRepository
	repository.UserRepository
	repository.CommentRepository
}

func (s *IPostService) Create(ctx context.Context, m *model.PostCreate) {
	verifier.Validate(m)

	previewImage, err := siss.RegisterImage(m.PreviewImage)
	if err != nil {
		panic(cerrors.ReqeustFailedError{ErrorMessage: err.Error()})
	}

	// TODO: token parsing -> get user id or name
	author, err := s.UserRepository.QueryByID(ctx, 4 /*toekn id 여기로*/)
	if err != nil {
		panic(cerrors.NotFoundError{ErrorMessage: err.Error()})
	}

	if _, err := s.PostRepository.Create(ctx, m.Title, m.Content, previewImage, m.IsPrivate, author); err != nil {
		panic(cerrors.NotCreateError{ErrorMessage: err.Error()})
	}
}

func (s *IPostService) SelectByID(ctx context.Context, m *model.PostSelecByID) *model.PostSelectByIDReturn {
	verifier.Validate(m)

	var returnModel model.PostSelectByIDReturn

	post, err := s.PostRepository.QueryByID(ctx, m.PostID)
	if err != nil {
		panic(cerrors.NotFoundError{ErrorMessage: err.Error()})
	}

	// post 전체
	returnModel.ID = post.ID
	returnModel.Title = post.Title
	returnModel.Content = post.Content
	returnModel.PreviewImage = post.PreviewImage
	returnModel.IsPrivate = post.IsPrivate
	returnModel.CreatedAt = post.CreatedAt
	returnModel.UpdatedAt = post.UpdatedAt

	author, err := s.UserRepository.QueryByPost(ctx, post)
	if err != nil {
		fmt.Println(author)
		panic(cerrors.NotFoundError{ErrorMessage: err.Error()})
	}

	returnModel.Author.ID = author.ID
	returnModel.Author.Name = author.Name
	returnModel.Author.ProfileImage = author.ProfileImage

	return &returnModel

}

func (s *IPostService) Update(ctx context.Context, m *model.PostUpdate) {
	verifier.Validate(m)

	prevPost, err := s.PostRepository.QueryByID(ctx, m.PostID)
	if err != nil {
		panic(cerrors.NotCreateError{ErrorMessage: err.Error()})
	}

	err = siss.RemoveImage(prevPost.PreviewImage)
	if err != nil {
		panic(cerrors.ReqeustFailedError{ErrorMessage: err.Error()})
	}

	previewImage, err := siss.RegisterImage(m.PreviewImage)
	if err != nil {
		panic(cerrors.ReqeustFailedError{ErrorMessage: err.Error()})
	}

	// TODO: token parsing -> get user id or name
	author, err := s.UserRepository.QueryByID(ctx, 4 /*toekn id 여기로*/)
	if err != nil {
		panic(cerrors.NotFoundError{ErrorMessage: err.Error()})
	}

	if _, err := s.PostRepository.UpdateByID(ctx, m.PostID, m.Title, m.Content, previewImage, m.IsPrivate, author); err != nil {
		panic(err)
	}

}

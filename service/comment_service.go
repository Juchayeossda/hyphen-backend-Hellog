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
	SelectByPost(ctx context.Context, m *model.CommentSelectByPost) *model.CommentSelectByPostReturn
	UpdateByID(ctx context.Context, m *model.CommentUpdateByID)
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

func (s *IPcommentService) SelectByPost(ctx context.Context, m *model.CommentSelectByPost) *model.CommentSelectByPostReturn {
	verifier.Validate(m)

	var returnModel model.CommentSelectByPostReturn

	post, err := s.PostRepository.QueryByID(ctx, m.PostID)
	if err != nil {
		panic(cerrors.NotFoundError{ErrorMessage: err.Error()})
	}

	// 댓글 전체
	parents, err := s.CommentRepository.QueryParentsByPost(ctx, post)
	if err != nil {
		panic(cerrors.NotFoundError{ErrorMessage: err.Error()})
	}

	for _, parent := range parents {
		// 해당 포스트의 부모 댓글
		var parentCommentModel model.Comment

		// 부모 댓글의 하위 댓글들
		childs, err := s.CommentRepository.QueryChilds(ctx, parent)
		if err != nil {
			panic(cerrors.NotFoundError{ErrorMessage: err.Error()})
		}

		// 부모 댓굴의 하위 댓글 무한 반복
		for _, child := range childs {
			// 부모 댓굴의 하위 댓글을 담을 변수
			var childComentModel model.CommentOfComment

			comment, err := s.CommentRepository.QueryByID(ctx, child.ID)
			if err != nil {
				panic(cerrors.NotFoundError{ErrorMessage: err.Error()})
			}

			childComentModel.ID = comment.ID
			childComentModel.Content = comment.Content

			childComentModel.CreatedAt = comment.CreatedAt
			childComentModel.UpdatedAt = comment.UpdatedAt

			author, err := s.UserRepository.QueryByPost(ctx, post)
			if err != nil {
				panic(cerrors.NotFoundError{ErrorMessage: err.Error()})
			}

			childComentModel.Author.ID = author.ID
			childComentModel.Author.Name = author.Name
			childComentModel.Author.ProfileImage = author.ProfileImage

			// 부모의 하위 댓글을 담는 변수에 하위 댓글 저장
			parentCommentModel.CommentOfComments = append(parentCommentModel.CommentOfComments, childComentModel)
		}

		// 이제 부모 댓글의 댓글 내용 채워줘
		parentComment, err := s.CommentRepository.QueryByID(ctx, parent.ID)
		if err != nil {
			panic(cerrors.NotFoundError{ErrorMessage: err.Error()})
		}

		parentCommentModel.ID = parentComment.ID
		parentCommentModel.Content = parentComment.Content

		parentCommentModel.CreatedAt = parentComment.CreatedAt
		parentCommentModel.UpdatedAt = parentComment.UpdatedAt

		parentAuthor, err := s.UserRepository.QueryByPost(ctx, post)
		if err != nil {
			panic(cerrors.NotFoundError{ErrorMessage: err.Error()})
		}

		parentCommentModel.Author.ID = parentAuthor.ID
		parentCommentModel.Author.Name = parentAuthor.Name
		parentCommentModel.Author.ProfileImage = parentAuthor.ProfileImage

		returnModel.Comments = append(returnModel.Comments, parentCommentModel)
	}

	return &returnModel
}

func (s *IPcommentService) UpdateByID(ctx context.Context, m *model.CommentUpdateByID) {
	verifier.Validate(m)

	if _, err := s.CommentRepository.UpdateByID(ctx, m.CommentID, m.Content); err != nil {
		panic(cerrors.NotUpdateError{ErrorMessage: err.Error()})
	}
}

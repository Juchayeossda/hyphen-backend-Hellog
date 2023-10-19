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
	Create(ctx context.Context, m *model.PostCreateUpdate)
	SelectByID(ctx context.Context, m *model.PostSelecByID) model.PostSelectByIDReturn
}

func NewPostRepository(postRepo repository.PostRepository, userRepo repository.UserRepository, commentRepo repository.CommentRepository) PostService {
	return &IPostService{postRepo, userRepo, commentRepo}
}

type IPostService struct {
	repository.PostRepository
	repository.UserRepository
	repository.CommentRepository
}

func (s *IPostService) Create(ctx context.Context, m *model.PostCreateUpdate) {
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

func CreateComment() {}

func (s *IPostService) SelectByID(ctx context.Context, m *model.PostSelecByID) model.PostSelectByIDReturn {
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

	return returnModel

}
func SelectPostComments()                                                     {}
func (s *IPostService) Update(ctx context.Context, m *model.PostCreateUpdate) {}
func UpdateComment()                                                          {}
func (s *IPostService) Delete(ctx context.Context, m *model.PostDelete)       {}
func DleteComment()                                                           {}

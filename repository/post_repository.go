package repository

import (
	"context"
	"hyphen-backend-hellog/ent"
	"time"
)

type PostRepository interface {
	Create(ctx context.Context, title string, content string, previewImage string, isPrivate bool, author *ent.User) (*ent.Post, error)
	QueryByID(ctx context.Context, id int) (*ent.Post, error)
	UpdateByID(ctx context.Context, id int, title string, content string, previewImage string, isPrivate bool, author *ent.User) (*ent.Post, error)
	DeleteByID(ctx context.Context, id int) error
}

func NewPostRepository(database *ent.Client) PostRepository {
	return &IPostRepository{client: database}
}

type IPostRepository struct {
	client *ent.Client
}

func (r *IPostRepository) Create(ctx context.Context, title string, content string, previewImage string, isPrivate bool, author *ent.User) (*ent.Post, error) {
	return r.client.Post.Create().
		SetTitle(title).
		SetContent(content).
		SetPreviewImage(previewImage).
		SetIsPrivate(isPrivate).
		SetAuthor(author).
		Save(ctx)
}

func (r *IPostRepository) QueryByID(ctx context.Context, id int) (*ent.Post, error) {
	return r.client.Post.Get(ctx, id)
}

func (r *IPostRepository) UpdateByID(ctx context.Context, id int, title string, content string, previewImage string, isPrivate bool, author *ent.User) (*ent.Post, error) {

	updatePost := r.client.Post.UpdateOneID(id).
		SetTitle(title).
		SetContent(content).
		SetPreviewImage(previewImage).
		SetIsPrivate(isPrivate).
		SetUpdatedAt(time.Now())

	if author != nil {
		updatePost.SetAuthor(author)
	}

	return updatePost.Save(ctx)
}

func (r *IPostRepository) DeleteByID(ctx context.Context, id int) error {
	return r.client.Post.DeleteOneID(id).
		Exec(ctx)
}

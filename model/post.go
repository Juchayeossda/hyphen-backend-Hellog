package model

import (
	"mime/multipart"
	"time"
)

type PostCreate struct {
	Title        string                `form:"title"  validate:"required"`
	Content      string                `form:"content" validate:"required"`
	PreviewImage *multipart.FileHeader `form:"preview_image" validate:"required"`
	IsPrivate    bool                  `form:"is_private" validate:"boolean"`
}

type PostUpdate struct {
	PostID       int                   `validate:"required"`
	Title        string                `form:"title"  validate:"required"`
	Content      string                `form:"content" validate:"required"`
	PreviewImage *multipart.FileHeader `form:"preview_image" validate:"required"`
	IsPrivate    bool                  `form:"is_private" validate:"boolean"`
}

type PostSelecByID struct {
	PostID int `json:"post_id"`
}

type PostSelectByIDReturn struct {
	Post `json:"post"`
}

type Post struct {
	ID           int    `json:"id"`
	Title        string `json:"title"`
	Content      string `json:"content"`
	PreviewImage string `json:"preview_image"`
	IsPrivate    bool   `json:"is_private"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	Author `json:"author"`
}

type Author struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	ProfileImage string `json:"profile_image"`
}

type PostDelete struct {
	PostID int `json:"post_id"`
}

type PostDeleteByID struct {
	PostID int `json:"post_id"`
}

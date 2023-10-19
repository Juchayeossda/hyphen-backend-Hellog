package model

import (
	"mime/multipart"
	"time"
)

type PostCreateUpdate struct {
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

	Author   `json:"author"`
	Comments []Comment `json:"comments"`
}

type Comment struct {
	ID      int    `json:"id"`
	Content string `json:"content"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	Author            `json:"author"`
	CommentOfComments []CommentOfComment `json:"coment_of_comments"`
}

type CommentOfComment struct {
	ID      int    `json:"id"`
	Content string `json:"content"`

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

package model

import "time"

type CommentCreate struct {
	PostID   int    `json:"post_id"`
	Content  string `json:"content"`
	ParentID int    `json:"parent_id"`
}

type CommentSelectByPost struct {
	PostID int `json:"comment_id"`
}

type CommentSelectByPostReturn struct {
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

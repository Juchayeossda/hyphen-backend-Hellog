package model

type CommentCreate struct {
	PostID   int    `json:"post_id"`
	Content  string `json:"content"`
	ParentID int    `json:"parent_id"`
}

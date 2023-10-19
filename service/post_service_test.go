package service

import (
	"hyphen-backend-hellog/initializer"
	"testing"
)

func TestCreatePost(t *testing.T) {

	component := initializer.NewCompnent("../.env")
	initializer.NewDatabase(component)

	// postRepo := repository.NewPostRepository(db)
	// userRepo := repository.NewUserRepository(db)
	// commentRepo := repository.NewCommentRepository(db)

	// postService := NewPostRepository(postRepo, userRepo, commentRepo)

}

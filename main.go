package main

import (
	"hyphen-backend-hellog/cerrors"
	"hyphen-backend-hellog/controller"
	"hyphen-backend-hellog/initializer"
	"hyphen-backend-hellog/repository"
	"hyphen-backend-hellog/service"

	"github.com/gofiber/fiber/v2"
)

func main() {
	component := initializer.NewCompnent()
	database := initializer.NewDatabase(component)

	userRepo := repository.NewUserRepository(database)
	postRepo := repository.NewPostRepository(database)
	commendRepo := repository.NewCommentRepository(database)

	postService := service.NewPostRepository(postRepo, userRepo, commendRepo)
	commentService := service.NewCommentService(commendRepo, userRepo, postRepo)

	commentCtl := controller.NewCommentController(commentService)
	postCtl := controller.NewPostController(postService)

	app := fiber.New(fiber.Config{
		ErrorHandler: cerrors.ErrorHandler,
	})

	// app.Use(recover.New())

	postCtl.Route(app)
	commentCtl.Route(app)

	app.Listen(":8080")

}

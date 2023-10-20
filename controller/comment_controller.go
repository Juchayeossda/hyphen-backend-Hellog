package controller

import (
	"hyphen-backend-hellog/cerrors/exception"
	"hyphen-backend-hellog/model"
	"hyphen-backend-hellog/service"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func NewCommentController(commentServ service.CommentService) *CommentController {
	return &CommentController{commentServ}
}

type CommentController struct {
	service.CommentService
}

func (ctr *CommentController) Route(app *fiber.App) {
	app.Post("/api/hellog/posts/:post_id/comments/comment", ctr.create)
	app.Get("/api/hellog/posts/:post_id/comments", ctr.selectByID)
	app.Patch("/api/hellog/posts/:post_id/comments/comment", ctr.updateByID)
	app.Delete("/api/hellog/posts/:post_id/comments/comment", ctr.deleteByID)
}

func (ctr *CommentController) create(c *fiber.Ctx) error {
	var clientRequest model.CommentCreate

	//content, PaerntID parsing
	err := c.BodyParser(&clientRequest)
	exception.Sniff(err)

	// is_private parsing
	clientRequest.PostID, err = strconv.Atoi(c.Params("post_id"))
	exception.Sniff(err)

	ctr.CommentService.Create(c.Context(), &clientRequest)

	return c.Status(fiber.StatusCreated).JSON(model.GeneralResponse{
		Code:    fiber.StatusCreated,
		Message: "Success",
		Data:    nil,
	})
}

func (ctr *CommentController) selectByID(c *fiber.Ctx) (err error) {
	var clientRequest model.CommentSelectByPost

	// id parsing
	clientRequest.PostID, err = strconv.Atoi(c.Params("post_id"))
	exception.Sniff(err)

	response := ctr.CommentService.SelectByPost(c.Context(), &clientRequest)

	return c.Status(fiber.StatusOK).JSON(model.GeneralResponse{
		Code:    fiber.StatusOK,
		Message: "Success Created",
		Data:    response,
	})
}

func (ctr *CommentController) updateByID(c *fiber.Ctx) error {
	var clientRequest model.CommentUpdateByID

	err := c.BodyParser(&clientRequest)
	exception.Sniff(err)

	ctr.CommentService.UpdateByID(c.Context(), &clientRequest)

	return c.Status(fiber.StatusOK).JSON(model.GeneralResponse{
		Code:    fiber.StatusOK,
		Message: "Success Updated",
		Data:    nil,
	})
}

func (ctr *CommentController) deleteByID(c *fiber.Ctx) error {
	var clientRequest model.CommentDeleteByID

	err := c.BodyParser(&clientRequest)
	exception.Sniff(err)

	ctr.CommentService.DeleteByID(c.Context(), &clientRequest)

	return c.Status(fiber.StatusOK).JSON(model.GeneralResponse{
		Code:    fiber.StatusOK,
		Message: "Success Deleted",
		Data:    nil,
	})
}

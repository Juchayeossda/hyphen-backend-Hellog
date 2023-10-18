package cerrors

import (
	"encoding/json"
	"hyphen-backend-hellog/cerrors/exception"
	"hyphen-backend-hellog/model"

	"github.com/gofiber/fiber/v2"
)

func ErrorHandler(ctx *fiber.Ctx, err error) error {
	_, validationError := err.(ValidationError)
	if validationError {
		data := err.Error()
		var messages []map[string]interface{}

		errJson := json.Unmarshal([]byte(data), &messages)
		exception.Sniff(errJson)
		return ctx.Status(fiber.StatusBadRequest).JSON(model.GeneralResponse{
			Code:    400,
			Message: "Bad Request",
			Data:    messages,
		})
	}

	_, notFoundError := err.(NotFoundError)
	if notFoundError {
		return ctx.Status(fiber.StatusNotFound).JSON(model.GeneralResponse{
			Code:    404,
			Message: "Not Found",
			Data:    err.Error(),
		})
	}

	_, notCraeteError := err.(NotCreateError)
	if notCraeteError {
		return ctx.Status(fiber.StatusNotFound).JSON(model.GeneralResponse{
			Code:    500,
			Message: "Not create",
			Data:    err.Error(),
		})
	}

	_, requestFailedError := err.(ReqeustFailedError)
	if requestFailedError {
		return ctx.Status(fiber.StatusNotFound).JSON(model.GeneralResponse{
			Code:    500,
			Message: "Not Reqeust Another Server",
			Data:    err.Error(),
		})
	}

	return ctx.Status(fiber.StatusInternalServerError).JSON(model.GeneralResponse{
		Code:    500,
		Message: "General Error",
		Data:    err.Error(),
	})
}

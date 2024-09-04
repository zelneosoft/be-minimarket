package products

import (
	"backend/constant"
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func IndexHandler(ctx *fiber.Ctx) error {
	var service = Service{
		Context: ctx,
	}

	var response = service.Find()

	return ctx.JSON(&fiber.Map{
		"data":   response,
		"status": constant.STATUS_SUCCESS,
	})
}

func DetailHandler(ctx *fiber.Ctx) error {
	var service = Service{
		Context: ctx,
	}

	IDStr := ctx.Params("id")
	ID, err := strconv.Atoi(IDStr)
	if err != nil {
		return fiber.ErrBadRequest
	}

	var response = service.FindByID(ID)

	return ctx.JSON(&fiber.Map{
		"data":   response,
		"status": constant.STATUS_SUCCESS,
	})
}

func InsertHandler(ctx *fiber.Ctx) error {
	var service = Service{
		Context: ctx,
	}

	request := ProductRequest{}
	if err := ctx.BodyParser(&request); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"error":   "Bad Request",
			"message": fmt.Sprintf("Invalid parsing data. Error: %v", err),
		})
	}

	response, err := service.Insert(request)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"error":   "Internal Server Error",
			"message": fmt.Sprintf("Failed to insert data. Error: %v", err),
		})
	}

	return ctx.JSON(&fiber.Map{
		"data": response,
		"code": constant.STATUS_SUCCESS,
	})
}

func UpdateHandler(ctx *fiber.Ctx) error {
	var service = Service{
		Context: ctx,
	}

	request := ProductRequest{}
	if err := ctx.BodyParser(&request); err != nil {
		fmt.Println(err)
		return ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"error":   "Bad Request",
			"message": "Invalid JSON format in request body",
		})
	}

	ID := ctx.Params("id")

	updateData, err := service.Update(ID, request)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"error":   "Internal Server Error",
			"message": fmt.Sprintf("Failed. Error: %v", err),
		})
	}

	return ctx.JSON(&fiber.Map{
		"data": updateData,
		"code": constant.STATUS_SUCCESS,
	})
}

package brand

import (
	"backend/constant"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func ListHandler(ctx *fiber.Ctx) error {
	var service = Service{
		Context: ctx,
	}

	search := ctx.Query("search")

	var response = service.Find(search)

	return ctx.JSON(&fiber.Map{
		"data":  response,
		"count": len(response),
		"code":  constant.STATUS_SUCCESS,
	})
}

func InsertHandler(ctx *fiber.Ctx) error {
	var service = Service{
		Context: ctx,
	}

	request := BrandRequest{}
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

	request := BrandRequest{}
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

func DeleteHandler(ctx *fiber.Ctx) error {
	var service = Service{
		Context: ctx,
	}

	ID := ctx.Params("id")

	err := service.Delete(ID)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"error":   "Internal Server Error",
			"message": fmt.Sprintf("Failed to delete. Error: %v", err),
		})
	}

	return ctx.JSON(&fiber.Map{
		"code":    constant.STATUS_SUCCESS,
		"message": fmt.Sprintf("Delete data where ID %s successfully", ID),
	})
}

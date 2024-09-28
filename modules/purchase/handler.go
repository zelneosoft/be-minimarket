package purchase

import (
	"backend/constant"
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func ListHandler(ctx *fiber.Ctx) error {
	var service = Service{
		Context: ctx,
	}

	search := ctx.Query("search")
	isActiveQuery := ctx.Query("is_active")

	// Convert `is_active` query parameter to a pointer to bool
	var isActive *bool
	if isActiveQuery != "" {
		parsedActive, err := strconv.ParseBool(isActiveQuery)
		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"error": "Invalid value for 'is_active'. It should be 'true' or 'false'.",
			})
		}
		isActive = &parsedActive
	}

	var response = service.Find(search, isActive)

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

	request := CreatePORequest{}
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

	request := WarehouseRequest{}
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

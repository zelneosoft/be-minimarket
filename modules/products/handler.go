package products

import (
	"backend/constant"

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

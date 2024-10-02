package shippingmethod

import (
	"backend/constant"

	"github.com/gofiber/fiber/v2"
)

func IndexHandler(ctx *fiber.Ctx) error {
	var service = Service{
		Context: ctx,
	}

	search := ctx.Query("search")

	var response = service.Find(search)

	return ctx.JSON(&fiber.Map{
		"data":   response,
		"status": constant.STATUS_SUCCESS,
	})
}

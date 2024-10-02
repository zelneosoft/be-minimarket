package paymentmethod

import (
	"backend/constant"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func IndexHandler(ctx *fiber.Ctx) error {
	var service = Service{
		Context: ctx,
	}

	search := ctx.Query("search")

	var isActivePurchase int
	if isActivePurchaseStr := ctx.Query("for-purchase"); isActivePurchaseStr != "" {
		var err error
		isActivePurchase, err = strconv.Atoi(isActivePurchaseStr)
		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"error": "invalid value for 'for-purchase' query parameter",
			})
		}
	}

	var isActiveSales int
	if isActiveSalesStr := ctx.Query("for-sales"); isActiveSalesStr != "" {
		var err error
		isActiveSales, err = strconv.Atoi(isActiveSalesStr)
		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"error": "invalid value for 'for-sales' query parameter",
			})
		}
	}

	var response = service.Find(search, isActivePurchase, isActiveSales)

	return ctx.JSON(&fiber.Map{
		"data":   response,
		"status": constant.STATUS_SUCCESS,
	})
}

package shippingmethod

import (
	"github.com/gofiber/fiber/v2"
)

func Register(app fiber.Router) {
	route := app.Group("/shipping-method")
	{
		route.Get("/", IndexHandler)
	}
}

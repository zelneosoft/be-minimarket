package paymentmethod

import (
	"github.com/gofiber/fiber/v2"
)

func Register(app fiber.Router) {
	route := app.Group("/payment-method")
	{
		route.Get("/", IndexHandler)
	}
}

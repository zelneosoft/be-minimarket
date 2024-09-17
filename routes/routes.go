package routes

import (
	"backend/middleware"
	"backend/modules/authentication"
	"backend/modules/branchs"
	"backend/modules/products"
	"backend/modules/supplier"
	"backend/modules/warehouse"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")
	v1 := api.Group("/v1")

	authentication.Register(v1)

	auth := v1.Group("/", middleware.AuthRequired)
	products.Register(auth)
	branchs.Register(auth)
	supplier.Register(auth)
	warehouse.Register(auth)
}

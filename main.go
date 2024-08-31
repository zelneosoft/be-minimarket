package main

import (
	"backend/config"
	"backend/routes"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	app := fiber.New()

	app.Use(cors.New())

	config.InitDatabase()

	app.Use(func(c *fiber.Ctx) error {
		c.Locals("db", config.DB)
		return c.Next()
	})

	routes.SetupRoutes(app)

	port := os.Getenv("PORT")
	if port == "" {
		port = "4000"
	}

	log.Fatal(app.Listen(fmt.Sprintf(":%s", port)))
}

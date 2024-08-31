package middleware

import (
	"backend/constant"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

func AuthRequired(ctx *fiber.Ctx) error {
	authHeader := ctx.Get("Authorization")
	if authHeader == "" {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":  constant.STATUS_ERROR,
			"message": "Missing or malformed JWT",
		})
	}

	tokenString := strings.TrimSpace(strings.Replace(authHeader, "Bearer", "", 1))
	if tokenString == "" {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":  constant.STATUS_ERROR,
			"message": "Missing or malformed JWT",
		})
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fiber.ErrUnauthorized
		}
		return []byte(os.Getenv("SECRET_KEY")), nil
	})

	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":  constant.STATUS_ERROR,
			"message": "Invalid or expired JWT",
		})
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		ctx.Locals("user_id", claims["user_id"])
		ctx.Locals("email", claims["email"])
		return ctx.Next()
	}

	return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
		"status":  constant.STATUS_ERROR,
		"message": "Invalid or expired JWT",
	})
}

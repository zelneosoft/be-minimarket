package authentication

import (
	"backend/constant"
	"backend/models"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func LoginHandler(ctx *fiber.Ctx) error {
	service := Service{
		Context: ctx,
	}

	var req LoginRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"code":    constant.STATUS_ERROR,
			"message": "Invalid request",
		})
	}

	user, err := service.FindUserByEmail(req.Email)
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"code":    constant.STATUS_ERROR,
			"message": "Invalid credentials",
		})
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"code":    constant.STATUS_ERROR,
			"message": "Invalid password",
		})
	}

	token, err := generateJWT(user)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"code":    constant.STATUS_ERROR,
			"message": "Could not generate token",
		})
	}

	return ctx.JSON(&fiber.Map{
		"code": constant.STATUS_SUCCESS,
		"data": LoginResponse{Token: token},
	})
}

func NewPasswordHandler(ctx *fiber.Ctx) error {
	var req NewPasswordRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"code":    constant.STATUS_ERROR,
			"message": "Invalid request",
		})
	}

	service := Service{
		Context: ctx,
	}

	if err := service.UpdatePassword(req.Email, req.NewPassword); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"code":    constant.STATUS_ERROR,
			"message": "Could not update password",
		})
	}

	return ctx.JSON(&fiber.Map{
		"code":    constant.STATUS_SUCCESS,
		"message": "Password updated successfully",
	})
}

func generateJWT(user *models.User) (string, error) {
	claims := jwt.MapClaims{
		"user_id": user.ID,
		"email":   user.Email,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("SECRET_KEY")))
}

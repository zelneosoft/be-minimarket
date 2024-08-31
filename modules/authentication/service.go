package authentication

import (
	"backend/models"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Service struct {
	Context *fiber.Ctx
}

func (s *Service) repo() *Repository {
	db := s.Context.Locals("db").(*gorm.DB)
	return Repo(db)
}

func (s *Service) FindUserByEmail(email string) (*models.User, error) {
	return s.repo().FindUserByEmail(email)
}

func (s *Service) UpdatePassword(email, newPassword string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	return s.repo().UpdatePassword(email, string(hashedPassword))
}

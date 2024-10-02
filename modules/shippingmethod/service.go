package shippingmethod

import (
	"backend/models"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Service struct {
	Context *fiber.Ctx
}

func (s *Service) repo() *Repository {
	db := s.Context.Locals("db").(*gorm.DB)
	return Repo(db)
}

func (s *Service) Find(search string) []models.ShippingMethod {
	return s.repo().Find(search)
}

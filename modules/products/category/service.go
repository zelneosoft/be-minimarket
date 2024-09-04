package category

import (
	"backend/models"
	"fmt"

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

func (s *Service) Find(search string, isActive *bool) []models.Category {
	return s.repo().Find(search, isActive)
}

func (s *Service) Insert(req CategoryRequest) ([]*models.Category, error) {
	data := &models.Category{
		Name:     req.Name,
		IsActive: req.IsActive,
		Color:    req.Color,
	}

	err := s.repo().Create(data)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return []*models.Category{data}, nil
}

func (s *Service) Update(id string, req CategoryRequest) (*models.Category, error) {
	updateData, err := s.repo().FindByID(id)
	if err != nil {
		return nil, err
	}

	updateData.Name = req.Name
	updateData.IsActive = req.IsActive
	updateData.Color = req.Color

	err = s.repo().Update(updateData)
	if err != nil {
		return nil, err
	}

	return updateData, nil
}

func (s *Service) Delete(id string) error {
	_, err := s.repo().FindByID(id)
	if err != nil {
		return err
	}

	err = s.repo().Delete(id)
	if err != nil {
		return err
	}

	return nil
}

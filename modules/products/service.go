package products

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

func (s *Service) Find() []models.Product {
	return s.repo().Find()
}

func (s *Service) Insert(req ProductRequest) ([]*models.Product, error) {
	data := &models.Product{
		Barcode:     req.Barcode,
		Name:        req.Name,
		Description: req.Description,
		CategoryID:  req.CategoryID,
		BrandID:     req.BrandID,
	}

	err := s.repo().Create(data)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return []*models.Product{data}, nil
}

func (s *Service) Update(id string, req ProductRequest) (*models.Product, error) {
	updateData, err := s.repo().FindByID(id)
	if err != nil {
		return nil, err
	}

	updateData.Barcode = req.Barcode
	updateData.Name = req.Name
	updateData.Description = req.Description
	updateData.CategoryID = req.CategoryID
	updateData.BrandID = req.BrandID

	err = s.repo().Update(updateData)
	if err != nil {
		return nil, err
	}

	return updateData, nil
}

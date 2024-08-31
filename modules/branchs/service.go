package branchs

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

func (s *Service) Find() []models.Branch {
	return s.repo().Find()
}

func (s *Service) Insert(req BranchRequest) ([]*models.Branch, error) {
	data := &models.Branch{
		Name:     req.Name,
		Address:  req.Address,
		Maps:     req.Maps,
		IsActive: req.IsActive,
	}

	err := s.repo().Create(data)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return []*models.Branch{data}, nil
}

func (s *Service) Update(id string, req BranchRequest) (*models.Branch, error) {
	updateData, err := s.repo().FindByID(id)
	if err != nil {
		return nil, err
	}

	updateData.Name = req.Name
	updateData.Address = req.Address
	updateData.Maps = req.Maps
	updateData.IsActive = req.IsActive

	err = s.repo().Update(updateData)
	if err != nil {
		return nil, err
	}

	return updateData, nil
}

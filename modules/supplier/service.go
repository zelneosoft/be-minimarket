package supplier

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

func (s *Service) Find(search string) []models.Branch {
	return s.repo().Find(search)
}

func (s *Service) FindByID(ID int) []models.Branch {
	branch, err := s.repo().FindByID(ID)
	if err != nil {
		return nil
	}
	return []models.Branch{*branch}
}

func (s *Service) Insert(req BranchRequest) ([]*models.Branch, error) {
	data := &models.Branch{
		Name:     req.Name,
		Address:  req.Address,
		Maps:     req.Maps,
		Email:    req.Email,
		Phone:    req.Phone,
		Pic:      req.Pic,
		IsActive: req.IsActive,
	}

	err := s.repo().Create(data)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return []*models.Branch{data}, nil
}

func (s *Service) Update(id int, req BranchRequest) (*models.Branch, error) {
	updateData, err := s.repo().FindByID(id)
	if err != nil {
		return nil, err
	}

	updateData.Name = req.Name
	updateData.Address = req.Address
	updateData.Maps = req.Maps
	updateData.Email = req.Email
	updateData.Phone = req.Phone
	updateData.Pic = req.Pic
	updateData.IsActive = req.IsActive

	err = s.repo().Update(updateData)
	if err != nil {
		return nil, err
	}

	return updateData, nil
}

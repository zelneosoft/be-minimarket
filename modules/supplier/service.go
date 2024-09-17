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

func (s *Service) Find(search string) []models.Supplier {
	return s.repo().Find(search)
}

func (s *Service) FindByID(ID int) []models.Supplier {
	supplier, err := s.repo().FindByID(ID)
	if err != nil {
		return nil
	}
	return []models.Supplier{*supplier}
}

func (s *Service) Insert(req SupplierRequest) ([]*models.Supplier, error) {
	data := &models.Supplier{
		Name:     req.Name,
		Address:  req.Address,
		Maps:     req.Maps,
		Phone:    req.Phone,
		IsActive: req.IsActive,
	}

	err := s.repo().Create(data)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return []*models.Supplier{data}, nil
}

func (s *Service) Update(id int, req SupplierRequest) (*models.Supplier, error) {
	updateData, err := s.repo().FindByID(id)
	if err != nil {
		return nil, err
	}

	updateData.Name = req.Name
	updateData.Address = req.Address
	updateData.Maps = req.Maps
	updateData.Phone = req.Phone
	updateData.IsActive = req.IsActive

	err = s.repo().Update(updateData)
	if err != nil {
		return nil, err
	}

	return updateData, nil
}

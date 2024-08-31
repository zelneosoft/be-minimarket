package products

import (
	"backend/models"

	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

func Repo(db *gorm.DB) *Repository {
	return &Repository{DB: db}
}

func (repo *Repository) Find() []models.Product {
	var products []models.Product
	repo.DB.Find(&products)
	return products
}

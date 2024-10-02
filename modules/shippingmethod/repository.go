package shippingmethod

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

func (repo *Repository) Find(search string) []models.ShippingMethod {
	var data []models.ShippingMethod
	query := repo.DB.Order("created_at desc")

	if search != "" {
		query = query.Where("name LIKE ?", "%"+search+"%")
	}

	query.Find(&data)
	return data
}

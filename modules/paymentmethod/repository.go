package paymentmethod

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

func (repo *Repository) Find(search string, forPurchase, forSales int) []models.PaymentMethod {
	var data []models.PaymentMethod
	query := repo.DB.Order("created_at desc")

	if search != "" {
		query = query.Where("name LIKE ?", "%"+search+"%")
	}

	if forPurchase == 1 {
		query = query.Where("is_active_for_purchase = ?", forPurchase)
	}

	if forSales == 1 {
		query = query.Where("is_active_for_purchase = ?", forPurchase)
	}

	query.Find(&data)
	return data
}

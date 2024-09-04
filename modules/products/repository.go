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
	repo.DB.
		Preload("Category").
		Preload("Brand").
		Order("created_at desc").
		Find(&products)
	return products
}

func (repo *Repository) FindDetailByID(ID int) []models.Product {
	var products []models.Product
	repo.DB.
		Where("id = ?", ID).
		Preload("Category").
		Preload("Brand").
		Order("created_at desc").
		Find(&products)
	return products
}

func (repo *Repository) Create(data *models.Product) error {
	return repo.DB.Create(data).Error
}

func (repo *Repository) FindByID(id string) (*models.Product, error) {
	var data models.Product
	if err := repo.DB.First(&data, id).Error; err != nil {
		return nil, err
	}
	return &data, nil
}

func (repo *Repository) Update(data *models.Product) error {
	return repo.DB.Save(data).Error
}

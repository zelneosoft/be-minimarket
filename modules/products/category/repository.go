package category

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

func (repo *Repository) Find(search string, isActive *bool) []models.Category {
	var data []models.Category
	query := repo.DB.Order("created_at desc")

	if isActive != nil {
		query = query.Where("is_active = ?", *isActive)
	}

	if search != "" {
		query = query.Where("name LIKE ?", "%"+search+"%")
	}

	query.Find(&data)
	return data
}

func (repo *Repository) FindByID(id string) (*models.Category, error) {
	var data models.Category
	if err := repo.DB.First(&data, id).Error; err != nil {
		return nil, err
	}
	return &data, nil
}

func (repo *Repository) Create(data *models.Category) error {
	return repo.DB.Create(data).Error
}

func (repo *Repository) Update(data *models.Category) error {
	return repo.DB.Save(data).Error
}

func (repo *Repository) Delete(id string) error {
	if err := repo.DB.Delete(&models.Category{}, id).Error; err != nil {
		return err
	}
	return nil
}

package purchase

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

func (repo *Repository) Find(search string, isActive *bool) []models.Warehouse {
	var data []models.Warehouse
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

func (repo *Repository) FindByID(id string) (*models.Warehouse, error) {
	var data models.Warehouse
	if err := repo.DB.First(&data, id).Error; err != nil {
		return nil, err
	}
	return &data, nil
}

func (repo *Repository) Create(data *models.Warehouse) error {
	return repo.DB.Create(data).Error
}

func (repo *Repository) Update(data *models.Warehouse) error {
	return repo.DB.Save(data).Error
}

func (repo *Repository) Delete(id string) error {
	if err := repo.DB.Delete(&models.Warehouse{}, id).Error; err != nil {
		return err
	}
	return nil
}

package brand

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

func (repo *Repository) Find(search string) []models.Brand {
	var data []models.Brand
	if search != "" {
		repo.DB.Where("name LIKE ?", "%"+search+"%").Find(&data)
	} else {
		repo.DB.Find(&data)
	}
	return data
}

func (repo *Repository) FindByID(id string) (*models.Brand, error) {
	var data models.Brand
	if err := repo.DB.First(&data, id).Error; err != nil {
		return nil, err
	}
	return &data, nil
}

func (repo *Repository) Create(data *models.Brand) error {
	return repo.DB.Create(data).Error
}

func (repo *Repository) Update(data *models.Brand) error {
	return repo.DB.Save(data).Error
}

func (repo *Repository) Delete(id string) error {
	if err := repo.DB.Delete(&models.Brand{}, id).Error; err != nil {
		return err
	}
	return nil
}

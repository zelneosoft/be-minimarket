package supplier

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

func (repo *Repository) Find(search string) []models.Supplier {
	var data []models.Supplier
	query := repo.DB.Order("created_at desc")

	if search != "" {
		query = query.Where("name LIKE ?", "%"+search+"%")
	}

	query.Find(&data)
	return data
}

func (repo *Repository) FindByID(id int) (*models.Supplier, error) {
	var data models.Supplier
	if err := repo.DB.First(&data, id).Error; err != nil {
		return nil, err
	}
	return &data, nil
}

func (repo *Repository) Create(data *models.Supplier) error {
	return repo.DB.Create(data).Error
}

func (repo *Repository) Update(data *models.Supplier) error {
	return repo.DB.Save(data).Error
}

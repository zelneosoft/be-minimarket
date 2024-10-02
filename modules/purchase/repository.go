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

func (repo *Repository) Find(search, status string) []models.PurchaseHeader {
	var data []models.PurchaseHeader
	query := repo.DB.Order("created_at desc").Preload("Supplier")

	if status != "" {
		query = query.Where("status = ?", status)
	}

	if search != "" {
		query = query.Where("id LIKE ?", "%"+search+"%")
	}

	query.Find(&data)
	return data
}

func (repo *Repository) FindProduct(search string) []models.Product {
	var data []models.Product
	query := repo.DB.Order("created_at desc")

	if search != "" {
		query = query.Where("name LIKE ?", "%"+search+"%")
	}

	query.Find(&data)

	for i := range data {
		var purchaseLine models.PurchaseLine
		err := repo.DB.Order("created_at desc").
			Where("item_id = ?", data[i].ID).
			First(&purchaseLine).Error

		if err == nil {
			data[i].PurchasePrice = &purchaseLine.ItemPrice
		} else {
			var _purchasePrice float64 = 0.0
			data[i].PurchasePrice = &_purchasePrice
		}
	}

	return data
}

func (repo *Repository) FindDetailPOByID(ID string) []models.PurchaseHeader {
	var data []models.PurchaseHeader
	repo.DB.
		Where("id = ?", ID).
		Preload("PurchaseLines.Item").
		Find(&data)
	return data
}

func (r *Repository) CreatePOHeader(tx *gorm.DB, header *models.PurchaseHeader) error {
	return tx.Create(header).Error
}

func (r *Repository) CreatePOLine(tx *gorm.DB, line *models.PurchaseLine) error {
	return tx.Create(line).Error
}

func (r *Repository) GetPOHeaderByID(tx *gorm.DB, id string, purchaseHeader *models.PurchaseHeader) error {
	return tx.Where("id = ?", id).First(purchaseHeader).Error
}

func (r *Repository) UpdatePOHeader(tx *gorm.DB, purchaseHeader *models.PurchaseHeader) error {
	return tx.Save(purchaseHeader).Error
}

func (r *Repository) DeletePOLinesByPurchaseID(tx *gorm.DB, purchaseID string) error {
	return tx.Where("purchase_id = ?", purchaseID).Delete(&models.PurchaseLine{}).Error
}

func (repo *Repository) Delete(id string) error {
	if err := repo.DB.Delete(&models.Warehouse{}, id).Error; err != nil {
		return err
	}
	return nil
}

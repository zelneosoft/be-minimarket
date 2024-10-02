package purchase

import (
	"backend/models"
	"fmt"
	"strconv"
	"time"

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

func (s *Service) Find(search, status string) []models.PurchaseHeader {
	return s.repo().Find(search, status)
}

func (s *Service) FindByID(ID string) []models.PurchaseHeader {
	return s.repo().FindDetailPOByID(ID)
}

func (s *Service) FindProduct(search string) []models.Product {
	return s.repo().FindProduct(search)
}

func (s *Service) Insert(req CreatePORequest) ([]*models.PurchaseHeader, error) {
	tx := s.repo().DB.Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}

	// Jika terjadi panic, rollback transaksi
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			panic(r)
		}
	}()

	purchaseID, err := generatePurchaseOrderID(tx)
	if err != nil {
		fmt.Println("Error generating Purchase Order ID:", err)
		tx.Rollback()
		return nil, err
	}

	// Konversi epoch time ke time.Time
	purchaseDate := time.Unix(req.PurchaseDate/1000, 0)

	// Membuat model PurchaseHeader dari request
	purchaseHeader := &models.PurchaseHeader{
		ID:               purchaseID,
		PurchaseDate:     purchaseDate,
		Status:           req.Status,
		SupplierID:       req.SupplierID,
		BranchID:         req.BranchID,
		PaymentMethodID:  req.PaymentMethodID,
		ShippingMethodID: req.ShippingMethodID,
		ShippingAmount:   req.ShippingAmount,
		DiscountAmount:   req.DiscountAmount,
		TotalAmount:      req.TotalAmount,
	}

	// Simpan PurchaseHeader ke database dalam transaksi
	err = s.repo().CreatePOHeader(tx, purchaseHeader)
	if err != nil {
		fmt.Println("Error creating PurchaseHeader:", err)
		tx.Rollback()
		return nil, err
	}

	// Membuat dan menyimpan item PurchaseLine
	for _, item := range req.Items {
		purchaseLine := &models.PurchaseLine{
			PurchaseID:   purchaseHeader.ID,
			ItemID:       item.ItemID,
			ItemPrice:    item.ItemPrice,
			ItemDiscount: item.ItemDiscount,
			ItemQty:      item.ItemQty,
			ItemTotal:    item.ItemTotal,
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
		}

		// Simpan setiap PurchaseLine ke database dalam transaksi
		err := s.repo().CreatePOLine(tx, purchaseLine)
		if err != nil {
			fmt.Println("Error creating PurchaseLine:", err)
			tx.Rollback()
			return nil, err
		}
	}

	// Commit transaksi jika semua operasi berhasil
	if err := tx.Commit().Error; err != nil {
		fmt.Println("Error committing transaction:", err)
		tx.Rollback()
		return nil, err
	}

	return []*models.PurchaseHeader{purchaseHeader}, nil
}

func (s *Service) Update(purchaseID string, req CreatePORequest) (*models.PurchaseHeader, error) {
	tx := s.repo().DB.Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			panic(r)
		}
	}()

	// Konversi epoch time ke time.Time
	purchaseDate := time.Unix(req.PurchaseDate/1000, 0)

	// Temukan PurchaseHeader berdasarkan ID
	var purchaseHeader models.PurchaseHeader
	err := s.repo().GetPOHeaderByID(tx, purchaseID, &purchaseHeader)
	if err != nil {
		fmt.Println("Error finding PurchaseHeader:", err)
		tx.Rollback()
		return nil, err
	}

	// Update field-field PurchaseHeader dengan data dari request
	purchaseHeader.PurchaseDate = purchaseDate
	purchaseHeader.Status = req.Status
	purchaseHeader.SupplierID = req.SupplierID
	purchaseHeader.BranchID = req.BranchID
	purchaseHeader.PaymentMethodID = req.PaymentMethodID
	purchaseHeader.ShippingMethodID = req.ShippingMethodID
	purchaseHeader.ShippingAmount = req.ShippingAmount
	purchaseHeader.DiscountAmount = req.DiscountAmount
	purchaseHeader.TotalAmount = req.TotalAmount

	// Update PurchaseHeader dalam transaksi
	err = s.repo().UpdatePOHeader(tx, &purchaseHeader)
	if err != nil {
		fmt.Println("Error updating PurchaseHeader:", err)
		tx.Rollback()
		return nil, err
	}

	// Hapus semua PurchaseLine terkait sebelum di-update
	err = s.repo().DeletePOLinesByPurchaseID(tx, purchaseHeader.ID)
	if err != nil {
		fmt.Println("Error deleting PurchaseLines:", err)
		tx.Rollback()
		return nil, err
	}

	// Insert ulang PurchaseLine yang baru dari request
	for _, item := range req.Items {
		purchaseLine := &models.PurchaseLine{
			PurchaseID:   purchaseHeader.ID,
			ItemID:       item.ItemID,
			ItemPrice:    item.ItemPrice,
			ItemDiscount: item.ItemDiscount,
			ItemQty:      item.ItemQty,
			ItemTotal:    item.ItemTotal,
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
		}

		// Simpan setiap PurchaseLine ke database dalam transaksi
		err := s.repo().CreatePOLine(tx, purchaseLine)
		if err != nil {
			fmt.Println("Error creating PurchaseLine:", err)
			tx.Rollback()
			return nil, err
		}
	}

	// Commit transaksi jika semua operasi berhasil
	if err := tx.Commit().Error; err != nil {
		fmt.Println("Error committing transaction:", err)
		tx.Rollback()
		return nil, err
	}

	return &purchaseHeader, nil
}

func generatePurchaseOrderID(tx *gorm.DB) (string, error) {
	now := time.Now()
	year := now.Year() % 100
	month := int(now.Month())

	baseID := fmt.Sprintf("PO%02d%02d", year, month)

	// Query untuk mendapatkan ID terakhir yang memiliki format yang sama (PO2409xxx)
	var lastHeader models.PurchaseHeader
	err := tx.Where("id LIKE ?", baseID+"%").Order("id DESC").First(&lastHeader).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return "", err
	}

	// Menentukan urutan (incremental) dari ID yang terakhir ditemukan
	sequence := 1
	if lastHeader.ID != "" {
		// Mengambil 3 digit terakhir dari ID sebelumnya (misalnya PO2409003 -> 003)
		lastSequenceStr := lastHeader.ID[len(lastHeader.ID)-3:]
		lastSequence, err := strconv.Atoi(lastSequenceStr)
		if err == nil {
			sequence = lastSequence + 1
		}
	}

	// Membuat ID baru dengan urutan yang bertambah
	newID := fmt.Sprintf("%s%03d", baseID, sequence)
	return newID, nil
}

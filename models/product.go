package models

import "time"

type Product struct {
	ID          uint      `gorm:"column:id;primaryKey" json:"id,omitempty"`
	Barcode     string    `gorm:"column:barcode" json:"barcode,omitempty"`
	Name        string    `gorm:"column:name" json:"name,omitempty"`
	Description string    `gorm:"column:description" json:"description,omitempty"`
	CategoryID  uint      `gorm:"column:category_id" json:"category_id,omitempty"`
	BrandID     uint      `gorm:"column:brand_id" json:"brand_id,omitempty"`
	CreatedAt   time.Time `gorm:"column:created_at" json:"created_at,omitempty"`
	UpdatedAt   time.Time `gorm:"column:updated_at" json:"updated_at,omitempty"`

	Category *Category `gorm:"foreignKey:CategoryID;references:ID" json:"category,omitempty"`
	Brand    *Brand    `gorm:"foreignKey:BrandID;references:ID" json:"brand,omitempty"`
}

package models

import "time"

type PurchaseLine struct {
	ID           uint      `gorm:"column:id;primaryKey" json:"id,omitempty"`
	PurchaseID   string    `gorm:"column:purchase_id" json:"purchase_id,omitempty"`
	ItemID       uint      `gorm:"column:item_id" json:"item_id,omitempty"`
	ItemPrice    float64   `gorm:"column:item_price" json:"item_price,omitempty"`
	ItemDiscount float64   `gorm:"column:item_discount" json:"item_discount,omitempty"`
	ItemQty      float64   `gorm:"column:item_qty" json:"item_qty,omitempty"`
	ItemTotal    float64   `gorm:"column:item_total" json:"item_total,omitempty"`
	CreatedAt    time.Time `gorm:"column:created_at" json:"created_at,omitempty"`
	UpdatedAt    time.Time `gorm:"column:updated_at" json:"updated_at,omitempty"`

	Item *Product `gorm:"foreignKey:ItemID;references:ID" json:"item,omitempty"`
}

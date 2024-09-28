package models

import "time"

type PaymentMethod struct {
	ID                  uint      `gorm:"column:id;primaryKey" json:"id,omitempty"`
	Name                string    `gorm:"column:name" json:"name,omitempty"`
	IsActiveForSales    int       `gorm:"column:is_active_for_sales" json:"is_active_for_sales,omitempty"`
	IsActiveForPurchase int       `gorm:"column:is_active_for_purchase" json:"is_active_for_purchase,omitempty"`
	CreatedAt           time.Time `gorm:"column:created_at" json:"created_at,omitempty"`
	UpdatedAt           time.Time `gorm:"column:updated_at" json:"updated_at,omitempty"`
}

package models

import "time"

type ShippingMethod struct {
	ID        uint      `gorm:"column:id;primaryKey" json:"id,omitempty"`
	Name      string    `gorm:"column:name" json:"name,omitempty"`
	IsActive  int       `gorm:"column:is_active" json:"is_active,omitempty"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at,omitempty"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at,omitempty"`
}

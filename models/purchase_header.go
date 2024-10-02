package models

import "time"

type PurchaseHeader struct {
	ID               string    `gorm:"column:id;primaryKey" json:"id,omitempty"`
	PurchaseDate     time.Time `gorm:"column:purchase_date" json:"purchase_date,omitempty"`
	Status           string    `gorm:"column:status" json:"status,omitempty"`
	SupplierID       uint      `gorm:"column:supplier_id" json:"supplier_id,omitempty"`
	BranchID         uint      `gorm:"column:branch_id" json:"branch_id,omitempty"`
	PaymentMethodID  uint      `gorm:"column:payment_method_id" json:"payment_method_id,omitempty"`
	ShippingMethodID uint      `gorm:"column:shipping_method_id" json:"shipping_method_id,omitempty"`
	ShippingAmount   float64   `gorm:"column:shipping_amount" json:"shipping_amount,omitempty"`
	DiscountAmount   float64   `gorm:"column:discount_amount" json:"discount_amount,omitempty"`
	TotalAmount      float64   `gorm:"column:total_amount" json:"total_amount,omitempty"`
	CreatedAt        time.Time `gorm:"column:created_at" json:"created_at,omitempty"`
	UpdatedAt        time.Time `gorm:"column:updated_at" json:"updated_at,omitempty"`

	Supplier       *Supplier       `gorm:"foreignKey:SupplierID;references:ID" json:"supplier,omitempty"`
	Branch         *Branch         `gorm:"foreignKey:BranchID;references:ID" json:"branch,omitempty"`
	PaymentMethod  *PaymentMethod  `gorm:"foreignKey:PaymentMethodID;references:ID" json:"payment_method,omitempty"`
	ShippingMethod *ShippingMethod `gorm:"foreignKey:ShippingMethodID;references:ID" json:"shipping_method,omitempty"`
	PurchaseLines  []PurchaseLine  `gorm:"foreignKey:PurchaseID" json:"purchase_lines,omitempty"`
}

package purchase

type CreatePORequest struct {
	ID               string          `json:"id"`
	PurchaseDate     int64           `json:"date"`
	Status           string          `json:"status"`
	SupplierID       uint            `json:"supplier_id"`
	BranchID         uint            `json:"branch_id"`
	PaymentMethodID  uint            `json:"payment_method_id"`
	ShippingMethodID uint            `json:"shipping_method_id"`
	ShippingAmount   float64         `json:"shipping_amount"`
	DiscountAmount   float64         `json:"discount_amount"`
	TotalAmount      float64         `json:"total_amount"`
	Items            []POLineRequest `json:"items"`
}
type POLineRequest struct {
	ItemID       uint    `json:"item_id"`
	ItemPrice    float64 `json:"item_price"`
	ItemDiscount float64 `json:"item_discount"`
	ItemQty      float64 `json:"item_qty"`
	ItemTotal    float64 `json:"item_total"`
}

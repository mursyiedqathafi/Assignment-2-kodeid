package models

type Item struct {
	ItemID      int    `gorm:"primaryKey" json:"lineItemId"`
	ItemCode    string `gorm:"not null;type:varchar" json:"itemCode"`
	Description string `gorm:"type:text" json:"description"`
	Quantity    int    `gorm:"not null" json:"quantity"`
	OrderID     int    `json:"-"`
}

package model

import "time"

type Product struct {
	ID        uint      `json:"id"      gorm:"primary_key"`
	Category  string    `json:"category"`
	Name      string    `json:"name"`
	Price     int       `json:"price"`
	UnitPrice int       `json:"unit_price"`
	Volume    string    `json:"volume"`
	ShopName  string    `json:"shop_name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

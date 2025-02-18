package models

import "time"

type Package struct {
	ID            uint      `json:"id"`
	Name          string    `json:"name"`
	NameSlug      string    `json:"name_slug"`
	Price         int       `json:"price"`
	DiscountPrice int       `json:"discount_price"`
	Features      string    `json:"features"`
	Image         string    `json:"image"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	ServiceID     uint      `json:"service_id"`
}

func (Package) TableName() string {
	return "packages"
}

package models

import (
	"time"
)

type Service struct {
	ID               uint      `json:"id"`
	Name             string    `json:"name"`
	NameSlug         string    `json:"name_slug"`
	ShortDescription string    `json:"short_description"`
	LongDescription  string    `json:"long_description"`
	Images           string    `json:"images"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
	Packages         []Package `json:"packages"`
}

func (Service) TableName() string {
	return "services"
}

package responses

import "time"

type PortfolioResponse struct {
	ID          uint              `json:"id"`
	Name        string            `json:"name"`
	NameSlug    string            `json:"name_slug"`
	Description string            `json:"description"`
	Url         string            `json:"url"`
	Image       string            `json:"image"`
	CreatedAt   time.Time         `json:"created_at"`
	UpdatedAt   time.Time         `json:"updated_at"`
	Services    []ServiceResponse `json:"services"`
}

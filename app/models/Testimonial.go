package models

import "time"

type Testimonial struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Job       string    `json:"job"`
	Message   string    `json:"message"`
	Image     string    `json:"image"`
	CreatedAt time.Time `json:"created_at"`
	UpdateAt  time.Time `json:"update_at"`
}

func (Testimonial) TableName() string {
	return "testimonials"
}

package models

type Pagination struct {
	CurrentPage int         `json:"current_page"`
	PerPage     int         `json:"per_page"`
	Total       int64       `json:"total"`
	TotalPages  int         `json:"total_pages"`
	Data        interface{} `json:"data"`
}

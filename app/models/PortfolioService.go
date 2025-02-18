package models

type PortfolioService struct {
	ID          uint `json:"id"`
	ServiceID   uint `json:"service_id"`
	PortfolioID uint `json:"portfolio_id"`
}

func (PortfolioService) TableName() string {
	return "portfolio_service"
}

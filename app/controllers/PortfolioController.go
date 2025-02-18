package controllers

import (
	"anis/app/enums"
	"anis/app/models"
	"anis/app/utils"
	"anis/database"
	"anis/responses"

	"github.com/gin-gonic/gin"
)

func GetAllPortfolio(ctx *gin.Context) {
	portfolios := new([]models.Portfolio)

	serviceID := ctx.DefaultQuery("service_id", "")

	query := database.DB.Model(&models.Portfolio{})

	if serviceID != "" {
		query = query.Joins("JOIN portfolio_service ps ON ps.portfolio_id = portfolios.id").
			Where("ps.service_id = ?", serviceID)
	}

	pagination, err := utils.Paginate(ctx, query.Preload("Services"), portfolios, enums.DEFAULT_PORTFOLIO_LIMIT_INTEGER, enums.DEFAULT_PORTFOLIO_PAGE_INTEGER)
	if err != nil {
		ctx.JSON(500, gin.H{"message": "Internal server error"})
		return
	}

	if len(*portfolios) == enums.EMPTY {
		ctx.JSON(404, gin.H{"message": "Portfolio not found"})
		return
	}

	var portfolioResponse []responses.PortfolioResponse

	for _, portfolio := range *portfolios {
		portfolioResponse = append(portfolioResponse, responses.PortfolioResponse{
			ID:          portfolio.ID,
			Name:        portfolio.Name,
			NameSlug:    portfolio.NameSlug,
			Description: portfolio.Description,
			Url:         portfolio.Url,
			Image:       portfolio.Image,
			CreatedAt:   portfolio.CreatedAt,
			UpdatedAt:   portfolio.UpdatedAt,
		})
	}

	pagination.Data = portfolioResponse

	ctx.JSON(200, pagination)
}

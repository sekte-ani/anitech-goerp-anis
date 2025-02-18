package utils

import (
	"anis/app/enums"
	"anis/app/models"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Paginate(ctx *gin.Context, db *gorm.DB, result interface{}, limitDefault int, pageDefault int) (*models.Pagination, error) {
	pageStr := ctx.DefaultQuery("page", strconv.Itoa(pageDefault))
	limitStr := ctx.DefaultQuery("limit", strconv.Itoa(limitDefault))

	page, err := strconv.Atoi(pageStr)
	if err != nil || page < enums.ONE {
		page = pageDefault
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit < enums.ONE {
		limit = limitDefault
	}

	offset := (page - enums.ONE) * limit

	var total int64
	err = db.Count(&total).Error
	if err != nil {
		return nil, err
	}

	err = db.Limit(limit).Offset(offset).Find(result).Error
	if err != nil {
		return nil, err
	}

	totalPages := int((total + int64(limit) - enums.ONE) / int64(limit))

	pagination := &models.Pagination{
		CurrentPage: page,
		PerPage:     limit,
		Total:       total,
		TotalPages:  totalPages,
		Data:        result,
	}

	return pagination, nil
}

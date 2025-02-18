package controllers

import (
	"anis/app/enums"
	"anis/app/models"
	"anis/app/utils"
	"anis/database"
	"anis/responses"

	"github.com/gin-gonic/gin"
)

func GetAllTestimonial(ctx *gin.Context) {
	testimonials := new([]models.Testimonial)

	query := database.DB.Model(&models.Testimonial{})

	pagination, err := utils.Paginate(ctx, query, testimonials, enums.DEFAULT_TESTIMONIAL_LIMIT_INTEGER, enums.DEFAULT_TESTIMONIAL_PAGE_INTEGER)
	if err != nil {
		ctx.JSON(500, gin.H{"message": "Internal server error"})
		return
	}

	if len(*testimonials) == enums.EMPTY {
		ctx.JSON(404, gin.H{"message": "Testimonial not found"})
		return
	}

	var testimonialResponse []responses.TestimonialResponse
	for _, testimonial := range *testimonials {
		testimonialResponse = append(testimonialResponse, responses.TestimonialResponse{
			ID:        testimonial.ID,
			Name:      testimonial.Name,
			Job:       testimonial.Job,
			Message:   testimonial.Message,
			Image:     testimonial.Image,
			CreatedAt: testimonial.CreatedAt,
			UpdateAt:  testimonial.UpdateAt,
		})
	}

	pagination.Data = testimonialResponse

	ctx.JSON(200, pagination)
}

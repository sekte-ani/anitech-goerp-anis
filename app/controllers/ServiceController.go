package controllers

import (
	"anis/app/enums"
	"anis/app/models"
	"anis/app/utils"
	"anis/database"
	"anis/responses"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetAllService(ctx *gin.Context) {
	services := new([]models.Service)

	err := database.DB.Find(&services).Error
	if err != nil {
		ctx.JSON(500, gin.H{"message": "Internal server error"})
		return
	}

	if len(*services) == enums.EMPTY {
		ctx.JSON(404, gin.H{"message": "Service not found"})
		return
	}

	var serviceResponse []responses.ServiceResponse

	for _, service := range *services {

		images, err := utils.DecodeImages(service.Images)
		if err != nil {
			ctx.JSON(500, gin.H{"message": "Error unmarshaling Images field"})
			return
		}

		serviceResponse = append(serviceResponse, responses.ServiceResponse{
			ID:               service.ID,
			Name:             service.Name,
			NameSlug:         service.NameSlug,
			ShortDescription: service.ShortDescription,
			LongDescription:  service.LongDescription,
			Images:           images,
			CreatedAt:        service.CreatedAt,
			UpdatedAt:        service.UpdatedAt,
		})
	}

	ctx.JSON(200, gin.H{"data": serviceResponse})
}

func GetByNameSlug(ctx *gin.Context) {
	serviceNameSlug := ctx.Param("name_slug")

	if serviceNameSlug == "" {
		ctx.JSON(400, gin.H{"message": "Param name_slug is required"})
		return
	}

	service := new(models.Service)
	err := database.DB.Preload("Packages").First(service, "name_slug = ?", serviceNameSlug).Error

	if err == gorm.ErrRecordNotFound {
		ctx.JSON(404, gin.H{"message": "Service not found"})
		return
	}

	if err != nil {
		ctx.JSON(500, gin.H{"message": "Internal server error"})
		return
	}

	var packageResponses []responses.PackageResponse
	if service.Packages != nil {
		for _, packageData := range service.Packages {
			packageResponses = append(packageResponses, responses.PackageResponse{
				ID:        packageData.ID,
				Name:      packageData.Name,
				NameSlug:  packageData.NameSlug,
				Features:  packageData.Features,
				Image:     packageData.Image,
				Price:     packageData.Price,
				CreatedAt: packageData.CreatedAt,
				UpdatedAt: packageData.UpdatedAt,
			})
		}
	}

	images, err := utils.DecodeImages(service.Images)
	if err != nil {
		ctx.JSON(500, gin.H{"message": "Error unmarshaling Images"})
		return
	}

	serviceResponse := responses.ServiceResponse{
		ID:               service.ID,
		Name:             service.Name,
		NameSlug:         service.NameSlug,
		ShortDescription: service.ShortDescription,
		LongDescription:  service.LongDescription,
		Images:           images,
		CreatedAt:        service.CreatedAt,
		UpdatedAt:        service.UpdatedAt,
		Packages:         packageResponses,
	}

	ctx.JSON(200, gin.H{"data": serviceResponse})
}

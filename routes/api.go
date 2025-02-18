package routes

import (
	"anis/app/controllers"
	appconfig "anis/config/app_config"

	"github.com/gin-gonic/gin"
)

func InitRoutes(app *gin.Engine) {
	route := app
	route.Static(appconfig.STATIC_ROUTE, appconfig.STATIC_DIR)

	api := route.Group("/api")
	{
		api.GET("/services", controllers.GetAllService)
		api.GET("/services/:name_slug", controllers.GetByNameSlug)
		api.GET("/portfolios", controllers.GetAllPortfolio)
		api.GET("/testimonials", controllers.GetAllTestimonial)
	}
}

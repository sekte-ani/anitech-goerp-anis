package bootstrap

import (
	"log"

	appconfig "anis/config/app_config"
	corsconfig "anis/config/cors_config"

	"anis/config"
	"anis/database"
	"anis/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func BootstrapApp() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Can't Connect to .env file")
	}

	config.InitConfig()
	database.ConnectDB()

	app := gin.Default()
	app.Use(corsconfig.CorsConfig())

	routes.InitRoutes(app)
	app.Run(":" + appconfig.PORT)
}

package database

import (
	dbconfig "anis/config/db_config"
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	var err error

	if dbconfig.DB_DRIVER == "mysql" {
		dsnMysql := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbconfig.DB_USER, dbconfig.DB_PASSWORD, dbconfig.DB_HOST, dbconfig.DB_PORT, dbconfig.DB_NAME)
		DB, err = gorm.Open(mysql.Open(dsnMysql), &gorm.Config{})
	}

	if err != nil {
		panic("Failed to Connect Database")
	}

	log.Println("Database Connected")
}

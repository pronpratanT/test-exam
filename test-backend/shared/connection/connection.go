package connection

import (
	"fmt"
	"log"
	"test-backend/shared/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var AppDB *gorm.DB

func ConnectDB() *gorm.DB {
	if AppDB != nil {
		return AppDB
	}

	config.LoadConfig()

	db, err := gorm.Open(postgres.Open(config.AppConfig.AppDB), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect postgres:", err)
	}
	fmt.Println("Connected to postgres successfully")

	AppDB = db
	return AppDB
}

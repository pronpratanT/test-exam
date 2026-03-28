package migrate

import (
	"log"
	db "test-backend/shared/connection"
	"test-backend/shared/models"
)

func AutoMigrate() error {
	database := db.ConnectDB()

	if err := database.AutoMigrate(
		&models.Test1{},
		&models.Test2{},
		&models.Test3{},
		&models.Test4{},
		&models.Test5{},
	); err != nil {
		log.Println("auto migrate failed:", err)
		return err
	}

	return nil
}

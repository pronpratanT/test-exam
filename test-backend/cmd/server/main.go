package main

import (
	"test-backend/internal/project/app/router"
	"test-backend/internal/project/handlers"
	"test-backend/internal/project/repository"
	"test-backend/internal/project/services"
	"test-backend/shared/config"
	db "test-backend/shared/connection"
	"test-backend/shared/middleware"
	"test-backend/shared/migrate"

	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadConfig()
	appDB := db.ConnectDB()

	if err := migrate.AutoMigrate(); err != nil {
		panic("Database migration failed: " + err.Error())
	}

	// เรียกคำสั่งเช็คและยัด Mockup Data ทันทีที่เชื่อมต่อ DB สำเร็จ (ในตอนเริ่มโปรเจกต์)
	db.SeedMockDataTest3(appDB)

	testRepo := repository.NewTestExamRepository(appDB)
	testService := services.NewTestService(testRepo)
	testHandler := handlers.NewHandler(testService)

	r := gin.Default()
	r.Use(middleware.CORSMiddleware()) // ใส่แค่นี้พอ
	router.SetupRouter(r, testHandler)

	r.Run(":7070")
}

package main

import (
	"test-backend/internal/project/app/router"
	"test-backend/internal/project/handlers"
	"test-backend/internal/project/repository"
	"test-backend/internal/project/services"
	"test-backend/shared/config"
	db "test-backend/shared/connection"

	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadConfig()
	appDB := db.ConntectDB()

	testRepo := repository.NewTestExamRepository(appDB)
	testService := services.NewTestService(testRepo)
	testHandler := handlers.NewHandler(testService)

	r := gin.Default()
	router.SetupRouter(r, testHandler)

	r.Run(":7070")
}

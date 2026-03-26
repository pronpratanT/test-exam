package main

import (
	"test-backend/shared/config"
	db "test-backend/shared/connection"

	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadConfig()
	db.ConntectDB()

	r := gin.Default()

	r.Run(":7070")
}

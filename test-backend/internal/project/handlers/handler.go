package handlers

import (
	"test-backend/internal/project/services"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	Service *services.TestService
}

func NewHandler(service *services.TestService) *Handler {
	return &Handler{Service: service}
}

func SetupRoutes(r *gin.RouterGroup, h *Handler) {
	// r.GET("/test", h.TestHandler)
}

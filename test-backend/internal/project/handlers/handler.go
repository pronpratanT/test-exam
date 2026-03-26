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
	//test 1
	r.GET("/test1/get-all", h.GetAllData)
	r.POST("/test1/create", h.CreateData)
	r.GET("/test1/get/:id", h.GetDetailByID)
}

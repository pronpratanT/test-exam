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
	//test 2
	// middleware.JWTAuthMiddleware(),
	r.POST("/test2/signup", h.Signup)
	r.POST("/test2/signin", h.Signin)
	//test 3
	r.GET("/test3/get-all", h.GetAllDataTest3)
	r.GET("/test3/get/:id", h.GetDetailByIDTest3)
	r.POST("/test3/approve", h.ApproveDataTest3)
	r.POST("/test3/reject", h.RejectDataTest3)
	r.POST("/test3/create", h.CreateDataTest3)
	//test 4
	r.POST("/test4/save", h.Test4SaveData)
	//test 5
	r.POST("/test5/create-ticket", h.CreateTicket)
	r.POST("/test5/clear-ticket/:queNumber", h.ClearTicket)
}

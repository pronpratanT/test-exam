package router

import (
	"test-backend/internal/project/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine, h *handlers.Handler) *gin.Engine {
	api := r.Group("/api")
	handlers.SetupRoutes(api.Group("/test"), h)
	return r
}

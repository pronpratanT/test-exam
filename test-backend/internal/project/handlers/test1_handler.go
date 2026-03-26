package handlers

import (
	"fmt"
	"net/http"
	"test-backend/internal/project/dto"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetAllData(c *gin.Context) {
	data, err := h.Service.GetAllData()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to retrieve data",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}

func (h *Handler) CreateData(c *gin.Context) {
	var dto dto.Test1DTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request body",
		})
		return
	}
	if err := h.Service.CreateData(dto); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create data",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Data created successfully",
	})
}

func (h *Handler) GetDetailByID(c *gin.Context) {
	idParam := c.Param("id")
	var id int
	_, err := fmt.Sscanf(idParam, "%d", &id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid ID parameter",
		})
		return
	}
	data, err := h.Service.GetDetailByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to retrieve data",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}

package handlers

import (
	"fmt"
	"net/http"

	"test-backend/internal/project/dto"
	"github.com/gin-gonic/gin"
)

func (h *Handler) GetAllDataTest3(c *gin.Context) {
	data, err := h.Service.GetDataAllTest3()
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

func (h *Handler) GetDetailByIDTest3(c *gin.Context) {
	idParam := c.Param("id")
	var id int
	_, err := fmt.Sscanf(idParam, "%d", &id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid ID parameter",
		})
		return
	}
	data, err := h.Service.GetDetailByIDTest3(id)
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

func (h *Handler) ApproveDataTest3(c *gin.Context) {
	idParam := c.Param("id")
	var id int
	_, err := fmt.Sscanf(idParam, "%d", &id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid ID parameter",
		})
		return
	}
	if err := h.Service.ApproveDataTest3(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to approve data",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Data approved successfully",
	})
}

func (h *Handler) RejectDataTest3(c *gin.Context) {
	idParam := c.Param("id")
	var id int
	_, err := fmt.Sscanf(idParam, "%d", &id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid ID parameter",
		})
		return
	}
	if err := h.Service.RejectDataTest3(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to reject data",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Data rejected successfully",
	})
}	

func (h *Handler) CreateDataTest3(c *gin.Context) {
	var req dto.Test3DTO
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request body",
		})
		return
	}
	if err := h.Service.CreateDataTest3(&req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create data",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Data created successfully",
	})
}
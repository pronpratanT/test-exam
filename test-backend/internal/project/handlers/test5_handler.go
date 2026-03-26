package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateTicket(c *gin.Context) {
	ticket, err := h.Service.CreateTicket()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create ticket: ",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Ticket created successfully",
		"ticket":  ticket,
	})
}

func (h *Handler) ClearTicket(c *gin.Context) {
	queNumber := c.Param("queNumber")
	if err := h.Service.ClearTicket(queNumber); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to clear ticket: ",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Ticket cleared successfully",
	})
}

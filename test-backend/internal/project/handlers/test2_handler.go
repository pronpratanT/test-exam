package handlers

import (
	"net/http"
	"test-backend/internal/project/dto"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Signup(c *gin.Context) {
	var req dto.Test2DTOSignup
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request",
		})
		return
	}
	if err := h.Service.Signup(&req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to signup",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Signup successful",
	})
}

func (h *Handler) Signin(c *gin.Context) {
	var req dto.Test2DTOSignin
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request",
		})
		return
	}
	resp, err := h.Service.Signin(&req)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid username or password",
		})
		return
	}
	c.JSON(http.StatusOK, resp)
}

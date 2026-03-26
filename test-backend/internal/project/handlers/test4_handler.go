package handlers

import (
	"encoding/base64"
	"io"
	"net/http"
	"test-backend/internal/project/dto"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Test4SaveData(c *gin.Context) {
	var inputDTO dto.Test4DTO

	inputDTO.FName = c.PostForm("fname")
	inputDTO.LName = c.PostForm("lname")
	inputDTO.Email = c.PostForm("email")
	inputDTO.Phone = c.PostForm("phone")
	inputDTO.Birthdate = c.PostForm("birthdate")
	inputDTO.Occupation = c.PostForm("occupation")

	file, err := c.FormFile("profile")
	if err == nil {
		f, _ := file.Open()
		defer f.Close()
		buf, _ := io.ReadAll(f)
		inputDTO.Profile = base64.StdEncoding.EncodeToString(buf)
	}

	id, err := h.Service.Test4SaveData(inputDTO)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Data saved successfully",
		"id":      id,
	})
}

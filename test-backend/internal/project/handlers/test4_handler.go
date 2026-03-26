package handlers

import (
	"encoding/base64"
	"io"
	"net/http"
	"test-backend/internal/project/dto"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Test4SaveData(c *gin.Context) {
	var dto dto.Test4DTO

	dto.FName = c.PostForm("fname")
	dto.LName = c.PostForm("lname")
	dto.Email = c.PostForm("email")
	dto.Phone = c.PostForm("phone")
	dto.Birthdate = c.PostForm("birthdate")
	dto.Occupation = c.PostForm("occupation")

	file, err := c.FormFile("profile")
	if err == nil {
		f, _ := file.Open()
		defer f.Close()
		buf, _ := io.ReadAll(f)
		dto.Profile = base64.StdEncoding.EncodeToString(buf)
	}

	if err := h.Service.Test4SaveData(dto); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to save data",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Data saved successfully",
	})
}

package repository

import (
	"test-backend/shared/models"
)

func (r *TestRepository) CreateTicket(queNumber string) (*models.Test5, error) {
	isCleared := false
	if queNumber == "Z9" {
		isCleared = true
	}

	ticket := models.Test5{
		QueNumber: queNumber,
		Iscleared: isCleared,
	}
	err := r.DB.Create(&ticket).Error
	if err != nil {
		return nil, err
	}
	return &ticket, nil
}

func (r *TestRepository) GetAllTicket() ([]models.Test5, error) {
	var tickets []models.Test5
	err := r.DB.Find(&tickets).Error
	if err != nil {
		return nil, err
	}
	return tickets, nil
}

func (r *TestRepository) ClearAllTicket(queNumber string) error {
	return r.DB.Model(&models.Test5{}).Where("que_number = ?", queNumber).Update("iscleared", true).Error
}

func (r *TestRepository) GetLastTicket() (*models.Test5, error) {
	var ticket models.Test5
	err := r.DB.Order("created_at desc").First(&ticket).Error
	if err != nil {
		return nil, err
	}
	return &ticket, nil
}

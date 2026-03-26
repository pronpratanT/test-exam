package repository

import (
	"test-backend/internal/project/dto"
	"test-backend/shared/models"
)

func (r *TestRepository) Signup(dto *dto.Test2DTOSignup) error {
	data := models.Test2{
		Username: dto.Username,
		Password: dto.Password,
	}
	return r.DB.Create(&data).Error
}

func (r *TestRepository) CheckDuplicateUsername(username string) (bool, error) {
	var count int64
	err := r.DB.Model(&models.Test2{}).
		Where("username = ?", username).
		Count(&count).Error
	return count > 0, err
}

func (r *TestRepository) Signin(username string) (*models.Test2, error) {
	var data models.Test2
	err := r.DB.Where("username = ?", username).First(&data).Error
	if err != nil {
		return nil, err
	}
	return &data, nil
}

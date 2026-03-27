package repository

import (
	"test-backend/internal/project/dto"
	"test-backend/shared/models"
)

func (r *TestRepository) Test4SaveData(data dto.Test4DTO) (int, error) {
	record := models.Test4{
		FName:      data.FName,
		LName:      data.LName,
		Email:      data.Email,
		Phone:      data.Phone,
		Profile:    data.Profile,
		Birthdate:  data.Birthdate,
		Occupation: data.Occupation,
		Sex:        data.Sex,
	}
	if err := r.DB.Create(&record).Error; err != nil {
		return 0, err
	}
	return record.ID, nil
}

func (r *TestRepository) CheckDuplicateEmail(email string) (bool, error) {
	var count int64
	err := r.DB.Model(&models.Test4{}).Where("email = ?", email).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

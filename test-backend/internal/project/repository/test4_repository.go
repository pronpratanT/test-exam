package repository

import "test-backend/internal/project/dto"

func (r *TestRepository) Test4SaveData(data dto.Test4DTO) error {
	data = dto.Test4DTO{
		FName:      data.FName,
		LName:      data.LName,
		Email:      data.Email,
		Phone:      data.Phone,
		Profile:    data.Profile,
		Birthdate:  data.Birthdate,
		Occupation: data.Occupation,
	}
	return r.DB.Create(&data).Error
}

func (r *TestRepository) CheckDuplicateEmail(email string) (bool, error) {
	var count int64
	err := r.DB.Model(&dto.Test4DTO{}).Where("email = ?", email).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

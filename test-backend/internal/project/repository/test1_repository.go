package repository

import (
	"test-backend/internal/project/dto"
	"test-backend/shared/models"
)

func (r *TestRepository) GetAllData() ([]models.Test1, error) {
	var data []models.Test1
	err := r.DB.Find(&data).Error
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (r *TestRepository) CreateData(dto dto.Test1DTO) error {
	data := models.Test1{
		FName:     dto.FName,
		LName:     dto.LName,
		BirthDate: dto.BirthDate,
		Age:       dto.Age,
		Address:   dto.Address,
	}
	return r.DB.Create(&data).Error
}

func (r *TestRepository) CheckDuplicateName(fname, lname string) (bool, error) {
	var count int64
	err := r.DB.Model(&models.Test1{}).
		Where("fname = ? AND lname = ?", fname, lname).
		Count(&count).Error
	return count > 0, err
}

func (r *TestRepository) GetDetailByID(id int) (*models.Test1, error) {
	var data models.Test1
	err := r.DB.First(&data, id).Error
	if err != nil {
		return nil, err
	}
	return &data, nil
}

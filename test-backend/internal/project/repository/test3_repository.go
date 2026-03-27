package repository

import (
	"test-backend/internal/project/dto"
	"test-backend/shared/models"
)

func (r *TestRepository) GetAllDataTest3() ([]models.Test3, error) {
	var data []models.Test3
	err := r.DB.Order("id ASC").Find(&data).Error
	return data, err
}

func (r *TestRepository) GetDetailByIDTest3(id int) (*models.Test3, error) {
	var data models.Test3
	err := r.DB.Where("id = ?", id).First(&data).Error
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (r *TestRepository) ApproveDataTest3(dto []dto.Test3ApproveDTO) error {
	for _, item := range dto {
		if err := r.DB.Model(&models.Test3{}).Where("id = ?", item.ID).Update("status", "approved").Update("reason", item.Reason).Error; err != nil {
			return err
		}
	}
	return nil
}

func (r *TestRepository) RejectDataTest3(dto []dto.Test3RejectDTO) error {
	for _, item := range dto {
		if err := r.DB.Model(&models.Test3{}).Where("id = ?", item.ID).Update("status", "rejected").Update("reason", item.Reason).Error; err != nil {
			return err
		}
	}
	return nil
}

func (r *TestRepository) CreateDataTest3(data *dto.Test3DTO) error {
	newModel := models.Test3{
		Name:   data.Name,
		Reason: data.Reason,
		Status: data.Status,
	}
	return r.DB.Create(&newModel).Error
}

package repository

import "test-backend/shared/models"

func (r *TestRepository) GetAllDataTest3() ([]models.Test3, error) {
	var data []models.Test3
	err := r.DB.Find(&data).Error
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

func (r *TestRepository) ApproveDataTest3(id int) error {
	return r.DB.Model(&models.Test3{}).Where("id = ?", id).Update("status", "approved").Error
}

func (r *TestRepository) RejectDataTest3(id int) error {
	return r.DB.Model(&models.Test3{}).Where("id = ?", id).Update("status", "rejected").Error
}

package services

import (
	"test-backend/internal/project/dto"
	"test-backend/shared/models"
)

func (s *TestService) GetDataAllTest3() ([]models.Test3, error) {
	return s.AppRepo.GetAllDataTest3()
}

func (s *TestService) GetDetailByIDTest3(id int) (*models.Test3, error) {
	return s.AppRepo.GetDetailByIDTest3(id)
}

func (s *TestService) ApproveDataTest3(id int, reason string) error {
	return s.AppRepo.ApproveDataTest3(id, reason)
}

func (s *TestService) RejectDataTest3(id int, reason string) error {
	return s.AppRepo.RejectDataTest3(id, reason)
}

func (s *TestService) CreateDataTest3(data *dto.Test3DTO) error {
	return s.AppRepo.CreateDataTest3(data)
}

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

func (s *TestService) ApproveDataTest3(dto []dto.Test3ApproveDTO) error {
	return s.AppRepo.ApproveDataTest3(dto)
}

func (s *TestService) RejectDataTest3(dto []dto.Test3RejectDTO) error {
	return s.AppRepo.RejectDataTest3(dto)
}

func (s *TestService) CreateDataTest3(data *dto.Test3DTO) error {
	return s.AppRepo.CreateDataTest3(data)
}

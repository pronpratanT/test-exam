package services

import (
	"errors"
	"test-backend/internal/project/dto"
	"test-backend/shared/models"
)

func (s *TestService) GetAllData() ([]models.Test1, error) {
	return s.AppRepo.GetAllData()
}

func (s *TestService) CreateData(dto dto.Test1DTO) error {
	duplicate, err := s.AppRepo.CheckDuplicateName(dto.FName, dto.LName)
	if err != nil {
		return err
	}
	if duplicate {
		return errors.New("ชื่อ-สกุลนี้มีอยู่แล้ว")
	}
	return s.AppRepo.CreateData(dto)
}

func (s *TestService) GetDetailByID(id int) (*models.Test1, error) {
	return s.AppRepo.GetDetailByID(id)
}

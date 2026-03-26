package services

import (
	"fmt"
	"test-backend/internal/project/dto"
)

func (s *TestService) Test4SaveData(data dto.Test4DTO) error {
	duplicate, err := s.AppRepo.CheckDuplicateEmail(data.Email)
	if err != nil {
		return err
	}
	if duplicate {
		return fmt.Errorf("email already exists")
	}
	return s.AppRepo.Test4SaveData(data)
}

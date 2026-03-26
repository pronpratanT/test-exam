package services

import (
	"fmt"
	"test-backend/internal/project/dto"
)

func (s *TestService) Test4SaveData(data dto.Test4DTO) (int, error) {
	duplicate, err := s.AppRepo.CheckDuplicateEmail(data.Email)
	if err != nil {
		return 0, err
	}
	if duplicate {
		return 0, fmt.Errorf("email already exists")
	}
	return s.AppRepo.Test4SaveData(data)
}

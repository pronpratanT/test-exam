package services

import "test-backend/internal/project/repository"

type TestService struct {
	AppRepo *repository.TestRepository
}

func NewTestService(appRepo *repository.TestRepository) *TestService {
	return &TestService{AppRepo: appRepo}
}

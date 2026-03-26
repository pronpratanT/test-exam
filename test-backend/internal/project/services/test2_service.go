package services

import (
	"errors"
	"test-backend/internal/project/dto"
	"test-backend/shared/auth"

	"golang.org/x/crypto/bcrypt"
)

func (s *TestService) Signup(dto *dto.Test2DTOSignup) error {
	duplicate, err := s.AppRepo.CheckDuplicateUsername(dto.Username)
	if err != nil {
		return err
	}
	if duplicate {
		return errors.New("Username already exists")
	}

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(dto.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	dto.Password = string(hashPassword)

	return s.AppRepo.Signup(dto)
}

func (s *TestService) Signin(signinDTO *dto.Test2DTOSignin) (*dto.SigninResponse, error) {
	user, err := s.AppRepo.Signin(signinDTO.Username)
	if err != nil {
		return nil, errors.New("Invalid username or password")
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(signinDTO.Password))
	if err != nil {
		return nil, errors.New("Invalid username or password")
	}
	token, err := auth.GenerateJWT(user.Username)
	if err != nil {
		return nil, errors.New("Failed to generate token")
	}
	return &dto.SigninResponse{
		Username: user.Username,
		Token:    token,
	}, nil
}

package services

import (
	"fmt"
	"test-backend/shared/models"

	"gorm.io/gorm"
)

func (s *TestService) CreateTicket() (*models.Test5, error) {
	tickets, err := s.AppRepo.GetAllTicket()
	if err != nil {
		return nil, err
	}

	var nextQue string

	if len(tickets) == 0 {
		nextQue = "A0"
	} else {
		last := tickets[len(tickets)-1]

		if last.Iscleared {
			nextQue = "A0"
		} else {
			nextQue, err = generateNextQue(last.QueNumber)
			if err != nil {
				return nil, err
			}
		}
	}

	return s.AppRepo.CreateTicket(nextQue)
}

func generateNextQue(current string) (string, error) {
	if len(current) != 2 {
		return "", fmt.Errorf("invalid que number format: %s", current)
	}

	letter := current[0]
	digit := current[1]

	if digit < '9' {
		// A0 → A1, A1 → A2, ...
		return string([]byte{letter, digit + 1}), nil
	} else if letter < 'Z' {
		// A9 → B0, B9 → C0, ...
		return string([]byte{letter + 1, '0'}), nil
	}

	return "", fmt.Errorf("queue is full: reached Z9")
}

func (s *TestService) ClearTicket(queNumber string) error {
	return s.AppRepo.ClearAllTicket(queNumber)
}

func (s *TestService) GetLastTicket() (*models.Test5, error) {
	last, err := s.AppRepo.GetLastTicket()
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return &models.Test5{QueNumber: "00"}, nil
		}
		return nil, err
	}
	if last.Iscleared == true {
		last.QueNumber = "00"
	}
	return last, nil

}

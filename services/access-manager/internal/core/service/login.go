package core

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type LoginInput struct {
	Username string
	Password string
}

func (s *AccessService) Login(input LoginInput) (*string, error) {
	err := input.verify()
	if err != nil {
		return nil, err
	}

	user, err := s.userRepo.Read(input.Username)
	if err != nil {
		return nil, err
	}

	if user.Password != input.Password {
		return nil, errors.New("invalid password")
	}

	token := uuid.New().String()

	s.cache.Set(token, "valid", time.Duration(24*time.Hour))

	return &token, nil
}

func (input LoginInput) verify() error {
	return nil
}

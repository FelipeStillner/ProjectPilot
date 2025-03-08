package core

import (
	"errors"
)

type LoginInput struct {
	Username string
	Password string
}

func (s *AccessService) Login(input LoginInput) (*uint32, error) {
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

	return &user.Id, nil
}

func (input LoginInput) verify() error {
	return nil
}

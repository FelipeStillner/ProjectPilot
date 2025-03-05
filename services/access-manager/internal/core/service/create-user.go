package core

import (
	core "github.com/FelipeStillner/ProjectPilot/services/access-manager/internal/core/struct"
)

type CreateUserInput struct {
	Username string
	Password string
	TeamId   uint32
}

func (s *AccessService) CreateUser(input CreateUserInput) (*core.User, error) {
	err := input.verify()
	if err != nil {
		return nil, err
	}

	user := core.User{
		Username: input.Username,
		Password: input.Password,
		TeamId:   input.TeamId,
	}

	_, err = s.userRepo.Create(user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (input CreateUserInput) verify() error {
	return nil
}

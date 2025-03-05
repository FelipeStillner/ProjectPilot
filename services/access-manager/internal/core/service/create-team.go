package core

import (
	core "github.com/FelipeStillner/ProjectPilot/services/access-manager/internal/core/struct"
	"github.com/google/uuid"
)

type CreateTeamInput struct {
	Name     string
	Username string
	Password string
}

func (s *AccessService) CreateTeam(input CreateTeamInput) (*core.Team, error) {
	err := input.verify()
	if err != nil {
		return nil, err
	}

	team := core.Team{
		Id:   uuid.New().ID(),
		Name: input.Name,
	}

	user := core.User{
		Username: input.Username,
		Password: input.Password,
		TeamId:   team.Id,
	}

	_, err = s.teamRepo.Create(team)
	if err != nil {
		return nil, err
	}

	_, err = s.userRepo.Create(user)
	if err != nil {
		return nil, err
	}

	return &team, nil
}

func (input CreateTeamInput) verify() error {
	return nil
}

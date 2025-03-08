package core

import "github.com/FelipeStillner/ProjectPilot/services/access-manager/internal/port"

type AccessService struct {
	userRepo port.UserRepositoryInterface
	teamRepo port.TeamRepositoryInterface
}

func NewAccessService(userRepo port.UserRepositoryInterface, teamRepo port.TeamRepositoryInterface) *AccessService {
	return &AccessService{userRepo, teamRepo}
}

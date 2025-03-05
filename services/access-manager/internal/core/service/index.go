package core

import "github.com/FelipeStillner/ProjectPilot/services/access-manager/internal/port"

type AccessService struct {
	userRepo port.UserRepositoryInterface
	teamRepo port.TeamRepositoryInterface
	cache    port.CacheInterface
}

func NewAccessService(userRepo port.UserRepositoryInterface, teamRepo port.TeamRepositoryInterface, cache port.CacheInterface) *AccessService {
	return &AccessService{userRepo, teamRepo, cache}
}

package port

import c "github.com/FelipeStillner/ProjectPilot/services/access-manager/internal/core/struct"

type UserRepositoryInterface interface {
	Create(tausersk c.User) (*c.User, error)
}

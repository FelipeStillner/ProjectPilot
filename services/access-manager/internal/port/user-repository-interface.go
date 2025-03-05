package port

import c "github.com/FelipeStillner/ProjectPilot/services/access-manager/internal/core/struct"

type UserRepositoryInterface interface {
	Create(user c.User) (*c.User, error)
	Read(username string) (*c.User, error)
}

package port

import c "github.com/FelipeStillner/ProjectPilot/services/access-manager/internal/core/struct"

type TeamRepositoryInterface interface {
	Create(team c.Team) (*c.Team, error)
}

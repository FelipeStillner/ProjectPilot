package port

import c "github.com/FelipeStillner/ProjectPilot/services/calendar-manager/internal/core/struct"

type IntegrationInterface interface {
	Create(event c.Event) error
	Update(event c.Event) error
	Delete(event c.Event) error
}

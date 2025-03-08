package port

import c "github.com/FelipeStillner/ProjectPilot/services/calendar-manager/internal/core/struct"

type EventRepositoryInterface interface {
	Create(event c.Event) (*c.Event, error)
	Read(id uint32) (*c.Event, error)
	Update(id uint32, event c.Event) (*c.Event, error)
	Delete(id uint32) error
}

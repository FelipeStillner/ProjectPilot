package core

import (
	"github.com/FelipeStillner/ProjectPilot/services/calendar-manager/internal/port"
)

type CalendarService struct {
	eventRepo    port.EventRepositoryInterface
	integrations []port.IntegrationInterface
}

func NewCalendarService(eventRepo port.EventRepositoryInterface, integrations []port.IntegrationInterface) *CalendarService {
	return &CalendarService{eventRepo, integrations}
}

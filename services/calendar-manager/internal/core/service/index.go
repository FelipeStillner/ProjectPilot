package core

import "github.com/FelipeStillner/ProjectPilot/services/calendar-manager/internal/port"

type CalendarService struct {
	eventRepo port.EventRepositoryInterface
}

func NewCalendarService(eventRepo port.EventRepositoryInterface) *CalendarService {
	return &CalendarService{eventRepo}
}
